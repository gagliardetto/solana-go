package token

import (
	"fmt"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

// Token contract interface

type Token struct {
	ProgramID string
	Mint      string
}

func New(programID string, mint string) *Token {
	return &Token{ProgramID: programID, Mint: mint}
}

type Account struct {
	Mint            solana.PublicKey
	Owner           solana.PublicKey
	Amount          bin.Uint64
	IsDelegateSet   uint32
	Delegate        solana.PublicKey
	IsInitialized   bool
	IsNative        bool
	Padding         [2]byte `json:"-"`
	DelegatedAmount bin.Uint64
}

type Multisig struct {
	M             byte
	N             byte
	IsInitialized bool
	Signers       [11]solana.PublicKey
}

type Mint struct {
	MintAuthorityOption   uint32
	MintAuthority         solana.PublicKey
	Supply                bin.Uint64
	Decimals              uint8
	IsInitialized         bool
	FreezeAuthorityOption uint32
	FreezeAuthority       solana.PublicKey
}

func DecodeMint(in []byte) (*Mint, error) {
	var m *Mint
	decoder := bin.NewDecoder(in)
	err := decoder.Decode(m)
	if err != nil {
		return nil, fmt.Errorf("unpack: %w", err)
	}
	return m, nil
}

type MintMeta struct {
	TokenSymbol string
	MintAddress solana.PublicKey
	TokenName   string
	IconURL     string `json:"icon"`
}
