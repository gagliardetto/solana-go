// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MigratePosition is the `migratePosition` instruction.
type MigratePosition struct {

	// [0] = [WRITE, SIGNER] positionV2
	//
	// [1] = [WRITE] positionV1
	//
	// [2] = [] lbPair
	//
	// [3] = [WRITE] binArrayLower
	//
	// [4] = [WRITE] binArrayUpper
	//
	// [5] = [WRITE, SIGNER] owner
	//
	// [6] = [] systemProgram
	//
	// [7] = [WRITE] rentReceiver
	//
	// [8] = [] eventAuthority
	//
	// [9] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMigratePositionInstructionBuilder creates a new `MigratePosition` instruction builder.
func NewMigratePositionInstructionBuilder() *MigratePosition {
	nd := &MigratePosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetPositionV2Account sets the "positionV2" account.
func (inst *MigratePosition) SetPositionV2Account(positionV2 ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(positionV2).WRITE().SIGNER()
	return inst
}

// GetPositionV2Account gets the "positionV2" account.
func (inst *MigratePosition) GetPositionV2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPositionV1Account sets the "positionV1" account.
func (inst *MigratePosition) SetPositionV1Account(positionV1 ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(positionV1).WRITE()
	return inst
}

// GetPositionV1Account gets the "positionV1" account.
func (inst *MigratePosition) GetPositionV1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *MigratePosition) SetLbPairAccount(lbPair ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(lbPair)
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *MigratePosition) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBinArrayLowerAccount sets the "binArrayLower" account.
func (inst *MigratePosition) SetBinArrayLowerAccount(binArrayLower ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(binArrayLower).WRITE()
	return inst
}

// GetBinArrayLowerAccount gets the "binArrayLower" account.
func (inst *MigratePosition) GetBinArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBinArrayUpperAccount sets the "binArrayUpper" account.
func (inst *MigratePosition) SetBinArrayUpperAccount(binArrayUpper ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(binArrayUpper).WRITE()
	return inst
}

// GetBinArrayUpperAccount gets the "binArrayUpper" account.
func (inst *MigratePosition) GetBinArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOwnerAccount sets the "owner" account.
func (inst *MigratePosition) SetOwnerAccount(owner ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *MigratePosition) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MigratePosition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MigratePosition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentReceiverAccount sets the "rentReceiver" account.
func (inst *MigratePosition) SetRentReceiverAccount(rentReceiver ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rentReceiver).WRITE()
	return inst
}

// GetRentReceiverAccount gets the "rentReceiver" account.
func (inst *MigratePosition) GetRentReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *MigratePosition) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *MigratePosition) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetProgramAccount sets the "program" account.
func (inst *MigratePosition) SetProgramAccount(program ag_solanago.PublicKey) *MigratePosition {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *MigratePosition) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst MigratePosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MigratePosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MigratePosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MigratePosition) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PositionV2 is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PositionV1 is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BinArrayLower is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BinArrayUpper is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.RentReceiver is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *MigratePosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MigratePosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    positionV2", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    positionV1", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        lbPair", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" binArrayLower", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" binArrayUpper", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  rentReceiver", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj MigratePosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *MigratePosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewMigratePositionInstruction declares a new MigratePosition instruction with the provided parameters and accounts.
func NewMigratePositionInstruction(
	// Accounts:
	positionV2 ag_solanago.PublicKey,
	positionV1 ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *MigratePosition {
	return NewMigratePositionInstructionBuilder().
		SetPositionV2Account(positionV2).
		SetPositionV1Account(positionV1).
		SetLbPairAccount(lbPair).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetOwnerAccount(owner).
		SetSystemProgramAccount(systemProgram).
		SetRentReceiverAccount(rentReceiver).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
