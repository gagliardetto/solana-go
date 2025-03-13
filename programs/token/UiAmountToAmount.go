// Copyright 2021 github.com/gagliardetto
// Copyright 2025 github.com/liquid-collective
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
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// Converts a `UiAmount` of tokens to a little-endian `u64` raw Amount,
// using the given mint. In this version of the program, the mint can
// only specify the number of decimals.
type UiAmountToAmount struct {
	// The `UiAmount` string to convert.
	UiAmount *string

	// [0] = [] mint
	// ··········· The mint.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *UiAmountToAmount) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts = ag_solanago.AccountMetaSlice(accounts)
	return nil
}

func (inst UiAmountToAmount) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	return
}

// NewUiAmountToAmountInstructionBuilder creates a new `UiAmountToAmount` instruction builder.
func NewUiAmountToAmountInstructionBuilder() *UiAmountToAmount {
	nd := &UiAmountToAmount{
		Accounts: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetUiAmount sets the "ui_amount" parameter.
// The `UiAmount` string to convert.
func (inst *UiAmountToAmount) SetUiAmount(uiAmount string) *UiAmountToAmount {
	inst.UiAmount = &uiAmount
	return inst
}

// SetMintAccount sets the "mint" account.
// The mint.
func (inst *UiAmountToAmount) SetMintAccount(mint ag_solanago.PublicKey) *UiAmountToAmount {
	inst.Accounts[0] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The mint.
func (inst *UiAmountToAmount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

func (inst UiAmountToAmount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UiAmountToAmount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UiAmountToAmount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UiAmountToAmount) Validate() error {
	// Check whether all (required) parameters are set:
	if inst.UiAmount == nil {
		return errors.New("UiAmount parameter is not set")
	}

	// Check whether all (required) accounts are set:
	if inst.Accounts[0] == nil {
		return errors.New("accounts.Accounts is not set")
	}

	return nil
}

func (inst *UiAmountToAmount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UiAmountToAmount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UiAmount", *inst.UiAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("mint", inst.Accounts[0]))
					})
				})
		})
}

func (inst UiAmountToAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UiAmount` param:
	err = encoder.Encode(inst.UiAmount)
	if err != nil {
		return err
	}
	return nil
}
func (inst *UiAmountToAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UiAmount`:
	err = decoder.Decode(&inst.UiAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewUiAmountToAmountInstruction declares a new UiAmountToAmount instruction with the provided parameters and accounts.
func NewUiAmountToAmountInstruction(
	// Parameters:
	uiAmount string,
	// Accounts:
	mint ag_solanago.PublicKey,
) *UiAmountToAmount {
	return NewUiAmountToAmountInstructionBuilder().
		SetUiAmount(uiAmount).
		SetMintAccount(mint)
}
