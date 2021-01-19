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
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/diff"
	"github.com/dfuse-io/solana-go/rpc"

	"github.com/stretchr/testify/assert"

	bin "github.com/dfuse-io/binary"

	"github.com/stretchr/testify/require"
)

func TestHas(t *testing.T) {

	cases := []struct {
		flag     uint8
		name     string
		expected string
	}{
		{
			name:     "2",
			flag:     uint8(2),
			expected: "OUT",
		},
		{
			name:     "1",
			flag:     uint8(1),
			expected: "FILL",
		},
		{
			name:     "6",
			flag:     uint8(6),
			expected: "OUT | BID",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := EventFlag(c.flag)
			require.Equal(t, f.String(), c.expected)
		})
	}

}

func TestRequestQueue_Decoder(t *testing.T) {
	bas64 := "c2VydW0JAAAAAAAAAAMAAAAAAAAAAAAAAAAAAACVlZMAAAAAAA0DBgEAAAAAECcAAAAAAABYvqFvAAAAAG1qbP//////TwcAAAAAAABbQ4ijQxgyr1dCuGPiALjnM85FHycAbicj2RmLNjNV0Y799DoWgU4WDQYGAQAAAAC4CwAAAAAAAMx8hiEAAAAAbGps//////9RBwAAAAAAAFtDiKNDGDKvV0K4Y+IAuOczzkUfJwBuJyPZGYs2M1XR6LcKOxaBThYNBwYBAAAAAIgTAAAAAAAAQILYNwAAAABramz//////1AHAAAAAAAAW0OIo0MYMq9XQrhj4gC45zPORR8nAG4nI9kZizYzVdE84P86FoFOFgYDAAAAAAAAjpWTAAAAAAAAAAAAAAAAAHtqbP//////TQcAAAAAAABbQ4ijQxgyr1dCuGPiALjnM85FHycAbicj2RmLNjNV0ZgF0ywMgU4WBgYAAAAAAACPlZMAAAAAAAAAAAAAAAAAemps//////9OBwAAAAAAAFtDiKNDGDKvV0K4Y+IAuOczzkUfJwBuJyPZGYs2M1XRkdfdLAyBThYGBwAAAAAAAJCVkwAAAAAAAAAAAAAAAACBamz//////08HAAAAAAAAW0OIo0MYMq9XQrhj4gC45zPORR8nAG4nI9kZizYzVdGFpMuDCYFOFgkABgEAAAAAuAsAAAAAAAAAAAAAAAAAAJGVkwAAAAAAWwcAAAAAAABbQ4ijQxgyr1dCuGPiALjnM85FHycAbicj2RmLNjNV0WFf6ToWgU4WAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcGFkZGluZw=="
	data, err := base64.StdEncoding.DecodeString(bas64)
	require.NoError(t, err)

	fmt.Println(hex.EncodeToString(data))

	var q *RequestQueue
	err = bin.NewDecoder(data).Decode(&q)
	require.NoError(t, err)

	assert.Equal(t, true, q.AccountFlags.Is(AccountFlagRequestQueue))
	assert.Equal(t, true, q.AccountFlags.Is(AccountFlagInitialized))
	assert.Equal(t, bin.Uint64(3), q.Head)
	assert.Equal(t, bin.Uint64(9672085), q.NextSeqNum)
	assert.Equal(t, bin.Uint64(0), q.Count)
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
