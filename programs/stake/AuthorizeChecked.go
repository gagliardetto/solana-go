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

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

const MAX_SIGNERS = 11

type AuthorizeChecked struct {
	// The new authority to be set.
	NewAuthority *ag_solanago.PublicKey

	// The type of authority to be set.
	AuthorityType *AuthorityType

	// [0] = [WRITE] StakeAccount
	// ··········· The stake account to be authorized.
	//
	// [1] = [] ClockSysvar
	// ··········· Clock sysvar account.
	//
	// [2] = [] OldAuthority
	// ··········· The old authority account.
	//
	// [3] = [SIGNER] NewAuthority
	// ··········· The new authority account.
	//
	// [4...] = [SIGNER] Signers
	// ··········· M signer accounts.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *AuthorizeChecked) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts, inst.Signers = ag_solanago.AccountMetaSlice(accounts).SplitFrom(4)
	return nil
}

func (inst AuthorizeChecked) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	accounts = append(accounts, inst.Signers...)
	return
}

// NewAuthorizeCheckedInstructionBuilder creates a new `AuthorizeChecked` instruction builder.
func NewAuthorizeCheckedInstructionBuilder() *AuthorizeChecked {
	nd := &AuthorizeChecked{
		Accounts: make(ag_solanago.AccountMetaSlice, 4),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

// SetNewAuthority sets the "new authority" parameter.
func (inst *AuthorizeChecked) SetNewAuthority(newAuthority ag_solanago.PublicKey) *AuthorizeChecked {
	inst.NewAuthority = &newAuthority
	return inst
}

// SetAuthorityType sets the "authority type" parameter.
func (inst *AuthorizeChecked) SetAuthorityType(authorityType AuthorityType) *AuthorizeChecked {
	inst.AuthorityType = &authorityType
	return inst
}

// SetStakeAccount sets the "stake account" account.
func (inst *AuthorizeChecked) SetStakeAccount(stakeAccount ag_solanago.PublicKey) *AuthorizeChecked {
	inst.Accounts[0] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// GetStakeAccount gets the "stake account" account.
func (inst *AuthorizeChecked) GetStakeAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[0]
}

// SetClockSysvarAccount sets the "clock sysvar" account.
func (inst *AuthorizeChecked) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *AuthorizeChecked {
	inst.Accounts[1] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clock sysvar" account.
func (inst *AuthorizeChecked) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[1]
}

// SetOldAuthorityAccount sets the "old authority" account.
func (inst *AuthorizeChecked) SetOldAuthorityAccount(oldAuthority ag_solanago.PublicKey) *AuthorizeChecked {
	inst.Accounts[2] = ag_solanago.Meta(oldAuthority)
	return inst
}

// GetOldAuthorityAccount gets the "old authority" account.
func (inst *AuthorizeChecked) GetOldAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[2]
}

// SetNewAuthorityAccount sets the "new authority" account.
func (inst *AuthorizeChecked) SetNewAuthorityAccount(newAuthority ag_solanago.PublicKey, multisigSigners ...ag_solanago.PublicKey) *AuthorizeChecked {
	inst.Accounts[3] = ag_solanago.Meta(newAuthority).SIGNER()
	for _, signer := range multisigSigners {
		inst.Signers = append(inst.Signers, ag_solanago.Meta(signer).SIGNER())
	}
	return inst
}

// GetNewAuthorityAccount gets the "new authority" account.
func (inst *AuthorizeChecked) GetNewAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.Accounts[3]
}

func (inst AuthorizeChecked) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_AuthorizeChecked, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AuthorizeChecked) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AuthorizeChecked) Validate() error {
	if inst.NewAuthority == nil {
		return errors.New("NewAuthority parameter is not set")
	}
	if inst.AuthorityType == nil {
		return errors.New("AuthorityType parameter is not set")
	}

	if inst.Accounts[0] == nil {
		return errors.New("accounts.StakeAccount is not set")
	}
	if inst.Accounts[1] == nil {
		return errors.New("accounts.ClockSysvar is not set")
	}
	if inst.Accounts[2] == nil {
		return errors.New("accounts.OldAuthority is not set")
	}
	if inst.Accounts[3] == nil {
		return errors.New("accounts.NewAuthority is not set")
	}
	if len(inst.Signers) > MAX_SIGNERS {
		return fmt.Errorf("too many signers; got %v, but max is 11", len(inst.Signers))
	}
	return nil
}

func (inst *AuthorizeChecked) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AuthorizeChecked")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewAuthority", *inst.NewAuthority))
						paramsBranch.Child(ag_format.Param("AuthorityType", *inst.AuthorityType))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("StakeAccount", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("ClockSysvar", inst.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("OldAuthority", inst.Accounts[2]))
						accountsBranch.Child(ag_format.Meta("NewAuthority", inst.Accounts[3]))

						signersBranch := accountsBranch.Child(fmt.Sprintf("Signers[len=%v]", len(inst.Signers)))
						for i, v := range inst.Signers {
							if len(inst.Signers) > 9 && i < 10 {
								signersBranch.Child(ag_format.Meta(fmt.Sprintf(" [%v]", i), v))
							} else {
								signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), v))
							}
						}
					})
				})
		})
}

func (inst AuthorizeChecked) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAuthority` param:
	err = encoder.Encode(inst.NewAuthority)
	if err != nil {
		return err
	}
	// Serialize `AuthorityType` param:
	err = encoder.Encode(inst.AuthorityType)
	if err != nil {
		return err
	}
	return nil
}

func (inst *AuthorizeChecked) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAuthority`:
	err = decoder.Decode(&inst.NewAuthority)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityType`:
	err = decoder.Decode(&inst.AuthorityType)
	if err != nil {
		return err
	}
	return nil
}

// NewAuthorizeCheckedInstruction declares a new AuthorizeChecked instruction with the provided parameters and accounts.
func NewAuthorizeCheckedInstruction(
	// Parameters:
	newAuthority ag_solanago.PublicKey,
	authorityType AuthorityType,
	// Accounts:
	stakeAccount ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	oldAuthority ag_solanago.PublicKey,
	newAuthorityAccount ag_solanago.PublicKey,
	multisigSigners []ag_solanago.PublicKey,
) *AuthorizeChecked {
	return NewAuthorizeCheckedInstructionBuilder().
		SetNewAuthority(newAuthority).
		SetAuthorityType(authorityType).
		SetStakeAccount(stakeAccount).
		SetClockSysvarAccount(clockSysvar).
		SetOldAuthorityAccount(oldAuthority).
		SetNewAuthorityAccount(newAuthorityAccount, multisigSigners...)
}
