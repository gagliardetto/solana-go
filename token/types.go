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

package token

import (
	"bytes"
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
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

func DecodeMint(in []byte) (*Mint, error) {
	var m Mint
	err := struc.Unpack(bytes.NewReader(in), &m)
	if err != nil {
		return nil, fmt.Errorf("unpack: %w", err)
	}
	return &m, nil
}

type MintMeta struct {
	TokenSymbol string
	MintAddress solana.PublicKey
	TokenName   string
	IconURL     string `json:"icon"`
}
