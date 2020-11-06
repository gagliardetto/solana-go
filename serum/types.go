package serum

import (
	"fmt"
	"io"

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
	SerumPadding           [5]byte          `json:"-",struc:"[5]pad"`
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
	EndPadding             [7]byte          `json:"-",struc:"[5]pad"`
}

type Orderbook struct {
	// ORDERBOOK_LAYOUT
	SerumPadding [5]byte    `json:"-",struc:"[5]pad"`
	AccountFlags solana.U64 `struc:"uint64,little"`
	// SLAB_LAYOUT
	// SLAB_HEADER_LAYOUT
	BumpIndex    uint32  `struc:"uint32,sizeof=Nodes"`
	ZeroPaddingA [4]byte `json:"-",struc:"[4]pad"`
	FreeListLen  uint32  `struc:"uint32,little"`
	ZeroPaddingB [4]byte `json:"-",struc:"[4]pad"`
	FreeListHead uint32  `struc:"uint32,little"`
	Root         uint32  `struc:"uint32,little"`
	LeafCount    uint32  `struc:"uint32,little"`
	ZeroPaddingC [4]byte `json:"-",struc:"[4]pad"`
	// SLAB_NODE_LAYOUT
	Nodes []SlabNode
}

var slabInstructionDef = solana.NewVariantDefinition([]solana.VariantType{
	{"uninitialized", (*SlabUninitialized)(nil)},
	{"innerNode", (*SlabInnerNode)(nil)},
	{"leafNode", (*SlackLeafNode)(nil)},
	{"freeNode", (*SlabFreeNode)(nil)},
	{"lastFreeNode", (*SlabLastFreeNode)(nil)},
})

type SlabNode struct {
	solana.BaseVariant
}

func (s *SlabNode) Unpack(r io.Reader, length int, opt *struc.Options) error {
	fmt.Println("Unpacking slab node")
	return s.BaseVariant.Unpack(slabInstructionDef, r, length, opt)
}

type SlabUninitialized struct {
}

type SlabInnerNode struct {
	PrefixLen uint32 `struc:"uint32,little"`
}

type SlackLeafNode struct {
	OwnerSlot uint8 `struc:"uint8,little"`
}

type SlabFreeNode struct {
	Next uint32 `struc:"uint32,little"`
}

type SlabLastFreeNode struct {
}
