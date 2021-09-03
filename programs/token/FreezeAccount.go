package token

import (
	"encoding/binary"
	"fmt"
	ag_binary "github.com/dfuse-io/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Freeze an Initialized account using the Mint's freeze_authority (if set).
type FreezeAccount struct {

	// [0] = [WRITE] account
	// ··········· The account to freeze.
	//
	// [1] = [] mint
	// ··········· The token mint.
	//
	// [2] = [] authority
	// ··········· The mint freeze authority.
	//
	// [3] = [SIGNER] signers
	// ··········· M signer accounts.
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewFreezeAccountInstructionBuilder creates a new `FreezeAccount` instruction builder.
func NewFreezeAccountInstructionBuilder() *FreezeAccount {
	nd := &FreezeAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetAccount sets the "account" account.
// The account to freeze.
func (inst *FreezeAccount) SetAccount(account ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(account).WRITE()
	return inst
}

// GetAccount gets the "account" account.
// The account to freeze.
func (inst *FreezeAccount) GetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
// The token mint.
func (inst *FreezeAccount) SetMintAccount(mint ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
// The token mint.
func (inst *FreezeAccount) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
// The mint freeze authority.
func (inst *FreezeAccount) SetAuthorityAccount(authority ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// The mint freeze authority.
func (inst *FreezeAccount) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSignersAccount sets the "signers" account.
// M signer accounts.
func (inst *FreezeAccount) SetSignersAccount(signers ag_solanago.PublicKey) *FreezeAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(signers).SIGNER()
	return inst
}

// GetSignersAccount gets the "signers" account.
// M signer accounts.
func (inst *FreezeAccount) GetSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst FreezeAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_FreezeAccount, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FreezeAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FreezeAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Account is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.Mint is not set")
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

func (inst *FreezeAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FreezeAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("account", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj FreezeAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *FreezeAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewFreezeAccountInstruction declares a new FreezeAccount instruction with the provided parameters and accounts.
func NewFreezeAccountInstruction(
	// Accounts:
	account ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *FreezeAccount {
	return NewFreezeAccountInstructionBuilder().
		SetAccount(account).
		SetMintAccount(mint).
		SetAuthorityAccount(authority).
		SetSignersAccount(signers)
}
