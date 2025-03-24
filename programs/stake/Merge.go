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

package stake

import (
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// Merge instruction
type Merge struct {
	// [0] = [WRITE] DestinationStakeAccount
	// ··········· The destination stake account.
	//
	// [1] = [WRITE] SourceStakeAccount
	// ··········· The source stake account.
	//
	// [2] = [] ClockSysvar
	// ··········· Clock sysvar account.
	//
	// [3] = [] StakeHistorySysvar
	// ··········· Stake history sysvar account.
	//
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *Merge) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return nil
}

func (inst Merge) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.AccountMetaSlice...)
	return
}

// NewMergeInstructionBuilder creates a new `Merge` instruction builder.
func NewMergeInstructionBuilder() *Merge {
	nd := &Merge{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetDestinationStakeAccount sets the "destination" account.
// The destination stake account.
func (inst *Merge) SetDestinationStakeAccount(destination ag_solanago.PublicKey) *Merge {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationStakeAccount gets the "destination" account.
// The destination stake account.
func (inst *Merge) GetDestinationStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetSourceStakeAccount sets the "source" account.
// The source stake account.
func (inst *Merge) SetSourceStakeAccount(source ag_solanago.PublicKey) *Merge {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceStakeAccount gets the "source" account.
// The source stake account.
func (inst *Merge) GetSourceStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetClockSysvarAccount sets the "clock" account.
// The clock sysvar account.
func (inst *Merge) SetClockSysvarAccount(clock ag_solanago.PublicKey) *Merge {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(clock)
	return inst
}

// GetClockSysvarAccount gets the "clock" account.
// The clock sysvar account.
func (inst *Merge) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetStakeHistorySysvarAccount sets the "stake history" account.
// The stake history sysvar account.
func (inst *Merge) SetStakeHistorySysvarAccount(stakeHistory ag_solanago.PublicKey) *Merge {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeHistory)
	return inst
}

// GetStakeHistorySysvarAccount gets the "stake history" account.
// The stake history sysvar account.
func (inst *Merge) GetStakeHistorySysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst Merge) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Merge, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Merge) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Merge) Validate() error {
	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

func (inst *Merge) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Merge")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("DestinationStakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     SourceStakeAccount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            ClockSysvar", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     StakeHistorySysvar", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (inst Merge) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (inst *Merge) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewMergeInstruction declares a new Merge instruction with the provided accounts.
func NewMergeInstruction(
	// Accounts:
	destination ag_solanago.PublicKey,
	source ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
) *Merge {
	return NewMergeInstructionBuilder().
		SetDestinationStakeAccount(destination).
		SetSourceStakeAccount(source).
		SetClockSysvarAccount(clock).
		SetStakeHistorySysvarAccount(stakeHistory)
}
