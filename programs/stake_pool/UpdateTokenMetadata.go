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

package stakepool

import (
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type UpdateTokenMetadata struct {
	Name   *string
	Symbol *string
	URI    *string
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] manager
	// [2] = [WRITE] stakePoolWithdrawAuthority
	// [3] = [WRITE] tokenMetadata
	// [4] = [] mplTokenMetadata
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewUpdateTokenMetadataInstruction(
	name string,
	symbol string,
	uri string,
	stakePool ag_solanago.PublicKey,
	manager ag_solanago.PublicKey,
	stakePoolWithdrawAuthority ag_solanago.PublicKey,
	tokenMetadata ag_solanago.PublicKey,
	mplTokenMetadata ag_solanago.PublicKey,
) *UpdateTokenMetadata {
	return NewUpdateTokenMetadataInstructionBuilder().
		SetName(name).
		SetSymbol(symbol).
		SetURI(uri).
		SetStakePool(stakePool).
		SetManager(manager).
		SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority).
		SetTokenMetadata(tokenMetadata).
		SetMplTokenMetadata(mplTokenMetadata)
}

func NewUpdateTokenMetadataInstructionBuilder() *UpdateTokenMetadata {
	return &UpdateTokenMetadata{
		Accounts: make(ag_solanago.AccountMetaSlice, 5),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (u *UpdateTokenMetadata) SetName(name string) *UpdateTokenMetadata {
	u.Name = &name
	return u
}

func (u *UpdateTokenMetadata) SetSymbol(symbol string) *UpdateTokenMetadata {
	u.Symbol = &symbol
	return u
}

func (u *UpdateTokenMetadata) SetURI(uri string) *UpdateTokenMetadata {
	u.URI = &uri
	return u
}

func (u *UpdateTokenMetadata) SetStakePool(stakePool ag_solanago.PublicKey) *UpdateTokenMetadata {
	u.Accounts[0] = ag_solanago.Meta(stakePool)
	return u
}

func (u *UpdateTokenMetadata) SetManager(manager ag_solanago.PublicKey) *UpdateTokenMetadata {
	u.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	u.Signers[0] = ag_solanago.Meta(manager).SIGNER()
	return u
}

func (u *UpdateTokenMetadata) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *UpdateTokenMetadata {
	u.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return u
}

func (u *UpdateTokenMetadata) SetTokenMetadata(tokenMetadata ag_solanago.PublicKey) *UpdateTokenMetadata {
	u.Accounts[3] = ag_solanago.Meta(tokenMetadata).WRITE()
	return u
}

func (u *UpdateTokenMetadata) SetMplTokenMetadata(mplTokenMetadata ag_solanago.PublicKey) *UpdateTokenMetadata {
	u.Accounts[4] = ag_solanago.Meta(mplTokenMetadata)
	return u
}

func (u *UpdateTokenMetadata) GetName() *string {
	return u.Name
}

func (u *UpdateTokenMetadata) GetSymbol() *string {
	return u.Symbol
}

func (u *UpdateTokenMetadata) GetURI() *string {
	return u.URI
}

func (u *UpdateTokenMetadata) GetStakePool() ag_solanago.PublicKey {
	return u.Accounts[0].PublicKey
}

func (u *UpdateTokenMetadata) GetManager() ag_solanago.PublicKey {
	return u.Accounts[1].PublicKey
}

func (u *UpdateTokenMetadata) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return u.Accounts[2].PublicKey
}

func (u *UpdateTokenMetadata) GetTokenMetadata() ag_solanago.PublicKey {
	return u.Accounts[3].PublicKey
}

func (u *UpdateTokenMetadata) GetMplTokenMetadata() ag_solanago.PublicKey {
	return u.Accounts[4].PublicKey
}

func (u *UpdateTokenMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return u.Build(), nil
}

func (u *UpdateTokenMetadata) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateTokenMetadata),
			Impl:   u,
		},
	}
}

func (u *UpdateTokenMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateTokenMetadata")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if u.Name != nil {
							paramsBranch.Child(ag_format.Param("Name", *u.Name))
						}
						if u.Symbol != nil {
							paramsBranch.Child(ag_format.Param("Symbol", *u.Symbol))
						}
						if u.URI != nil {
							paramsBranch.Child(ag_format.Param("URI", *u.URI))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range u.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(u.Signers)))
						for j, signer := range u.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (u *UpdateTokenMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if u.Name != nil {
		if err := encoder.Encode(u.Name); err != nil {
			return err
		}
	}
	if u.Symbol != nil {
		if err := encoder.Encode(u.Symbol); err != nil {
			return err
		}
	}
	if u.URI != nil {
		if err := encoder.Encode(u.URI); err != nil {
			return err
		}
	}
	for _, account := range u.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (u *UpdateTokenMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if u.Name != nil {
		if err := decoder.Decode(u.Name); err != nil {
			return err
		}
	}
	if u.Symbol != nil {
		if err := decoder.Decode(u.Symbol); err != nil {
			return err
		}
	}
	if u.URI != nil {
		if err := decoder.Decode(u.URI); err != nil {
			return err
		}
	}
	for i := range u.Accounts {
		if err := decoder.Decode(u.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *UpdateTokenMetadata) Validate() error {
	if u.Name == nil {
		return errors.New("name is not set")
	}
	if u.Symbol == nil {
		return errors.New("symbol is not set")
	}
	if u.URI == nil {
		return errors.New("uri is not set")
	}
	for i, account := range u.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(u.Signers) == 0 || !u.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}
