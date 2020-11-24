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
	"fmt"

	bin "github.com/dfuse-io/binary"
	solana "github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/text"
)

var PROGRAM_ID = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")

func init() {
	solana.RegisterInstructionDecoder(PROGRAM_ID, registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, rawInstruction *solana.CompiledInstruction) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, rawInstruction)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, compiledInstruction *solana.CompiledInstruction) (*Instruction, error) {
	var inst *Instruction
	if err := bin.NewDecoder(compiledInstruction.Data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountSettable); ok {
		err := v.SetAccounts(accounts, compiledInstruction.Accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return inst, nil
}

type Instruction struct {
	bin.BaseVariant
}

func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

var InstructionImplDef = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"create_account", (*CreateAccount)(nil)},
	{"assign", (*Assign)(nil)},
	{"transfer", (*Transfer)(nil)},
})

func (i *Instruction) UnmarshalBinary(decoder *bin.Decoder) error {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

type CreateAccountAccounts struct {
	From *solana.AccountMeta `text:"linear,notype"`
	New  *solana.AccountMeta `text:"linear,notype"`
}

type CreateAccount struct {
	// prefixed with byte 0x00
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey
	Accounts *CreateAccountAccounts `bin:"-"`
}

func (i *CreateAccount) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	i.Accounts = &CreateAccountAccounts{
		From: accounts[instructionActIdx[0]],
		New:  accounts[instructionActIdx[1]],
	}
	return nil
}

type Assign struct {
	// prefixed with byte 0x01
	Owner solana.PublicKey
}

type Transfer struct {
	// Prefixed with byte 0x02
	Lamports bin.Uint64
	Accounts *TransferAccounts `bin:"-"`
}

type TransferAccounts struct {
	From *solana.AccountMeta `text:"linear,notype"`
	To   *solana.AccountMeta `text:"linear,notype"`
}

func (i *Transfer) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	i.Accounts = &TransferAccounts{
		From: accounts[instructionActIdx[0]],
		To:   accounts[instructionActIdx[1]],
	}
	return nil
}

type CreateAccountWithSeed struct {
	// Prefixed with byte 0x03
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
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
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Space    bin.Uint64
	Owner    solana.PublicKey
}

type AssignWithSeed struct {
	// Prefixed with byte 0x0a
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey
}
