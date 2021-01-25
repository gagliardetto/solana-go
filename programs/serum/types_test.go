package serum

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountFlag_Decoder(t *testing.T) {
	hexStr := "0300000000000000"
	data, err := hex.DecodeString(hexStr)
	require.NoError(t, err)

	var f *AccountFlag
	err = bin.NewDecoder(data).Decode(&f)
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
	err = bin.NewDecoder(data).Decode(&f2)
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
	err = bin.NewDecoder(data).Decode(&m)
	require.NoError(t, err)

	assert.Equal(t, true, m.AccountFlags.Is(AccountFlagInitialized))
	assert.Equal(t, true, m.AccountFlags.Is(AccountFlagMarket))
	assert.Equal(t, false, m.AccountFlags.Is(AccountFlagEventQueue))
	assert.Equal(t, solana.MustPublicKeyFromBase58("13iGJcA4w5hcJZDjJbJQor1zUiDLE4jv2rMW9HkD5Eo1"), m.EventQueue)
}

func TestDecoder_Orderbook(t *testing.T) {
	cnt, err := ioutil.ReadFile("./testdata/orderbook.hex")
	require.NoError(t, err)

	data, err := hex.DecodeString(string(cnt))
	require.NoError(t, err)

	decoder := bin.NewDecoder(data)
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
			TypeID: 1,
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
			TypeID: 3,
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
			TypeID: 2,
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
					TypeID: 1,
					Impl: &SlabInnerNode{
						PrefixLen: 53,
						Key: bin.Uint128{
							Lo: 1345,
							Hi: 18446744073703983873,
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
					TypeID: 2,
					Impl: &SlabLeafNode{
						OwnerSlot: 20,
						FeeTier:   6,
						Padding:   [2]byte{0x00, 0x00},
						Key: bin.Uint128{
							Lo: 1827,
							Hi: 18446744073703640754,
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
					TypeID: 3,
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

			decoder := bin.NewDecoder(cnt)
			var slab *Slab
			err = decoder.Decode(&slab)
			require.NoError(t, err)

			assert.Equal(t, test.expectSlab, slab)
		})

	}
}

func pad(count uint) []byte {
	out := make([]byte, count)
	for i := 0; i < int(count); i++ {
		out[i] = 0x00
	}
	return out
}
