package token

import (
	"encoding/binary"
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Revokes the delegate's authority.
type Revoke struct {

	// [0] = [WRITE] source
	// ··········· The source account.
	//
	// [1] = [] owner
	// ··········· The source account's owner.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRevokeInstructionBuilder creates a new `Revoke` instruction builder.
func NewRevokeInstructionBuilder() *Revoke {
	nd := &Revoke{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetSourceAccount sets the "source" account.
// The source account.
func (inst *Revoke) SetSourceAccount(source ag_solanago.PublicKey) *Revoke {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
// The source account.
func (inst *Revoke) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOwnerAccount sets the "owner" account.
// The source account's owner.
func (inst *Revoke) SetOwnerAccount(owner ag_solanago.PublicKey) *Revoke {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The source account's owner.
func (inst *Revoke) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst Revoke) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Revoke, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Revoke) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Revoke) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Source is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
	}
	return nil
}

func (inst *Revoke) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Revoke")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("source", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj Revoke) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Revoke) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRevokeInstruction declares a new Revoke instruction with the provided parameters and accounts.
func NewRevokeInstruction(
	// Accounts:
	source ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *Revoke {
	return NewRevokeInstructionBuilder().
		SetSourceAccount(source).
		SetOwnerAccount(owner)
}
