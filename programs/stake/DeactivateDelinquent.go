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
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// DeactivateDelinquent instruction
type DeactivateDelinquent struct {
	// [0] = [WRITE] StakeAccount
	// ··········· The stake account to deactivate.
	//
	// [1] = [] DelinquentVoteAccount
	// ··········· The delinquent vote account.
	//
	// [2] = [] ReferenceVoteAccount
	// ··········· The reference vote account.
	//
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *DeactivateDelinquent) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return nil
}

func (inst DeactivateDelinquent) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.AccountMetaSlice...)
	return
}

// NewDeactivateDelinquentInstructionBuilder creates a new `DeactivateDelinquent` instruction builder.
func NewDeactivateDelinquentInstructionBuilder() *DeactivateDelinquent {
	nd := &DeactivateDelinquent{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStakeAccount sets the "StakeAccount" account.
// The stake account to deactivate.
func (inst *DeactivateDelinquent) SetStakeAccount(stakeAccount ag_solanago.PublicKey) *DeactivateDelinquent {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccount gets the "StakeAccount" account.
// The stake account to deactivate.
func (inst *DeactivateDelinquent) GetStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDelinquentVoteAccount sets the "DelinquentVoteAccount" account.
// The delinquent vote account.
func (inst *DeactivateDelinquent) SetDelinquentVoteAccount(delinquentVoteAccount ag_solanago.PublicKey) *DeactivateDelinquent {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(delinquentVoteAccount)
	return inst
}

// GetDelinquentVoteAccount gets the "DelinquentVoteAccount" account.
// The delinquent vote account.
func (inst *DeactivateDelinquent) GetDelinquentVoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetReferenceVoteAccount sets the "ReferenceVoteAccount" account.
// The reference vote account.
func (inst *DeactivateDelinquent) SetReferenceVoteAccount(referenceVoteAccount ag_solanago.PublicKey) *DeactivateDelinquent {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(referenceVoteAccount)
	return inst
}

// GetReferenceVoteAccount gets the "ReferenceVoteAccount" account.
// The reference vote account.
func (inst *DeactivateDelinquent) GetReferenceVoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst DeactivateDelinquent) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_DeactivateDelinquent, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeactivateDelinquent) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeactivateDelinquent) Validate() error {
	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

func (inst *DeactivateDelinquent) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeactivateDelinquent")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("DelinquentVoteAccount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("ReferenceVoteAccount", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (inst DeactivateDelinquent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (inst *DeactivateDelinquent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeactivateDelinquentInstruction declares a new DeactivateDelinquent instruction with the provided accounts.
func NewDeactivateDelinquentInstruction(
	// Accounts:
	stakeAccount ag_solanago.PublicKey,
	delinquentVoteAccount ag_solanago.PublicKey,
	referenceVoteAccount ag_solanago.PublicKey,
) *DeactivateDelinquent {
	return NewDeactivateDelinquentInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetDelinquentVoteAccount(delinquentVoteAccount).
		SetReferenceVoteAccount(referenceVoteAccount)
}
