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
	"fmt"
	"math/big"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"go.uber.org/zap"
)

type MarketV2 struct {
	SerumPadding           [5]byte `json:"-"`
	AccountFlags           bin.Uint64
	OwnAddress             solana.PublicKey
	VaultSignerNonce       bin.Uint64
	BaseMint               solana.PublicKey
	QuoteMint              solana.PublicKey
	BaseVault              solana.PublicKey
	BaseDepositsTotal      bin.Uint64
	BaseFeesAccrued        bin.Uint64
	QuoteVault             solana.PublicKey
	QuoteDepositsTotal     bin.Uint64
	QuoteFeesAccrued       bin.Uint64
	QuoteDustThreshold     bin.Uint64
	RequestQueue           solana.PublicKey
	EventQueue             solana.PublicKey
	Bids                   solana.PublicKey
	Asks                   solana.PublicKey
	BaseLotSize            bin.Uint64
	QuoteLotSize           bin.Uint64
	FeeRateBPS             bin.Uint64
	ReferrerRebatesAccrued bin.Uint64
	EndPadding             [7]byte `json:"-"`
}

func (m *MarketV2) Decode(in []byte) error {
	decoder := bin.NewDecoder(in)
	err := decoder.Decode(&m)
	if err != nil {
		return fmt.Errorf("unpack: %w", err)
	}
	return nil
}

type Orderbook struct {
	SerumPadding [5]byte `json:"-"`
	AccountFlags uint64
	BumpIndex    uint32  `bin:"sizeof=Nodes"`
	ZeroPaddingA [4]byte `json:"-"`
	FreeListLen  uint32
	ZeroPaddingB [4]byte `json:"-"`
	FreeListHead uint32
	Root         uint32
	LeafCount    uint32
	ZeroPaddingC [4]byte `json:"-"`
	Nodes        []*Slab
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

type SlabUninitialized struct {
	Padding  [4]byte  `json:"-"`
	PaddingA [64]byte `json:"-"` // ensure variant is 68 bytes
}

type SlabInnerNode struct {
	PrefixLen uint32
	Key       bin.Uint128
	Children  [2]uint32
	Padding   [40]byte `json:"-"` // ensure variant is 68 bytes
}

type SlabFreeNode struct {
	Next    uint32
	Padding [64]byte `json:"-"` // ensure variant is 68 bytes
}

type SlabLastFreeNode struct {
	Padding  [4]byte  `json:"-"`
	PaddingA [64]byte `json:"-"` // ensure variant is 68 bytes
}

type SlabLeafNode struct {
	OwnerSlot     uint8
	FeeTier       uint8
	Padding       [2]byte `json:"-"`
	Key           bin.Uint128
	Owner         solana.PublicKey
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

type EventQueueHeader struct {
	Serum        [5]byte
	AccountFlags uint64
	Head         uint64
	Count        uint64
	SeqNum       uint64
}

type EventFlag uint8

const (
	EventFlagFill  EventFlag = 0x1
	EventFlagOut   EventFlag = 0x2
	EventFlagBid   EventFlag = 0x4
	EventFlagMaker EventFlag = 0x8
)

type EventSide string

const (
	EventSideAsk EventSide = "ASK"
	EventSideBid EventSide = "BID"
)

type Event struct {
	Flag              EventFlag
	OwnerSlot         uint8
	FeeTier           uint8
	Padding           [5]uint8
	NativeQtyReleased uint64
	NativeQtyPaid     uint64
	NativeFeeOrRebate uint64
	OrderID           bin.Uint128
	Owner             solana.PublicKey
	ClientOrderID     uint64
}

func (e *Event) Side() EventSide {
	if Has(uint8(e.Flag), uint8(EventFlagBid)) {
		return EventSideBid
	}
	return EventSideAsk
}

func (e *Event) Filled() bool {
	return Has(uint8(e.Flag), uint8(EventFlagFill))
}

func Has(b, flag uint8) bool { return b&flag != 0 }

type EventQueue struct {
	Header *EventQueueHeader
	Events []*Event
}

func (q *EventQueue) DecodeFromBase64(b64 string) error {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return fmt.Errorf("event queue: from base64: %w", err)
	}

	return q.Decode(data)
}

const EventDataLength = 88

func (q *EventQueue) Decode(data []byte) error {
	decoder := bin.NewDecoder(data)
	err := decoder.Decode(&q.Header)
	if err != nil {
		return fmt.Errorf("event queue: decode header: %w", err)
	}

	//remainingData := decoder.Remaining()
	//if remainingData%EventDataLength != 0 {
	//	return fmt.Errorf("event queue: wrong event data length %d", remainingData)
	//}
	//
	//eventCount := remainingData / EventDataLength

	for decoder.Remaining() >= 88 {
		var e *Event
		err = decoder.Decode(&e)
		if err != nil {
			return fmt.Errorf("event queue: decode events: %w", err)
		}
		q.Events = append(q.Events, e)
	}

	return nil
}
