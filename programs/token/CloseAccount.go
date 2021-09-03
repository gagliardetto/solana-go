package token

import (
	"encoding/binary"
	"fmt"
	ag_binary "github.com/dfuse-io/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Close an account by transferring all its SOL to the destination account.
// Non-native accounts may only be closed if its token amount is zero.
type CloseAccount struct {

	// [0] = [WRITE] account
	// ··········· The account to close.
	//
	// [1] = [WRITE] destination
	// ··········· The destination account.
	//
	// [2] = [] owner
	// ··········· The account's owner.
	//
	// [3] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCloseAccountInstructionBuilder creates a new `CloseAccount` instruction builder.
func NewCloseAccountInstructionBuilder() *CloseAccount {
	nd := &CloseAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// The account to close.
func (inst *CloseAccount) SetAccount(account ag_solanago.PublicKey) *CloseAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *CloseAccount) GetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// The destination account.
func (inst *CloseAccount) SetDestinationAccount(destination ag_solanago.PublicKey) *CloseAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destination).WRITE()
	return inst
}

func (inst *CloseAccount) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// The account's owner.
func (inst *CloseAccount) SetOwnerAccount(owner ag_solanago.PublicKey) *CloseAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner)
	return inst
}

func (inst *CloseAccount) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// M signer accounts.
func (inst *CloseAccount) SetSignersAccount(signers ag_solanago.PublicKey) *CloseAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

func (inst *CloseAccount) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst CloseAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_CloseAccount, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Account is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Destination is not set")
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

func (inst *CloseAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("account", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("destination", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj CloseAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseAccountInstruction declares a new CloseAccount instruction with the provided parameters and accounts.
func NewCloseAccountInstruction(
	// Accounts:
	account ag_solanago.PublicKey,
	destination ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *CloseAccount {
	return NewCloseAccountInstructionBuilder().
		SetAccount(account).
		SetDestinationAccount(destination).
		SetOwnerAccount(owner).
		SetSignersAccount(signers)
}
