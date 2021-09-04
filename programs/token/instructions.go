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
	"fmt"

	"github.com/gagliardetto/solana-go/text"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

var TOKEN_PROGRAM_ID = solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

func init() {
	solana.RegisterInstructionDecoder(TOKEN_PROGRAM_ID, registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	var inst Instruction
	if err := bin.NewBinDecoder(data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return &inst, nil
}

var InstructionDefVariant = bin.NewVariantDefinition(bin.Uint8TypeIDEncoding, []bin.VariantType{
	{"initialize_mint", (*InitializeMint)(nil)},
	{"initialize_account", (*InitializeAccount)(nil)},
	{"InitializeMultisig", (*InitializeMultisig)(nil)},
	{"Transfer", (*Transfer)(nil)},
	{"Approve", (*Approve)(nil)},
	{"Revoke", (*Revoke)(nil)},
	{"SetAuthority", (*SetAuthority)(nil)},
	{"MintTo", (*MintTo)(nil)},
	{"Burn", (*Burn)(nil)},
	{"CloseAccount", (*CloseAccount)(nil)},
	{"FreezeAccount", (*FreezeAccount)(nil)},
	{"ThawAccount", (*ThawAccount)(nil)},
	{"TransferChecked", (*TransferChecked)(nil)},
	{"ApproveChecked", (*ApproveChecked)(nil)},
	{"MintToChecked", (*MintToChecked)(nil)},
	{"BurnChecked", (*BurnChecked)(nil)},
})

type Instruction struct {
	bin.BaseVariant
}

var _ bin.EncoderDecoder = &Instruction{}

func (i *Instruction) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionDefVariant)
}

func (i *Instruction) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.WriteUint8(i.TypeID.Uint8())
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(i.Impl)
}
func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

type InitializeMultisigAccounts struct {
}
type InitializeMultisig struct {
	Accounts *InitializeMultisigAccounts
}

type InitializeMintAccounts struct {
}
type InitializeMint struct {
	Accounts *InitializeMintAccounts
}

type TransferAccounts struct {
}
type Transfer struct {
	Accounts *TransferAccounts
}

type ApproveAccounts struct {
}
type Approve struct {
	Accounts *ApproveAccounts
}

type RevokeAccounts struct {
}
type Revoke struct {
	Accounts *RevokeAccounts
}

type SetAuthorityAccounts struct {
}
type SetAuthority struct {
	Accounts *SetAuthorityAccounts
}

type MintToAccounts struct {
}
type MintTo struct {
	Accounts *MintToAccounts
}

type BurnAccounts struct {
}
type Burn struct {
	Accounts *BurnAccounts
}

type CloseAccountAccounts struct {
}
type CloseAccount struct {
	Accounts *CloseAccountAccounts
}

type FreezeAccountAccounts struct {
}
type FreezeAccount struct {
	Accounts *FreezeAccountAccounts
}

type ThawAccountAccounts struct {
}
type ThawAccount struct {
	Accounts *ThawAccountAccounts
}

type TransferCheckedAccounts struct {
}
type TransferChecked struct {
	Accounts *TransferCheckedAccounts
}

type ApproveCheckedAccounts struct {
}
type ApproveChecked struct {
	Accounts *ApproveCheckedAccounts
}

type MintToCheckedAccounts struct {
}
type MintToChecked struct {
	Accounts *MintToCheckedAccounts
}

type BurnCheckedAccounts struct {
}
type BurnChecked struct {
	Accounts *BurnCheckedAccounts
}

type InitializeAccountAccounts struct {
	Account    *solana.AccountMeta `text:"linear,notype"`
	Mint       *solana.AccountMeta `text:"linear,notype"`
	Owner      *solana.AccountMeta `text:"linear,notype"`
	RentSysvar *solana.AccountMeta `text:"linear,notype"`
}

type InitializeAccount struct {
	Accounts *InitializeAccountAccounts `bin:"-"`
}

func (i *InitializeAccount) SetAccounts(accounts []*solana.AccountMeta) error {
	i.Accounts = &InitializeAccountAccounts{
		Account:    accounts[0],
		Mint:       accounts[1],
		Owner:      accounts[2],
		RentSysvar: accounts[3],
	}
	return nil
}
