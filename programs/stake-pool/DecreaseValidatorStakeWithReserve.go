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

type DecreaseValidatorStakeWithReserve struct {
	Lamports           *uint64
	TransientStakeSeed *uint64
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] stakePoolWithdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] reserveStakeAccount
	// [5] = [WRITE] canonicalStakeAccount
	// [6] = [WRITE] transientStakeAccount
	// [7] = [] sysvarClock
	// [8] = [] sysvarStakeHistory
	// [9] = [] systemProgram
	// [10] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewDecreaseValidatorStakeWithReserveWithReserveInstruction(
	// Parameters:
	lamports uint64,
	transientStakeSeed uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	stakePoolWithdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStakeAccount ag_solanago.PublicKey,
	canonicalStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	sysvarClock ag_solanago.PublicKey,
	sysvarStakeHistory ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *DecreaseValidatorStakeWithReserve {
	return NewDecreaseValidatorStakeWithReserveWithReserveInstructionBuilder().
		SetLamports(lamports).
		SetTransientStakeSeed(transientStakeSeed).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority).
		SetValidatorList(validatorList).
		SetReserveStakeAccount(reserveStakeAccount).
		SetCanonicalStakeAccount(canonicalStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetSysvarClock(sysvarClock).
		SetSysvarStakeHistory(sysvarStakeHistory).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewDecreaseValidatorStakeWithReserveWithReserveInstructionBuilder() *DecreaseValidatorStakeWithReserve {
	return &DecreaseValidatorStakeWithReserve{
		Accounts: make(ag_solanago.AccountMetaSlice, 11),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *DecreaseValidatorStakeWithReserve) SetLamports(lamports uint64) *DecreaseValidatorStakeWithReserve {
	inst.Lamports = &lamports
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetTransientStakeSeed(transientStakeSeed uint64) *DecreaseValidatorStakeWithReserve {
	inst.TransientStakeSeed = &transientStakeSeed
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetStakePool(pool ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[0] = ag_solanago.Meta(pool)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetStaker(staker ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetValidatorList(validatorList ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetReserveStakeAccount(reserveStakeAccount ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[4] = ag_solanago.Meta(reserveStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetCanonicalStakeAccount(canonicalStakeAccount ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[5] = ag_solanago.Meta(canonicalStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[6] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetSysvarClock(sysvarClock ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[7] = ag_solanago.Meta(sysvarClock)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetSysvarStakeHistory(sysvarStakeHist ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[8] = ag_solanago.Meta(sysvarStakeHist)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetSystemProgram(systemProgram ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[9] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *DecreaseValidatorStakeWithReserve {
	inst.Accounts[10] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *DecreaseValidatorStakeWithReserve) GetLamports() *uint64 {
	return inst.Lamports
}

func (inst *DecreaseValidatorStakeWithReserve) GetTransientStakeSeed() *uint64 {
	return inst.TransientStakeSeed
}

func (inst *DecreaseValidatorStakeWithReserve) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetReserveStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetCanonicalStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetSysvarClock() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetSysvarStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *DecreaseValidatorStakeWithReserve) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DecreaseValidatorStakeWithReserve) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DecreaseValidatorStakeWithReserve),
			Impl:   inst,
		},
	}
}

func (inst *DecreaseValidatorStakeWithReserve) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecreaseValidatorStakeWithReserve")).
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

func (inst *DecreaseValidatorStakeWithReserve) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
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

func (inst *DecreaseValidatorStakeWithReserve) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
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
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *DecreaseValidatorStakeWithReserve) Validate() error {
	if inst.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if inst.TransientStakeSeed == nil {
		return errors.New("transientStakeSeed is not set")
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
