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

type DecreaseValidatorStake struct {
	Lamports           *uint64
	TransientStakeSeed *uint64
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] stakePoolWithdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] canonicalStakeAccount
	// [5] = [WRITE] transientStakeAccount
	// [6] = [] sysvarClock
	// [7] = [] sysvarRent
	// [8] = [] systemProgram
	// [9] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewDecreaseValidatorStakeInstruction(
	// Parameters:
	lamports uint64,
	transientStakeSeed uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	stakePoolWithdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	canonicalStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	sysvarClock ag_solanago.PublicKey,
	sysvarRent ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *DecreaseValidatorStake {
	return NewDecreaseValidatorStakeInstructionBuilder().
		SetLamports(lamports).
		SetTransientStakeSeed(transientStakeSeed).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority).
		SetValidatorList(validatorList).
		SetCanonicalStakeAccount(canonicalStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetSysvarClock(sysvarClock).
		SetSysvarRent(sysvarRent).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewDecreaseValidatorStakeInstructionBuilder() *DecreaseValidatorStake {
	return &DecreaseValidatorStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 10),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (d *DecreaseValidatorStake) SetLamports(lamports uint64) *DecreaseValidatorStake {
	d.Lamports = &lamports
	return d
}

func (d *DecreaseValidatorStake) SetTransientStakeSeed(transientStakeSeed uint64) *DecreaseValidatorStake {
	d.TransientStakeSeed = &transientStakeSeed
	return d
}

func (d *DecreaseValidatorStake) SetStakePool(pool ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[0] = ag_solanago.Meta(pool)
	return d
}

func (d *DecreaseValidatorStake) SetStaker(staker ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	d.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return d
}

func (d *DecreaseValidatorStake) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return d
}

func (d *DecreaseValidatorStake) SetValidatorList(validatorList ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return d
}

func (d *DecreaseValidatorStake) SetCanonicalStakeAccount(canonicalStakeAccount ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[4] = ag_solanago.Meta(canonicalStakeAccount).WRITE()
	return d
}

func (d *DecreaseValidatorStake) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[5] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return d
}

func (d *DecreaseValidatorStake) SetSysvarClock(sysvarClock ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[6] = ag_solanago.Meta(sysvarClock)
	return d
}

func (d *DecreaseValidatorStake) SetSysvarRent(sysvarRent ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[7] = ag_solanago.Meta(sysvarRent)
	return d
}

func (d *DecreaseValidatorStake) SetSystemProgram(systemProgram ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[8] = ag_solanago.Meta(systemProgram)
	return d
}

func (d *DecreaseValidatorStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *DecreaseValidatorStake {
	d.Accounts[9] = ag_solanago.Meta(stakeProgram)
	return d
}

func (d *DecreaseValidatorStake) GetLamports() *uint64 {
	return d.Lamports
}

func (d *DecreaseValidatorStake) GetTransientStakeSeed() *uint64 {
	return d.TransientStakeSeed
}

func (d *DecreaseValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return d.Accounts[0].PublicKey
}

func (d *DecreaseValidatorStake) GetStaker() ag_solanago.PublicKey {
	return d.Accounts[1].PublicKey
}

func (d *DecreaseValidatorStake) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return d.Accounts[2].PublicKey
}

func (d *DecreaseValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return d.Accounts[3].PublicKey
}

func (d *DecreaseValidatorStake) GetCanonicalStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[4].PublicKey
}

func (d *DecreaseValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[5].PublicKey
}

func (d *DecreaseValidatorStake) GetSysvarClock() ag_solanago.PublicKey {
	return d.Accounts[6].PublicKey
}

func (d *DecreaseValidatorStake) GetSysvarRent() ag_solanago.PublicKey {
	return d.Accounts[7].PublicKey
}

func (d *DecreaseValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return d.Accounts[8].PublicKey
}

func (d *DecreaseValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return d.Accounts[9].PublicKey
}

func (d *DecreaseValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d.Build(), nil
}

func (d *DecreaseValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DecreaseValidatorStake),
			Impl:   d,
		},
	}
}

func (d *DecreaseValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecreaseValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if d.Lamports != nil {
							paramsBranch.Child(ag_format.Param("Lamports", *d.Lamports))
						}
						if d.TransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("TransientStakeSeed", *d.TransientStakeSeed))
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

func (d *DecreaseValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if d.Lamports != nil {
		if err := encoder.Encode(d.Lamports); err != nil {
			return err
		}
	}
	if d.TransientStakeSeed != nil {
		if err := encoder.Encode(d.TransientStakeSeed); err != nil {
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

func (d *DecreaseValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if d.Lamports != nil {
		if err := decoder.Decode(d.Lamports); err != nil {
			return err
		}
	}
	if d.TransientStakeSeed != nil {
		if err := decoder.Decode(d.TransientStakeSeed); err != nil {
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

func (d *DecreaseValidatorStake) Validate() error {
	if d.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if d.TransientStakeSeed == nil {
		return errors.New("transientStakeSeed is not set")
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
