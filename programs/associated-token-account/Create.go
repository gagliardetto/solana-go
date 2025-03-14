// Copyright 2021 github.com/gagliardetto
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

	bin "github.com/gagliardetto/binary"
	treeout "github.com/gagliardetto/treeout"

	solana "github.com/gagliardetto/solana-go"
	format "github.com/gagliardetto/solana-go/text/format"
)

type Create struct {
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
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCreateInstructionBuilder creates a new `Create` instruction builder.
func NewCreateInstructionBuilder() *Create {
	return &Create{}
}

func (inst *Create) SetPayer(payer solana.PublicKey) *Create {
	inst.AccountMetaSlice[0] = solana.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *Create) SetWallet(wallet solana.PublicKey) *Create {
	inst.AccountMetaSlice[2] = solana.Meta(wallet)
	return inst
}

func (inst *Create) SetMint(mint solana.PublicKey) *Create {
	inst.AccountMetaSlice[3] = solana.Meta(mint)
	return inst
}

func (inst *Create) SetAccounts(payer, wallet, mint solana.PublicKey) *Create {
	associatedTokenAddress, _, _ := solana.FindAssociatedTokenAddress(wallet, mint)

	inst.AccountMetaSlice = []*solana.AccountMeta{
		{
			PublicKey:  payer,
			IsSigner:   true,
			IsWritable: true,
		},
		{
			PublicKey:  associatedTokenAddress,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PublicKey:  wallet,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  mint,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  solana.SystemProgramID,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  solana.TokenProgramID,
			IsSigner:   false,
			IsWritable: false,
		},
	}

	return inst
}

func (inst Create) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.NoTypeIDDefaultID,
	}}
}

func (inst Create) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Create) Validate() error {
	if inst.AccountMetaSlice.Get(0).PublicKey.IsZero() {
		return errors.New("Payer not set")
	}
	if inst.AccountMetaSlice.Get(2).PublicKey.IsZero() {
		return errors.New("Wallet not set")
	}
	if inst.AccountMetaSlice.Get(3).PublicKey.IsZero() {
		return errors.New("Mint not set")
	}
	_, _, err := solana.FindAssociatedTokenAddress(
		inst.AccountMetaSlice.Get(2).PublicKey,
		inst.AccountMetaSlice.Get(3).PublicKey,
	)
	if err != nil {
		return fmt.Errorf("error while FindAssociatedTokenAddress: %w", err)
	}
	return nil
}

func (inst *Create) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Create")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("                 payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("associatedTokenAddress", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(format.Meta("                wallet", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(format.Meta("             tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(format.Meta("         systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (inst Create) MarshalWithEncoder(encoder *bin.Encoder) error {
	return nil
}

func (inst *Create) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	return nil
}

func NewCreateInstruction(
	payer solana.PublicKey,
	walletAddress solana.PublicKey,
	splTokenMintAddress solana.PublicKey,
) *Create {
	return NewCreateInstructionBuilder().
		SetPayer(payer).
		SetWallet(walletAddress).
		SetMint(splTokenMintAddress)
}
