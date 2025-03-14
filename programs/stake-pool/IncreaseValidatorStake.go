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

type IncreaseValidatorStake struct {
	Lamports           *uint64
	TransientStakeSeed *uint64
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] stakePoolWithdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] stakePoolReserveStake
	// [5] = [WRITE] transientStakeAccount
	// [6] = [] validatorStakeAccount
	// [7] = [] validatorVoteAccountToDelegateTo
	// [8] = [] sysvarClock
	// [9] = [] sysvarRent
	// [10] = [] stakeHistorySysvar
	// [11] = [] stakeConfigSysvar
	// [12] = [] systemProgram
	// [13] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewIncreaseValidatorStakeInstruction(
	// Parameters
	lamports uint64,
	transientStakeSeed uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	stakePoolWithdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	stakePoolReserveStake ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	validatorVoteAccountToDelegateTo ag_solanago.PublicKey,
	sysvarClock ag_solanago.PublicKey,
	sysvarRent ag_solanago.PublicKey,
	stakeHistorySysvar ag_solanago.PublicKey,
	stakeConfigSysvar ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *IncreaseValidatorStake {
	return NewIncreaseValidatorStakeInstructionBuilder().
		SetLamports(lamports).
		SetTransientStakeSeed(transientStakeSeed).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority).
		SetValidatorList(validatorList).
		SetStakePoolReserveStake(stakePoolReserveStake).
		SetTransientStakeAccount(transientStakeAccount).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetValidatorVoteAccountToDelegateTo(validatorVoteAccountToDelegateTo).
		SetSysvarClock(sysvarClock).
		SetSysvarRent(sysvarRent).
		SetStakeHistorySysvar(stakeHistorySysvar).
		SetStakeConfigSysvar(stakeConfigSysvar).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewIncreaseValidatorStakeInstructionBuilder() *IncreaseValidatorStake {
	return &IncreaseValidatorStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 14),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *IncreaseValidatorStake) SetLamports(lamports uint64) *IncreaseValidatorStake {
	inst.Lamports = &lamports
	return inst
}

func (inst *IncreaseValidatorStake) SetTransientStakeSeed(seed uint64) *IncreaseValidatorStake {
	inst.TransientStakeSeed = &seed
	return inst
}

func (inst *IncreaseValidatorStake) SetStakePool(pool ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[0] = ag_solanago.Meta(pool)
	return inst
}

func (inst *IncreaseValidatorStake) SetStaker(staker ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *IncreaseValidatorStake) SetStakePoolWithdrawAuthority(authority ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[2] = ag_solanago.Meta(authority)
	return inst
}

func (inst *IncreaseValidatorStake) SetValidatorList(list ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[3] = ag_solanago.Meta(list).WRITE()
	return inst
}

func (inst *IncreaseValidatorStake) SetStakePoolReserveStake(stake ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[4] = ag_solanago.Meta(stake).WRITE()
	return inst
}

func (inst *IncreaseValidatorStake) SetTransientStakeAccount(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[5] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *IncreaseValidatorStake) SetValidatorStakeAccount(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[6] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetValidatorVoteAccountToDelegateTo(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[7] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetSysvarClock(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[8] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetSysvarRent(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[9] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetStakeHistorySysvar(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[10] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetStakeConfigSysvar(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[11] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetSystemProgram(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[12] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) SetStakeProgram(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	inst.Accounts[13] = ag_solanago.Meta(account)
	return inst
}

func (inst *IncreaseValidatorStake) GetLamports() *uint64 {
	return inst.Lamports
}

func (inst *IncreaseValidatorStake) GetTransientStakeSeed() *uint64 {
	return inst.TransientStakeSeed
}

func (inst *IncreaseValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *IncreaseValidatorStake) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *IncreaseValidatorStake) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *IncreaseValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *IncreaseValidatorStake) GetStakePoolReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *IncreaseValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *IncreaseValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *IncreaseValidatorStake) GetValidatorVoteAccountToDelegateTo() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *IncreaseValidatorStake) GetSysvarClock() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *IncreaseValidatorStake) GetSysvarRent() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *IncreaseValidatorStake) GetStakeHistorySysvar() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *IncreaseValidatorStake) GetStakeConfigSysvar() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *IncreaseValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *IncreaseValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[13].PublicKey
}

func (inst *IncreaseValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *IncreaseValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_IncreaseValidatorStake),
			Impl:   inst,
		},
	}
}

func (inst *IncreaseValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Lamports != nil {
							paramsBranch.Child(ag_format.Param("Lamports", *inst.Lamports))
						}
						if inst.TransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("TransientStakeSeed", *inst.TransientStakeSeed))
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

func (inst *IncreaseValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Lamports != nil {
		if err := encoder.Encode(inst.Lamports); err != nil {
			return err
		}
	}
	if inst.TransientStakeSeed != nil {
		if err := encoder.Encode(inst.TransientStakeSeed); err != nil {
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

func (inst *IncreaseValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Lamports != nil {
		if err := decoder.Decode(inst.Lamports); err != nil {
			return err
		}
	}
	if inst.TransientStakeSeed != nil {
		if err := decoder.Decode(inst.TransientStakeSeed); err != nil {
			return err
		}
	}
	for j := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[j]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *IncreaseValidatorStake) Validate() error {
	if inst.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if inst.TransientStakeSeed == nil {
		return errors.New("TransientStakeSeed is not set")
	}
	for j, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", j)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
