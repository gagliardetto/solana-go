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

type WithdrawStake struct {
	Amount *uint64
	// [0] = [WRITE] stakePool
	// [1] = [WRITE] validatorStakeList
	// [2] = [] stakePoolWithdrawAuthority
	// [3] = [WRITE] validatorAccount
	// [4] = [WRITE] uninitializedStakeAccount
	// [5] = [] userAccount
	// [6] = [SIGNER] userTransferAuthority
	// [7] = [WRITE] userAccountWithPoolTokensToBurnFrom
	// [8] = [WRITE] accountToReceivePoolFeeTokens
	// [9] = [WRITE] poolTokenMintAccount
	// [10] = [] sysvarClockAccount
	// [11] = [] poolTokenProgramID
	// [12] = [] stakeProgramID
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewWithdrawStakeInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	validatorStakeList ag_solanago.PublicKey,
	stakePoolWithdrawAuthority ag_solanago.PublicKey,
	validatorAccount ag_solanago.PublicKey,
	uninitializedStakeAccount ag_solanago.PublicKey,
	userAccount ag_solanago.PublicKey,
	userTransferAuthority ag_solanago.PublicKey,
	userAccountWithPoolTokensToBurnFrom ag_solanago.PublicKey,
	accountToReceivePoolFeeTokens ag_solanago.PublicKey,
	poolTokenMintAccount ag_solanago.PublicKey,
	sysvarClockAccount ag_solanago.PublicKey,
	poolTokenProgramID ag_solanago.PublicKey,
	stakeProgramID ag_solanago.PublicKey,
) *WithdrawStake {
	return NewWithdrawStakeBuilder().
		SetAmount(amount).
		SetStakePool(stakePool).
		SetValidatorStakeList(validatorStakeList).
		SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority).
		SetValidatorAccount(validatorAccount).
		SetUninitializedStakeAccount(uninitializedStakeAccount).
		SetUserAccount(userAccount).
		SetUserTransferAuthority(userTransferAuthority).
		SetUserAccountWithPoolTokensToBurnFrom(userAccountWithPoolTokensToBurnFrom).
		SetAccountToReceivePoolFeeTokens(accountToReceivePoolFeeTokens).
		SetPoolTokenMintAccount(poolTokenMintAccount).
		SetSysvarClockAccount(sysvarClockAccount).
		SetPoolTokenProgramID(poolTokenProgramID).
		SetStakeProgramID(stakeProgramID)
}

func NewWithdrawStakeBuilder() *WithdrawStake {
	return &WithdrawStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 13),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (inst *WithdrawStake) SetAmount(amount uint64) *WithdrawStake {
	inst.Amount = &amount
	return inst
}

func (inst *WithdrawStake) SetStakePool(stakePool ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return inst
}

func (inst *WithdrawStake) SetValidatorStakeList(validatorStakeList ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[1] = ag_solanago.Meta(validatorStakeList).WRITE()
	return inst
}

func (inst *WithdrawStake) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return inst
}

func (inst *WithdrawStake) SetValidatorAccount(validatorAccount ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[3] = ag_solanago.Meta(validatorAccount).WRITE()
	return inst
}

func (inst *WithdrawStake) SetUninitializedStakeAccount(uninitializedStakeAccount ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[4] = ag_solanago.Meta(uninitializedStakeAccount).WRITE()
	return inst
}

func (inst *WithdrawStake) SetUserAccount(userAccount ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[5] = ag_solanago.Meta(userAccount)
	return inst
}

func (inst *WithdrawStake) SetUserTransferAuthority(userTransferAuthority ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[6] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return inst
}

func (inst *WithdrawStake) SetUserAccountWithPoolTokensToBurnFrom(userAccountWithPoolTokensToBurnFrom ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[7] = ag_solanago.Meta(userAccountWithPoolTokensToBurnFrom).WRITE()
	return inst
}

func (inst *WithdrawStake) SetAccountToReceivePoolFeeTokens(accountToReceivePoolFeeTokens ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[8] = ag_solanago.Meta(accountToReceivePoolFeeTokens).WRITE()
	return inst
}

func (inst *WithdrawStake) SetPoolTokenMintAccount(poolTokenMintAccount ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[9] = ag_solanago.Meta(poolTokenMintAccount).WRITE()
	return inst
}

func (inst *WithdrawStake) SetSysvarClockAccount(sysvarClockAccount ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[10] = ag_solanago.Meta(sysvarClockAccount)
	return inst
}

func (inst *WithdrawStake) SetPoolTokenProgramID(poolTokenProgramID ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[11] = ag_solanago.Meta(poolTokenProgramID)
	return inst
}

func (inst *WithdrawStake) SetStakeProgramID(stakeProgramID ag_solanago.PublicKey) *WithdrawStake {
	inst.Accounts[12] = ag_solanago.Meta(stakeProgramID)
	return inst
}

func (inst *WithdrawStake) GetAmount() *uint64 {
	return inst.Amount
}

func (inst *WithdrawStake) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *WithdrawStake) GetValidatorStakeList() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *WithdrawStake) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *WithdrawStake) GetValidatorAccount() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *WithdrawStake) GetUninitializedStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *WithdrawStake) GetUserAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *WithdrawStake) GetUserTransferAuthority() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *WithdrawStake) GetUserAccountWithPoolTokensToBurnFrom() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *WithdrawStake) GetAccountToReceivePoolFeeTokens() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *WithdrawStake) GetPoolTokenMintAccount() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *WithdrawStake) GetSysvarClockAccount() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *WithdrawStake) GetPoolTokenProgramID() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *WithdrawStake) GetStakeProgramID() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *WithdrawStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawStake),
			Impl:   inst,
		},
	}
}

func (inst *WithdrawStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Amount != nil {
							paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
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

func (inst *WithdrawStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Amount != nil {
		if err := encoder.Encode(inst.Amount); err != nil {
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

func (inst *WithdrawStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Amount != nil {
		if err := decoder.Decode(inst.Amount); err != nil {
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

func (inst *WithdrawStake) Validate() error {
	if inst.Amount == nil {
		return errors.New("amount is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.userTransferAuthority should be a signer")
	}
	return nil
}
