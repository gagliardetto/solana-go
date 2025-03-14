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

func (d *DecreaseAdditionalValidatorStake) SetArgs(args *AdditionalValidatorStakeArgs) *DecreaseAdditionalValidatorStake {
	d.Args = args
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetStakePool(stakePool ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[0] = ag_solanago.Meta(stakePool)
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetStaker(staker ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	d.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetValidatorList(validatorList ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetReserveStake(reserveStake ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[5] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetEphemeralStakeAccount(ephemeralStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[6] = ag_solanago.Meta(ephemeralStakeAccount).WRITE()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[7] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetClock(clock ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[8] = ag_solanago.Meta(clock)
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[9] = ag_solanago.Meta(stakeHistory)
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetSystemProgram(systemProgram ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[10] = ag_solanago.Meta(systemProgram)
	return d
}

func (d *DecreaseAdditionalValidatorStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *DecreaseAdditionalValidatorStake {
	d.Accounts[11] = ag_solanago.Meta(stakeProgram)
	return d
}

func (d *DecreaseAdditionalValidatorStake) GetArgs() *AdditionalValidatorStakeArgs {
	return d.Args
}

func (d *DecreaseAdditionalValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return d.Accounts[0].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetStaker() ag_solanago.PublicKey {
	return d.Accounts[1].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetWithdrawAuthority() ag_solanago.PublicKey {
	return d.Accounts[2].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return d.Accounts[3].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetReserveStake() ag_solanago.PublicKey {
	return d.Accounts[4].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[5].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetEphemeralStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[6].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[7].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetClock() ag_solanago.PublicKey {
	return d.Accounts[8].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetStakeHistory() ag_solanago.PublicKey {
	return d.Accounts[9].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return d.Accounts[10].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return d.Accounts[11].PublicKey
}

func (d *DecreaseAdditionalValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d.Build(), nil
}

func (d *DecreaseAdditionalValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DecreaseAdditionalValidatorStake),
			Impl:   d,
		},
	}
}

func (d *DecreaseAdditionalValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecreaseAdditionalValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if d.Args != nil {
							paramsBranch.Child(ag_format.Param("Args", d.Args))
						}
					})

					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range d.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(d.Signers)))
						for j, signer := range d.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (d *DecreaseAdditionalValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if d.Args != nil {
		if err := encoder.Encode(d.Args); err != nil {
			return err
		}
	}
	for _, account := range d.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (d *DecreaseAdditionalValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if d.Args != nil {
		if err := decoder.Decode(d.Args); err != nil {
			return err
		}
	}
	for i := range d.Accounts {
		if err := decoder.Decode(d.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (d *DecreaseAdditionalValidatorStake) Validate() error {
	if d.Args == nil {
		return errors.New("args is not set")
	}

	for i, account := range d.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(d.Signers) == 0 || !d.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
