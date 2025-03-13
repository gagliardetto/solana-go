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

// Convert an Amount of tokens to a `UiAmount` string, using the given mint.
type AmountToUiAmount struct {
	// The amount of tokens to convert.
	Amount *uint64

	// [0] = [] mint
	// ··········· The mint.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *AmountToUiAmount) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts = accounts
	return nil
}

func (inst AmountToUiAmount) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	return
}

func NewAmountToUiAmountInstructionBuilder() *AmountToUiAmount {
	nd := &AmountToUiAmount{
		Accounts: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of tokens to convert.
func (inst *AmountToUiAmount) SetAmount(amount uint64) *AmountToUiAmount {
	inst.Amount = &amount
	return inst
}

// SetMintAccount sets the "mint" account.
// The mint.
func (inst *AmountToUiAmount) SetMintAccount(mint ag_solanago.PublicKey) *AmountToUiAmount {
	inst.Accounts[0] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The mint.
func (inst *AmountToUiAmount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

func (inst AmountToUiAmount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_AmountToUiAmount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AmountToUiAmount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AmountToUiAmount) Validate() error {
	// Check whether all (required) parameters are set:
	if inst.Amount == nil {
		return errors.New("amount parameter is not set")
	}

	// Check whether all (required) accounts are set:
	if inst.Accounts[0] == nil {
		return errors.New("accounts.Accounts is not set")
	}

	return nil
}

func (inst *AmountToUiAmount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AmountToUiAmount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("mint", inst.Accounts[0]))
					})
				})
		})
}

func (inst AmountToUiAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(inst.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (inst *AmountToUiAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&inst.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewAmountToUiAmountInstruction declares a new AmountToUiAmount instruction with the provided parameters and accounts.
func NewAmountToUiAmountInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	mint ag_solanago.PublicKey,
) *AmountToUiAmount {
	return NewAmountToUiAmountInstructionBuilder().
		SetAmount(amount).
		SetMintAccount(mint)
}
