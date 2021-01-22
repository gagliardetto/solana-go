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
	solana "github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/text"
)

var PROGRAM_ID = solana.MustPublicKeyFromBase58("11111111111111111111111111111111")

func init() {
	solana.RegisterInstructionDecoder(PROGRAM_ID, registryDecodeInstruction)
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

	if v, ok := inst.Impl.(solana.AccountSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return inst, nil
}

func NewCreateAccountInstruction(lamports uint64, space uint64, owner, from, to solana.PublicKey) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{
			TypeID: 0,
			Impl: &CreateAccount{
				Lamports: bin.Uint64(lamports),
				Space:    bin.Uint64(space),
				Owner:    owner,
				Accounts: &CreateAccountAccounts{
					From: &solana.AccountMeta{PublicKey: from, IsSigner: true, IsWritable: true},
					New:  &solana.AccountMeta{PublicKey: to, IsSigner: true, IsWritable: true},
				},
			},
		},
	}
}

type Instruction struct {
	bin.BaseVariant
}

func (i *Instruction) Accounts() (out []*solana.AccountMeta) {
	switch i.TypeID {
	case 0:
		accounts := i.Impl.(*CreateAccount).Accounts
		out = []*solana.AccountMeta{accounts.From, accounts.New}
	case 1:
		// no account here
	case 2:
		accounts := i.Impl.(*Transfer).Accounts
		out = []*solana.AccountMeta{accounts.From, accounts.To}
	}
	return
}

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

var InstructionImplDef = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"create_account", (*CreateAccount)(nil)},
	{"assign", (*Assign)(nil)},
	{"transfer", (*Transfer)(nil)},
})

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

type CreateAccountAccounts struct {
	From *solana.AccountMeta `text:"linear,notype"`
	New  *solana.AccountMeta `text:"linear,notype"`
}

type CreateAccount struct {
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey
	Accounts *CreateAccountAccounts `bin:"-"`
}

func (i *CreateAccount) SetAccounts(accounts []*solana.AccountMeta) error {
	i.Accounts = &CreateAccountAccounts{
		From: accounts[0],
		New:  accounts[1],
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

func (i *Transfer) SetAccounts(accounts []*solana.AccountMeta) error {
	i.Accounts = &TransferAccounts{
		From: accounts[0],
		To:   accounts[1],
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
