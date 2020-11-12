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

package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/dfuse-io/solana-go"
)

type SystemInstruction struct {
	//Type    bin.Varuint16
	//Variant interface{}
	bin.BaseVariant
}

var SystemInstructionImplDef = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"create_account", (*CreateAccount)(nil)},
	{"assign", (*Assign)(nil)},
	{"transfer", (*Transfer)(nil)},
})

func (i *SystemInstruction) UnmarshalBinary(decoder *bin.Decoder) error {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, SystemInstructionImplDef)
}

type CreateAccount struct {
	// prefixed with byte 0x00
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey
}

type Assign struct {
	// prefixed with byte 0x01
	Owner solana.PublicKey
}

type Transfer struct {
	// Prefixed with byte 0x02
	Lamports bin.Uint64
}

type CreateAccountWithSeed struct {
	// Prefixed with byte 0x03
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey
}

type AdvanceNonceAccount struct {
	// Prefix with 0x04
}

type WithdrawNonceAccount struct {
	// Prefix with 0x05
	Lamports bin.Uint64
}

type InitializeNonceAccount struct {
	// Prefix with 0x06
	AuthorizedAccount solana.PublicKey
}

type AuthorizeNonceAccount struct {
	// Prefix with 0x07
	AuthorizeAccount solana.PublicKey
}

type Allocate struct {
	// Prefix with 0x08
	Space bin.Uint64
}

type AllocateWithSeed struct {
	// Prefixed with byte 0x09
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Space    bin.Uint64
	Owner    solana.PublicKey
}

type AssignWithSeed struct {
	// Prefixed with byte 0x0a
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey
}
