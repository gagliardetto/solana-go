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
)

var PROGRAM_ID = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")

func init() {
	solana.RegisterInstructionDecoder(PROGRAM_ID, registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []solana.PublicKey, rawInstruction *solana.CompiledInstruction) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, rawInstruction)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []solana.PublicKey, compiledInstruction *solana.CompiledInstruction) (*Instruction, error) {
	var inst *Instruction
	if err := bin.NewDecoder(compiledInstruction.Data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountSettable); ok {
		v.SetAccounts(accounts)
	}

	return inst, nil
}

type Instruction struct {
	bin.BaseVariant
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%s", i.Impl)
}

var InstructionImplDef = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"create_account", (*CreateAccount)(nil)},
	{"assign", (*Assign)(nil)},
	{"transfer", (*Transfer)(nil)},
})

func (i *Instruction) UnmarshalBinary(decoder *bin.Decoder) error {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

type CreateAccount struct {
	// prefixed with byte 0x00
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey
}

func (i *CreateAccount) String() string {
	out := "Create Account\n"
	out += fmt.Sprintf("Lamports: %d, space: %d, Owner: %s", i.Lamports, i.Space, i.Owner.String())
	return out
}

type Assign struct {
	// prefixed with byte 0x01
	Owner solana.PublicKey
}

func (i *Assign) String() string {
	out := "Assign\n"
	out += fmt.Sprintf("Owner: %s", i.Owner.String())
	return out
}

type Transfer struct {
	// Prefixed with byte 0x02
	Lamports bin.Uint64
}

func (t *Transfer) String() string {
	out := "Transfer\n"
	out += fmt.Sprintf("Lamports: %d", t.Lamports)
	return out
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

func (i *CreateAccountWithSeed) String() string {
	out := "Create Account With Seed\n"
	out += fmt.Sprintf("Base: %s SeedSize: %d Seed: %s Lamports: %d Space: %d Owner: %s", i.Base, i.SeedSize, i.Seed, i.Lamports, i.Space, i.Owner.String())
	return out
}

type AdvanceNonceAccount struct {
	// Prefix with 0x04
}

func (i *AdvanceNonceAccount) String() string {
	out := "Advance Nonce Account\n"
	out += "Accounts:"
	return out
}

type WithdrawNonceAccount struct {
	// Prefix with 0x05
	Lamports bin.Uint64
}

func (i *WithdrawNonceAccount) String() string {
	out := "Withdraw Nonce Account\n"
	out += fmt.Sprintf("Lamports: %d", i.Lamports)
	return out
}

type InitializeNonceAccount struct {
	// Prefix with 0x06
	AuthorizedAccount solana.PublicKey
}

func (i *InitializeNonceAccount) String() string {
	out := "Initialize Nonce Account\n"
	out += fmt.Sprintf("AuthorizedAccount: %s", i.AuthorizedAccount.String())
	return out
}

type AuthorizeNonceAccount struct {
	// Prefix with 0x07
	AuthorizeAccount solana.PublicKey
}

func (i *AuthorizeNonceAccount) String() string {
	out := "Authorize Nonce Account\n"
	out += fmt.Sprintf("AuthorizedAccount: %s", i.AuthorizeAccount.String())
	return out
}

type Allocate struct {
	// Prefix with 0x08
	Space bin.Uint64
}

func (i *Allocate) String() string {
	out := "Allocate"
	out += fmt.Sprintf("Space: %d", i.Space)
	return out
}

type AllocateWithSeed struct {
	// Prefixed with byte 0x09
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Space    bin.Uint64
	Owner    solana.PublicKey
}

func (i *AllocateWithSeed) String() string {
	out := "Allocate With Seed\n"
	out += fmt.Sprintf("Base: %s SeedSize: %d Seed: %s Space: %d Owner: %s", i.Base, i.SeedSize, i.Seed, i.Owner)
	return out
}

type AssignWithSeed struct {
	// Prefixed with byte 0x0a
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey
}

func (i *AssignWithSeed) String() string {
	out := "Assign With Seed\n"
	out += fmt.Sprintf("Base: %s SeedSize: %d Seed: %s Owner: %s", i.Base, i.SeedSize, i.Seed, i.Owner)
	return out
}
