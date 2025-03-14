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

func (c *CreateTokenMetadata) SetName(name string) *CreateTokenMetadata {
	c.Name = &name
	return c
}

func (c *CreateTokenMetadata) SetSymbol(symbol string) *CreateTokenMetadata {
	c.Symbol = &symbol
	return c
}

func (c *CreateTokenMetadata) SetURI(uri string) *CreateTokenMetadata {
	c.URI = &uri
	return c
}

func (c *CreateTokenMetadata) SetStakePool(stakePool ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[0] = ag_solanago.Meta(stakePool)
	return c
}

func (c *CreateTokenMetadata) SetManager(manager ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	c.Signers[0] = ag_solanago.Meta(manager).SIGNER()
	return c
}

func (c *CreateTokenMetadata) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return c
}

func (c *CreateTokenMetadata) SetPoolMint(poolMint ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[3] = ag_solanago.Meta(poolMint)
	return c
}

func (c *CreateTokenMetadata) SetPayer(payer ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[4] = ag_solanago.Meta(payer).WRITE()
	return c
}

func (c *CreateTokenMetadata) SetTokenMetadata(tokenMetadata ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[5] = ag_solanago.Meta(tokenMetadata).WRITE()
	return c
}

func (c *CreateTokenMetadata) SetMplTokenMetadata(mplTokenMetadata ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[6] = ag_solanago.Meta(mplTokenMetadata)
	return c
}

func (c *CreateTokenMetadata) SetSystemProgram(systemProgram ag_solanago.PublicKey) *CreateTokenMetadata {
	c.Accounts[7] = ag_solanago.Meta(systemProgram)
	return c
}

func (c *CreateTokenMetadata) GetName() *string {
	return c.Name
}

func (c *CreateTokenMetadata) GetSymbol() *string {
	return c.Symbol
}

func (c *CreateTokenMetadata) GetURI() *string {
	return c.URI
}

func (c *CreateTokenMetadata) GetStakePool() ag_solanago.PublicKey {
	return c.Accounts[0].PublicKey
}

func (c *CreateTokenMetadata) GetManager() ag_solanago.PublicKey {
	return c.Accounts[1].PublicKey
}

func (c *CreateTokenMetadata) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return c.Accounts[2].PublicKey
}

func (c *CreateTokenMetadata) GetPoolMint() ag_solanago.PublicKey {
	return c.Accounts[3].PublicKey
}

func (c *CreateTokenMetadata) GetPayer() ag_solanago.PublicKey {
	return c.Accounts[4].PublicKey
}

func (c *CreateTokenMetadata) GetTokenMetadata() ag_solanago.PublicKey {
	return c.Accounts[5].PublicKey
}

func (c *CreateTokenMetadata) GetMplTokenMetadata() ag_solanago.PublicKey {
	return c.Accounts[6].PublicKey
}

func (c *CreateTokenMetadata) GetSystemProgram() ag_solanago.PublicKey {
	return c.Accounts[7].PublicKey
}

func (c *CreateTokenMetadata) ValidateAndBuild() (*Instruction, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c.Build(), nil
}

func (c *CreateTokenMetadata) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateTokenMetadata),
			Impl:   c,
		},
	}
}

func (c *CreateTokenMetadata) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateTokenMetadata")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if c.Name != nil {
							paramsBranch.Child(ag_format.Param("Name", *c.Name))
						}
						if c.Symbol != nil {
							paramsBranch.Child(ag_format.Param("Symbol", *c.Symbol))
						}
						if c.URI != nil {
							paramsBranch.Child(ag_format.Param("URI", *c.URI))
						}
					})

					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range c.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(c.Signers)))
						for j, signer := range c.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (c *CreateTokenMetadata) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if c.Name != nil {
		if err := encoder.Encode(c.Name); err != nil {
			return err
		}
	}
	if c.Symbol != nil {
		if err := encoder.Encode(c.Symbol); err != nil {
			return err
		}
	}
	if c.URI != nil {
		if err := encoder.Encode(c.URI); err != nil {
			return err
		}
	}
	for _, account := range c.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (c *CreateTokenMetadata) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if c.Name != nil {
		if err := decoder.Decode(c.Name); err != nil {
			return err
		}
	}
	if c.Symbol != nil {
		if err := decoder.Decode(c.Symbol); err != nil {
			return err
		}
	}
	if c.URI != nil {
		if err := decoder.Decode(c.URI); err != nil {
			return err
		}
	}
	for i := range c.Accounts {
		if err := decoder.Decode(c.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (c *CreateTokenMetadata) Validate() error {
	if c.Name == nil {
		return errors.New("name is not set")
	}
	if c.Symbol == nil {
		return errors.New("symbol is not set")
	}
	if c.URI == nil {
		return errors.New("uri is not set")
	}

	for i, account := range c.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(c.Signers) == 0 || !c.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}
