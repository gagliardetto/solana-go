package token

import "github.com/dfuse-io/solana-go"

// Token contract interface

type Token struct {
	ProgramID string
	Mint      string
}

func New(programID string, mint string) *Token {
	return &Token{ProgramID: programID, Mint: mint}
}

type Account struct {
	Mint            solana.PublicKey `struc:"[32]byte"`
	Owner           solana.PublicKey `struc:"[32]byte"`
	Amount          solana.U64       `struc:"uint64,little"`
	IsDelegateSet   uint32           `struc:"uint32,little"`
	Delegate        solana.PublicKey `struc:"[32]byte"`
	IsInitialized   bool
	IsNative        bool
	Padding         [2]byte `json:"-",struc:"[2]pad"`
	DelegatedAmount solana.U64
}

type Multisig struct {
	M             byte
	N             byte
	IsInitialized bool
	Signers       [11]solana.PublicKey
}

//type Mint struct {
//	OwnerOption   uint32           `struc:"uint32,little"`
//	Owner         solana.PublicKey `struc:"[32]byte"`
//	Decimals      byte
//	IsInitialized bool
//	Padding1      uint16 `json:"-",struct:"[2]pad"`
//}

type Mint struct {
	MintAuthorityOption   uint32           `struc:"uint32,little"`
	MintAuthority         solana.PublicKey `struc:"[32]byte"`
	Supply                solana.U64       `struc:"uint64,little"`
	Decimals              uint8            `struc:"uint8,little"`
	IsInitialized         bool
	FreezeAuthorityOption uint32           `struc:"uint32,little"`
	FreezeAuthority       solana.PublicKey `struc:"[32]byte"`
}
