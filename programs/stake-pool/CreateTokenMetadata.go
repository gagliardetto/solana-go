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

type CreateTokenMetadata struct {
	Name   *string
	Symbol *string
	URI    *string
	// [0] = [] stakePool
	// [1] = [SIGNER] manager
	// [2] = [] stakePoolWithdrawAuthority
	// [3] = [] poolMint
	// [4] = [SIGNER WRITE] payer
	// [5] = [WRITE] tokenMetadata
	// [6] = [] mplTokenMetadata
	// [7] = [] systemProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewCreateTokenMetadataInstructionBuilder() *CreateTokenMetadata {
	return &CreateTokenMetadata{
		Accounts: make(ag_solanago.AccountMetaSlice, 8),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (inst *CreateTokenMetadata) SetName(name string) *CreateTokenMetadata {
	inst.Name = &name
	return inst
}

func (inst *CreateTokenMetadata) SetSymbol(symbol string) *CreateTokenMetadata {
	inst.Symbol = &symbol
	return inst
}

func (inst *CreateTokenMetadata) SetURI(uri string) *CreateTokenMetadata {
	inst.URI = &uri
	return inst
}

func (inst *CreateTokenMetadata) SetStakePool(stakePool ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[0] = ag_solanago.Meta(stakePool)
	return inst
}

func (inst *CreateTokenMetadata) SetManager(manager ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(manager).SIGNER()
	return inst
}

func (inst *CreateTokenMetadata) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return inst
}

func (inst *CreateTokenMetadata) SetPoolMint(poolMint ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[3] = ag_solanago.Meta(poolMint)
	return inst
}

func (inst *CreateTokenMetadata) SetPayer(payer ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[4] = ag_solanago.Meta(payer).WRITE()
	return inst
}

func (inst *CreateTokenMetadata) SetTokenMetadata(tokenMetadata ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[5] = ag_solanago.Meta(tokenMetadata).WRITE()
	return inst
}

func (inst *CreateTokenMetadata) SetMplTokenMetadata(mplTokenMetadata ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[6] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

func (inst *CreateTokenMetadata) SetSystemProgram(systemProgram ag_solanago.PublicKey) *CreateTokenMetadata {
	inst.Accounts[7] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *CreateTokenMetadata) GetName() *string {
	return inst.Name
}

func (inst *CreateTokenMetadata) GetSymbol() *string {
	return inst.Symbol
}

func (inst *CreateTokenMetadata) GetURI() *string {
	return inst.URI
}

func (inst *CreateTokenMetadata) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *CreateTokenMetadata) GetManager() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *CreateTokenMetadata) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *CreateTokenMetadata) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *CreateTokenMetadata) GetPayer() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *CreateTokenMetadata) GetTokenMetadata() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *CreateTokenMetadata) GetMplTokenMetadata() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *CreateTokenMetadata) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *CreateTokenMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateTokenMetadata) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateTokenMetadata),
			Impl:   inst,
		},
	}
}

func (inst *CreateTokenMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateTokenMetadata")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Name != nil {
							paramsBranch.Child(ag_format.Param("Name", *inst.Name))
						}
						if inst.Symbol != nil {
							paramsBranch.Child(ag_format.Param("Symbol", *inst.Symbol))
						}
						if inst.URI != nil {
							paramsBranch.Child(ag_format.Param("URI", *inst.URI))
						}
					})

					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range inst.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(inst.Signers)))
						for j, signer := range inst.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (inst *CreateTokenMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Name != nil {
		if err := encoder.Encode(inst.Name); err != nil {
			return err
		}
	}
	if inst.Symbol != nil {
		if err := encoder.Encode(inst.Symbol); err != nil {
			return err
		}
	}
	if inst.URI != nil {
		if err := encoder.Encode(inst.URI); err != nil {
			return err
		}
	}
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *CreateTokenMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Name != nil {
		if err := decoder.Decode(inst.Name); err != nil {
			return err
		}
	}
	if inst.Symbol != nil {
		if err := decoder.Decode(inst.Symbol); err != nil {
			return err
		}
	}
	if inst.URI != nil {
		if err := decoder.Decode(inst.URI); err != nil {
			return err
		}
	}
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *CreateTokenMetadata) Validate() error {
	if inst.Name == nil {
		return errors.New("name is not set")
	}
	if inst.Symbol == nil {
		return errors.New("symbol is not set")
	}
	if inst.URI == nil {
		return errors.New("uri is not set")
	}

	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}
