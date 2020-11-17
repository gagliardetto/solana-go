package serum

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecoder_Market(t *testing.T) {
	b64 := `c2VydW0DAAAAAAAAAF4kKlwSa8cc6xshYrDN0SrwrDDLBBUwemtddQHhfjgKAQAAAAAAAACL34duLBe2W5K3QFyI1rhNSESYe+cR/nc2UqvgE9x1VMb6evO+2606PWXzaqvJdDGxu+TC0vbg5HymAgNFL11habDAgiZH59TQw5/Y/52i1DhnPZFYOUB4C3G0hhSSXiRAZw8oAwAAAAAAAAAAAAAANvvq/rQwheCOf85MPshRgZEhXzDFAUh3IjalXs/zJ3I5cTmQBAAAABoqGA0AAAAAZAAAAAAAAACuBhNqk2KYdlbj/V5jbAGnnybh+XBss48/P00r053wbACx0Z1WrY+X9jL+huHdyUdpKzL/JScDimaQlNfzjpWANi1Nu6kEazO0bu0NkhnKFyQt2psF0SRCimAVpNimaOjou1Esrd0dKTtLbedHvt62Vi1bRJYveY74GEP6vkH/qBAnAAAAAAAACgAAAAAAAAAAAAAAAAAAAMsrAAAAAAAAcGFkZGluZw==`

	data, err := base64.StdEncoding.DecodeString(b64)
	require.NoError(t, err)

	var m *MarketV2
	err = bin.NewDecoder(data).Decode(&m)
	require.NoError(t, err)

	fmt.Println("market event queue:", m.EventQueue.String())

}
func TestDecoder_Event(t *testing.T) {

	b64 := `c2VydW0RAAAAAAAAAKMoAAAAAAAAAAAAAAAAAAAD+gUAAAAAAAYJAAAAAAAAP9MRAAAAAACwUAY6AAAAAAAAAAAAAAAA/Af4//////8aBQAAAAAAAJGeNN64UdRK+szEsGLeTBiPnrTkfJaOEzsacSpRiAYs6zimHOimKK8CAQAAAAAAAEC6CycAAAAAAAAAAAAAAAAAAAAAAAAAAP/3BwAAAAAAIAUAAAAAAACRnjTeuFHUSvrMxLBi3kwYj5605HyWjhM7GnEqUYgGLH2wFQ01ceby`

	data, err := base64.StdEncoding.DecodeString(b64)
	headerData := data[0:38]
	eventData := data[37:]
	require.NoError(t, err)

	fmt.Println("Size of data:", len(data), hex.EncodeToString(data))

	//read the header
	//read the rest ...
	var h *EventQueueHeader
	err = bin.NewDecoder(headerData).Decode(&h)

	var events [2]*Event
	err = bin.NewDecoder(eventData).Decode(&events)

	require.NoError(t, err)

	for _, e := range events {
		fmt.Println("serum?:", string(h.Serum[:]))
		fmt.Println("Type:", e.Flag)
		fmt.Println("event owner:", e.Owner)
	}

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

	jsonCnt, err := json.MarshalIndent(ob, "", " ")
	require.NoError(t, err)
	fmt.Println(string(jsonCnt))

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
					Lo: 18446744073702344907,
					Hi: 1858,
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
					Lo: 18446744073702358592,
					Hi: 1820,
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
					TypeID: 2,
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
