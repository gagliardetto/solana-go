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

type Initialize struct {
	Fee           *Fee
	WithdrawalFee *Fee
	DepositFee    *Fee
	ReferralFee   *uint8
	MaxValidators *uint32
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] manager
	// [2] = [] staker
	// [3] = [] withdrawAuthority
	// [4] = [WRITE] validatorList
	// [5] = [] reserveStake
	// [6] = [WRITE] poolMint
	// [7] = [WRITE] managerFeeAccount
	// [8] = [] tokenProgram
	// [9] = [] depositAuthority
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (i *Initialize) SetAccounts(accounts []*ag_solanago.AccountMeta) *Initialize {
	i.Accounts = accounts
	return i
}

func (i *Initialize) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, i.Accounts...)
	accounts = append(accounts, i.Signers...)
	return
}

func NewInitializeInstructionBuilder() *Initialize {
	return &Initialize{
		Accounts: make(ag_solanago.AccountMetaSlice, 10),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (i *Initialize) SetSigners(signers []*ag_solanago.AccountMeta) *Initialize {
	i.Signers = signers
	return i
}

func (i *Initialize) SetStakePoolAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[0] = ag_solanago.Meta(account).WRITE()
	return i
}

func (i *Initialize) GetStakePoolAccount() *ag_solanago.AccountMeta {
	return i.Accounts[0]
}

func (i *Initialize) SetManagerAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[1] = ag_solanago.Meta(account).SIGNER()
	i.Signers[0] = ag_solanago.Meta(account).SIGNER()
	return i
}

func (i *Initialize) GetManagerAccount() *ag_solanago.AccountMeta {
	return i.Accounts[1]
}

func (i *Initialize) SetStakerAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[2] = ag_solanago.Meta(account)
	return i
}

func (i *Initialize) GetStakerAccount() *ag_solanago.AccountMeta {
	return i.Accounts[2]
}

func (i *Initialize) SetWithdrawAuthorityAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[3] = ag_solanago.Meta(account)
	return i
}

func (i *Initialize) GetWithdrawAuthorityAccount() *ag_solanago.AccountMeta {
	return i.Accounts[3]
}

func (i *Initialize) SetValidatorListAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[4] = ag_solanago.Meta(account).WRITE()
	return i
}

func (i *Initialize) GetValidatorListAccount() *ag_solanago.AccountMeta {
	return i.Accounts[4]
}

func (i *Initialize) SetReserveStakeAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[5] = ag_solanago.Meta(account)
	return i
}

func (i *Initialize) GetReserveStakeAccount() *ag_solanago.AccountMeta {
	return i.Accounts[5]
}

func (i *Initialize) SetPoolMintAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[6] = ag_solanago.Meta(account).WRITE()
	return i
}

func (i *Initialize) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return i.Accounts[6]
}

func (i *Initialize) SetManagerFeeAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[7] = ag_solanago.Meta(account).WRITE()
	return i
}

func (i *Initialize) GetManagerFeeAccount() *ag_solanago.AccountMeta {
	return i.Accounts[7]
}

func (i *Initialize) SetTokenProgramAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[8] = ag_solanago.Meta(account)
	return i
}

func (i *Initialize) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return i.Accounts[8]
}

func (i *Initialize) SetDepositAuthorityAccount(account ag_solanago.PublicKey) *Initialize {
	i.Accounts[9] = ag_solanago.Meta(account)
	return i
}

func (i *Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}
	return i.Build(), nil
}

func (i *Initialize) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_Initialize),
			Impl:   i,
		},
	}
}

func (i *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Fee", *i.Fee))
						paramsBranch.Child(ag_format.Param("WithdrawalFee", *i.WithdrawalFee))
						paramsBranch.Child(ag_format.Param("DepositFee", *i.DepositFee))
						paramsBranch.Child(ag_format.Param("ReferralFee", *i.ReferralFee))
						paramsBranch.Child(ag_format.Param("MaxValidators", *i.MaxValidators))
					})

					// Accounts of the instructions:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("stakePool", i.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("manager", i.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("staker", i.Accounts[2]))
						accountsBranch.Child(ag_format.Meta("withdrawAuthority", i.Accounts[3]))
						accountsBranch.Child(ag_format.Meta("validatorList", i.Accounts[4]))
						accountsBranch.Child(ag_format.Meta("reserveStake", i.Accounts[5]))
						accountsBranch.Child(ag_format.Meta("poolMint", i.Accounts[6]))
						accountsBranch.Child(ag_format.Meta("managerFeeAccount", i.Accounts[7]))
						accountsBranch.Child(ag_format.Meta("tokenProgram", i.Accounts[8]))
						accountsBranch.Child(ag_format.Meta("depositAuthority", i.Accounts[9]))

						// Signers of the instructions:
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(i.Signers)))
						for j, signer := range i.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (i *Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if err := encoder.Encode(i.Fee); err != nil {
		return err
	}
	if err := encoder.Encode(i.WithdrawalFee); err != nil {
		return err
	}
	if err := encoder.Encode(i.DepositFee); err != nil {
		return err
	}
	if err := encoder.Encode(i.ReferralFee); err != nil {
		return err
	}
	if err := encoder.Encode(i.MaxValidators); err != nil {
		return err
	}
	for _, account := range i.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (i *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if err := decoder.Decode(&i.Fee); err != nil {
		return err
	}
	if err := decoder.Decode(&i.WithdrawalFee); err != nil {
		return err
	}
	if err := decoder.Decode(&i.DepositFee); err != nil {
		return err
	}
	if err := decoder.Decode(&i.ReferralFee); err != nil {
		return err
	}
	if err := decoder.Decode(&i.MaxValidators); err != nil {
		return err
	}
	for j := range i.Accounts {
		if err := decoder.Decode(&i.Accounts[j]); err != nil {
			return err
		}
	}
	return nil
}

func (i *Initialize) Validate() error {
	if i.Fee == nil {
		return errors.New("fee is not set")
	}
	if i.WithdrawalFee == nil {
		return errors.New("withdrawalFee is not set")
	}
	if i.DepositFee == nil {
		return errors.New("depositFee is not set")
	}
	if i.ReferralFee == nil {
		return errors.New("referralFee is not set")
	}
	if i.MaxValidators == nil {
		return errors.New("maxValidators is not set")
	}
	for j, account := range i.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", j)
		}
	}
	if len(i.Signers) == 0 || !i.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}

func (i *Initialize) SetFee(fee Fee) *Initialize {
	i.Fee = &fee
	return i
}

func (i *Initialize) SetWithdrawalFee(withdrawalFee Fee) *Initialize {
	i.WithdrawalFee = &withdrawalFee
	return i
}

func (i *Initialize) SetDepositFee(depositFee Fee) *Initialize {
	i.DepositFee = &depositFee
	return i
}

func (i *Initialize) SetReferralFee(referralFee uint8) *Initialize {
	i.ReferralFee = &referralFee
	return i
}

func (i *Initialize) SetMaxValidators(maxValidators uint32) *Initialize {
	i.MaxValidators = &maxValidators
	return i
}

func NewInitializeInstruction(
	// Parameters:
	fee Fee,
	withdrawalFee Fee,
	depositFee Fee,
	referralFee uint8,
	maxValidators uint32,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	manager ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	depositAuthority ag_solanago.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetFee(fee).
		SetWithdrawalFee(withdrawalFee).
		SetDepositFee(depositFee).
		SetReferralFee(referralFee).
		SetMaxValidators(maxValidators).
		SetStakePoolAccount(stakePool).
		SetManagerAccount(manager).
		SetStakerAccount(staker).
		SetWithdrawAuthorityAccount(withdrawAuthority).
		SetValidatorListAccount(validatorList).
		SetReserveStakeAccount(reserveStake).
		SetPoolMintAccount(poolMint).
		SetManagerFeeAccount(managerFeeAccount).
		SetTokenProgramAccount(tokenProgram).
		SetDepositAuthorityAccount(depositAuthority)
}
