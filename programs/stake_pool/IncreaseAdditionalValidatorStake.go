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

type IncreaseAdditionalValidatorStake struct {
	Args *AdditionalValidatorStakeArgs
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] withdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] reserveStake
	// [5] = [WRITE] ephemeralStakeAccount
	// [6] = [WRITE] transientStakeAccount
	// [7] = [] validatorStakeAccount
	// [8] = [] voteAccount
	// [9] = [] clock
	// [10] = [] stakeHistory
	// [11] = [] stakeConfig
	// [12] = [] systemProgram
	// [13] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewIncreaseAdditionalValidatorStakeInstruction(
	// Parameters:
	args *AdditionalValidatorStakeArgs,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	ephemeralStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	voteAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	stakeConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *IncreaseAdditionalValidatorStake {
	return NewIncreaseAdditionalValidatorStakeInstructionBuilder().
		SetArgs(args).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetReserveStake(reserveStake).
		SetEphemeralStakeAccount(ephemeralStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetVoteAccount(voteAccount).
		SetClock(clock).
		SetStakeHistory(stakeHistory).
		SetStakeConfig(stakeConfig).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewIncreaseAdditionalValidatorStakeInstructionBuilder() *IncreaseAdditionalValidatorStake {
	return &IncreaseAdditionalValidatorStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 14),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (i *IncreaseAdditionalValidatorStake) SetArgs(args *AdditionalValidatorStakeArgs) *IncreaseAdditionalValidatorStake {
	i.Args = args
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetStakePool(stakePool ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[0] = ag_solanago.Meta(stakePool)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetStaker(staker ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	i.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetValidatorList(validatorList ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetReserveStake(reserveStake ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetEphemeralStakeAccount(ephemeralStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[5] = ag_solanago.Meta(ephemeralStakeAccount).WRITE()
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[6] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[7] = ag_solanago.Meta(validatorStakeAccount)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetVoteAccount(voteAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[8] = ag_solanago.Meta(voteAccount)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetClock(clock ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[9] = ag_solanago.Meta(clock)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[10] = ag_solanago.Meta(stakeHistory)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetStakeConfig(stakeConfig ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[11] = ag_solanago.Meta(stakeConfig)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetSystemProgram(systemProgram ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[12] = ag_solanago.Meta(systemProgram)
	return i
}

func (i *IncreaseAdditionalValidatorStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	i.Accounts[13] = ag_solanago.Meta(stakeProgram)
	return i
}

func (i *IncreaseAdditionalValidatorStake) GetArgs() *AdditionalValidatorStakeArgs {
	return i.Args
}

func (i *IncreaseAdditionalValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return i.Accounts[0].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetStaker() ag_solanago.PublicKey {
	return i.Accounts[1].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetWithdrawAuthority() ag_solanago.PublicKey {
	return i.Accounts[2].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return i.Accounts[3].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetReserveStake() ag_solanago.PublicKey {
	return i.Accounts[4].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetEphemeralStakeAccount() ag_solanago.PublicKey {
	return i.Accounts[5].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return i.Accounts[6].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return i.Accounts[7].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetVoteAccount() ag_solanago.PublicKey {
	return i.Accounts[8].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetClock() ag_solanago.PublicKey {
	return i.Accounts[9].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetStakeHistory() ag_solanago.PublicKey {
	return i.Accounts[10].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetStakeConfig() ag_solanago.PublicKey {
	return i.Accounts[11].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return i.Accounts[12].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return i.Accounts[13].PublicKey
}

func (i *IncreaseAdditionalValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}
	return i.Build(), nil
}

func (i *IncreaseAdditionalValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_IncreaseAdditionalValidatorStake),
			Impl:   i,
		},
	}
}

func (i *IncreaseAdditionalValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseAdditionalValidatorStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if i.Args != nil {
							paramsBranch.Child(ag_format.Param("Args", i.Args))
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

func (i *IncreaseAdditionalValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if i.Args != nil {
		if err := encoder.Encode(i.Args); err != nil {
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

func (i *IncreaseAdditionalValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if i.Args != nil {
		if err := decoder.Decode(i.Args); err != nil {
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

func (i *IncreaseAdditionalValidatorStake) Validate() error {
	if i.Args == nil {
		return errors.New("args is not set")
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
