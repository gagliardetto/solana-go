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
	"encoding/binary"
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

func (inst *IncreaseAdditionalValidatorStake) GetAccounts() []*ag_solanago.AccountMeta {
	return inst.Accounts
}

func (inst *IncreaseAdditionalValidatorStake) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts = accounts
	return nil
}

func (inst *IncreaseAdditionalValidatorStake) SetArgs(args *AdditionalValidatorStakeArgs) *IncreaseAdditionalValidatorStake {
	inst.Args = args
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetStakePool(stakePool ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[0] = ag_solanago.Meta(stakePool)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetStaker(staker ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetValidatorList(validatorList ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetReserveStake(reserveStake ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetEphemeralStakeAccount(ephemeralStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[5] = ag_solanago.Meta(ephemeralStakeAccount).WRITE()
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[6] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[7] = ag_solanago.Meta(validatorStakeAccount)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetVoteAccount(voteAccount ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[8] = ag_solanago.Meta(voteAccount)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetClock(clock ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[9] = ag_solanago.Meta(clock)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[10] = ag_solanago.Meta(stakeHistory)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetStakeConfig(stakeConfig ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[11] = ag_solanago.Meta(stakeConfig)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetSystemProgram(systemProgram ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[12] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *IncreaseAdditionalValidatorStake {
	inst.Accounts[13] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *IncreaseAdditionalValidatorStake) GetArgs() *AdditionalValidatorStakeArgs {
	return inst.Args
}

func (inst *IncreaseAdditionalValidatorStake) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetEphemeralStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetVoteAccount() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetStakeConfig() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[13].PublicKey
}

func (inst *IncreaseAdditionalValidatorStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *IncreaseAdditionalValidatorStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_IncreaseAdditionalValidatorStake),
			Impl:   inst,
		},
	}
}

func (inst *IncreaseAdditionalValidatorStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseAdditionalValidatorStake")).
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

func (inst *IncreaseAdditionalValidatorStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	return encoder.Encode(inst.Args)
}

func (inst *IncreaseAdditionalValidatorStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return decoder.Decode(&inst.Args)
}

func (inst *IncreaseAdditionalValidatorStake) Validate() error {
	if inst.Args == nil {
		return errors.New("args is not set")
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

func (inst *IncreaseAdditionalValidatorStake) FindEphemeralAccount(programID, stakePoolAddress ag_solanago.PublicKey, seed uint64) (ag_solanago.PublicKey, uint8, error) {
	seedBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(seedBytes, seed)

	seeds := [][]byte{
		[]byte("ephemeral"),
		stakePoolAddress.Bytes(),
		seedBytes,
	}

	// Find Program Address (PDA)
	return ag_solanago.FindProgramAddress(seeds, programID)
}
