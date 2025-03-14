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

type UpdateValidatorListBalance struct {
	StartIndex *uint32
	NoMerge    *bool
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] withdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] reserveStake
	// [5] = [WRITE] transientStakeAccount
	// [6] = [] validatorStakeAccount
	// [7] = [] validatorVoteAccount
	// [8] = [] clock
	// [9] = [] rent
	// [10] = [] stakeHistory
	// [11] = [] stakeConfig
	// [12] = [] systemProgram
	// [13] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewUpdateValidatorListBalanceInstruction(
	// Parameters:
	startIndex uint32,
	noMerge bool,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	validatorVoteAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	stakeConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *UpdateValidatorListBalance {
	return NewUpdateValidatorListBalanceInstructionBuilder().
		SetStartIndex(startIndex).
		SetNoMerge(noMerge).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetReserveStake(reserveStake).
		SetTransientStakeAccount(transientStakeAccount).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetValidatorVoteAccount(validatorVoteAccount).
		SetClock(clock).
		SetRent(rent).
		SetStakeHistory(stakeHistory).
		SetStakeConfig(stakeConfig).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewUpdateValidatorListBalanceInstructionBuilder() *UpdateValidatorListBalance {
	return &UpdateValidatorListBalance{
		Accounts: make(ag_solanago.AccountMetaSlice, 9),
	}
}

func (inst *UpdateValidatorListBalance) SetStartIndex(index uint32) *UpdateValidatorListBalance {
	inst.StartIndex = &index
	return inst
}

func (inst *UpdateValidatorListBalance) SetNoMerge(noMerge bool) *UpdateValidatorListBalance {
	inst.NoMerge = &noMerge
	return inst
}

func (inst *UpdateValidatorListBalance) SetStakePool(pool ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[0] = ag_solanago.Meta(pool)
	return inst
}

func (inst *UpdateValidatorListBalance) SetStaker(staker ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *UpdateValidatorListBalance) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *UpdateValidatorListBalance) SetValidatorList(validatorList ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *UpdateValidatorListBalance) SetReserveStake(reserveStake ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return inst
}

func (inst *UpdateValidatorListBalance) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[5] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return inst
}

func (inst *UpdateValidatorListBalance) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[6] = ag_solanago.Meta(validatorStakeAccount)
	return inst
}

func (inst *UpdateValidatorListBalance) SetValidatorVoteAccount(validatorVoteAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[7] = ag_solanago.Meta(validatorVoteAccount)
	return inst
}

func (inst *UpdateValidatorListBalance) SetClock(clock ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[8] = ag_solanago.Meta(clock)
	return inst
}

func (inst *UpdateValidatorListBalance) SetRent(rent ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[9] = ag_solanago.Meta(rent)
	return inst
}

func (inst *UpdateValidatorListBalance) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[10] = ag_solanago.Meta(stakeHistory)
	return inst
}

func (inst *UpdateValidatorListBalance) SetStakeConfig(stakeConfig ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[11] = ag_solanago.Meta(stakeConfig)
	return inst
}

func (inst *UpdateValidatorListBalance) SetSystemProgram(systemProgram ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[12] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *UpdateValidatorListBalance) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *UpdateValidatorListBalance {
	inst.Accounts[13] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *UpdateValidatorListBalance) GetStartIndex() *uint32 {
	return inst.StartIndex
}

func (inst *UpdateValidatorListBalance) GetNoMerge() *bool {
	return inst.NoMerge
}

func (inst *UpdateValidatorListBalance) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *UpdateValidatorListBalance) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *UpdateValidatorListBalance) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *UpdateValidatorListBalance) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *UpdateValidatorListBalance) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *UpdateValidatorListBalance) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *UpdateValidatorListBalance) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *UpdateValidatorListBalance) GetValidatorVoteAccount() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *UpdateValidatorListBalance) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *UpdateValidatorListBalance) GetRent() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *UpdateValidatorListBalance) GetStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *UpdateValidatorListBalance) GetStakeConfig() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *UpdateValidatorListBalance) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *UpdateValidatorListBalance) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[13].PublicKey
}

func (inst *UpdateValidatorListBalance) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateValidatorListBalance) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateValidatorListBalance),
			Impl:   inst,
		},
	}
}

func (inst *UpdateValidatorListBalance) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateValidatorListBalance")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.StartIndex != nil {
							paramsBranch.Child(ag_format.Param("StartIndex", *inst.StartIndex))
						}
						if inst.NoMerge != nil {
							paramsBranch.Child(ag_format.Param("NoMerge", *inst.NoMerge))
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

func (inst *UpdateValidatorListBalance) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.StartIndex != nil {
		if err := encoder.Encode(inst.StartIndex); err != nil {
			return err
		}
	}
	if inst.NoMerge != nil {
		if err := encoder.Encode(inst.NoMerge); err != nil {
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

func (inst *UpdateValidatorListBalance) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.StartIndex != nil {
		if err := decoder.Decode(inst.StartIndex); err != nil {
			return err
		}
	}
	if inst.NoMerge != nil {
		if err := decoder.Decode(inst.NoMerge); err != nil {
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

func (inst *UpdateValidatorListBalance) Validate() error {
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
