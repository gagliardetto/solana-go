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

func (w *WithdrawStake) SetAmount(amount uint64) *WithdrawStake {
	w.Amount = &amount
	return w
}

func (w *WithdrawStake) SetStakePool(stakePool ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return w
}

func (w *WithdrawStake) SetValidatorStakeList(validatorStakeList ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[1] = ag_solanago.Meta(validatorStakeList).WRITE()
	return w
}

func (w *WithdrawStake) SetStakePoolWithdrawAuthority(stakePoolWithdrawAuthority ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[2] = ag_solanago.Meta(stakePoolWithdrawAuthority)
	return w
}

func (w *WithdrawStake) SetValidatorAccount(validatorAccount ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[3] = ag_solanago.Meta(validatorAccount).WRITE()
	return w
}

func (w *WithdrawStake) SetUninitializedStakeAccount(uninitializedStakeAccount ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[4] = ag_solanago.Meta(uninitializedStakeAccount).WRITE()
	return w
}

func (w *WithdrawStake) SetUserAccount(userAccount ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[5] = ag_solanago.Meta(userAccount)
	return w
}

func (w *WithdrawStake) SetUserTransferAuthority(userTransferAuthority ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[6] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	w.Signers[0] = ag_solanago.Meta(userTransferAuthority).SIGNER()
	return w
}

func (w *WithdrawStake) SetUserAccountWithPoolTokensToBurnFrom(userAccountWithPoolTokensToBurnFrom ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[7] = ag_solanago.Meta(userAccountWithPoolTokensToBurnFrom).WRITE()
	return w
}

func (w *WithdrawStake) SetAccountToReceivePoolFeeTokens(accountToReceivePoolFeeTokens ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[8] = ag_solanago.Meta(accountToReceivePoolFeeTokens).WRITE()
	return w
}

func (w *WithdrawStake) SetPoolTokenMintAccount(poolTokenMintAccount ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[9] = ag_solanago.Meta(poolTokenMintAccount).WRITE()
	return w
}

func (w *WithdrawStake) SetSysvarClockAccount(sysvarClockAccount ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[10] = ag_solanago.Meta(sysvarClockAccount)
	return w
}

func (w *WithdrawStake) SetPoolTokenProgramID(poolTokenProgramID ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[11] = ag_solanago.Meta(poolTokenProgramID)
	return w
}

func (w *WithdrawStake) SetStakeProgramID(stakeProgramID ag_solanago.PublicKey) *WithdrawStake {
	w.Accounts[12] = ag_solanago.Meta(stakeProgramID)
	return w
}

func (w *WithdrawStake) GetAmount() *uint64 {
	return w.Amount
}

func (w *WithdrawStake) GetStakePool() ag_solanago.PublicKey {
	return w.Accounts[0].PublicKey
}

func (w *WithdrawStake) GetValidatorStakeList() ag_solanago.PublicKey {
	return w.Accounts[1].PublicKey
}

func (w *WithdrawStake) GetStakePoolWithdrawAuthority() ag_solanago.PublicKey {
	return w.Accounts[2].PublicKey
}

func (w *WithdrawStake) GetValidatorAccount() ag_solanago.PublicKey {
	return w.Accounts[3].PublicKey
}

func (w *WithdrawStake) GetUninitializedStakeAccount() ag_solanago.PublicKey {
	return w.Accounts[4].PublicKey
}

func (w *WithdrawStake) GetUserAccount() ag_solanago.PublicKey {
	return w.Accounts[5].PublicKey
}

func (w *WithdrawStake) GetUserTransferAuthority() ag_solanago.PublicKey {
	return w.Accounts[6].PublicKey
}

func (w *WithdrawStake) GetUserAccountWithPoolTokensToBurnFrom() ag_solanago.PublicKey {
	return w.Accounts[7].PublicKey
}

func (w *WithdrawStake) GetAccountToReceivePoolFeeTokens() ag_solanago.PublicKey {
	return w.Accounts[8].PublicKey
}

func (w *WithdrawStake) GetPoolTokenMintAccount() ag_solanago.PublicKey {
	return w.Accounts[9].PublicKey
}

func (w *WithdrawStake) GetSysvarClockAccount() ag_solanago.PublicKey {
	return w.Accounts[10].PublicKey
}

func (w *WithdrawStake) GetPoolTokenProgramID() ag_solanago.PublicKey {
	return w.Accounts[11].PublicKey
}

func (w *WithdrawStake) GetStakeProgramID() ag_solanago.PublicKey {
	return w.Accounts[12].PublicKey
}

func (w *WithdrawStake) ValidateAndBuild() (*Instruction, error) {
	if err := w.Validate(); err != nil {
		return nil, err
	}
	return w.Build(), nil
}

func (w *WithdrawStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawStake),
			Impl:   w,
		},
	}
}

func (w *WithdrawStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if w.Amount != nil {
							paramsBranch.Child(ag_format.Param("Amount", *w.Amount))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range w.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(w.Signers)))
						for j, signer := range w.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (w *WithdrawStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if w.Amount != nil {
		if err := encoder.Encode(w.Amount); err != nil {
			return err
		}
	}
	for _, account := range w.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (w *WithdrawStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if w.Amount != nil {
		if err := decoder.Decode(w.Amount); err != nil {
			return err
		}
	}
	for i := range w.Accounts {
		if err := decoder.Decode(w.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (w *WithdrawStake) Validate() error {
	if w.Amount == nil {
		return errors.New("amount is not set")
	}
	for i, account := range w.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(w.Signers) == 0 || !w.Signers[0].IsSigner {
		return errors.New("accounts.userTransferAuthority should be a signer")
	}
	return nil
}
