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
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountFlag_Decoder(t *testing.T) {
	hexStr := "0300000000000000"
	data, err := hex.DecodeString(hexStr)
	require.NoError(t, err)

	var f *AccountFlag
	err = bin.NewBinDecoder(data).Decode(&f)
	require.NoError(t, err)

	assert.Equal(t, f.Is(AccountFlagInitialized), true, "initialized")
	assert.Equal(t, f.Is(AccountFlagMarket), true, "market")
	assert.Equal(t, f.Is(AccountFlagOpenOrders), false, "openOrders")
	assert.Equal(t, f.Is(AccountFlagRequestQueue), false, "requestQueue")
	assert.Equal(t, f.Is(AccountFlagEventQueue), false, "eventQueue")
	assert.Equal(t, f.Is(AccountFlagBids), false, "bids")
	assert.Equal(t, f.Is(AccountFlagAsks), false, "asks")
	assert.Equal(t, f.Is(AccountFlagDisabled), false, "disabled")

	hexStr = "0900000000000000"
	data, err = hex.DecodeString(hexStr)
	require.NoError(t, err)

	var f2 *AccountFlag
	err = bin.NewBinDecoder(data).Decode(&f2)
	require.NoError(t, err)

	assert.Equal(t, f2.Is(AccountFlagInitialized), true, "initialized")
	assert.Equal(t, f2.Is(AccountFlagMarket), false, "market")
	assert.Equal(t, f2.Is(AccountFlagOpenOrders), false, "openOrders")
	assert.Equal(t, f2.Is(AccountFlagRequestQueue), true, "requestQueue")
	assert.Equal(t, f2.Is(AccountFlagEventQueue), false, "eventQueue")
	assert.Equal(t, f2.Is(AccountFlagBids), false, "bids")
	assert.Equal(t, f2.Is(AccountFlagAsks), false, "asks")
	assert.Equal(t, f2.Is(AccountFlagDisabled), false, "disabled")
}

func TestDecoder_Market(t *testing.T) {
	b64 := `c2VydW0DAAAAAAAAAF4kKlwSa8cc6xshYrDN0SrwrDDLBBUwemtddQHhfjgKAQAAAAAAAACL34duLBe2W5K3QFyI1rhNSESYe+cR/nc2UqvgE9x1VMb6evO+2606PWXzaqvJdDGxu+TC0vbg5HymAgNFL11habDAgiZH59TQw5/Y/52i1DhnPZFYOUB4C3G0hhSSXiRAZw8oAwAAAAAAAAAAAAAANvvq/rQwheCOf85MPshRgZEhXzDFAUh3IjalXs/zJ3I5cTmQBAAAABoqGA0AAAAAZAAAAAAAAACuBhNqk2KYdlbj/V5jbAGnnybh+XBss48/P00r053wbACx0Z1WrY+X9jL+huHdyUdpKzL/JScDimaQlNfzjpWANi1Nu6kEazO0bu0NkhnKFyQt2psF0SRCimAVpNimaOjou1Esrd0dKTtLbedHvt62Vi1bRJYveY74GEP6vkH/qBAnAAAAAAAACgAAAAAAAAAAAAAAAAAAAMsrAAAAAAAAcGFkZGluZw==`

	data, err := base64.StdEncoding.DecodeString(b64)
	require.NoError(t, err)
	fmt.Println(hex.EncodeToString(data))

	var m *MarketV2
	err = bin.NewBinDecoder(data).Decode(&m)
	require.NoError(t, err)

	assert.Equal(t, true, m.AccountFlags.Is(AccountFlagInitialized))
	assert.Equal(t, true, m.AccountFlags.Is(AccountFlagMarket))
	assert.Equal(t, false, m.AccountFlags.Is(AccountFlagEventQueue))
	assert.Equal(t, solana.MustPublicKeyFromBase58("13iGJcA4w5hcJZDjJbJQor1zUiDLE4jv2rMW9HkD5Eo1"), m.EventQueue)
}

func TestDecoder_Orderbook(t *testing.T) {
	t.Skip("long running test")
	cnt, err := ioutil.ReadFile("./testdata/orderbook.hex")
	require.NoError(t, err)

	data, err := hex.DecodeString(string(cnt))
	require.NoError(t, err)

	decoder := bin.NewBinDecoder(data)
	var ob *Orderbook
	err = decoder.Decode(&ob)
	require.NoError(t, err)

	assert.Equal(t, uint32(101), ob.BumpIndex)
	assert.Equal(t, uint32(68), ob.FreeListLen)
	assert.Equal(t, uint32(37), ob.FreeListHead)
	assert.Equal(t, uint32(17), ob.LeafCount)
	assert.Equal(t, 101, len(ob.Nodes))
	assert.Equal(t, &Slab{
		BaseVariant: bin.BaseVariant{
			TypeID: bin.TypeIDFromUint32(1, binary.LittleEndian),
			Impl: &SlabInnerNode{
				PrefixLen: 57,
				Key: bin.Uint128{
					Lo: 1858,
					Hi: 18446744073702344907,
				},
				Children: [2]uint32{55, 56},
				Padding: [40]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}, ob.Nodes[0])
	assert.Equal(t, &Slab{
		BaseVariant: bin.BaseVariant{
			TypeID: bin.TypeIDFromUint32(3, binary.LittleEndian),
			Impl: &SlabFreeNode{
				Next: 2,
				Padding: [64]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}, ob.Nodes[1])
	assert.Equal(t, &Slab{
		BaseVariant: bin.BaseVariant{
			TypeID: bin.TypeIDFromUint32(2, binary.LittleEndian),
			Impl: &SlabLeafNode{
				OwnerSlot: 1,
				FeeTier:   5,
				Padding:   [2]byte{0x00, 0x00},
				Key: bin.Uint128{
					Lo: 1820,
					Hi: 18446744073702358592,
				},
				Owner:         solana.MustPublicKeyFromBase58("77jtrBDbUvwsdNKeq1ERUBcg8kk2hNTzf5E4iRihNgTh"),
				Quantity:      38918,
				ClientOrderId: 14114313590397044635,
			},
		},
	}, ob.Nodes[5])

}

func TestDecoder_Slabs(t *testing.T) {
	tests := []struct {
		name       string
		slabData   string
		expectSlab *Slab
	}{
		{
			name:     "inner_node",
			slabData: "0100000035000000010babffffffffff4105000000000000400000003f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expectSlab: &Slab{
				BaseVariant: bin.BaseVariant{
					TypeID: bin.TypeIDFromUint32(1, binary.LittleEndian),
					Impl: &SlabInnerNode{
						PrefixLen: 53,
						Key: bin.Uint128{
							Lo: 18446744073703983873,
							Hi: 1345,
						},
						Children: [2]uint32{
							64,
							63,
						},
					},
				},
			},
		},

		{
			name:     "leaf_node",
			slabData: "0200000014060000b2cea5ffffffffff23070000000000005ae01b52d00a090c6dc6fce8e37a225815cff2223a99c6dfdad5aae56d3db670e62c000000000000140b0fadcf8fcebf",
			expectSlab: &Slab{
				BaseVariant: bin.BaseVariant{
					TypeID: bin.TypeIDFromUint32(2, binary.LittleEndian),
					Impl: &SlabLeafNode{
						OwnerSlot: 20,
						FeeTier:   6,
						Padding:   [2]byte{0x00, 0x00},
						Key: bin.Uint128{
							Lo: 18446744073703640754,
							Hi: 1827,
						},
						Owner:         solana.MustPublicKeyFromBase58("77jtrBDbUvwsdNKeq1ERUBcg8kk2hNTzf5E4iRihNgTh"),
						Quantity:      11494,
						ClientOrderId: 13821142428571077396,
					},
				},
			},
		},
		{
			name:     "free_node",
			slabData: "030000003400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expectSlab: &Slab{
				BaseVariant: bin.BaseVariant{
					TypeID: bin.TypeIDFromUint32(3, binary.LittleEndian),
					Impl: &SlabFreeNode{
						Next: 52,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cnt, err := hex.DecodeString(test.slabData)
			require.NoError(t, err)

			decoder := bin.NewBinDecoder(cnt)
			var slab *Slab
			err = decoder.Decode(&slab)
			require.NoError(t, err)

			assert.Equal(t, test.expectSlab, slab)
		})

	}
}

func TestOrderID(t *testing.T) {
	orderID, err := NewOrderID("000000000000b868ffffffffff8ee7a0")
	require.NoError(t, err)
	assert.Equal(t, uint64(0xffffffffff8ee7a0), orderID.Lo)
	assert.Equal(t, uint64(0xb868), orderID.Hi)

	assert.Equal(t, uint64(47208), orderID.Price())
	assert.Equal(t, uint64(7411807), orderID.SeqNum(SideBid))

	assert.Equal(t, "000000000000b868ffffffffff8ee7a0", orderID.HexString(false))
	assert.Equal(t, "0x000000000000b868ffffffffff8ee7a0", orderID.HexString(true))

	orderID, err = NewOrderID("00000000000193c100000000000fbcc2")
	require.NoError(t, err)
	assert.Equal(t, uint64(0xfbcc2), orderID.Lo)
	assert.Equal(t, uint64(0x193c1), orderID.Hi)

	assert.Equal(t, uint64(103361), orderID.Price())
	assert.Equal(t, uint64(1031362), orderID.SeqNum(SideAsk))

	assert.Equal(t, "00000000000193c100000000000fbcc2", orderID.HexString(false))
	assert.Equal(t, "0x00000000000193c100000000000fbcc2", orderID.HexString(true))

}

func Test_OpenOrder_GetOrder(t *testing.T) {
	openOrderData := "testdata/serum-open-orders-new.hex"

	openOrders := &OpenOrders{}
	require.NoError(t, openOrders.Decode(readHexFile(t, openOrderData)))
	o := openOrders.GetOrder(20)
	assert.Equal(t, &Order{
		ID: OrderID{
			Hi: 0x0000000000000840,
			Lo: 0xffffffffffacdefd,
		},
		Side: SideBid,
	}, o)
	assert.Equal(t, o.SeqNum(), uint64(5447938))
	assert.Equal(t, o.Price(), uint64(2112))
}

func TestIsBitZero(t *testing.T) {
	tests := []struct {
		name        string
		value       bin.Uint128
		index       uint32
		expect      bool
		expectError bool
	}{
		{
			name: "Index 0, bit is 1",
			value: bin.Uint128{
				Hi: 0x0000000000000000,
				Lo: 0x0000000000000001,
			},
			index:  0,
			expect: false,
		},
		{
			name: "Index 0, bit is 0",
			value: bin.Uint128{
				Hi: 0x0000000000000000,
				Lo: 0x0000000000000000,
			},
			index:  0,
			expect: true,
		},
		{
			name: "Index less then 64, bit is 1",
			value: bin.Uint128{
				Hi: 0x0000000000000000,
				Lo: 0x0000000000100000,
			},
			index:  20,
			expect: false,
		},
		{
			name: "Index less then 64, bit is 1",
			value: bin.Uint128{
				Hi: 0xffffffffffffffff,
				Lo: 0xffffffffffefffff,
			},
			index:  20,
			expect: true,
		},
		{
			name: "Index 64, bit is 1",
			value: bin.Uint128{
				Hi: 0x0000000000000001,
				Lo: 0x0000000000000000,
			},
			index:  64,
			expect: false,
		},
		{
			name: "Index 64, bit is 0",
			value: bin.Uint128{
				Hi: 0x0000000000000000,
				Lo: 0x0000000000000000,
			},
			index:  64,
			expect: true,
		},
		{
			name: "Index greater 64, bit is 1",
			value: bin.Uint128{
				Hi: 0x0000002000000000,
				Lo: 0x0000000000000000,
			},
			index:  101,
			expect: false,
		},
		{
			name: "Index greater 64, bit is 0",
			value: bin.Uint128{
				Hi: 0xffffffdfffffffff,
				Lo: 0xffffffffffffffff,
			},
			index:  101,
			expect: true,
		},
		{
			name:        "Index 128",
			value:       bin.Uint128{},
			index:       128,
			expectError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := IsBitZero(test.value, test.index)
			if test.expectError {
				require.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, f, test.expect)
			}
		})
	}
}
