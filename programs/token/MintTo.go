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

// Mints new tokens to an account.  The native mint does not support
// minting.
type MintTo struct {
	// The amount of new tokens to mint.
	Amount *uint64

	// [0] = [WRITE] mint
	// ··········· The mint.
	//
	// [1] = [WRITE] destination
	// ··········· The account to mint tokens to.
	//
	// [2] = [] authority
	// ··········· The mint's minting authority.
	//
	// [3] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewMintToInstructionBuilder creates a new `MintTo` instruction builder.
func NewMintToInstructionBuilder() *MintTo {
	nd := &MintTo{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
// The amount of new tokens to mint.
func (inst *MintTo) SetAmount(amount uint64) *MintTo {
	inst.Amount = &amount
	return inst
}

// SetMintAccount sets the "mint" account.
// The mint.
func (inst *MintTo) SetMintAccount(mint ag_solanago.PublicKey) *MintTo {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
// The mint.
func (inst *MintTo) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDestinationAccount sets the "destination" account.
// The account to mint tokens to.
func (inst *MintTo) SetDestinationAccount(destination ag_solanago.PublicKey) *MintTo {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
// The account to mint tokens to.
func (inst *MintTo) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
// The mint's minting authority.
func (inst *MintTo) SetAuthorityAccount(authority ag_solanago.PublicKey) *MintTo {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The mint's minting authority.
func (inst *MintTo) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSignersAccount sets the "signers" account.
// M signer accounts.
func (inst *MintTo) SetSignersAccount(signers ag_solanago.PublicKey) *MintTo {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

// GetSignersAccount gets the "signers" account.
// M signer accounts.
func (inst *MintTo) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst MintTo) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_MintTo, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintTo) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintTo) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return fmt.Errorf("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return fmt.Errorf("accounts.Signers is not set")
		}
	}
	return nil
}

func (inst *MintTo) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintTo")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("mint", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("destination", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj MintTo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintTo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewMintToInstruction declares a new MintTo instruction with the provided parameters and accounts.
func NewMintToInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	mint ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *MintTo {
	return NewMintToInstructionBuilder().
		SetAmount(amount).
		SetMintAccount(mint).
		SetDestinationAccount(destination).
		SetAuthorityAccount(authority).
		SetSignersAccount(signers)
}
