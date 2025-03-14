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

type SetManager struct {
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] manager
	// [2] = [SIGNER] newManager
	// [3] = [] newManagerFeeAccount
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewSetManagerInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	manager ag_solanago.PublicKey,
	newManager ag_solanago.PublicKey,
	newManagerFeeAccount ag_solanago.PublicKey,
) *SetManager {
	return NewSetManagerInstructionBuilder().
		SetStakePool(stakePool).
		SetManager(manager).
		SetNewManager(newManager).
		SetNewManagerFeeAccount(newManagerFeeAccount)
}

func NewSetManagerInstructionBuilder() *SetManager {
	return &SetManager{
		Accounts: make(ag_solanago.AccountMetaSlice, 4),
		Signers:  make(ag_solanago.AccountMetaSlice, 2),
	}
}

func (s *SetManager) SetStakePool(pool ag_solanago.PublicKey) *SetManager {
	s.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return s
}

func (s *SetManager) SetManager(manager ag_solanago.PublicKey) *SetManager {
	s.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	return s
}

func (s *SetManager) SetNewManager(newManager ag_solanago.PublicKey) *SetManager {
	s.Accounts[2] = ag_solanago.Meta(newManager).SIGNER()
	s.Signers[0] = ag_solanago.Meta(newManager).SIGNER()
	return s
}

func (s *SetManager) SetNewManagerFeeAccount(newManagerFeeAccount ag_solanago.PublicKey) *SetManager {
	s.Accounts[3] = ag_solanago.Meta(newManagerFeeAccount)
	return s
}

func (s *SetManager) GetStakePool() ag_solanago.PublicKey {
	return s.Accounts[0].PublicKey
}

func (s *SetManager) GetManager() ag_solanago.PublicKey {
	return s.Accounts[1].PublicKey
}

func (s *SetManager) GetNewManager() ag_solanago.PublicKey {
	return s.Accounts[2].PublicKey
}

func (s *SetManager) GetNewManagerFeeAccount() ag_solanago.PublicKey {
	return s.Accounts[3].PublicKey
}

func (s *SetManager) ValidateAndBuild() (*Instruction, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s.Build(), nil
}

func (s *SetManager) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_SetManager),
			Impl:   s,
		},
	}
}

func (s *SetManager) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetManager")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
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

func (s *SetManager) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range s.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (s *SetManager) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range s.Accounts {
		if err := decoder.Decode(s.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (s *SetManager) Validate() error {
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
