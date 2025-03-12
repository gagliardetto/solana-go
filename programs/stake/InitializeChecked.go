// Copyright 2021 github.com/gagliardetto
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

package stake

import (
	"errors"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/treeout"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
)

type InitializeChecked struct {
	// Authorization settings for stake account
	Authorized *Authorized

	// Lockup settings for stake account
	Lockup *Lockup

	// [0] = [WRITE SIGNER] StakeAccount
	// ··········· Stake account getting initialized
	//
	// [1] = [] RentSysvar
	// ··········· RentSysvar account
	//
	// [2] = [] StakeAuthority
	// ··········· Stake authority account
	//
	// [3] = [] WithdrawAuthority
	// ··········· Withdraw authority account
	//
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *InitializeChecked) UnmarshalWithDecoder(dec *bin.Decoder) error {
	{
		err := dec.Decode(&inst.Authorized)
		if err != nil {
			return err
		}
	}
	{
		err := dec.Decode(&inst.Lockup)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *InitializeChecked) MarshalWithEncoder(encoder *bin.Encoder) error {
	{
		err := encoder.Encode(*inst.Authorized)
		if err != nil {
			return err
		}
	}
	{
		err := encoder.Encode(*inst.Lockup)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *InitializeChecked) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Authorized == nil {
			return errors.New("authorized parameter is not set")
		}
		err := inst.Authorized.Validate()
		if err != nil {
			return err
		}
	}
	{
		if inst.Lockup == nil {
			return errors.New("lockup parameter is not set")
		}
		err := inst.Lockup.Validate()
		if err != nil {
			return err
		}
	}

	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

// Stake account account
func (inst *InitializeChecked) SetStakeAccount(stakeAccount solana.PublicKey) *InitializeChecked {
	inst.AccountMetaSlice[0] = solana.Meta(stakeAccount).WRITE().SIGNER()
	return inst
}

// Rent sysvar account
func (inst *InitializeChecked) SetRentSysvarAccount(rentSysvar solana.PublicKey) *InitializeChecked {
	inst.AccountMetaSlice[1] = solana.Meta(rentSysvar)
	return inst
}

// Stake authority account
func (inst *InitializeChecked) SetStakeAuthorityAccount(stakeAuthority solana.PublicKey) *InitializeChecked {
	inst.AccountMetaSlice[2] = solana.Meta(stakeAuthority)
	return inst
}

// Withdraw authority account
func (inst *InitializeChecked) SetWithdrawAuthorityAccount(withdrawAuthority solana.PublicKey) *InitializeChecked {
	inst.AccountMetaSlice[3] = solana.Meta(withdrawAuthority)
	return inst
}

func (inst *InitializeChecked) GetStakeAccount() *solana.AccountMeta { return inst.AccountMetaSlice[0] }
func (inst *InitializeChecked) GetRentSysvarAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[1]
}
func (inst *InitializeChecked) GetStakeAuthorityAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[2]
}
func (inst *InitializeChecked) GetWithdrawAuthorityAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst *InitializeChecked) SetStaker(staker solana.PublicKey) *InitializeChecked {
	inst.Authorized.Staker = &staker
	return inst
}

func (inst *InitializeChecked) SetWithdrawer(withdrawer solana.PublicKey) *InitializeChecked {
	inst.Authorized.Withdrawer = &withdrawer
	return inst
}

func (inst *InitializeChecked) SetLockupTimestamp(unixTimestamp int64) *InitializeChecked {
	inst.Lockup.UnixTimestamp = &unixTimestamp
	return inst
}

func (inst *InitializeChecked) SetLockupEpoch(epoch uint64) *InitializeChecked {
	inst.Lockup.Epoch = &epoch
	return inst
}

func (inst *InitializeChecked) SetCustodian(custodian solana.PublicKey) *InitializeChecked {
	inst.Lockup.Custodian = &custodian
	return inst
}

func (inst InitializeChecked) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_InitializeChecked, bin.LE),
	}}
}

func (inst *InitializeChecked) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("InitializeChecked")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child("Authorized").ParentFunc(func(authBranch treeout.Branches) {
							authBranch.Child(format.Account("    Staker", *inst.Authorized.Staker))
							authBranch.Child(format.Account("Withdrawer", *inst.Authorized.Withdrawer))
						})
						paramsBranch.Child("Lockup").ParentFunc(func(authBranch treeout.Branches) {
							authBranch.Child(format.Param("UnixTimestamp", inst.Lockup.UnixTimestamp))
							authBranch.Child(format.Param("        Epoch", inst.Lockup.Epoch))
							authBranch.Child(format.Account("    Custodian", *inst.Lockup.Custodian))
						})
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("RentSysvar", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(format.Meta("StakeAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(format.Meta("WithdrawAuthority", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

// NewInitializeCheckedInstructionBuilder creates a new `InitializeChecked` instruction builder.
func NewInitializeCheckedInstructionBuilder() *InitializeChecked {
	nd := &InitializeChecked{
		AccountMetaSlice: make(solana.AccountMetaSlice, 4),
		Authorized:       &Authorized{},
		Lockup:           &Lockup{},
	}
	return nd
}

// NewInitializeCheckedInstruction declares a new InitializeChecked instruction with the provided parameters and accounts.
func NewInitializeCheckedInstruction(
	// parameters:
	staker solana.PublicKey,
	withdrawer solana.PublicKey,
	// Accounts:
	stakeAccount solana.PublicKey,
) *InitializeChecked {
	return NewInitializeCheckedInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetRentSysvarAccount(solana.SysVarRentPubkey).
		SetStakeAuthorityAccount(staker).
		SetWithdrawAuthorityAccount(withdrawer).
		SetStaker(staker).
		SetWithdrawer(withdrawer).
		SetLockupEpoch(0).
		SetLockupTimestamp(0).
		SetCustodian(solana.SystemProgramID)
}
