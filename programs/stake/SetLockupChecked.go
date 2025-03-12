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

type SetLockupChecked struct {
	// Lockup settings for stake account
	Lockup *Lockup

	// [0] = [WRITE] StakeAccount
	// ··········· Stake account to set lockup
	//
	// [1] = [] ClockSysvar
	// ··········· ClockSysvar account
	//
	// [2] = [] Custodian
	// ··········· Optional: Lockup custodian
	//
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *SetLockupChecked) UnmarshalWithDecoder(dec *bin.Decoder) error {
	err := dec.Decode(&inst.Lockup)
	if err != nil {
		return err
	}

	return nil
}

func (inst *SetLockupChecked) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.Encode(*inst.Lockup)
	if err != nil {
		return err
	}

	return nil
}

func (inst *SetLockupChecked) Validate() error {
	// Check whether all (required) parameters are set:
	if inst.Lockup == nil {
		return errors.New("lockup parameter is not set")
	}
	err := inst.Lockup.Validate()
	if err != nil {
		return err
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
func (inst *SetLockupChecked) SetStakeAccount(stakeAccount solana.PublicKey) *SetLockupChecked {
	inst.AccountMetaSlice[0] = solana.Meta(stakeAccount).WRITE()
	return inst
}

// Clock sysvar account
func (inst *SetLockupChecked) SetClockSysvarAccount(clockSysvar solana.PublicKey) *SetLockupChecked {
	inst.AccountMetaSlice[1] = solana.Meta(clockSysvar)
	return inst
}

// Optional: Lockup custodian
func (inst *SetLockupChecked) SetCustodian(custodian solana.PublicKey) *SetLockupChecked {
	inst.AccountMetaSlice[2] = solana.Meta(custodian)
	return inst
}

// Lockup settings for stake account
func (inst *SetLockupChecked) SetLockup(lockup Lockup) *SetLockupChecked {
	inst.Lockup = &lockup
	return inst
}

func (inst *SetLockupChecked) GetStakeAccount() *solana.AccountMeta { return inst.AccountMetaSlice[0] }
func (inst *SetLockupChecked) GetClockSysvarAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[1]
}
func (inst *SetLockupChecked) GetCustodian() *solana.AccountMeta { return inst.AccountMetaSlice[2] }

func (inst SetLockupChecked) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_SetLockupChecked, bin.LE),
	}}
}

func (inst *SetLockupChecked) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("SetLockupChecked")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child("Lockup").ParentFunc(func(authBranch treeout.Branches) {
							authBranch.Child(format.Param("UnixTimestamp", inst.Lockup.UnixTimestamp))
							authBranch.Child(format.Param("        Epoch", inst.Lockup.Epoch))
							authBranch.Child(format.Account("    Custodian", *inst.Lockup.Custodian))
						})
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("ClockSysvar", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(format.Meta("Custodian", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

// NewSetLockupCheckedInstructionBuilder creates a new `SetLockupChecked` instruction builder.
func NewSetLockupCheckedInstructionBuilder() *SetLockupChecked {
	nd := &SetLockupChecked{
		AccountMetaSlice: make(solana.AccountMetaSlice, 3),
		Lockup:           &Lockup{},
	}
	return nd
}

// NewSetLockupCheckedInstruction declares a new SetLockupChecked instruction with the provided parameters and accounts.
func NewSetLockupCheckedInstruction(
	// parameters:
	lockup Lockup,
	// Accounts:
	stakeAccount solana.PublicKey,
	clockSysvar solana.PublicKey,
	custodian solana.PublicKey,
) *SetLockupChecked {
	return NewSetLockupCheckedInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetClockSysvarAccount(clockSysvar).
		SetCustodian(custodian).
		SetLockup(lockup)
}
