// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/diff"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecoder_EventQueue_Diff(t *testing.T) {
	//t.Skip("diff event queue test")

	oldDataFile := "testdata/serum-event-queue-old.bin.zst"
	newDataFile := "testdata/serum-event-queue-new.bin.zst"

	olDataJSONFile := strings.ReplaceAll(oldDataFile, ".bin.zst", ".json")
	newDataJSONFile := strings.ReplaceAll(newDataFile, ".bin.zst", ".json")

	if os.Getenv("TESTDATA_UPDATE") == "true" {
		client := rpc.New("http://api.mainnet-beta.solana.com:80/rpc")
		ctx := context.Background()
		account := solana.MustPublicKeyFromBase58("13iGJcA4w5hcJZDjJbJQor1zUiDLE4jv2rMW9HkD5Eo1")

		info, err := client.GetAccountInfo(ctx, account)
		require.NoError(t, err)
		writeCompressedFile(t, oldDataFile, info.Value.Data.GetBinary())

		oldQueue := &EventQueue{}
		require.NoError(t, oldQueue.Decode(info.Value.Data.GetBinary()))
		writeJSONFile(t, olDataJSONFile, oldQueue)

		time.Sleep(900 * time.Millisecond)

		info, err = client.GetAccountInfo(ctx, account)
		require.NoError(t, err)
		writeCompressedFile(t, newDataFile, info.Value.Data.GetBinary())

		newQueue := &EventQueue{}
		require.NoError(t, newQueue.Decode(info.Value.Data.GetBinary()))
		writeJSONFile(t, newDataJSONFile, newQueue)
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
			{OrderID: OrderID(bin.Uint128{Lo: 1})},
			{OrderID: OrderID(bin.Uint128{Lo: 2})},
		},
		EndPadding: [7]byte{},
	}

	newQueue := &EventQueue{
		Head:   120,
		Count:  13,
		SeqNum: 25,
		Events: []*Event{
			{OrderID: OrderID(bin.Uint128{Lo: 1})},
			{OrderID: OrderID(bin.Uint128{Lo: 4})},
			{OrderID: OrderID(bin.Uint128{Lo: 5})},
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
