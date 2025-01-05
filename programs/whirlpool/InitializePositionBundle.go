// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializePositionBundle is the `initializePositionBundle` instruction.
type InitializePositionBundle struct {

	// [0] = [WRITE] positionBundle
	//
	// [1] = [WRITE, SIGNER] positionBundleMint
	//
	// [2] = [WRITE] positionBundleTokenAccount
	//
	// [3] = [] positionBundleOwner
	//
	// [4] = [WRITE, SIGNER] funder
	//
	// [5] = [] tokenProgram
	//
	// [6] = [] systemProgram
	//
	// [7] = [] rent
	//
	// [8] = [] associatedTokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializePositionBundleInstructionBuilder creates a new `InitializePositionBundle` instruction builder.
func NewInitializePositionBundleInstructionBuilder() *InitializePositionBundle {
	nd := &InitializePositionBundle{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetPositionBundleAccount sets the "positionBundle" account.
func (inst *InitializePositionBundle) SetPositionBundleAccount(positionBundle ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(positionBundle).WRITE()
	return inst
}

// GetPositionBundleAccount gets the "positionBundle" account.
func (inst *InitializePositionBundle) GetPositionBundleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPositionBundleMintAccount sets the "positionBundleMint" account.
func (inst *InitializePositionBundle) SetPositionBundleMintAccount(positionBundleMint ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(positionBundleMint).WRITE().SIGNER()
	return inst
}

// GetPositionBundleMintAccount gets the "positionBundleMint" account.
func (inst *InitializePositionBundle) GetPositionBundleMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionBundleTokenAccountAccount sets the "positionBundleTokenAccount" account.
func (inst *InitializePositionBundle) SetPositionBundleTokenAccountAccount(positionBundleTokenAccount ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(positionBundleTokenAccount).WRITE()
	return inst
}

// GetPositionBundleTokenAccountAccount gets the "positionBundleTokenAccount" account.
func (inst *InitializePositionBundle) GetPositionBundleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPositionBundleOwnerAccount sets the "positionBundleOwner" account.
func (inst *InitializePositionBundle) SetPositionBundleOwnerAccount(positionBundleOwner ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(positionBundleOwner)
	return inst
}

// GetPositionBundleOwnerAccount gets the "positionBundleOwner" account.
func (inst *InitializePositionBundle) GetPositionBundleOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializePositionBundle) SetFunderAccount(funder ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializePositionBundle) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *InitializePositionBundle) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *InitializePositionBundle) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializePositionBundle) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializePositionBundle) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializePositionBundle) SetRentAccount(rent ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializePositionBundle) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *InitializePositionBundle) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *InitializePositionBundle {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *InitializePositionBundle) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst InitializePositionBundle) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializePositionBundle,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializePositionBundle) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializePositionBundle) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PositionBundle is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PositionBundleMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PositionBundleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PositionBundleOwner is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *InitializePositionBundle) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializePositionBundle")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        positionBundle", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    positionBundleMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   positionBundleToken", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   positionBundleOwner", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                funder", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj InitializePositionBundle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializePositionBundle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializePositionBundleInstruction declares a new InitializePositionBundle instruction with the provided parameters and accounts.
func NewInitializePositionBundleInstruction(
	// Accounts:
	positionBundle ag_solanago.PublicKey,
	positionBundleMint ag_solanago.PublicKey,
	positionBundleTokenAccount ag_solanago.PublicKey,
	positionBundleOwner ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *InitializePositionBundle {
	return NewInitializePositionBundleInstructionBuilder().
		SetPositionBundleAccount(positionBundle).
		SetPositionBundleMintAccount(positionBundleMint).
		SetPositionBundleTokenAccountAccount(positionBundleTokenAccount).
		SetPositionBundleOwnerAccount(positionBundleOwner).
		SetFunderAccount(funder).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
