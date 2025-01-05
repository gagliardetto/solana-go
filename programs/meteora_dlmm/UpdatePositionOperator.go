// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdatePositionOperator is the `updatePositionOperator` instruction.
type UpdatePositionOperator struct {
	Operator *ag_solanago.PublicKey

	// [0] = [WRITE] position
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [] eventAuthority
	//
	// [3] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdatePositionOperatorInstructionBuilder creates a new `UpdatePositionOperator` instruction builder.
func NewUpdatePositionOperatorInstructionBuilder() *UpdatePositionOperator {
	nd := &UpdatePositionOperator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetOperator sets the "operator" parameter.
func (inst *UpdatePositionOperator) SetOperator(operator ag_solanago.PublicKey) *UpdatePositionOperator {
	inst.Operator = &operator
	return inst
}

// SetPositionAccount sets the "position" account.
func (inst *UpdatePositionOperator) SetPositionAccount(position ag_solanago.PublicKey) *UpdatePositionOperator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *UpdatePositionOperator) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *UpdatePositionOperator) SetOwnerAccount(owner ag_solanago.PublicKey) *UpdatePositionOperator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *UpdatePositionOperator) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *UpdatePositionOperator) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *UpdatePositionOperator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *UpdatePositionOperator) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProgramAccount sets the "program" account.
func (inst *UpdatePositionOperator) SetProgramAccount(program ag_solanago.PublicKey) *UpdatePositionOperator {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *UpdatePositionOperator) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst UpdatePositionOperator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdatePositionOperator,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdatePositionOperator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdatePositionOperator) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Operator == nil {
			return errors.New("Operator parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *UpdatePositionOperator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdatePositionOperator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Operator", *inst.Operator))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj UpdatePositionOperator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Operator` param:
	err = encoder.Encode(obj.Operator)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdatePositionOperator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Operator`:
	err = decoder.Decode(&obj.Operator)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdatePositionOperatorInstruction declares a new UpdatePositionOperator instruction with the provided parameters and accounts.
func NewUpdatePositionOperatorInstruction(
	// Parameters:
	operator ag_solanago.PublicKey,
	// Accounts:
	position ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *UpdatePositionOperator {
	return NewUpdatePositionOperatorInstructionBuilder().
		SetOperator(operator).
		SetPositionAccount(position).
		SetOwnerAccount(owner).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
