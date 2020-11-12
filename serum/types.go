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
	"fmt"
	"math/big"

	"go.uber.org/zap"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

type MarketV2 struct {
	/*
	   blob(5),
	   accountFlagsLayout('accountFlags'),
	   publicKeyLayout('ownAddress'),
	   u64('vaultSignerNonce'),
	   publicKeyLayout('baseMint'),
	   publicKeyLayout('quoteMint'),
	   publicKeyLayout('baseVault'),
	   u64('baseDepositsTotal'),
	   u64('baseFeesAccrued'),
	   publicKeyLayout('quoteVault'),
	   u64('quoteDepositsTotal'),
	   u64('quoteFeesAccrued'),
	   u64('quoteDustThreshold'),
	   publicKeyLayout('requestQueue'),
	   publicKeyLayout('eventQueue'),
	   publicKeyLayout('bids'),

	   publicKeyLayout('asks'),
	   u64('baseLotSize'),
	   u64('quoteLotSize'),
	   u64('feeRateBps'),
	   u64('referrerRebatesAccrued'),
	   blob(7),
	*/
	SerumPadding           [5]byte          `json:"-" struc:"[5]pad"`
	AccountFlags           bin.Uint64       `struc:"uint64,little"`
	OwnAddress             solana.PublicKey `struc:"[32]byte"`
	VaultSignerNonce       bin.Uint64       `struc:"uint64,little"`
	BaseMint               solana.PublicKey `struc:"[32]byte"`
	QuoteMint              solana.PublicKey `struc:"[32]byte"`
	BaseVault              solana.PublicKey `struc:"[32]byte"`
	BaseDepositsTotal      bin.Uint64       `struc:"uint64,little"`
	BaseFeesAccrued        bin.Uint64       `struc:"uint64,little"`
	QuoteVault             solana.PublicKey `struc:"[32]byte"`
	QuoteDepositsTotal     bin.Uint64       `struc:"uint64,little"`
	QuoteFeesAccrued       bin.Uint64       `struc:"uint64,little"`
	QuoteDustThreshold     bin.Uint64       `struc:"uint64,little"`
	RequestQueue           solana.PublicKey `struc:"[32]byte"`
	EventQueue             solana.PublicKey `struc:"[32]byte"`
	Bids                   solana.PublicKey `struc:"[32]byte"`
	Asks                   solana.PublicKey `struc:"[32]byte"`
	BaseLotSize            bin.Uint64       `struc:"uint64,little"`
	QuoteLotSize           bin.Uint64       `struc:"uint64,little"`
	FeeRateBPS             bin.Uint64       `struc:"uint64,little"`
	ReferrerRebatesAccrued bin.Uint64       `struc:"uint64,little"`
	EndPadding             [7]byte          `json:"-" struc:"[7]pad"`
}

func (m *MarketV2) Decode(in []byte) error {
	decoder := bin.NewDecoder(in)
	err := decoder.Decode(&m)
	if err != nil {
		return fmt.Errorf("unpack: %w", err)
	}
	return nil
}

type SlabUninitialized struct {
	Padding [4]byte `json:"-"`
}

type SlabInnerNode struct {
	//    u32('prefixLen'),
	//    u128('key'),
	//    seq(u32(), 2, 'children'),
	PrefixLen uint32
	Key       bin.Uint128
	Children  [2]uint32
}

type SlabLeafNode struct {
	OwnerSlot     uint8
	FeeTier       uint8
	Padding       [2]byte `json:"-"`
	Key           bin.Uint128
	Owner         PublicKey
	Quantity      bin.Uint64
	ClientOrderId bin.Uint64
}

func (s *SlabLeafNode) GetPrice() *big.Int {
	raw := s.Key.BigInt().Bytes()
	if len(raw) <= 8 {
		return big.NewInt(0)
	}
	v := new(big.Int).SetBytes(raw[0 : len(raw)-8])
	return v
}

type SlabFreeNode struct {
	Next uint32
}

type SlabLastFreeNode struct {
	//Padding [68]byte `json:"-"`
	Padding [4]byte `json:"-"`
}

type PublicKey [32]byte

var SlabFactoryImplDef = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"uninitialized", (*SlabUninitialized)(nil)},
	{"inner_node", (*SlabInnerNode)(nil)},
	{"leaf_node", (*SlabLeafNode)(nil)},
	{"free_node", (*SlabFreeNode)(nil)},
	{"last_free_node", (*SlabLastFreeNode)(nil)},
})

type Slab struct {
	bin.BaseVariant
}

func (s *Slab) UnmarshalBinary(decoder *bin.Decoder) error {
	return s.BaseVariant.UnmarshalBinaryVariant(decoder, SlabFactoryImplDef)
}
func (s *Slab) MarshalBinary(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(s.TypeID)
	if err != nil {
		return err
	}
	return encoder.Encode(s.Impl)
}

type Orderbook struct {
	// ORDERBOOK_LAYOUT
	SerumPadding [5]byte `json:"-"`
	AccountFlags uint64
	// SLAB_LAYOUT
	// SLAB_HEADER_LAYOUT
	BumpIndex    uint32  `bin:"sizeof=Nodes"`
	ZeroPaddingA [4]byte `json:"-"`
	FreeListLen  uint32
	ZeroPaddingB [4]byte `json:"-"`
	FreeListHead uint32
	Root         uint32
	LeafCount    uint32
	ZeroPaddingC [4]byte `json:"-"`

	// SLAB_NODE_LAYOUT
	Nodes []*Slab
}

func (o *Orderbook) Items(descending bool, f func(node *SlabLeafNode) error) error {
	if o.LeafCount == 0 {
		return nil
	}

	index := uint32(0)
	stack := []uint32{o.Root}
	for len(stack) > 0 {
		index, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if traceEnabled {
			zlog.Debug("looking at slab index", zap.Int("index", int(index)))
		}
		slab := o.Nodes[index]
		impl := slab.Impl
		switch s := impl.(type) {
		case *SlabInnerNode:
			if descending {
				stack = append(stack, s.Children[0], s.Children[1])
			} else {
				stack = append(stack, s.Children[1], s.Children[0])
			}
		case *SlabLeafNode:
			if traceEnabled {
				zlog.Debug("found leaf", zap.Int("leaf", int(index)))
			}
			f(s)
		}
	}
	return nil
}
