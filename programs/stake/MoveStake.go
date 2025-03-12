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
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MoveStake moves stake from one account to another.
type MoveStake struct {
	// The amount of lamports to move.
	Lamports *uint64

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

func (inst *MoveStake) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts = ag_solanago.AccountMetaSlice(accounts)
	return nil
}

// NewMoveStakeInstructionBuilder creates a new `MoveStake` instruction builder.
func NewMoveStakeInstructionBuilder() *MoveStake {
	nd := &MoveStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetLamports sets the "lamports" parameter.
// The amount of lamports to move.
func (inst *MoveStake) SetLamports(lamports uint64) *MoveStake {
	inst.Lamports = &lamports
	return inst
}

// SetSourceStakeAccount sets the "source" account.
// The source stake account.
func (inst *MoveStake) SetSourceStakeAccount(source ag_solanago.PublicKey) *MoveStake {
	inst.Accounts[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceStakeAccount gets the "source" account.
// The source stake account.
func (inst *MoveStake) GetSourceStakeAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

// SetDestinationStakeAccount sets the "destination" account.
// The destination stake account.
func (inst *MoveStake) SetDestinationStakeAccount(destination ag_solanago.PublicKey) *MoveStake {
	inst.Accounts[1] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationStakeAccount gets the "destination" account.
// The destination stake account.
func (inst *MoveStake) GetDestinationStakeAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[1]
}

// SetStakeAuthority sets the "stake authority" account.
// The stake authority.
func (inst *MoveStake) SetStakeAuthority(stakeAuthority ag_solanago.PublicKey) *MoveStake {
	inst.Accounts[2] = ag_solanago.Meta(stakeAuthority).SIGNER()
	return inst
}

// GetStakeAuthority gets the "stake authority" account.
// The stake authority.
func (inst *MoveStake) GetStakeAuthority() *ag_solanago.AccountMeta {
	return inst.Accounts[2]
}

func (inst MoveStake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_MoveStake, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MoveStake) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MoveStake) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Lamports == nil {
			return errors.New("Lamports parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.Accounts[0] == nil {
			return errors.New("accounts.SourceStakeAccount is not set")
		}
		if inst.Accounts[1] == nil {
			return errors.New("accounts.DestinationStakeAccount is not set")
		}
		if inst.Accounts[2] == nil {
			return errors.New("accounts.StakeAuthority is not set")
		}
	}
	return nil
}

func (inst *MoveStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MoveStake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Lamports", *inst.Lamports))
					})

					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("SourceStakeAccount", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("DestinationStakeAccount", inst.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("StakeAuthority", inst.Accounts[2]))
					})
				})
		})
}

func (inst MoveStake) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Lamports` param:
	err = encoder.Encode(inst.Lamports)
	if err != nil {
		return err
	}
	return nil
}

func (inst *MoveStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Lamports`:
	err = decoder.Decode(&inst.Lamports)
	if err != nil {
		return err
	}
	return nil
}

// NewMoveStakeInstruction declares a new MoveStake instruction with the provided parameters and accounts.
func NewMoveStakeInstruction(
	// Parameters:
	lamports uint64,
	// Accounts:
	source ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	stakeAuthority ag_solanago.PublicKey,
) *MoveStake {
	return NewMoveStakeInstructionBuilder().
		SetLamports(lamports).
		SetSourceStakeAccount(source).
		SetDestinationStakeAccount(destination).
		SetStakeAuthority(stakeAuthority)
}
