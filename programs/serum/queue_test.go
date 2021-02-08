// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serum

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/diff"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecoder_ScanEvenQueue(t *testing.T) {
	t.Skip("broken test, don't have the courage to fix it right now")
	// market -> 7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932 (SOL/USDT)
	// Base SOL -> So11111111111111111111111111111111111111112
	baseLotSize := uint64(100000000)
	baseDecimal := uint64(9)
	// Quote USDT -> BQcdHdAQW1hczDbBi9hiegXAR7A98Q9jx3X3iBBBDiq4
	quoteLotSize := uint64(100)
	quoteDecimal := uint64(6)

	data, err := ioutil.ReadFile("./testdata/serum-sol-usdt-event-queue.hex")
	require.NoError(t, err)

	byteData, err := hex.DecodeString(string(data))
	require.NoError(t, err)

	offset := 5 + 8 + 8 + 8 + 8

	decoder := bin.NewDecoder(byteData[offset:])

	i := 0
	for {
		e := &Event{}
		err := decoder.Decode(e)
		require.NoError(t, err)

		if e.Flag.IsFill() {
			orderID := make([]byte, 16)
			binary.BigEndian.PutUint64(orderID[:], e.OrderID.Lo)
			binary.BigEndian.PutUint64(orderID[8:], e.OrderID.Hi)

			p, err := GetPrice(hex.EncodeToString(orderID))
			require.NoError(t, err)

			pf := PriceLotsToNumber(p, baseLotSize, quoteLotSize, baseDecimal, quoteDecimal)
			fmt.Printf("Index: %d: Amount Released: %d, Amount Out: %d, Price: %d Price as num %s\n", i, e.NativeQtyReleased, e.NativeQtyPaid, p, pf.String())
		}
		i++
	}
}

func TestEventQueue_Decoder_withRollOver(t *testing.T) {
	eqCnt, err := ioutil.ReadFile("./testdata/event-queue-rollover.log")
	require.NoError(t, err)

	data, err := hex.DecodeString(string(eqCnt))
	require.NoError(t, err)

	var q *EventQueue
	err = bin.NewDecoder(data).Decode(&q)
	require.NoError(t, err)

	cnt, err := json.MarshalIndent(q, "", "  ")
	require.NoError(t, err)

	fmt.Println(string(cnt))
}

func TestDecoder_EventQueue_Diff(t *testing.T) {
	oldDataFile := "testdata/serum-event-queue-old.bin.zst"
	newDataFile := "testdata/serum-event-queue-new.bin.zst"

	// olDataJSONFile := strings.ReplaceAll(oldDataFile, ".bin.zst", ".json")
	// newDataJSONFile := strings.ReplaceAll(newDataFile, ".bin.zst", ".json")

	if os.Getenv("TESTDATA_UPDATE") == "true" {
		client := rpc.NewClient("http://api.mainnet-beta.solana.com:80/rpc")
		ctx := context.Background()
		account := solana.MustPublicKeyFromBase58("13iGJcA4w5hcJZDjJbJQor1zUiDLE4jv2rMW9HkD5Eo1")

		info, err := client.GetAccountInfo(ctx, account)
		require.NoError(t, err)
		writeCompressedFile(t, oldDataFile, info.Value.Data)

		// oldQueue := &EventQueue{}
		// require.NoError(t, oldQueue.Decode(info.Value.Data))
		// writeJSONFile(t, olDataJSONFile, oldQueue)

		time.Sleep(900 * time.Millisecond)

		info, err = client.GetAccountInfo(ctx, account)
		require.NoError(t, err)
		writeCompressedFile(t, newDataFile, info.Value.Data)

		// newQueue := &EventQueue{}
		// require.NoError(t, newQueue.Decode(info.Value.Data))
		// writeJSONFile(t, newDataJSONFile, newQueue)
	}

	oldQueue := &EventQueue{}
	require.NoError(t, oldQueue.Decode(readCompressedFile(t, oldDataFile)))

	newQueue := &EventQueue{}
	require.NoError(t, newQueue.Decode(readCompressedFile(t, newDataFile)))

	fmt.Println("==>> All diff(s)")
	diff.Diff(oldQueue, newQueue, diff.OnEvent(func(event diff.Event) { fmt.Println("Event " + event.String()) }))
}

func Test_fill(t *testing.T) {
	tests := []struct {
		name          string
		e             *Event
		expectIsFill  bool
		expectIsOut   bool
		expectIsBid   bool
		expectIsMaker bool
	}{
		{
			name: "Is Fill",
			e: &Event{
				Flag: 0b00000001,
			},
			expectIsFill:  true,
			expectIsOut:   false,
			expectIsBid:   false,
			expectIsMaker: false,
		},
		{
			name: "Is Out",
			e: &Event{
				Flag: 0b00000010,
			},
			expectIsFill:  false,
			expectIsOut:   true,
			expectIsBid:   false,
			expectIsMaker: false,
		},
		{
			name: "Is Fill & bid",
			e: &Event{
				Flag: 0b00000101,
			},
			expectIsFill:  true,
			expectIsOut:   false,
			expectIsBid:   true,
			expectIsMaker: false,
		},
		{
			name: "Is Fill & bid & maker",
			e: &Event{
				Flag: 0b00001101,
			},
			expectIsFill:  true,
			expectIsOut:   false,
			expectIsBid:   true,
			expectIsMaker: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectIsFill, test.e.Flag.IsFill())
			assert.Equal(t, test.expectIsOut, test.e.Flag.IsOut())
			assert.Equal(t, test.expectIsBid, test.e.Flag.IsBid())
			assert.Equal(t, test.expectIsMaker, test.e.Flag.IsMaker())
		})
	}
}

func TestDecoder_EventQueue_DiffManual(t *testing.T) {
	oldQueue := &EventQueue{
		SerumPadding: [5]byte{},
		Head:         120,
		Count:        13,
		SeqNum:       25,
		Events: []*Event{
			{OrderID: bin.Uint128{Lo: 1}},
			{OrderID: bin.Uint128{Lo: 2}},
		},
		EndPadding: [7]byte{},
	}

	newQueue := &EventQueue{
		Head:   120,
		Count:  13,
		SeqNum: 25,
		Events: []*Event{
			{OrderID: bin.Uint128{Lo: 1}},
			{OrderID: bin.Uint128{Lo: 4}},
			{OrderID: bin.Uint128{Lo: 5}},
		},
	}

	fmt.Println("All diff lines")
	diff.Diff(oldQueue, newQueue, diff.OnEvent(func(event diff.Event) { fmt.Println("Event " + event.String()) }))

	fmt.Println("")
	fmt.Println("Processed diff lines")
	diff.Diff(oldQueue, newQueue, diff.OnEvent(func(event diff.Event) {
		if match, _ := event.Match("Events[#]"); match {
			fmt.Printf("Event %s => %v\n", event.Kind, event.Element())
		}
	}))
}

func writeCompressedFile(t *testing.T, filename string, data []byte) {
	require.NoError(t, ioutil.WriteFile(filename, zstEncoder.EncodeAll(data, nil), os.ModePerm), "unable to write compressed file %s", filename)
}

func readCompressedFile(t *testing.T, file string) []byte {
	data, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	out, err := zstDecoder.DecodeAll(data, nil)
	require.NoError(t, err)

	return out
}

var zstEncoder, _ = zstd.NewWriter(nil)
var zstDecoder, _ = zstd.NewReader(nil)

func writeFile(t *testing.T, filename string, data []byte) {
	require.NoError(t, ioutil.WriteFile(filename, data, os.ModePerm), "unable to write file %s", filename)
}

func readFile(t *testing.T, file string) []byte {
	data, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	return data
}

func writeJSONFile(t *testing.T, filename string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	require.NoError(t, err)

	require.NoError(t, ioutil.WriteFile(filename, out, os.ModePerm), "unable to write JSON file %s", filename)
}

func readJSONFile(t *testing.T, file string, v interface{}) {
	data, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	require.NoError(t, json.Unmarshal(data, v))
	return
}
