package serum

import "github.com/dfuse-io/solana-go"

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
	SerumPadding [5]byte          `json:"-",struc:"[5]pad"`
	AccountFlags solana.U64       `struc:"uint64,little"`
	OwnAddress   solana.PublicKey `struc:"[32]byte"`

	VaultSignerNonce solana.U64 `struc:"uint64,little"`

	BaseMint  solana.PublicKey `struc:"[32]byte"`
	QuoteMint solana.PublicKey `struc:"[32]byte"`

	BaseVault         solana.PublicKey `struc:"[32]byte"`
	BaseDepositsTotal solana.U64       `struc:"uint64,little"`
	BaseFeesAccrued   solana.U64       `struc:"uint64,little"`

	QuoteVault         solana.PublicKey `struc:"[32]byte"`
	QuoteDepositsTotal solana.U64       `struc:"uint64,little"`
	QuoteFeesAccrued   solana.U64       `struc:"uint64,little"`

	QuoteDustThreshold solana.U64 `struc:"uint64,little"`

	RequestQueue solana.PublicKey `struc:"[32]byte"`
	EventQueue   solana.PublicKey `struc:"[32]byte"`

	Bids solana.PublicKey `struc:"[32]byte"`
	Asks solana.PublicKey `struc:"[32]byte"`

	BaseLotSize  solana.U64 `struc:"uint64,little"`
	QuoteLotSize solana.U64 `struc:"uint64,little"`

	FeeRateBPS solana.U64 `struc:"uint64,little"`

	ReferrerRebatesAccrued solana.U64 `struc:"uint64,little"`

	EndPadding [7]byte `json:"-",struc:"[5]pad"`
}


type Orderbook struct {
	BumpIndex uint32 `struc:"uint32,sizeof=Nodes"`
	Padding [4]byte
	FreeListLen uint32
	Padding [4]byte
	FreeListHead uint32
	Root uint32
	LeafCount uint32

	Nodes []SlabNode
}

type SlabNode struct {
}
