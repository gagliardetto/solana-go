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

// Approves a delegate.  A delegate is given the authority over tokens on
// behalf of the source account's owner.
type Approve struct {
	// The amount of tokens the delegate is approved for.
	Amount *uint64

	// [0] = [WRITE] source
	// ··········· The source account.
	//
	// [1] = [] delegate
	// ··········· The delegate.
	//
	// [2] = [] owner
	// ··········· The source account owner.
	//
	// [3] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewApproveInstructionBuilder creates a new `Approve` instruction builder.
func NewApproveInstructionBuilder() *Approve {
	nd := &Approve{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of tokens the delegate is approved for.
func (inst *Approve) SetAmount(amount uint64) *Approve {
	inst.Amount = &amount
	return inst
}

// SetSourceAccount sets the "source" account.
// The source account.
func (inst *Approve) SetSourceAccount(source ag_solanago.PublicKey) *Approve {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
// The source account.
func (inst *Approve) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDelegateAccount sets the "delegate" account.
// The delegate.
func (inst *Approve) SetDelegateAccount(delegate ag_solanago.PublicKey) *Approve {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(delegate)
	return inst
}

// GetDelegateAccount gets the "delegate" account.
// The delegate.
func (inst *Approve) GetDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetOwnerAccount sets the "owner" account.
// The source account owner.
func (inst *Approve) SetOwnerAccount(owner ag_solanago.PublicKey) *Approve {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The source account owner.
func (inst *Approve) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSignersAccount sets the "signers" account.
// M signer accounts.
func (inst *Approve) SetSignersAccount(signers ag_solanago.PublicKey) *Approve {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

// GetSignersAccount gets the "signers" account.
// M signer accounts.
func (inst *Approve) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst Approve) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Approve, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Approve) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Approve) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Source is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Delegate is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return fmt.Errorf("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return fmt.Errorf("accounts.Signers is not set")
		}
	}
	return nil
}

func (inst *Approve) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Approve")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("source", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("delegate", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj Approve) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Approve) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewApproveInstruction declares a new Approve instruction with the provided parameters and accounts.
func NewApproveInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	source ag_solanago.PublicKey,
	delegate ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *Approve {
	return NewApproveInstructionBuilder().
		SetAmount(amount).
		SetSourceAccount(source).
		SetDelegateAccount(delegate).
		SetOwnerAccount(owner).
		SetSignersAccount(signers)
}
