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

type SetPreferredValidator struct {
	ValidatorType        PreferredValidatorType
	ValidatorVoteAddress *ag_solanago.PublicKey
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] validatorList
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewSetPreferredValidatorInstruction(
	// Parameters:
	validatorType PreferredValidatorType,
	validatorVoteAddress *ag_solanago.PublicKey,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
) *SetPreferredValidator {
	return NewSetPreferredValidatorInstructionBuilder().
		SetValidatorType(validatorType).
		SetValidatorVoteAddress(validatorVoteAddress).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetValidatorList(validatorList)
}

func NewSetPreferredValidatorInstructionBuilder() *SetPreferredValidator {
	return &SetPreferredValidator{
		Accounts: make(ag_solanago.AccountMetaSlice, 3),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *SetPreferredValidator) SetValidatorType(validatorType PreferredValidatorType) *SetPreferredValidator {
	inst.ValidatorType = validatorType
	return inst
}

func (inst *SetPreferredValidator) SetValidatorVoteAddress(validatorVoteAddress *ag_solanago.PublicKey) *SetPreferredValidator {
	inst.ValidatorVoteAddress = validatorVoteAddress
	return inst
}

func (inst *SetPreferredValidator) SetStakePool(stakePool ag_solanago.PublicKey) *SetPreferredValidator {
	inst.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return inst
}

func (inst *SetPreferredValidator) SetStaker(staker ag_solanago.PublicKey) *SetPreferredValidator {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *SetPreferredValidator) SetValidatorList(validatorList ag_solanago.PublicKey) *SetPreferredValidator {
	inst.Accounts[2] = ag_solanago.Meta(validatorList)
	return inst
}

func (inst *SetPreferredValidator) GetValidatorType() PreferredValidatorType {
	return inst.ValidatorType
}

func (inst *SetPreferredValidator) GetValidatorVoteAddress() *ag_solanago.PublicKey {
	return inst.ValidatorVoteAddress
}

func (inst *SetPreferredValidator) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *SetPreferredValidator) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *SetPreferredValidator) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *SetPreferredValidator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetPreferredValidator) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_SetPreferredValidator),
			Impl:   inst,
		},
	}
}

func (inst *SetPreferredValidator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetPreferredValidator")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ValidatorType", inst.ValidatorType))
						if inst.ValidatorVoteAddress != nil {
							paramsBranch.Child(ag_format.Param("ValidatorVoteAddress", *inst.ValidatorVoteAddress))
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

func (inst *SetPreferredValidator) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if err := encoder.Encode(inst.ValidatorType); err != nil {
		return err
	}
	if inst.ValidatorVoteAddress != nil {
		if err := encoder.Encode(inst.ValidatorVoteAddress); err != nil {
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

func (inst *SetPreferredValidator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if err := decoder.Decode(&inst.ValidatorType); err != nil {
		return err
	}
	if inst.ValidatorVoteAddress != nil {
		if err := decoder.Decode(inst.ValidatorVoteAddress); err != nil {
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

func (inst *SetPreferredValidator) Validate() error {
	if inst.ValidatorVoteAddress == nil {
		return errors.New("validatorVoteAddress is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
