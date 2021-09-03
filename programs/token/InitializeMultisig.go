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

// Initializes a multisignature account with N provided signers.
//
// Multisignature accounts can used in place of any single owner/delegate
// accounts in any token instruction that require an owner/delegate to be
// present.  The variant field represents the number of signers (M)
// required to validate this multisignature account.
//
// The `InitializeMultisig` instruction requires no signers and MUST be
// included within the same Transaction as the system program's
// `CreateAccount` instruction that creates the account being initialized.
// Otherwise another party can acquire ownership of the uninitialized
// account.
type InitializeMultisig struct {
	// The number of signers (M) required to validate this multisignature
	// account.
	M *uint8

	// [0] = [WRITE] account
	// ··········· The multisignature account to initialize.
	//
	// [1] = [] $(SysVarRentPubkey)
	// ··········· Rent sysvar.
	//
	// [2...] = [SIGNER] signers
	// ··········· ..2+N The signer accounts, must equal to N where 1 <= N <=11
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeMultisigInstructionBuilder creates a new `InitializeMultisig` instruction builder.
func NewInitializeMultisigInstructionBuilder() *InitializeMultisig {
	nd := &InitializeMultisig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(ag_solanago.SysVarRentPubkey)
	return nd
}

// The number of signers (M) required to validate this multisignature
// account.
func (inst *InitializeMultisig) SetM(m uint8) *InitializeMultisig {
	inst.M = &m
	return inst
}

// The multisignature account to initialize.
func (inst *InitializeMultisig) SetAccount(account ag_solanago.PublicKey) *InitializeMultisig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *InitializeMultisig) GetAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// Rent sysvar.
func (inst *InitializeMultisig) SetSysVarRentPubkeyAccount(SysVarRentPubkey ag_solanago.PublicKey) *InitializeMultisig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(SysVarRentPubkey)
	return inst
}

func (inst *InitializeMultisig) GetSysVarRentPubkeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// ..2+N The signer accounts, must equal to N where 1 <= N <=11
func (inst *InitializeMultisig) SetSigners(signers ...ag_solanago.PublicKey) *InitializeMultisig {
	for _, signer := range signers {
		inst.AccountMetaSlice = append(inst.AccountMetaSlice, ag_solanago.Meta(signer).SIGNER())
	}
	return inst
}

func (inst *InitializeMultisig) GetSigners() []*ag_solanago.AccountMeta {
	if len(inst.AccountMetaSlice) == 2 {
		return nil
	}
	return inst.AccountMetaSlice[2:]
}

func (inst InitializeMultisig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_InitializeMultisig, binary.LittleEndian),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeMultisig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeMultisig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.M == nil {
			return errors.New("M parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return fmt.Errorf("accounts.Account is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return fmt.Errorf("accounts.SysVarRentPubkey is not set")
		}
		if len(inst.AccountMetaSlice) == 2 {
			return fmt.Errorf("Signers are not set")
		}
	}
	return nil
}

func (inst *InitializeMultisig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeMultisig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("M", *inst.M))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("account", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("$(SysVarRentPubkey)", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("signers", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj InitializeMultisig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `M` param:
	err = encoder.Encode(obj.M)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeMultisig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `M`:
	err = decoder.Decode(&obj.M)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeMultisigInstruction declares a new InitializeMultisig instruction with the provided parameters and accounts.
func NewInitializeMultisigInstruction(
	// Parameters:
	m uint8,
	// Accounts:
	account ag_solanago.PublicKey,
	SysVarRentPubkey ag_solanago.PublicKey,
	signers ag_solanago.PublicKey) *InitializeMultisig {
	return NewInitializeMultisigInstructionBuilder().
		SetM(m).
		SetAccount(account).
		SetSysVarRentPubkeyAccount(SysVarRentPubkey).
		SetSignersAccount(signers)
}
