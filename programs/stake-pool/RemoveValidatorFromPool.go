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

type RemoveValidatorFromPool struct {
	// [0]  = [WRITE] stakePool
	// [1]  = [SIGNER] staker
	// [2]  = [] withdrawAuthority
	// [3]  = [WRITE] validatorList
	// [4]  = [WRITE] validatorStakeAccount
	// [5]  = [WRITE] transientStakeAccount
	// [6]  = [] clock
	// [7]  = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewRemoveValidatorFromPoolInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *RemoveValidatorFromPool {
	return NewRemoveValidatorFromPoolInstructionBuilder().
		SetStakePool(stakePool).
		SetStaker(staker).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetClock(clock).
		SetStakeProgram(stakeProgram)
}

func NewRemoveValidatorFromPoolInstructionBuilder() *RemoveValidatorFromPool {
	return &RemoveValidatorFromPool{
		Accounts: make(ag_solanago.AccountMetaSlice, 8),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *RemoveValidatorFromPool) SetStakePool(pool ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *RemoveValidatorFromPool) SetStaker(staker ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *RemoveValidatorFromPool) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *RemoveValidatorFromPool) SetValidatorList(validatorList ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *RemoveValidatorFromPool) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[4] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return inst
}

func (inst *RemoveValidatorFromPool) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[5] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return inst
}

func (inst *RemoveValidatorFromPool) SetClock(clock ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[6] = ag_solanago.Meta(clock)
	return inst
}

func (inst *RemoveValidatorFromPool) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *RemoveValidatorFromPool {
	inst.Accounts[7] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *RemoveValidatorFromPool) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *RemoveValidatorFromPool) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *RemoveValidatorFromPool) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *RemoveValidatorFromPool) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *RemoveValidatorFromPool) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *RemoveValidatorFromPool) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *RemoveValidatorFromPool) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *RemoveValidatorFromPool) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *RemoveValidatorFromPool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveValidatorFromPool) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_RemoveValidatorFromPool),
			Impl:   inst,
		},
	}
}

func (inst *RemoveValidatorFromPool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveValidatorFromPool")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
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

func (inst *RemoveValidatorFromPool) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *RemoveValidatorFromPool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *RemoveValidatorFromPool) Validate() error {
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
