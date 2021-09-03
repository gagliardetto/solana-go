package token

import (
	"encoding/binary"
	"errors"
	"fmt"
	ag_binary "github.com/dfuse-io/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets a new authority of a mint or account.
type SetAuthority struct {
	// The type of authority to update.
	AuthorityType *AuthorityType

	// The new authority.
	NewAuthority *ag_solanago.PublicKey `bin:"optional"`

	// [0] = [WRITE] subject
	// ··········· The mint or account to change the authority of.
	//
	// [1] = [] authority
	// ··········· The current authority of the mint or account.
	//
	// [2] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetAuthorityInstructionBuilder creates a new `SetAuthority` instruction builder.
func NewSetAuthorityInstructionBuilder() *SetAuthority {
	nd := &SetAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// The type of authority to update.
func (inst *SetAuthority) SetAuthorityType(authority_type AuthorityType) *SetAuthority {
	inst.AuthorityType = &authority_type
	return inst
}

// The new authority.
func (inst *SetAuthority) SetNewAuthority(new_authority ag_solanago.PublicKey) *SetAuthority {
	inst.NewAuthority = &new_authority
	return inst
}

// The mint or account to change the authority of.
func (inst *SetAuthority) SetSubjectAccount(subject ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(subject).WRITE()
	return inst
}

func (inst *SetAuthority) GetSubjectAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// The current authority of the mint or account.
func (inst *SetAuthority) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

func (inst *SetAuthority) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// M signer accounts.
func (inst *SetAuthority) SetSignersAccount(signers ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

func (inst *SetAuthority) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_SetAuthority, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AuthorityType == nil {
			return errors.New("AuthorityType parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Subject is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return fmt.Errorf("accounts.Signers is not set")
		}
	}
	return nil
}

func (inst *SetAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("AuthorityType", *inst.AuthorityType))
						paramsBranch.Child(ag_format.Param("NewAuthority (OPTIONAL)", inst.NewAuthority))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("subject", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj SetAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AuthorityType` param:
	err = encoder.Encode(obj.AuthorityType)
	if err != nil {
		return err
	}
	// Serialize `NewAuthority` param (optional):
	{
		if obj.NewAuthority == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.NewAuthority)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *SetAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AuthorityType`:
	err = decoder.Decode(&obj.AuthorityType)
	if err != nil {
		return err
	}
	// Deserialize `NewAuthority` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.NewAuthority)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewSetAuthorityInstruction declares a new SetAuthority instruction with the provided parameters and accounts.
func NewSetAuthorityInstruction(
	// Parameters:
	authority_type AuthorityType,
	new_authority ag_solanago.PublicKey,
	// Accounts:
	subject ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *SetAuthority {
	return NewSetAuthorityInstructionBuilder().
		SetAuthorityType(authority_type).
		SetNewAuthority(new_authority).
		SetSubjectAccount(subject).
		SetAuthorityAccount(authority).
		SetSignersAccount(signers)
}
