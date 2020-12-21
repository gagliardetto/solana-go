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
	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

type RequestQueue struct {
	SerumPadding [5]byte `json:"-"`

	AccountFlags AccountFlag
	Head         bin.Uint64
	Count        bin.Uint64 `bin:"sizeof=Requests"`
	NextSeqNum   bin.Uint64
	Requests     []*Request

	EndPadding [7]byte `json:"-"`
}

func (r *RequestQueue) Decode(data []byte) error {
	decoder := bin.NewDecoder(data)
	return decoder.Decode(&r)
}

type RequestFlag uint8

const (
	RequestFlagNewOrder = RequestFlag(1 << iota)
	RequestFlagCancelOrder
	RequestFlagBid
	RequestFlagPostOnly
	RequestFlagImmediateOrCancel
	RequestFlagDecrementTakeOnSelfTrade
)

type Request struct {
	RequestFlags         RequestFlag
	OwnerSlot            uint8
	FeeTier              uint8
	SelfTradeBehavior    uint8
	Padding              [4]byte `json:"-"`
	MaxCoinQtyOrCancelId bin.Uint64
	NativePCQtyLocked    bin.Uint64
	OrderID              bin.Uint128
	OpenOrders           [4]bin.Uint64 // this is the openOrder address
	ClientOrderID        bin.Uint64
}

func (r *Request) Equal(other *Request) bool {
	//return (r.OrderID.Hi == other.OrderID.Hi && r.OrderID.Lo == other.OrderID.Lo) &&
	//	(r.MaxCoinQtyOrCancelId == other.MaxCoinQtyOrCancelId) &&
	//	(r.NativePCQtyLocked == other.NativePCQtyLocked)
	return (r.OrderID.Hi == other.OrderID.Hi && r.OrderID.Lo == other.OrderID.Lo) &&
		(r.MaxCoinQtyOrCancelId == other.MaxCoinQtyOrCancelId) &&
		(r.NativePCQtyLocked == other.NativePCQtyLocked)
}

type EventQueue struct {
	SerumPadding [5]byte `json:"-"`

	AccountFlags AccountFlag
	Head         bin.Uint64
	Count        bin.Uint64 `bin:"sizeof=Events"`
	SeqNum       bin.Uint64
	Events       []*Event

	EndPadding [7]byte `json:"-"`
}

func (q *EventQueue) Decode(data []byte) error {
	decoder := bin.NewDecoder(data)
	return decoder.Decode(&q)
}

type EventFlag uint8

const (
	EventFlagFill = EventFlag(1 << iota)
	EventFlagOut
	EventFlagBid
	EventFlagMaker
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

func (e *Event) Equal(other *Event) bool {
	return e.OrderID.Hi == other.OrderID.Hi && e.OrderID.Lo == other.OrderID.Lo
}

func (e *Event) Side() Side {
	if Has(uint8(e.Flag), uint8(EventFlagBid)) {
		return SideBid
	}
	return SideAsk
}

func (e *Event) Filled() bool {
	return Has(uint8(e.Flag), uint8(EventFlagFill))
}

func Has(b, flag uint8) bool { return b&flag != 0 }
