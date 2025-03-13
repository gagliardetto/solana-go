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
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// AuthorizeCheckedWithSeed instruction
type AuthorizeCheckedWithSeed struct {
	// The new authority's public key
	NewAuthority *ag_solanago.PublicKey

	// The authority type
	AuthorityType *uint32

	// The base public key
	Base *ag_solanago.PublicKey

	// The seed
	Seed *string

	// The owner public key
	Owner *ag_solanago.PublicKey

	// [0] = [WRITE] StakeAccount
	// ··········· The stake account to authorize
	//
	// [1] = [] ClockSysvar
	// ··········· Clock sysvar account
	//
	// [2] = [] AuthorityBase
	// ··········· The base public key account
	//
	// [3...] = [SIGNER] Signers
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *AuthorizeCheckedWithSeed) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts, inst.Signers = ag_solanago.AccountMetaSlice(accounts).SplitFrom(3)
	return nil
}

func (inst AuthorizeCheckedWithSeed) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	accounts = append(accounts, inst.Signers...)
	return
}

// NewAuthorizeCheckedWithSeedInstructionBuilder creates a new `AuthorizeCheckedWithSeed` instruction builder.
func NewAuthorizeCheckedWithSeedInstructionBuilder() *AuthorizeCheckedWithSeed {
	nd := &AuthorizeCheckedWithSeed{
		Accounts: make(ag_solanago.AccountMetaSlice, 3),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

// SetNewAuthority sets the "new authority" parameter.
func (inst *AuthorizeCheckedWithSeed) SetNewAuthority(newAuthority ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.NewAuthority = &newAuthority
	return inst
}

// SetAuthorityType sets the "authority type" parameter.
func (inst *AuthorizeCheckedWithSeed) SetAuthorityType(authorityType uint32) *AuthorizeCheckedWithSeed {
	inst.AuthorityType = &authorityType
	return inst
}

// SetBase sets the "base" parameter.
func (inst *AuthorizeCheckedWithSeed) SetBase(base ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.Base = &base
	return inst
}

// SetSeed sets the "seed" parameter.
func (inst *AuthorizeCheckedWithSeed) SetSeed(seed string) *AuthorizeCheckedWithSeed {
	inst.Seed = &seed
	return inst
}

// SetOwner sets the "owner" parameter.
func (inst *AuthorizeCheckedWithSeed) SetOwner(owner ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.Owner = &owner
	return inst
}

// SetStakeAccount sets the "stake account" account.
func (inst *AuthorizeCheckedWithSeed) SetStakeAccount(stakeAccount ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.Accounts[0] = ag_solanago.Meta(stakeAccount).WRITE()
	return inst
}

// SetClockSysvarAccount sets the "clock sysvar" account.
func (inst *AuthorizeCheckedWithSeed) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.Accounts[1] = ag_solanago.Meta(clockSysvar)
	return inst
}

// SetAuthorityBaseAccount sets the "authority base" account.
func (inst *AuthorizeCheckedWithSeed) SetAuthorityBaseAccount(authorityBase ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	inst.Accounts[2] = ag_solanago.Meta(authorityBase)
	return inst
}

// SetSigners sets the "signers" accounts.
func (inst *AuthorizeCheckedWithSeed) SetSigners(signers ...ag_solanago.PublicKey) *AuthorizeCheckedWithSeed {
	for _, signer := range signers {
		inst.Signers = append(inst.Signers, ag_solanago.Meta(signer).SIGNER())
	}
	return inst
}

func (inst AuthorizeCheckedWithSeed) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_AuthorizeCheckedWithSeed, ag_binary.LE),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AuthorizeCheckedWithSeed) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AuthorizeCheckedWithSeed) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewAuthority == nil {
			return errors.New("NewAuthority parameter is not set")
		}
		if inst.AuthorityType == nil {
			return errors.New("AuthorityType parameter is not set")
		}
		if inst.Base == nil {
			return errors.New("Base parameter is not set")
		}
		if inst.Seed == nil {
			return errors.New("Seed parameter is not set")
		}
		if inst.Owner == nil {
			return errors.New("Owner parameter is not set")
		}
	}

	{
		if inst.Accounts[0] == nil {
			return errors.New("accounts.StakeAccount is not set")
		}
		if inst.Accounts[1] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
		if inst.Accounts[2] == nil {
			return errors.New("accounts.AuthorityBase is not set")
		}
		if len(inst.Signers) == 0 {
			return errors.New("accounts.Signers is not set")
		}
	}
	return nil
}

func (inst *AuthorizeCheckedWithSeed) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AuthorizeCheckedWithSeed")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" NewAuthority", *inst.NewAuthority))
						paramsBranch.Child(ag_format.Param("AuthorityType", *inst.AuthorityType))
						paramsBranch.Child(ag_format.Param("         Base", *inst.Base))
						paramsBranch.Child(ag_format.Param("         Seed", *inst.Seed))
						paramsBranch.Child(ag_format.Param("        Owner", *inst.Owner))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("StakeAccount", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta(" ClockSysvar", inst.Accounts[1]))
						signersBranch := accountsBranch.Child(fmt.Sprintf("Signers[len=%v]", len(inst.Signers)))
						for i, v := range inst.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), v))
						}
					})
				})
		})
}

func (inst AuthorizeCheckedWithSeed) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Base` param:
	err = encoder.Encode(inst.Base)
	if err != nil {
		return err
	}
	// Serialize `Seed` param:
	err = encoder.Encode(inst.Seed)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(inst.Owner)
	if err != nil {
		return err
	}
	return nil
}

func (inst *AuthorizeCheckedWithSeed) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Base`:
	err = decoder.Decode(&inst.Base)
	if err != nil {
		return err
	}
	// Deserialize `Seed`:
	err = decoder.Decode(&inst.Seed)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&inst.Owner)
	if err != nil {
		return err
	}
	return nil
}

// NewAuthorizeCheckedWithSeedInstruction declares a new AuthorizeCheckedWithSeed instruction with the provided parameters and accounts.
func NewAuthorizeCheckedWithSeedInstruction(
	// Parameters:
	newAuthority ag_solanago.PublicKey,
	authorityType uint32,
	base ag_solanago.PublicKey,
	seed string,
	owner ag_solanago.PublicKey,
	// Accounts:
	stakeAccount ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	authorityBase ag_solanago.PublicKey,
	signers []ag_solanago.PublicKey,
) *AuthorizeCheckedWithSeed {
	return NewAuthorizeCheckedWithSeedInstructionBuilder().
		SetNewAuthority(newAuthority).
		SetAuthorityType(authorityType).
		SetBase(base).
		SetSeed(seed).
		SetOwner(owner).
		SetStakeAccount(stakeAccount).
		SetClockSysvarAccount(clockSysvar).
		SetAuthorityBaseAccount(authorityBase).
		SetSigners(signers...)
}
