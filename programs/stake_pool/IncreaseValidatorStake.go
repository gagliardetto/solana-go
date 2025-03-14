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

func (i *IncreaseValidatorStake) SetLamports(lamports uint64) *IncreaseValidatorStake {
	i.Lamports = &lamports
	return i
}

func (i *IncreaseValidatorStake) SetTransientStakeSeed(seed uint64) *IncreaseValidatorStake {
	i.TransientStakeSeed = &seed
	return i
}

func (i *IncreaseValidatorStake) SetStakePool(pool ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[0] = ag_solanago.Meta(pool)
	return i
}

func (i *IncreaseValidatorStake) SetStaker(staker ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	i.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return i
}

func (i *IncreaseValidatorStake) SetStakePoolWithdrawAuthority(authority ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[2] = ag_solanago.Meta(authority)
	return i
}

func (i *IncreaseValidatorStake) SetValidatorList(list ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[3] = ag_solanago.Meta(list).WRITE()
	return i
}

func (i *IncreaseValidatorStake) SetStakePoolReserveStake(stake ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[4] = ag_solanago.Meta(stake).WRITE()
	return i
}

func (i *IncreaseValidatorStake) SetTransientStakeAccount(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[5] = ag_solanago.Meta(account).WRITE()
	return i
}

func (i *IncreaseValidatorStake) SetValidatorStakeAccount(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[6] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetValidatorVoteAccountToDelegateTo(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[7] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetSysvarClock(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[8] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetSysvarRent(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[9] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetStakeHistorySysvar(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[10] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetStakeConfigSysvar(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[11] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetSystemProgram(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[12] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) SetStakeProgram(account ag_solanago.PublicKey) *IncreaseValidatorStake {
	i.Accounts[13] = ag_solanago.Meta(account)
	return i
}

func (i *IncreaseValidatorStake) GetLamports() *uint64 {
	return i.Lamports
}

func (i *IncreaseValidatorStake) GetTransientStakeSeed() *uint64 {
	return i.TransientStakeSeed
}

func (i *IncreaseValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return i.Accounts[0].PublicKey
}

func (i *IncreaseValidatorStake) GetStaker() ag_solanago.PublicKey {
	return i.Accounts[1].PublicKey
}

func (i *IncreaseValidatorStake) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return i.Accounts[2].PublicKey
}

func (i *IncreaseValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return i.Accounts[3].PublicKey
}

func (i *IncreaseValidatorStake) GetStakePoolReserveStake() ag_solanago.PublicKey {
	return i.Accounts[4].PublicKey
}

func (i *IncreaseValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return i.Accounts[5].PublicKey
}

func (i *IncreaseValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return i.Accounts[6].PublicKey
}

func (i *IncreaseValidatorStake) GetValidatorVoteAccountToDelegateTo() ag_solanago.PublicKey {
	return i.Accounts[7].PublicKey
}

func (i *IncreaseValidatorStake) GetSysvarClock() ag_solanago.PublicKey {
	return i.Accounts[8].PublicKey
}

func (i *IncreaseValidatorStake) GetSysvarRent() ag_solanago.PublicKey {
	return i.Accounts[9].PublicKey
}

func (i *IncreaseValidatorStake) GetStakeHistorySysvar() ag_solanago.PublicKey {
	return i.Accounts[10].PublicKey
}

func (i *IncreaseValidatorStake) GetStakeConfigSysvar() ag_solanago.PublicKey {
	return i.Accounts[11].PublicKey
}

func (i *IncreaseValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return i.Accounts[12].PublicKey
}

func (i *IncreaseValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return i.Accounts[13].PublicKey
}

func (i *IncreaseValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}
	return i.Build(), nil
}

func (i *IncreaseValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_IncreaseValidatorStake),
			Impl:   i,
		},
	}
}

func (i *IncreaseValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if i.Lamports != nil {
							paramsBranch.Child(ag_format.Param("Lamports", *i.Lamports))
						}
						if i.TransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("TransientStakeSeed", *i.TransientStakeSeed))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range i.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(i.Signers)))
						for j, signer := range i.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (i *IncreaseValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if i.Lamports != nil {
		if err := encoder.Encode(i.Lamports); err != nil {
			return err
		}
	}
	if i.TransientStakeSeed != nil {
		if err := encoder.Encode(i.TransientStakeSeed); err != nil {
			return err
		}
	}
	for _, account := range i.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (i *IncreaseValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if i.Lamports != nil {
		if err := decoder.Decode(i.Lamports); err != nil {
			return err
		}
	}
	if i.TransientStakeSeed != nil {
		if err := decoder.Decode(i.TransientStakeSeed); err != nil {
			return err
		}
	}
	for j := range i.Accounts {
		if err := decoder.Decode(i.Accounts[j]); err != nil {
			return err
		}
	}
	return nil
}

func (i *IncreaseValidatorStake) Validate() error {
	if i.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if i.TransientStakeSeed == nil {
		return errors.New("TransientStakeSeed is not set")
	}
	for j, account := range i.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", j)
		}
	}
	if len(i.Signers) == 0 || !i.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
