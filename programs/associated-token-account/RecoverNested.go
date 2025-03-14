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
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type RecoverNested struct {
	// [0] = [WRITE] NestedToken
	// ··········· Nested associated token account, must be owned by `3`
	//
	// [1] = [] TokenMint
	// ··········· Token mint for the nested associated token account.
	//
	// [2] = [WRITE] AssociatedToken
	// ··········· Wallet's associated token account.
	//
	// [3] = [] OwnerAssociatedToken
	// ··········· Owner associated token account address, must be owned by `5`
	//
	// [4] = [] TokenMintForOwner
	// ··········· Token mint for the owner associated token account.
	//
	// [5] = [WRITE, SIGNER] OwnerAssociatedToken
	// ··········· Wallet address for the owner associated token account.
	//
	// [6] = [] SplTokenProgram
	// ··········· SPL Token program.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *RecoverNested) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	if len(accounts) != 7 {
		return fmt.Errorf("expected 7 accounts, got %v", len(accounts))
	}
	inst.Accounts = accounts
	return nil
}

func (inst RecoverNested) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	return
}

// NewRecoverNestedInstructionBuilder creates a new `RecoverNested` instruction builder.
func NewRecoverNestedInstructionBuilder() *RecoverNested {
	return &RecoverNested{
		Accounts: make(ag_solanago.AccountMetaSlice, 7),
	}
}

// SetNestedAccount sets the "nested" account.
func (inst *RecoverNested) SetNestedAccount(nested ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[0] = ag_solanago.Meta(nested).WRITE()
	return inst
}

// GetNestedAccount gets the "nested" account.
func (inst *RecoverNested) GetNestedAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

// SetTokenMintForNestedAccount sets the "token mint for nested" account.
func (inst *RecoverNested) SetTokenMintForNestedAccount(tokenMint ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[1] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintForNestedAccount gets the "token mint for nested" account.
func (inst *RecoverNested) GetTokenMintForNestedAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[1]
}

// SetAssociatedTokenAccount sets the "wallet's associated token" account.
func (inst *RecoverNested) SetAssociatedTokenAccount(walletAssociatedToken ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[2] = ag_solanago.Meta(walletAssociatedToken).WRITE()
	return inst
}

// GetAssociatedTokenAccount gets the "wallet's associated token" account.
func (inst *RecoverNested) GetAssociatedTokenAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[2]
}

// SetOwnerAssociatedTokenAccount sets the "owner associated token" account.
func (inst *RecoverNested) SetOwnerAssociatedTokenAccount(ownerAssociatedToken ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[3] = ag_solanago.Meta(ownerAssociatedToken)
	return inst
}

// GetOwnerAssociatedTokenAccount gets the "owner associated token" account.
func (inst *RecoverNested) GetOwnerAssociatedTokenAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[3]
}

// SetTokenMintForOwnerAccount sets the "token mint for owner" account.
func (inst *RecoverNested) SetTokenMintForOwnerAccount(tokenMint ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[4] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintForOwnerAccount gets the "token mint for owner" account.
func (inst *RecoverNested) GetTokenMintForOwnerAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[4]
}

// SetWalletAccount sets the "wallet" account.
func (inst *RecoverNested) SetWalletAccount(wallet ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[5] = ag_solanago.Meta(wallet).WRITE().SIGNER()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *RecoverNested) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[5]
}

// SetSplTokenProgram sets the "token program" account.
func (inst *RecoverNested) SetSplTokenProgram(tokenProgram ag_solanago.PublicKey) *RecoverNested {
	inst.Accounts[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetSplTokenProgram gets the "token program" account.
func (inst *RecoverNested) GetSplTokenProgram() *ag_solanago.AccountMeta {
	return inst.Accounts[6]
}

func (inst RecoverNested) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RecoverNested),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RecoverNested) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RecoverNested) Validate() error {
	// Check whether all accounts are set:
	for i, acc := range inst.Accounts {
		if acc == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	return nil
}

func (inst *RecoverNested) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RecoverNested")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          nestedToken", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("            tokenMint", inst.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("     associatedToken", inst.Accounts[2]))
						accountsBranch.Child(ag_format.Meta("ownerAssociatedToken", inst.Accounts[3]))
						accountsBranch.Child(ag_format.Meta("      tokenMintOwner", inst.Accounts[4]))
						accountsBranch.Child(ag_format.Meta("ownerAssociatedToken", inst.Accounts[5]))
						accountsBranch.Child(ag_format.Meta("     splTokenProgram", inst.Accounts[6]))
					})
				})
		})
}

func (inst RecoverNested) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (inst *RecoverNested) UnmarshalWithDecoder(_ *ag_binary.Decoder) (err error) {
	return nil
}

// NewRecoverNestedInstruction declares a new RecoverNested instruction with the provided parameters and accounts.
func NewRecoverNestedInstruction(
// Accounts:
	nested ag_solanago.PublicKey,
	tokenMintForNested ag_solanago.PublicKey,
	walletAssociatedToken ag_solanago.PublicKey,
	ownerAssociatedToken ag_solanago.PublicKey,
	tokenMintForOwner ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *RecoverNested {
	return NewRecoverNestedInstructionBuilder().
		SetNestedAccount(nested).
		SetTokenMintForNestedAccount(tokenMintForNested).
		SetAssociatedTokenAccount(walletAssociatedToken).
		SetOwnerAssociatedTokenAccount(ownerAssociatedToken).
		SetTokenMintForOwnerAccount(tokenMintForOwner).
		SetWalletAccount(wallet).
		SetSplTokenProgram(tokenProgram)
}
