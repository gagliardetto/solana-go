// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// initialize new vault
type Initialize struct {

	// [0] = [WRITE] vault
	// ··········· This is base account for all vault
	// ··········· No need base key now because we only allow 1 vault per token now
	// ··········· Vault account
	//
	// [1] = [WRITE, SIGNER] payer
	// ··········· Payer can be anyone
	//
	// [2] = [WRITE] tokenVault
	// ··········· Token vault account
	//
	// [3] = [] tokenMint
	// ··········· Token mint account
	//
	// [4] = [WRITE] lpMint
	//
	// [5] = [] rent
	// ··········· rent
	//
	// [6] = [] tokenProgram
	// ··········· token_program
	//
	// [7] = [] systemProgram
	// ··········· system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetVaultAccount sets the "vault" account.
// This is base account for all vault
// No need base key now because we only allow 1 vault per token now
// Vault account
func (inst *Initialize) SetVaultAccount(vault ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
// This is base account for all vault
// No need base key now because we only allow 1 vault per token now
// Vault account
func (inst *Initialize) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
// Payer can be anyone
func (inst *Initialize) SetPayerAccount(payer ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer can be anyone
func (inst *Initialize) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenVaultAccount sets the "tokenVault" account.
// Token vault account
func (inst *Initialize) SetTokenVaultAccount(tokenVault ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenVault).WRITE()
	return inst
}

// GetTokenVaultAccount gets the "tokenVault" account.
// Token vault account
func (inst *Initialize) GetTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintAccount sets the "tokenMint" account.
// Token mint account
func (inst *Initialize) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
// Token mint account
func (inst *Initialize) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLpMintAccount sets the "lpMint" account.
func (inst *Initialize) SetLpMintAccount(lpMint ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(lpMint).WRITE()
	return inst
}

// GetLpMintAccount gets the "lpMint" account.
func (inst *Initialize) GetLpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentAccount sets the "rent" account.
// rent
func (inst *Initialize) SetRentAccount(rent ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// rent
func (inst *Initialize) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// token_program
func (inst *Initialize) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// token_program
func (inst *Initialize) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// system_program
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// system_program
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LpMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        vault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   tokenVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       lpMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Accounts:
	vault ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenVault ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	lpMint ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetVaultAccount(vault).
		SetPayerAccount(payer).
		SetTokenVaultAccount(tokenVault).
		SetTokenMintAccount(tokenMint).
		SetLpMintAccount(lpMint).
		SetRentAccount(rent).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram)
}
