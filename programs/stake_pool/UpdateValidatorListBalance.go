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

func (u *UpdateValidatorListBalance) SetStartIndex(index uint32) *UpdateValidatorListBalance {
	u.StartIndex = &index
	return u
}

func (u *UpdateValidatorListBalance) SetNoMerge(noMerge bool) *UpdateValidatorListBalance {
	u.NoMerge = &noMerge
	return u
}

func (u *UpdateValidatorListBalance) SetStakePool(pool ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[0] = ag_solanago.Meta(pool)
	return u
}

func (u *UpdateValidatorListBalance) SetStaker(staker ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	u.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return u
}

func (u *UpdateValidatorListBalance) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return u
}

func (u *UpdateValidatorListBalance) SetValidatorList(validatorList ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return u
}

func (u *UpdateValidatorListBalance) SetReserveStake(reserveStake ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return u
}

func (u *UpdateValidatorListBalance) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[5] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return u
}

func (u *UpdateValidatorListBalance) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[6] = ag_solanago.Meta(validatorStakeAccount)
	return u
}

func (u *UpdateValidatorListBalance) SetValidatorVoteAccount(validatorVoteAccount ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[7] = ag_solanago.Meta(validatorVoteAccount)
	return u
}

func (u *UpdateValidatorListBalance) SetClock(clock ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[8] = ag_solanago.Meta(clock)
	return u
}

func (u *UpdateValidatorListBalance) SetRent(rent ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[9] = ag_solanago.Meta(rent)
	return u
}

func (u *UpdateValidatorListBalance) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[10] = ag_solanago.Meta(stakeHistory)
	return u
}

func (u *UpdateValidatorListBalance) SetStakeConfig(stakeConfig ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[11] = ag_solanago.Meta(stakeConfig)
	return u
}

func (u *UpdateValidatorListBalance) SetSystemProgram(systemProgram ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[12] = ag_solanago.Meta(systemProgram)
	return u
}

func (u *UpdateValidatorListBalance) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *UpdateValidatorListBalance {
	u.Accounts[13] = ag_solanago.Meta(stakeProgram)
	return u
}

func (u *UpdateValidatorListBalance) GetStartIndex() *uint32 {
	return u.StartIndex
}

func (u *UpdateValidatorListBalance) GetNoMerge() *bool {
	return u.NoMerge
}

func (u *UpdateValidatorListBalance) GetStakePool() ag_solanago.PublicKey {
	return u.Accounts[0].PublicKey
}

func (u *UpdateValidatorListBalance) GetStaker() ag_solanago.PublicKey {
	return u.Accounts[1].PublicKey
}

func (u *UpdateValidatorListBalance) GetWithdrawAuthority() ag_solanago.PublicKey {
	return u.Accounts[2].PublicKey
}

func (u *UpdateValidatorListBalance) GetValidatorList() ag_solanago.PublicKey {
	return u.Accounts[3].PublicKey
}

func (u *UpdateValidatorListBalance) GetReserveStake() ag_solanago.PublicKey {
	return u.Accounts[4].PublicKey
}

func (u *UpdateValidatorListBalance) GetTransientStakeAccount() ag_solanago.PublicKey {
	return u.Accounts[5].PublicKey
}

func (u *UpdateValidatorListBalance) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return u.Accounts[6].PublicKey
}

func (u *UpdateValidatorListBalance) GetValidatorVoteAccount() ag_solanago.PublicKey {
	return u.Accounts[7].PublicKey
}

func (u *UpdateValidatorListBalance) GetClock() ag_solanago.PublicKey {
	return u.Accounts[8].PublicKey
}

func (u *UpdateValidatorListBalance) GetRent() ag_solanago.PublicKey {
	return u.Accounts[9].PublicKey
}

func (u *UpdateValidatorListBalance) GetStakeHistory() ag_solanago.PublicKey {
	return u.Accounts[10].PublicKey
}

func (u *UpdateValidatorListBalance) GetStakeConfig() ag_solanago.PublicKey {
	return u.Accounts[11].PublicKey
}

func (u *UpdateValidatorListBalance) GetSystemProgram() ag_solanago.PublicKey {
	return u.Accounts[12].PublicKey
}

func (u *UpdateValidatorListBalance) GetStakeProgram() ag_solanago.PublicKey {
	return u.Accounts[13].PublicKey
}

func (u *UpdateValidatorListBalance) ValidateAndBuild() (*Instruction, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return u.Build(), nil
}

func (u *UpdateValidatorListBalance) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateValidatorListBalance),
			Impl:   u,
		},
	}
}

func (u *UpdateValidatorListBalance) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateValidatorListBalance")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if u.StartIndex != nil {
							paramsBranch.Child(ag_format.Param("StartIndex", *u.StartIndex))
						}
						if u.NoMerge != nil {
							paramsBranch.Child(ag_format.Param("NoMerge", *u.NoMerge))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range u.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(u.Signers)))
						for j, signer := range u.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (u *UpdateValidatorListBalance) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if u.StartIndex != nil {
		if err := encoder.Encode(u.StartIndex); err != nil {
			return err
		}
	}
	if u.NoMerge != nil {
		if err := encoder.Encode(u.NoMerge); err != nil {
			return err
		}
	}
	for _, account := range u.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (u *UpdateValidatorListBalance) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if u.StartIndex != nil {
		if err := decoder.Decode(u.StartIndex); err != nil {
			return err
		}
	}
	if u.NoMerge != nil {
		if err := decoder.Decode(u.NoMerge); err != nil {
			return err
		}
	}
	for i := range u.Accounts {
		if err := decoder.Decode(u.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (u *UpdateValidatorListBalance) Validate() error {
	for i, account := range u.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(u.Signers) == 0 || !u.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
