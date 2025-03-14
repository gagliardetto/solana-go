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
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type UpdateStakePoolBalance struct {
	// [0] = [WRITE] stakePool
	// [1] = [] withdrawAuthority
	// [2] = [WRITE] validatorList
	// [3] = [] reserveStake
	// [4] = [WRITE] managerFeeAccount
	// [5] = [WRITE] poolMint
	// [6] = [] tokenProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewUpdateStakePoolBalanceInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *UpdateStakePoolBalance {
	return NewUpdateStakePoolBalanceInstructionBuilder().
		SetStakePool(stakePool).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetReserveStake(reserveStake).
		SetManagerFeeAccount(managerFeeAccount).
		SetPoolMint(poolMint).
		SetTokenProgram(tokenProgram)
}

func NewUpdateStakePoolBalanceInstructionBuilder() *UpdateStakePoolBalance {
	return &UpdateStakePoolBalance{
		Accounts: make(ag_solanago.AccountMetaSlice, 7),
	}
}

func (inst *UpdateStakePoolBalance) SetStakePool(pool ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *UpdateStakePoolBalance) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[1] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *UpdateStakePoolBalance) SetValidatorList(validatorList ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[2] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *UpdateStakePoolBalance) SetReserveStake(reserveStake ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[3] = ag_solanago.Meta(reserveStake)
	return inst
}

func (inst *UpdateStakePoolBalance) SetManagerFeeAccount(managerFeeAccount ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[4] = ag_solanago.Meta(managerFeeAccount).WRITE()
	return inst
}

func (inst *UpdateStakePoolBalance) SetPoolMint(poolMint ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[5] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

func (inst *UpdateStakePoolBalance) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *UpdateStakePoolBalance {
	inst.Accounts[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

func (inst *UpdateStakePoolBalance) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *UpdateStakePoolBalance) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *UpdateStakePoolBalance) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *UpdateStakePoolBalance) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *UpdateStakePoolBalance) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *UpdateStakePoolBalance) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *UpdateStakePoolBalance) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *UpdateStakePoolBalance) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateStakePoolBalance) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateStakePoolBalance),
			Impl:   inst,
		},
	}
}

func (inst *UpdateStakePoolBalance) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateStakePoolBalance")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range inst.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
					})
				})
		})
}

func (inst *UpdateStakePoolBalance) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *UpdateStakePoolBalance) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *UpdateStakePoolBalance) Validate() error {
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	return nil
}
