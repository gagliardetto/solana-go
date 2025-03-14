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

type AddValidatorToPool struct {
	OptionalSeed *uint32
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] staker
	// [2] = [WRITE] reserveStake
	// [3] = [] withdrawAuthority
	// [4] = [WRITE] validatorList
	// [5] = [WRITE] validatorStakeAccount
	// [6] = [] voteAccount
	// [7] = [] rent
	// [8] = [] clock
	// [9] = [] stakeHistory
	// [10] = [] stakeConfig
	// [11] = [] systemProgram
	// [12] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (a *AddValidatorToPool) SetAccounts(accounts ag_solanago.AccountMetaSlice) *AddValidatorToPool {
	a.Accounts = accounts
	return a
}

func (a *AddValidatorToPool) SetSigners(signers ag_solanago.AccountMetaSlice) *AddValidatorToPool {
	a.Signers = signers
	return a
}

func (a *AddValidatorToPool) GetSigners() []*ag_solanago.AccountMeta {
	return a.Signers
}

func (a *AddValidatorToPool) GetAccounts() []*ag_solanago.AccountMeta {
	return a.Accounts
}

func NewAddValidatorToPoolInstruction(
	// Parameters:
	optionalSeed *uint32,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	voteAccount ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	stakeConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *AddValidatorToPool {
	return NewAddValidatorToPoolBuilder().
		SetOptionalSeed(optionalSeed).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetReserveStake(reserveStake).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetVoteAccount(voteAccount).
		SetRent(rent).
		SetClock(clock).
		SetStakeHistory(stakeHistory).
		SetStakeConfig(stakeConfig).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewAddValidatorToPoolBuilder() *AddValidatorToPool {
	return &AddValidatorToPool{
		Accounts: make(ag_solanago.AccountMetaSlice, 13),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (a *AddValidatorToPool) SetOptionalSeed(seed *uint32) *AddValidatorToPool {
	a.OptionalSeed = seed
	return a
}

func (a *AddValidatorToPool) SetStakePool(pool ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return a
}

func (a *AddValidatorToPool) SetStaker(staker ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	a.Signers[0] = a.Accounts[1]
	return a
}

func (a *AddValidatorToPool) SetReserveStake(reserveStake ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[2] = ag_solanago.Meta(reserveStake).WRITE()
	return a
}

func (a *AddValidatorToPool) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[3] = ag_solanago.Meta(withdrawAuthority)
	return a
}

func (a *AddValidatorToPool) SetValidatorList(validatorList ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[4] = ag_solanago.Meta(validatorList).WRITE()
	return a
}

func (a *AddValidatorToPool) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[5] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return a
}

func (a *AddValidatorToPool) SetVoteAccount(voteAccount ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[6] = ag_solanago.Meta(voteAccount)
	return a
}

func (a *AddValidatorToPool) SetRent(rent ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[7] = ag_solanago.Meta(rent)
	return a
}

func (a *AddValidatorToPool) SetClock(clock ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[8] = ag_solanago.Meta(clock)
	return a
}

func (a *AddValidatorToPool) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[9] = ag_solanago.Meta(stakeHistory)
	return a
}

func (a *AddValidatorToPool) SetStakeConfig(stakeConfig ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[10] = ag_solanago.Meta(stakeConfig)
	return a
}

func (a *AddValidatorToPool) SetSystemProgram(systemProgram ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[11] = ag_solanago.Meta(systemProgram)
	return a
}

func (a *AddValidatorToPool) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *AddValidatorToPool {
	a.Accounts[12] = ag_solanago.Meta(stakeProgram)
	return a
}

func (a *AddValidatorToPool) GetOptionalSeed() *uint32 {
	return a.OptionalSeed
}

func (a *AddValidatorToPool) GetStakePool() ag_solanago.PublicKey {
	return a.Accounts[0].PublicKey
}

func (a *AddValidatorToPool) GetStaker() ag_solanago.PublicKey {
	return a.Accounts[1].PublicKey
}

func (a *AddValidatorToPool) GetReserveStake() ag_solanago.PublicKey {
	return a.Accounts[2].PublicKey
}

func (a *AddValidatorToPool) GetWithdrawAuthority() ag_solanago.PublicKey {
	return a.Accounts[3].PublicKey
}

func (a *AddValidatorToPool) GetValidatorList() ag_solanago.PublicKey {
	return a.Accounts[4].PublicKey
}

func (a *AddValidatorToPool) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return a.Accounts[5].PublicKey
}

func (a *AddValidatorToPool) GetVoteAccount() ag_solanago.PublicKey {
	return a.Accounts[6].PublicKey
}

func (a *AddValidatorToPool) GetRent() ag_solanago.PublicKey {
	return a.Accounts[7].PublicKey
}

func (a *AddValidatorToPool) GetClock() ag_solanago.PublicKey {
	return a.Accounts[8].PublicKey
}

func (a *AddValidatorToPool) GetStakeHistory() ag_solanago.PublicKey {
	return a.Accounts[9].PublicKey
}

func (a *AddValidatorToPool) GetStakeConfig() ag_solanago.PublicKey {
	return a.Accounts[10].PublicKey
}

func (a *AddValidatorToPool) GetSystemProgram() ag_solanago.PublicKey {
	return a.Accounts[11].PublicKey
}

func (a *AddValidatorToPool) GetStakeProgram() ag_solanago.PublicKey {
	return a.Accounts[12].PublicKey
}

func (a *AddValidatorToPool) ValidateAndBuild() (*Instruction, error) {
	if err := a.Validate(); err != nil {
		return nil, err
	}
	return a.Build(), nil
}

func (a *AddValidatorToPool) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_AddValidatorToPool),
			Impl:   a,
		},
	}
}

func (a *AddValidatorToPool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddValidatorToPool")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if a.OptionalSeed != nil {
							paramsBranch.Child(ag_format.Param("OptionalSeed", *a.OptionalSeed))
						}
					})

					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range a.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(a.Signers)))
						for j, signer := range a.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (a *AddValidatorToPool) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if a.OptionalSeed != nil {
		if err := encoder.Encode(a.OptionalSeed); err != nil {
			return err
		}
	}
	for _, account := range a.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (a *AddValidatorToPool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	if a.OptionalSeed != nil {
		err = decoder.Decode(a.OptionalSeed)
		if err != nil {
			return err
		}
	}
	for i := range a.Accounts {
		err = decoder.Decode(a.Accounts[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AddValidatorToPool) Validate() error {
	for i, account := range a.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(a.Signers) == 0 || !a.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
