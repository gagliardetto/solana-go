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

type DecreaseAdditionalValidatorStake struct {
	Args *AdditionalValidatorStakeArgs
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] withdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] reserveStake
	// [5] = [WRITE] validatorStakeAccount
	// [6] = [WRITE] ephemeralStakeAccount
	// [7] = [WRITE] transientStakeAccount
	// [8] = [] clock
	// [9] = [] stakeHistory
	// [10] = [] systemProgram
	// [11] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewDecreaseAdditionalValidatorStakeInstruction(
	// Parameters:
	args *AdditionalValidatorStakeArgs,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	ephemeralStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *DecreaseAdditionalValidatorStake {
	return NewDecreaseAdditionalValidatorStakeBuilder().
		SetArgs(args).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetReserveStake(reserveStake).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetEphemeralStakeAccount(ephemeralStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetClock(clock).
		SetStakeHistory(stakeHistory).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewDecreaseAdditionalValidatorStakeBuilder() *DecreaseAdditionalValidatorStake {
	return &DecreaseAdditionalValidatorStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 12),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}

}

func (inst *DecreaseAdditionalValidatorStake) SetArgs(args *AdditionalValidatorStakeArgs) *DecreaseAdditionalValidatorStake {
	inst.Args = args
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetStakePool(stakePool ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[0] = ag_solanago.Meta(stakePool)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetStaker(staker ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetValidatorList(validatorList ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetReserveStake(reserveStake ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[5] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetEphemeralStakeAccount(ephemeralStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[6] = ag_solanago.Meta(ephemeralStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[7] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetClock(clock ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[8] = ag_solanago.Meta(clock)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[9] = ag_solanago.Meta(stakeHistory)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetSystemProgram(systemProgram ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[10] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	inst.Accounts[11] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *DecreaseAdditionalValidatorStake) GetArgs() *AdditionalValidatorStakeArgs {
	return inst.Args
}

func (inst *DecreaseAdditionalValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetEphemeralStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *DecreaseAdditionalValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DecreaseAdditionalValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DecreaseAdditionalValidatorStake),
			Impl:   inst,
		},
	}
}

func (inst *DecreaseAdditionalValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecreaseAdditionalValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Args != nil {
							paramsBranch.Child(ag_format.Param("Args", inst.Args))
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

func (inst *DecreaseAdditionalValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Args != nil {
		if err := encoder.Encode(inst.Args); err != nil {
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

func (inst *DecreaseAdditionalValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Args != nil {
		if err := decoder.Decode(inst.Args); err != nil {
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

func (inst *DecreaseAdditionalValidatorStake) Validate() error {
	if inst.Args == nil {
		return errors.New("args is not set")
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
