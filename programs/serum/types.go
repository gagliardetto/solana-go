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
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
	"go.uber.org/zap"
)

type AccountFlag uint64

const (
	AccountFlagInitialized = AccountFlag(1 << iota)
	AccountFlagMarket
	AccountFlagOpenOrders
	AccountFlagRequestQueue
	AccountFlagEventQueue
	AccountFlagBids
	AccountFlagAsks
	AccountFlagDisabled
)

func (a *AccountFlag) Is(flag AccountFlag) bool { return *a&flag != 0 }
func (a *AccountFlag) String() string {
	status := "unknown"
	account_type := "unknown"
	if a.Is(AccountFlagInitialized) {
		status = "initialized"
	} else if a.Is(AccountFlagDisabled) {
		status = "disabled"
	}
	if a.Is(AccountFlagMarket) {
		account_type = "market"
	} else if a.Is(AccountFlagOpenOrders) {
		account_type = "open orders"
	} else if a.Is(AccountFlagRequestQueue) {
		account_type = "request queue"
	} else if a.Is(AccountFlagEventQueue) {
		account_type = "event queue"
	} else if a.Is(AccountFlagBids) {
		account_type = "bids"
	} else if a.Is(AccountFlagAsks) {
		account_type = "asks"
	}
	return fmt.Sprintf("%s %s", status, account_type)
}

type Side uint32

const (
	SideBid = iota
	SideAsk
)

type MarketV2 struct {
	SerumPadding           [5]byte `json:"-"`
	AccountFlags           AccountFlag
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
	decoder := bin.NewBinDecoder(in)
	err := decoder.Decode(&m)
	if err != nil {
		return fmt.Errorf("unpack: %w", err)
	}
	return nil
}

type Orderbook struct {
	SerumPadding [5]byte `json:"-"`
	AccountFlags AccountFlag
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
	{Name: "uninitialized", Type: (*SlabUninitialized)(nil)},
	{Name: "inner_node", Type: (*SlabInnerNode)(nil)},
	{Name: "leaf_node", Type: (*SlabLeafNode)(nil)},
	{Name: "free_node", Type: (*SlabFreeNode)(nil)},
	{Name: "last_free_node", Type: (*SlabLastFreeNode)(nil)},
})

type Slab struct {
	bin.BaseVariant
}

var _ bin.EncoderDecoder = &Slab{}

func (s *Slab) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	return s.BaseVariant.UnmarshalBinaryVariant(decoder, SlabFactoryImplDef)
}

func (s *Slab) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(s.TypeID.Uint32(), binary.LittleEndian)
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

type OrderID bin.Uint128

func NewOrderID(orderID string) (OrderID, error) {
	d, err := hex.DecodeString(orderID)
	if err != nil {
		return OrderID(bin.Uint128{
			Lo: 0,
			Hi: 0,
		}), fmt.Errorf("unable to decode order ID: %w", err)
	}

	if len(d) < 16 {
		return OrderID(bin.Uint128{
			Lo: 0,
			Hi: 0,
		}), fmt.Errorf("order ID too short expecting at least 16 bytes got %d", len(d))
	}

	return OrderID(bin.Uint128{
		Lo: binary.BigEndian.Uint64(d[8:]),
		Hi: binary.BigEndian.Uint64(d[:8]),
	}), nil
}

func (o OrderID) Price() uint64 {
	return o.Hi
}

func (o OrderID) HexString(withPrefix bool) string {
	number := make([]byte, 16)
	binary.BigEndian.PutUint64(number[:], o.Hi)  // old price
	binary.BigEndian.PutUint64(number[8:], o.Lo) // old seq_number
	str := hex.EncodeToString(number)
	if withPrefix {
		return "0x" + str
	}
	return str
}

func (o OrderID) SeqNum(side Side) uint64 {
	if side == SideBid {
		return ^o.Lo
	}

	return o.Lo
}

type OpenOrders struct {
	SerumPadding           [5]byte `json:"-"`
	AccountFlags           AccountFlag
	Market                 solana.PublicKey
	Owner                  solana.PublicKey
	NativeBaseTokenFree    bin.Uint64
	NativeBaseTokenTotal   bin.Uint64
	NativeQuoteTokenFree   bin.Uint64
	NativeQuoteTokenTotal  bin.Uint64
	FreeSlotBits           bin.Uint128
	IsBidBits              bin.Uint128 // Bids is 1,  Ask is 0
	Orders                 [128]OrderID
	ClientOrderIDs         [128]bin.Uint64
	ReferrerRebatesAccrued bin.Uint64
	EndPadding             [7]byte `json:"-"`
}

type Order struct {
	ID   OrderID
	Side Side
}

func (o *Order) SeqNum() uint64 {
	return o.ID.SeqNum(o.Side)
}

func (o *Order) Price() uint64 {
	return o.ID.Price()
}

func (o *OpenOrders) GetOrder(index uint32) *Order {
	order := &Order{
		ID:   o.Orders[index],
		Side: SideBid,
	}
	isZero, err := IsBitZero(o.IsBidBits, index)
	if err != nil {
		panic("this should never happen")
	}
	if isZero {
		order.Side = SideAsk
	}
	return order
}

func (o *OpenOrders) Decode(in []byte) error {
	decoder := bin.NewBinDecoder(in)
	err := decoder.Decode(&o)
	if err != nil {
		return fmt.Errorf("unpack: %w", err)
	}
	return nil
}

func IsBitZero(v bin.Uint128, bitIndex uint32) (bool, error) {
	if bitIndex > 127 {
		return false, fmt.Errorf("bit index out of range")
	}

	if bitIndex > 63 {
		bitIndex = bitIndex - 64
		mask := uint64(1 << bitIndex)
		return (v.Hi&mask == 0), nil
	}
	mask := uint64(1 << bitIndex)
	return (v.Lo&mask == 0), nil
}
