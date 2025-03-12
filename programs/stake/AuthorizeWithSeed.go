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

type AuthorizeWithSeed struct {
	// The public key of the base account used to derive the authority address
	Base *solana.PublicKey

	// The seed used to derive the authority address
	Seed *string

	// The public key of the owner program used to derive the authority address
	Owner *solana.PublicKey

	// The public key of the new authority
	NewAuthority *solana.PublicKey

	// The type of authority to authorize
	AuthorityType *AuthorityType

	// [0] = [WRITE] StakeAccount
	// ··········· The stake account to authorize
	//
	// [1] = [] BaseAccount
	// ··········· The base account used to derive the authority address
	//
	// [2] = [] ClockSysvar
	// ··········· Clock sysvar account
	//
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *AuthorizeWithSeed) UnmarshalWithDecoder(dec *bin.Decoder) error {
	if err := dec.Decode(&inst.Base); err != nil {
		return err
	}
	if err := dec.Decode(&inst.Seed); err != nil {
		return err
	}
	if err := dec.Decode(&inst.Owner); err != nil {
		return err
	}
	if err := dec.Decode(&inst.NewAuthority); err != nil {
		return err
	}
	if err := dec.Decode(&inst.AuthorityType); err != nil {
		return err
	}
	return nil
}

func (inst *AuthorizeWithSeed) MarshalWithEncoder(encoder *bin.Encoder) error {
	if err := encoder.Encode(*inst.Base); err != nil {
		return err
	}
	if err := encoder.Encode(*inst.Seed); err != nil {
		return err
	}
	if err := encoder.Encode(*inst.Owner); err != nil {
		return err
	}
	if err := encoder.Encode(*inst.NewAuthority); err != nil {
		return err
	}
	if err := encoder.Encode(*inst.AuthorityType); err != nil {
		return err
	}
	return nil
}

func (inst *AuthorizeWithSeed) Validate() error {
	if inst.Base == nil {
		return errors.New("base parameter is not set")
	}
	if inst.Seed == nil {
		return errors.New("seed parameter is not set")
	}
	if inst.Owner == nil {
		return errors.New("owner parameter is not set")
	}
	if inst.NewAuthority == nil {
		return errors.New("new authority parameter is not set")
	}
	if inst.AuthorityType == nil {
		return errors.New("authority type parameter is not set")
	}

	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

// Stake account account
func (inst *AuthorizeWithSeed) SetStakeAccount(stakeAccount solana.PublicKey) *AuthorizeWithSeed {
	inst.AccountMetaSlice[0] = solana.Meta(stakeAccount).WRITE()
	return inst
}

// Base account
func (inst *AuthorizeWithSeed) SetBaseAccount(baseAccount solana.PublicKey) *AuthorizeWithSeed {
	inst.AccountMetaSlice[1] = solana.Meta(baseAccount)
	return inst
}

// Clock sysvar account
func (inst *AuthorizeWithSeed) SetClockSysvarAccount(clockSysvar solana.PublicKey) *AuthorizeWithSeed {
	inst.AccountMetaSlice[2] = solana.Meta(clockSysvar)
	return inst
}

func (inst *AuthorizeWithSeed) GetStakeAccount() *solana.AccountMeta { return inst.AccountMetaSlice[0] }
func (inst *AuthorizeWithSeed) GetBaseAccount() *solana.AccountMeta  { return inst.AccountMetaSlice[1] }
func (inst *AuthorizeWithSeed) GetClockSysvarAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst *AuthorizeWithSeed) SetBase(base solana.PublicKey) *AuthorizeWithSeed {
	inst.Base = &base
	return inst
}

func (inst *AuthorizeWithSeed) SetSeed(seed string) *AuthorizeWithSeed {
	inst.Seed = &seed
	return inst
}

func (inst *AuthorizeWithSeed) SetOwner(owner solana.PublicKey) *AuthorizeWithSeed {
	inst.Owner = &owner
	return inst
}

func (inst *AuthorizeWithSeed) SetNewAuthority(newAuthority solana.PublicKey) *AuthorizeWithSeed {
	inst.NewAuthority = &newAuthority
	return inst
}

func (inst *AuthorizeWithSeed) SetAuthorityType(authorityType AuthorityType) *AuthorizeWithSeed {
	inst.AuthorityType = &authorityType
	return inst
}

func (inst AuthorizeWithSeed) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.TypeIDFromUint32(Instruction_AuthorizeWithSeed, bin.LE),
	}}
}

func (inst *AuthorizeWithSeed) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("AuthorizeWithSeed")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Account("Base", *inst.Base))
						paramsBranch.Child(format.Param("Seed", *inst.Seed))
						paramsBranch.Child(format.Account("Owner", *inst.Owner))
						paramsBranch.Child(format.Account("NewAuthority", *inst.NewAuthority))
						paramsBranch.Child(format.Param("AuthorityType", *inst.AuthorityType))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(format.Meta("BaseAccount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(format.Meta("ClockSysvar", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

// NewAuthorizeWithSeedInstructionBuilder creates a new `AuthorizeWithSeed` instruction builder.
func NewAuthorizeWithSeedInstructionBuilder() *AuthorizeWithSeed {
	nd := &AuthorizeWithSeed{
		AccountMetaSlice: make(solana.AccountMetaSlice, 3),
	}
	return nd
}

// NewAuthorizeWithSeedInstruction declares a new AuthorizeWithSeed instruction with the provided parameters and accounts.
func NewAuthorizeWithSeedInstruction(
	// parameters:
	base solana.PublicKey,
	seed string,
	owner solana.PublicKey,
	newAuthority solana.PublicKey,
	authorityType AuthorityType,
	// Accounts:
	stakeAccount solana.PublicKey,
	baseAccount solana.PublicKey,
	clockSysvar solana.PublicKey,
) *AuthorizeWithSeed {
	return NewAuthorizeWithSeedInstructionBuilder().
		SetBase(base).
		SetSeed(seed).
		SetOwner(owner).
		SetNewAuthority(newAuthority).
		SetAuthorityType(authorityType).
		SetStakeAccount(stakeAccount).
		SetBaseAccount(baseAccount).
		SetClockSysvarAccount(clockSysvar)
}
