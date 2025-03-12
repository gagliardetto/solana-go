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

type SetLockup struct {
	// Lockup settings for stake account
	Lockup *Lockup

	// [0] = [WRITE] StakeAccount
	// ··········· Stake account to set lockup for
	//
	// [1] = [] ClockSysvar
	// ··········· ClockSysvar account
	//
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *SetLockup) UnmarshalWithDecoder(dec *bin.Decoder) error {
	return dec.Decode(&inst.Lockup)
}

func (inst *SetLockup) MarshalWithEncoder(encoder *bin.Encoder) error {
	return encoder.Encode(*inst.Lockup)
}

func (inst *SetLockup) Validate() error {
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

func (inst *SetLockup) SetStakeAccount(stakeAccount solana.PublicKey) *SetLockup {
	inst.AccountMetaSlice[0] = solana.Meta(stakeAccount).WRITE()
	return inst
}

func (inst *SetLockup) SetClockSysvarAccount(clockSysvar solana.PublicKey) *SetLockup {
	inst.AccountMetaSlice[1] = solana.Meta(clockSysvar)
	return inst
}

func (inst *SetLockup) GetStakeAccount() *solana.AccountMeta       { return inst.AccountMetaSlice[0] }
func (inst *SetLockup) GetClockSysvarAccount() *solana.AccountMeta { return inst.AccountMetaSlice[1] }

func (inst *SetLockup) SetLockupTimestamp(unixTimestamp int64) *SetLockup {
	inst.Lockup.UnixTimestamp = &unixTimestamp
	return inst
}

func (inst *SetLockup) SetLockupEpoch(epoch uint64) *SetLockup {
	inst.Lockup.Epoch = &epoch
	return inst
}

func (inst *SetLockup) SetCustodian(custodian solana.PublicKey) *SetLockup {
	inst.Lockup.Custodian = &custodian
	return inst
}

func (inst SetLockup) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_SetLockup, bin.LE),
	}}
}

func (inst *SetLockup) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("SetLockup")).
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
					})
				})
		})
}

// NewSetLockupInstructionBuilder creates a new `SetLockup` instruction builder.
func NewSetLockupInstructionBuilder() *SetLockup {
	nd := &SetLockup{
		AccountMetaSlice: make(solana.AccountMetaSlice, 2),
		Lockup:           &Lockup{},
	}
	return nd
}

// NewSetLockupInstruction declares a new SetLockup instruction with the provided parameters and accounts.
func NewSetLockupInstruction(
	// parameters:
	unixTimestamp int64,
	epoch uint64,
	custodian solana.PublicKey,
	// Accounts:
	stakeAccount solana.PublicKey,
) *SetLockup {
	return NewSetLockupInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetClockSysvarAccount(solana.SysVarClockPubkey).
		SetLockupTimestamp(unixTimestamp).
		SetLockupEpoch(epoch).
		SetCustodian(custodian)
}
