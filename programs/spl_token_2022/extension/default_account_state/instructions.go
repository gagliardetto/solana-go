// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package default_account_state

import (
	"errors"

	binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go/programs/common"
	spltoken2022 "github.com/gagliardetto/solana-go/programs/spl_token_2022"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Initialize Instruction
type Initialize struct {
	State *spltoken2022.AccountState
	// [0] = [WRITE] mint `The mint to initialize.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	return &Initialize{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewInitializeInstruction
//
// Parameters:
//
//	state:
//	mint: The mint to initialize.
func NewInitializeInstruction(
	state spltoken2022.AccountState,
	mint common.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetState(state).
		SetMintAccount(mint)
}

// SetState sets the "state" parameter.
func (obj *Initialize) SetState(state spltoken2022.AccountState) *Initialize {
	obj.State = &state
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint to initialize.
func (obj *Initialize) SetMintAccount(mint common.PublicKey, multiSigners ...common.PublicKey) *Initialize {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[0] = common.Meta(mint)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	}
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint to initialize.
func (obj *Initialize) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *Initialize) SetProgramId(programId *common.PublicKey) *Initialize {
	obj._programId = programId
	return obj
}

func (obj *Initialize) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Initialize}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Initialize) Validate() error {
	if obj.State == nil {
		return errors.New("[Initialize] state param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Initialize] accounts.mint is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Initialize) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.State); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.State); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Initialize")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("State", *obj.State))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("mint", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}

// Update Instruction
type Update struct {
	State *spltoken2022.AccountState
	// [0] = [WRITE] mint `The mint.`
	// [1] = [SIGNER] mintAuthority `The mint freeze authority.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewUpdateInstructionBuilder creates a new `Update` instruction builder.
func NewUpdateInstructionBuilder() *Update {
	return &Update{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewUpdateInstruction
//
// Parameters:
//
//	state:
//	mint: The mint.
//	mintAuthority: The mint freeze authority.
func NewUpdateInstruction(
	state spltoken2022.AccountState,
	mint common.PublicKey,
	mintAuthority common.PublicKey,
) *Update {
	return NewUpdateInstructionBuilder().
		SetState(state).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority)
}

// SetState sets the "state" parameter.
func (obj *Update) SetState(state spltoken2022.AccountState) *Update {
	obj.State = &state
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint.
func (obj *Update) SetMintAccount(mint common.PublicKey) *Update {
	obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint.
func (obj *Update) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetMintAuthorityAccount sets the "mintAuthority" parameter.
// The mint freeze authority.
func (obj *Update) SetMintAuthorityAccount(mintAuthority common.PublicKey, multiSigners ...common.PublicKey) *Update {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(mintAuthority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(mintAuthority).SIGNER()
	}
	return obj
}

// GetMintAuthorityAccount gets the "mintAuthority" parameter.
// The mint freeze authority.
func (obj *Update) GetMintAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *Update) SetProgramId(programId *common.PublicKey) *Update {
	obj._programId = programId
	return obj
}

func (obj *Update) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Update}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Update) Validate() error {
	if obj.State == nil {
		return errors.New("[Update] state param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Update] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Update] accounts.mintAuthority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Update) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Update) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.State); err != nil {
		return err
	}
	return nil
}

func (obj *Update) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.State); err != nil {
		return err
	}
	return nil
}

func (obj *Update) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Update")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("State", *obj.State))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("         mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("mintAuthority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}
