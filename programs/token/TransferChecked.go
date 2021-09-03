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

// Transfers tokens from one account to another either directly or via a
// delegate.  If this account is associated with the native mint then equal
// amounts of SOL and Tokens will be transferred to the destination
// account.
//
// This instruction differs from Transfer in that the token mint and
// decimals value is checked by the caller.  This may be useful when
// creating transactions offline or within a hardware wallet.
type TransferChecked struct {
	// The amount of tokens to transfer.
	Amount *uint64

	// Expected number of base 10 digits to the right of the decimal place.
	Decimals *uint8

	// [0] = [WRITE] source
	// ··········· The source account.
	//
	// [1] = [] mint
	// ··········· The token mint.
	//
	// [2] = [WRITE] destination
	// ··········· The destination account.
	//
	// [3] = [] owner
	// ··········· The source account's owner/delegate.
	//
	// [4] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewTransferCheckedInstructionBuilder creates a new `TransferChecked` instruction builder.
func NewTransferCheckedInstructionBuilder() *TransferChecked {
	nd := &TransferChecked{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of tokens to transfer.
func (inst *TransferChecked) SetAmount(amount uint64) *TransferChecked {
	inst.Amount = &amount
	return inst
}

// SetDecimals sets the "decimals" parameter.
// Expected number of base 10 digits to the right of the decimal place.
func (inst *TransferChecked) SetDecimals(decimals uint8) *TransferChecked {
	inst.Decimals = &decimals
	return inst
}

// SetSourceAccount sets the "source" account.
// The source account.
func (inst *TransferChecked) SetSourceAccount(source ag_solanago.PublicKey) *TransferChecked {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(source).WRITE()
	return inst
}

// GetSourceAccount gets the "source" account.
// The source account.
func (inst *TransferChecked) GetSourceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
// The token mint.
func (inst *TransferChecked) SetMintAccount(mint ag_solanago.PublicKey) *TransferChecked {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The token mint.
func (inst *TransferChecked) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetDestinationAccount sets the "destination" account.
// The destination account.
func (inst *TransferChecked) SetDestinationAccount(destination ag_solanago.PublicKey) *TransferChecked {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
// The destination account.
func (inst *TransferChecked) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetOwnerAccount sets the "owner" account.
// The source account's owner/delegate.
func (inst *TransferChecked) SetOwnerAccount(owner ag_solanago.PublicKey) *TransferChecked {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The source account's owner/delegate.
func (inst *TransferChecked) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSignersAccount sets the "signers" account.
// M signer accounts.
func (inst *TransferChecked) SetSignersAccount(signers ag_solanago.PublicKey) *TransferChecked {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

// GetSignersAccount gets the "signers" account.
// M signer accounts.
func (inst *TransferChecked) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst TransferChecked) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_TransferChecked, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferChecked) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferChecked) Validate() error {
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
			return fmt.Errorf("accounts.Source is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return fmt.Errorf("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return fmt.Errorf("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return fmt.Errorf("accounts.Signers is not set")
		}
	}
	return nil
}

func (inst *TransferChecked) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferChecked")).
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
						accountsBranch.Child(ag_format.Meta("destination", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj TransferChecked) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *TransferChecked) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewTransferCheckedInstruction declares a new TransferChecked instruction with the provided parameters and accounts.
func NewTransferCheckedInstruction(
	// Parameters:
	amount uint64,
	decimals uint8,
	// Accounts:
	source ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *TransferChecked {
	return NewTransferCheckedInstructionBuilder().
		SetAmount(amount).
		SetDecimals(decimals).
		SetSourceAccount(source).
		SetMintAccount(mint).
		SetDestinationAccount(destination).
		SetOwnerAccount(owner).
		SetSignersAccount(signers)
}
