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

type SetFundingAuthority struct {
	Auth *FundingType
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] manager
	// [2] = [] newFundingAuthority
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewSetFundingAuthorityInstruction(
	// Parameters:
	auth *FundingType,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	manager ag_solanago.PublicKey,
	newFundingAuthority ag_solanago.PublicKey,
) *SetFundingAuthority {
	return NewSetFundingAuthorityInstructionBuilder().
		SetAuth(auth).
		SetStakePool(stakePool).
		SetManager(manager).
		SetNewFundingAuthority(newFundingAuthority)
}

func NewSetFundingAuthorityInstructionBuilder() *SetFundingAuthority {
	return &SetFundingAuthority{
		Accounts: make(ag_solanago.AccountMetaSlice, 3),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (s *SetFundingAuthority) SetAuth(auth *FundingType) *SetFundingAuthority {
	s.Auth = auth
	return s
}

func (s *SetFundingAuthority) SetStakePool(stakePool ag_solanago.PublicKey) *SetFundingAuthority {
	s.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return s
}

func (s *SetFundingAuthority) SetManager(manager ag_solanago.PublicKey) *SetFundingAuthority {
	s.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	s.Signers[0] = ag_solanago.Meta(manager).SIGNER()
	return s
}

func (s *SetFundingAuthority) SetNewFundingAuthority(newFundingAuthority ag_solanago.PublicKey) *SetFundingAuthority {
	s.Accounts[2] = ag_solanago.Meta(newFundingAuthority)
	return s
}

func (s *SetFundingAuthority) GetAuth() *FundingType {
	return s.Auth
}

func (s *SetFundingAuthority) GetStakePool() ag_solanago.PublicKey {
	return s.Accounts[0].PublicKey
}

func (s *SetFundingAuthority) GetManager() ag_solanago.PublicKey {
	return s.Accounts[1].PublicKey
}

func (s *SetFundingAuthority) GetNewFundingAuthority() ag_solanago.PublicKey {
	return s.Accounts[2].PublicKey
}

func (s *SetFundingAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s.Build(), nil
}

func (s *SetFundingAuthority) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_SetFundingAuthority),
			Impl:   s,
		},
	}
}

func (s *SetFundingAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetFundingAuthority")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if s.Auth != nil {
							paramsBranch.Child(ag_format.Param("Auth", *s.Auth))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range s.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(s.Signers)))
						for j, signer := range s.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (s *SetFundingAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if s.Auth != nil {
		if err := encoder.Encode(s.Auth); err != nil {
			return err
		}
	}
	for _, account := range s.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (s *SetFundingAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if s.Auth != nil {
		if err := decoder.Decode(s.Auth); err != nil {
			return err
		}
	}
	for i := range s.Accounts {
		if err := decoder.Decode(s.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (s *SetFundingAuthority) Validate() error {
	if s.Auth == nil {
		return errors.New("auth is not set")
	}
	for i, account := range s.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(s.Signers) == 0 || !s.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}
