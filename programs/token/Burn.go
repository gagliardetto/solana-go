package token

import (
	"encoding/binary"
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Burns tokens by removing them from an account.  `Burn` does not support
// accounts associated with the native mint, use `CloseAccount` instead.
type Burn struct {
	// The amount of tokens to burn.
	Amount *uint64

	// [0] = [WRITE] source
	// ··········· The account to burn from.
	//
	// [1] = [WRITE] mint
	// ··········· The token mint.
	//
	// [2] = [] owner
	// ··········· The account's owner/delegate.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewBurnInstructionBuilder creates a new `Burn` instruction builder.
func NewBurnInstructionBuilder() *Burn {
	nd := &Burn{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of tokens to burn.
func (inst *Burn) SetAmount(amount uint64) *Burn {
	inst.Amount = &amount
	return inst
}

// SetSourceAccount sets the "source" account.
// The account to burn from.
func (inst *Burn) SetSourceAccount(source ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
// The account to burn from.
func (inst *Burn) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
// The token mint.
func (inst *Burn) SetMintAccount(mint ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
// The token mint.
func (inst *Burn) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetOwnerAccount sets the "owner" account.
// The account's owner/delegate.
func (inst *Burn) SetOwnerAccount(owner ag_solanago.PublicKey) *Burn {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The account's owner/delegate.
func (inst *Burn) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst Burn) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Burn, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Burn) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Burn) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Source is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
	}
	return nil
}

func (inst *Burn) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Burn")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("source", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj Burn) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Burn) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewBurnInstruction declares a new Burn instruction with the provided parameters and accounts.
func NewBurnInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	source ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *Burn {
	return NewBurnInstructionBuilder().
		SetAmount(amount).
		SetSourceAccount(source).
		SetMintAccount(mint).
		SetOwnerAccount(owner)
}
