package serum

import (
	"bytes"
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
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
	AccountFlags           solana.U64       `struc:"uint64,little"`
	OwnAddress             solana.PublicKey `struc:"[32]byte"`
	VaultSignerNonce       solana.U64       `struc:"uint64,little"`
	BaseMint               solana.PublicKey `struc:"[32]byte"`
	QuoteMint              solana.PublicKey `struc:"[32]byte"`
	BaseVault              solana.PublicKey `struc:"[32]byte"`
	BaseDepositsTotal      solana.U64       `struc:"uint64,little"`
	BaseFeesAccrued        solana.U64       `struc:"uint64,little"`
	QuoteVault             solana.PublicKey `struc:"[32]byte"`
	QuoteDepositsTotal     solana.U64       `struc:"uint64,little"`
	QuoteFeesAccrued       solana.U64       `struc:"uint64,little"`
	QuoteDustThreshold     solana.U64       `struc:"uint64,little"`
	RequestQueue           solana.PublicKey `struc:"[32]byte"`
	EventQueue             solana.PublicKey `struc:"[32]byte"`
	Bids                   solana.PublicKey `struc:"[32]byte"`
	Asks                   solana.PublicKey `struc:"[32]byte"`
	BaseLotSize            solana.U64       `struc:"uint64,little"`
	QuoteLotSize           solana.U64       `struc:"uint64,little"`
	FeeRateBPS             solana.U64       `struc:"uint64,little"`
	ReferrerRebatesAccrued solana.U64       `struc:"uint64,little"`
	EndPadding             [7]byte          `json:"-" struc:"[7]pad"`
}

type Orderbook struct {
	// ORDERBOOK_LAYOUT
	SerumPadding [5]byte    `json:"-" struc:"[5]pad"`
	AccountFlags solana.U64 `struc:"uint64,little"`
	// SLAB_LAYOUT
	// SLAB_HEADER_LAYOUT
	BumpIndex    uint32  `struc:"uint32,little,sizeof=Nodes"`
	ZeroPaddingA [4]byte `json:"-" struc:"[4]pad"`
	FreeListLen  uint32  `struc:"uint32,little"`
	ZeroPaddingB [4]byte `json:"-" struc:"[4]pad"`
	FreeListHead uint32  `struc:"uint32,little"`
	Root         uint32  `struc:"uint32,little"`
	LeafCount    uint32  `struc:"uint32,little"`
	ZeroPaddingC [4]byte `json:"-" struc:"[4]pad"`
	// SLAB_NODE_LAYOUT
	Nodes []SlabNode `struc:""`
}

func (o *Orderbook) Items(descending bool, f func(node *SlabLeafNode) error) error {
	if o.LeafCount == 0 {
		return nil
	}

	index := uint32(0)
	stack := []uint32{o.Root}
	for len(stack) > 0 {
		index, stack = stack[len(stack)-1], stack[:len(stack)-1]
		slab := o.Nodes[index]
		impl, err := slab.Variant()
		if err != nil {
			return err
		}
		switch s := impl.(type) {
		case *SlabInnerNode:
			if descending {
				stack = append(stack, s.Children[0], s.Children[1])
			} else {
				stack = append(stack, s.Children[1], s.Children[0])
			}
		case *SlabLeafNode:
			f(s)
		}
	}
	return nil
}

type SlabNode struct {
	Type uint32 `struc:"uint32,little"`
	Blob []byte `struc:"[68]byte"`
}

func (s *SlabNode) Variant() (interface{}, error) {
	var el interface{}
	switch s.Type {
	case 0:
		el = &SlabUninitialized{}
	case 1:
		el = &SlabInnerNode{}
	case 2:
		el = &SlabLeafNode{}
	case 3:
		el = &SlabFreeNode{}
	case 4:
		el = &SlabLastFreeNode{}
	default:
		return nil, fmt.Errorf("unsupported SlabNode variant %d", s.Type)
	}
	if err := struc.Unpack(bytes.NewReader(s.Blob), el); err != nil {
		return nil, err
	}
	return el, nil
}

type SlabUninitialized struct {
}

type SlabInnerNode struct {
	PrefixLen uint32 `struc:"uint32,little"`
	// this corresponds to the uint128 key
	KeySeq   solana.U64 `struc:"uint64,little"`
	KeyPrice solana.U64 `struc:"uint64,little"`
	Children []uint32   `struc:"[2]uint32,little"`
}

type SlabLeafNode struct {
	//u8('ownerSlot'), // Index into OPEN_ORDERS_LAYOUT.orders
	//u8('feeTier'),
	//blob(2),
	//u128('key'), // (price, seqNum)
	//publicKeyLayout('owner'), // Open orders account
	//u64('quantity'), // In units of lot size
	//u64('clientOrderId'),
	OwnerSlot uint8   `struc:"uint8,little"`
	FeeTier   uint8   `struc:"uint8,little"`
	Padding   [2]byte `json:"-" struc:"[2]pad"`

	Price solana.U128 `struc:"[16]byte,little"`
	//KeySeq   solana.U64 `struc:"uint64",little"`
	//KeyPrice solana.U64 `struc:"uint64",little"`

	Owner         solana.PublicKey
	Quantity      solana.U64 `struc:"uint64,little"`
	ClientOrderId solana.U64 `struc:"uint64,little"`
}

type SlabFreeNode struct {
	Next uint32 `struc:"uint32,little"`
}

type SlabLastFreeNode struct {
}
