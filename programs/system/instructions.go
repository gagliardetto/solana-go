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
	"bytes"
	"encoding/binary"
	"fmt"

	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
)

var PROGRAM_ID = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")

func init() {
	solana.RegisterInstructionDecoder(PROGRAM_ID, registryDecodeInstruction)
}

const (
	// Create a new account.
	Instruction_CreateAccount uint32 = iota

	// Assign account to a program.
	Instruction_Assign

	// Transfer lamports.
	Instruction_Transfer

	// Create a new account at an address derived from a base pubkey and a seed.
	Instruction_CreateAccountWithSeed

	// Consumes a stored nonce, replacing it with a successor.
	Instruction_AdvanceNonceAccount

	// Withdraw funds from a nonce account.
	Instruction_WithdrawNonceAccount

	// Drive state of Uninitalized nonce account to Initialized, setting the nonce value.
	Instruction_InitializeNonceAccount

	// Change the entity authorized to execute nonce instruction_s on the account.
	Instruction_AuthorizeNonceAccount

	// Allocate space in a (possibly new) account without funding.
	Instruction_Allocate

	// Allocate space for and assign an account at an address derived from a base public key and a seed.
	Instruction_AllocateWithSeed

	// Assign account to a program based on a seed.
	Instruction_AssignWithSeed

	// Transfer lamports from a derived address.
	Instruction_TransferWithSeed
)

type Instruction struct {
	bin.BaseVariant
}

var (
	// TODO: each instruction must be here:
	_ solana.AccountsGettable = &CreateAccount{}
	_ solana.AccountsSettable = &CreateAccount{}

	_ solana.AccountsGettable = &Assign{}
	_ solana.AccountsSettable = &Assign{}

	_ solana.AccountsGettable = &Transfer{}
	_ solana.AccountsSettable = &Transfer{}

	_ solana.AccountsGettable = &CreateAccountWithSeed{}
	_ solana.AccountsSettable = &CreateAccountWithSeed{}

	_ solana.AccountsGettable = &AdvanceNonceAccount{}
	_ solana.AccountsSettable = &AdvanceNonceAccount{}
)

func (ins *Instruction) Accounts() (out []*solana.AccountMeta) {
	return ins.Impl.(solana.AccountsGettable).GetAccounts()
}

// InstructionImplDef is used for deciding binary,
// encoding and decoding json.
var InstructionImplDef = bin.NewVariantDefinition(
	bin.Uint32TypeIDEncoding,
	[]bin.VariantType{
		// TODO:
		{"create_account", (*CreateAccount)(nil)},
		{"assign", (*Assign)(nil)},
		{"transfer", (*Transfer)(nil)},
		{"create_account_with_seed", (*CreateAccountWithSeed)(nil)},
		{"advance_nonce_account", (*AdvanceNonceAccount)(nil)},
	})

func (i *Instruction) ProgramID() solana.PublicKey {
	return PROGRAM_ID
}

func (i *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewEncoder(buf).Encode(i); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

func (i *Instruction) UnmarshalBinary(decoder *bin.Decoder) error {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (i *Instruction) MarshalBinary(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(i.TypeID, binary.LittleEndian)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(i.Impl)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	var inst *Instruction
	if err := bin.NewDecoder(data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return inst, nil
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
