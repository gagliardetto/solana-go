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

package associatedtokenaccount

import (
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type CreateIdempotent struct {
	// [0] = [WRITE, SIGNER] Payer
	// ··········· Funding account
	//
	// [1] = [WRITE] AssociatedTokenAccount
	// ··········· Associated token account address to be created
	//
	// [2] = [] Wallet
	// ··········· Wallet address for the new associated token account
	//
	// [3] = [] TokenMint
	// ··········· The token mint for the new associated token account
	//
	// [4] = [] SystemProgram
	// ··········· System program ID
	//
	// [5] = [] TokenProgram
	// ··········· SPL token program ID
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *CreateIdempotent) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	if len(accounts) != 6 {
		return fmt.Errorf("expected 6 accounts, got %v", len(accounts))
	}
	inst.AccountMetaSlice = accounts
	return nil
}

func (inst CreateIdempotent) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.AccountMetaSlice...)
	return
}

// NewCreateIdempotentInstructionBuilder creates a new `CreateIdempotent` instruction builder.
func NewCreateIdempotentInstructionBuilder() *CreateIdempotent {
	return &CreateIdempotent{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
}

// SetFundingAccount sets the "funding" account.
func (inst *CreateIdempotent) SetFundingAccount(funding ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(funding).WRITE().SIGNER()
	return inst
}

// GetFundingAccount gets the "funding" account.
func (inst *CreateIdempotent) GetFundingAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAssociatedTokenAccount sets the "associated token" account.
func (inst *CreateIdempotent) SetAssociatedTokenAccount(associatedToken ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(associatedToken).WRITE()
	return inst
}

// GetAssociatedTokenAccount gets the "associated token" account.
func (inst *CreateIdempotent) GetAssociatedTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetWalletAccount sets the "wallet" account.
func (inst *CreateIdempotent) SetWalletAccount(wallet ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *CreateIdempotent) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetTokenMintAccount sets the "token mint" account.
func (inst *CreateIdempotent) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "token mint" account.
func (inst *CreateIdempotent) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSystemProgramAccount sets the "system program" account.
func (inst *CreateIdempotent) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system program" account.
func (inst *CreateIdempotent) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetTokenProgramAccount sets the "token program" account.
func (inst *CreateIdempotent) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token program" account.
func (inst *CreateIdempotent) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSysVarRentAccount sets the "sys var rent" account.
func (inst *CreateIdempotent) SetSysVarRentAccount(sysVarRent ag_solanago.PublicKey) *CreateIdempotent {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sysVarRent)
	return inst
}

// GetSysVarRentAccount gets the "sys var rent" account.
func (inst *CreateIdempotent) GetSysVarRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst CreateIdempotent) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateIdempotent),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateIdempotent) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateIdempotent) Validate() error {
	if inst.AccountMetaSlice.Get(0).PublicKey.IsZero() {
		return errors.New("Payer not set")
	}
	if inst.AccountMetaSlice.Get(2).PublicKey.IsZero() {
		return errors.New("Wallet not set")
	}
	if inst.AccountMetaSlice.Get(3).PublicKey.IsZero() {
		return errors.New("Mint not set")
	}
	_, _, err := ag_solanago.FindAssociatedTokenAddress(
		inst.AccountMetaSlice.Get(2).PublicKey,
		inst.AccountMetaSlice.Get(3).PublicKey,
	)
	if err != nil {
		return fmt.Errorf("error while FindAssociatedTokenAddress: %w", err)
	}
	return nil
}

func (inst *CreateIdempotent) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateIdempotent")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("associatedTokenAddress", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                wallet", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (inst CreateIdempotent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return encoder.WriteBytes([]byte{}, false)
}

func (inst *CreateIdempotent) UnmarshalWithDecoder(_ *ag_binary.Decoder) (err error) {
	return nil
}

// NewCreateIdempotentInstruction declares a new CreateIdempotent instruction with the provided parameters and accounts.
func NewCreateIdempotentInstruction(
	// Accounts:
	funding ag_solanago.PublicKey,
	associatedToken ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *CreateIdempotent {
	return NewCreateIdempotentInstructionBuilder().
		SetFundingAccount(funding).
		SetAssociatedTokenAccount(associatedToken).
		SetWalletAccount(wallet).
		SetTokenMintAccount(tokenMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
