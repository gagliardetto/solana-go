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
	"strings"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type RequestQueue struct {
	SerumPadding [5]byte `json:"-"`

	AccountFlags AccountFlag
	Head         bin.Uint64 `bin:"sliceoffsetof=Requests,80"`

	Count      bin.Uint64 `bin:"sizeof=Requests"`
	NextSeqNum bin.Uint64
	Requests   []*Request

	EndPadding [7]byte `json:"-"`
}

func (r *RequestQueue) Decode(data []byte) error {
	decoder := bin.NewDecoder(data)
	return decoder.Decode(&r)
}

func (q *RequestQueue) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	if err = decoder.SkipBytes(5); err != nil {
		return err
	}

	if err = decoder.Decode(&q.AccountFlags); err != nil {
		return err
	}
	if err = decoder.Decode(&q.Head); err != nil {
		return err
	}

	if err = decoder.Decode(&q.Count); err != nil {
		return err
	}

	if err = decoder.Decode(&q.NextSeqNum); err != nil {
		return err
	}

	ringbufStartByte := decoder.Position()
	ringbufByteSize := uint(decoder.Remaining() - 7)
	ringbugLength := ringbufByteSize / EVENT_BYTE_SIZE

	q.Requests = make([]*Request, q.Count)

	for i := uint(0); i < uint(q.Count); i++ {
		itemIndex := ((uint(q.Head) + i) % ringbugLength)
		offset := ringbufStartByte + (itemIndex * EVENT_BYTE_SIZE)
		if err = decoder.SetPosition(offset); err != nil {
			return err
		}

		if err = decoder.Decode(&q.Requests[i]); err != nil {
			return err
		}
	}

	return nil
}

// TODO: fill up later
func (q *RequestQueue) MarshalBinary(encoder *bin.Encoder) error {
	return nil
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

func (f RequestFlag) String() string {
	var flags []string

	if f.IsNewOrder() {
		flags = append(flags, "NEW_ORDER")
	}
	if f.IsCancelOrder() {
		flags = append(flags, "CANCEL_ORDER")
	}
	if f.IsBid() {
		flags = append(flags, "BID")
	} else {
		flags = append(flags, "ASK")
	}
	if f.IsPostOnly() {
		flags = append(flags, "POST_ONLY")
	}
	if f.IsImmediateOrCancel() {
		flags = append(flags, "IMMEDIATE_OR_CANCEL")
	}
	if f.IsDecrementTakeOnSelfTrade() {
		flags = append(flags, "DECR_TAKE_ON_SELF")
	}
	return strings.Join(flags, " | ")
}

func (r RequestFlag) IsNewOrder() bool {
	return Has(uint8(r), uint8(RequestFlagNewOrder))
}

func (r RequestFlag) IsCancelOrder() bool {
	return Has(uint8(r), uint8(RequestFlagCancelOrder))
}

func (r RequestFlag) IsBid() bool {
	return Has(uint8(r), uint8(RequestFlagBid))
}

func (r RequestFlag) IsPostOnly() bool {
	return Has(uint8(r), uint8(RequestFlagPostOnly))
}

func (r RequestFlag) IsImmediateOrCancel() bool {
	return Has(uint8(r), uint8(RequestFlagImmediateOrCancel))
}

func (r RequestFlag) IsDecrementTakeOnSelfTrade() bool {
	return Has(uint8(r), uint8(RequestFlagDecrementTakeOnSelfTrade))
}

// Size 80 byte
type Request struct {
	RequestFlags         RequestFlag
	OwnerSlot            uint8
	FeeTier              uint8
	SelfTradeBehavior    uint8
	Padding              [4]byte    `json:"-"`
	MaxCoinQtyOrCancelId bin.Uint64 //the max amount you wish to buy or sell
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

// -> 262144 + 12 bytes
// -> 12 bytes of serum padding -> 262144
// -> 262144 = HEADER + RING-BUFFER
// -> HEADER = 32 bytes
// -> RING BUFF = (262144 - 32)  262112
// -> ring buf (262144 - 32) -> 262112 -> max number of event
type EventQueue struct {
	SerumPadding [5]byte `json:"-"`

	AccountFlags AccountFlag
	Head         bin.Uint64 `bin:"sliceoffsetof=Events,88"`
	Count        bin.Uint64 `bin:"sizeof=Events"`
	SeqNum       bin.Uint64
	Events       []*Event

	EndPadding [7]byte `json:"-"`
}

func (q *EventQueue) Decode(data []byte) error {
	decoder := bin.NewDecoder(data)
	return decoder.Decode(&q)
}

func (q *EventQueue) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	if err = decoder.SkipBytes(5); err != nil {
		return err
	}

	if err = decoder.Decode(&q.AccountFlags); err != nil {
		return err
	}
	if err = decoder.Decode(&q.Head); err != nil {
		return err
	}

	if err = decoder.Decode(&q.Count); err != nil {
		return err
	}

	if err = decoder.Decode(&q.SeqNum); err != nil {
		return err
	}

	ringbufStartByte := decoder.Position()
	ringbufByteSize := uint(decoder.Remaining() - 7)
	ringbugLength := ringbufByteSize / EVENT_BYTE_SIZE

	q.Events = make([]*Event, q.Count)

	for i := uint(0); i < uint(q.Count); i++ {
		itemIndex := ((uint(q.Head) + i) % ringbugLength)
		offset := ringbufStartByte + (itemIndex * EVENT_BYTE_SIZE)
		if err = decoder.SetPosition(offset); err != nil {
			return err
		}

		if err = decoder.Decode(&q.Events[i]); err != nil {
			return err
		}
	}

	return nil
}

// TODO: fill up later
func (q *EventQueue) MarshalBinary(encoder *bin.Encoder) error {
	return nil
}

type EventFlag uint8

const (
	EventFlagFill  = 0x1
	EventFlagOut   = 0x2
	EventFlagBid   = 0x4
	EventFlagMaker = 0x8

	// Added in DEX v3

	EventFlagReleaseFunds = 0x10
)

func (e EventFlag) IsFill() bool {
	return Has(uint8(e), uint8(EventFlagFill))
}

func (e EventFlag) IsOut() bool {
	return Has(uint8(e), uint8(EventFlagOut))
}

func (e EventFlag) IsBid() bool {
	return Has(uint8(e), uint8(EventFlagBid))
}

func (e EventFlag) IsMaker() bool {
	return Has(uint8(e), uint8(EventFlagMaker))
}

func (e EventFlag) IsReleaseFunds() bool {
	return Has(uint8(e), uint8(EventFlagReleaseFunds))
}

func (e EventFlag) String() string {
	var flags []string
	if e.IsFill() {
		flags = append(flags, "FILL")
	}
	if e.IsOut() {
		flags = append(flags, "OUT")
	}
	if e.IsBid() {
		flags = append(flags, "BID")
	}
	if e.IsMaker() {
		flags = append(flags, "MAKER")
	}
	if e.IsReleaseFunds() {
		flags = append(flags, "RELEASE_FUNDS")
	}

	return strings.Join(flags, " | ")
}

const EVENT_BYTE_SIZE = uint(88)

type Event struct {
	Flag              EventFlag
	OwnerSlot         uint8
	FeeTier           uint8
	Padding           [5]uint8
	NativeQtyReleased uint64 // the amount you should release (free to settle)
	NativeQtyPaid     uint64 // The amount out of your account
	NativeFeeOrRebate uint64 // maker etc...
	OrderID           OrderID
	Owner             solana.PublicKey // OpenOrder Account address NOT trader
	ClientOrderID     uint64
}

/* Fill Event*/
// BID: Buying Coin paying in PC
// NativeQtyPaid will deplete PC
// NativeQtyReleased: (native_qty_received) amount of Coin you received
//	will incremet natice_coin_total, native_coin_free

// ASK: Buying PC paying in COIN
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

func Has(b, flag uint8) bool { return b&flag == flag }
