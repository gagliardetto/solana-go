package token

import (
	"encoding/binary"
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Approves a delegate.  A delegate is given the authority over tokens on
// behalf of the source account's owner.
//
// This instruction differs from Approve in that the token mint and
// decimals value is checked by the caller.  This may be useful when
// creating transactions offline or within a hardware wallet.
type ApproveChecked struct {
	// The amount of tokens the delegate is approved for.
	Amount *uint64

	// Expected number of base 10 digits to the right of the decimal place.
	Decimals *uint8

	// [0] = [WRITE] source
	// ··········· The source account.
	//
	// [1] = [] mint
	// ··········· The token mint.
	//
	// [2] = [] delegate
	// ··········· The delegate.
	//
	// [3] = [] owner
	// ··········· The source account owner.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewApproveCheckedInstructionBuilder creates a new `ApproveChecked` instruction builder.
func NewApproveCheckedInstructionBuilder() *ApproveChecked {
	nd := &ApproveChecked{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of tokens the delegate is approved for.
func (inst *ApproveChecked) SetAmount(amount uint64) *ApproveChecked {
	inst.Amount = &amount
	return inst
}

// SetDecimals sets the "decimals" parameter.
// Expected number of base 10 digits to the right of the decimal place.
func (inst *ApproveChecked) SetDecimals(decimals uint8) *ApproveChecked {
	inst.Decimals = &decimals
	return inst
}

// SetSourceAccount sets the "source" account.
// The source account.
func (inst *ApproveChecked) SetSourceAccount(source ag_solanago.PublicKey) *ApproveChecked {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
// The source account.
func (inst *ApproveChecked) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
// The token mint.
func (inst *ApproveChecked) SetMintAccount(mint ag_solanago.PublicKey) *ApproveChecked {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The token mint.
func (inst *ApproveChecked) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetDelegateAccount sets the "delegate" account.
// The delegate.
func (inst *ApproveChecked) SetDelegateAccount(delegate ag_solanago.PublicKey) *ApproveChecked {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(delegate)
	return inst
}

// GetDelegateAccount gets the "delegate" account.
// The delegate.
func (inst *ApproveChecked) GetDelegateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetOwnerAccount sets the "owner" account.
// The source account owner.
func (inst *ApproveChecked) SetOwnerAccount(owner ag_solanago.PublicKey) *ApproveChecked {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The source account owner.
func (inst *ApproveChecked) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst ApproveChecked) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_ApproveChecked, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveChecked) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveChecked) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Decimals == nil {
			return errors.New("Decimals parameter is not set")
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
			return errors.New("accounts.Delegate is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Owner is not set")
		}
	}
	return nil
}

func (inst *ApproveChecked) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveChecked")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("Decimals", *inst.Decimals))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("source", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("delegate", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj ApproveChecked) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ApproveChecked) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}

// NewApproveCheckedInstruction declares a new ApproveChecked instruction with the provided parameters and accounts.
func NewApproveCheckedInstruction(
	// Parameters:
	amount uint64,
	decimals uint8,
	// Accounts:
	source ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	delegate ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *ApproveChecked {
	return NewApproveCheckedInstructionBuilder().
		SetAmount(amount).
		SetDecimals(decimals).
		SetSourceAccount(source).
		SetMintAccount(mint).
		SetDelegateAccount(delegate).
		SetOwnerAccount(owner)
}
