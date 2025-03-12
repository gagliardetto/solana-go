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

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// MoveLamports instruction
type MoveLamports struct {
	// The amount of lamports to move.
	Amount *uint64

	// [0] = [WRITE] SourceStakeAccount
	// ··········· The source stake account.
	//
	// [1] = [WRITE] DestinationStakeAccount
	// ··········· The destination stake account.
	//
	// [2] = [SIGNER] StakeAuthority
	// ··········· The stake authority.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *MoveLamports) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	if len(accounts) != 3 {
		return errors.New("incorrect number of accounts")
	}

	inst.Accounts = ag_solanago.AccountMetaSlice(accounts)
	return nil
}

func (inst MoveLamports) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	return inst.Accounts
}

// NewMoveLamportsInstructionBuilder creates a new `MoveLamports` instruction builder.
func NewMoveLamportsInstructionBuilder() *MoveLamports {
	nd := &MoveLamports{
		Accounts: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of lamports to move.
func (inst *MoveLamports) SetAmount(amount uint64) *MoveLamports {
	inst.Amount = &amount
	return inst
}

// SetSourceStakeAccount sets the "source" account.
// The source stake account.
func (inst *MoveLamports) SetSourceStakeAccount(source ag_solanago.PublicKey) *MoveLamports {
	inst.Accounts[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceStakeAccount gets the "source" account.
// The source stake account.
func (inst *MoveLamports) GetSourceStakeAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

// SetDestinationStakeAccount sets the "destination" account.
// The destination stake account.
func (inst *MoveLamports) SetDestinationStakeAccount(destination ag_solanago.PublicKey) *MoveLamports {
	inst.Accounts[1] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationStakeAccount gets the "destination" account.
// The destination stake account.
func (inst *MoveLamports) GetDestinationStakeAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[1]
}

// SetStakeAuthorityAccount sets the "stake authority" account.
// The stake authority.
func (inst *MoveLamports) SetStakeAuthorityAccount(stakeAuthority ag_solanago.PublicKey) *MoveLamports {
	inst.Accounts[2] = ag_solanago.Meta(stakeAuthority).SIGNER()
	return inst
}

// GetStakeAuthorityAccount gets the "stake authority" account.
// The stake authority.
func (inst *MoveLamports) GetStakeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[2]
}

func (inst MoveLamports) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_MoveLamports, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MoveLamports) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MoveLamports) Validate() error {
	if inst.Amount == nil {
		return errors.New("Amount parameter is not set")
	}
	if inst.Accounts[0] == nil {
		return errors.New("accounts.SourceStakeAccount is not set")
	}
	if inst.Accounts[1] == nil {
		return errors.New("accounts.DestinationStakeAccount is not set")
	}
	if inst.Accounts[2] == nil {
		return errors.New("accounts.StakeAuthority is not set")
	}
	return nil
}

func (inst *MoveLamports) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MoveLamports")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("SourceStakeAccount", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("DestinationStakeAccount", inst.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("StakeAuthority", inst.Accounts[2]))
					})
				})
		})
}

func (inst MoveLamports) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(inst.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (inst *MoveLamports) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&inst.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewMoveLamportsInstruction declares a new MoveLamports instruction with the provided parameters and accounts.
func NewMoveLamportsInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	source ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	stakeAuthority ag_solanago.PublicKey,
) *MoveLamports {
	return NewMoveLamportsInstructionBuilder().
		SetAmount(amount).
		SetSourceStakeAccount(source).
		SetDestinationStakeAccount(destination).
		SetStakeAuthorityAccount(stakeAuthority)
}
