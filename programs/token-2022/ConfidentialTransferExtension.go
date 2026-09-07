package token2022

import (
	"encoding"
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConfidentialTransfer sub-instruction IDs.
const (
	ConfidentialTransfer_InitializeMint uint8 = iota
	ConfidentialTransfer_UpdateMint
	ConfidentialTransfer_ConfigureAccount
	ConfidentialTransfer_ApproveAccount
	ConfidentialTransfer_EmptyAccount
	ConfidentialTransfer_Deposit
	ConfidentialTransfer_Withdraw
	ConfidentialTransfer_Transfer
	ConfidentialTransfer_ApplyPendingBalance
	ConfidentialTransfer_EnableConfidentialCredits
	ConfidentialTransfer_DisableConfidentialCredits
	ConfidentialTransfer_EnableNonConfidentialCredits
	ConfidentialTransfer_DisableNonConfidentialCredits
	ConfidentialTransfer_TransferWithFee
	ConfidentialTransfer_ConfigureAccountWithRegistry
)

const (
	// Deprecated: sub-instruction 13 is TransferWithFee; use
	// ConfidentialTransfer_TransferWithFee.
	ConfidentialTransfer_TransferWithSplitProofs = ConfidentialTransfer_TransferWithFee
	// Deprecated: sub-instruction 14 is ConfigureAccountWithRegistry; use
	// ConfidentialTransfer_ConfigureAccountWithRegistry.
	ConfidentialTransfer_TransferWithSplitProofsInParallel = ConfidentialTransfer_ConfigureAccountWithRegistry
)

// ConfidentialTransferInstructionData is the data of a ConfidentialTransfer
// sub-instruction, implemented by the ConfidentialTransfer*Data structs.
type ConfidentialTransferInstructionData interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	bytes() []byte
}

// ctSubInstructions describes every ConfidentialTransfer sub-instruction.
var ctSubInstructions = [...]struct {
	name    string
	newData func() ConfidentialTransferInstructionData
}{
	ConfidentialTransfer_InitializeMint:      {"InitializeMint", func() ConfidentialTransferInstructionData { return &ConfidentialTransferInitializeMintData{} }},
	ConfidentialTransfer_UpdateMint:          {"UpdateMint", func() ConfidentialTransferInstructionData { return &ConfidentialTransferUpdateMintData{} }},
	ConfidentialTransfer_ConfigureAccount:    {"ConfigureAccount", func() ConfidentialTransferInstructionData { return &ConfidentialTransferConfigureAccountData{} }},
	ConfidentialTransfer_ApproveAccount:      {"ApproveAccount", func() ConfidentialTransferInstructionData { return &ConfidentialTransferApproveAccountData{} }},
	ConfidentialTransfer_EmptyAccount:        {"EmptyAccount", func() ConfidentialTransferInstructionData { return &ConfidentialTransferEmptyAccountData{} }},
	ConfidentialTransfer_Deposit:             {"Deposit", func() ConfidentialTransferInstructionData { return &ConfidentialTransferDepositData{} }},
	ConfidentialTransfer_Withdraw:            {"Withdraw", func() ConfidentialTransferInstructionData { return &ConfidentialTransferWithdrawData{} }},
	ConfidentialTransfer_Transfer:            {"Transfer", func() ConfidentialTransferInstructionData { return &ConfidentialTransferTransferData{} }},
	ConfidentialTransfer_ApplyPendingBalance: {"ApplyPendingBalance", func() ConfidentialTransferInstructionData { return &ConfidentialTransferApplyPendingBalanceData{} }},
	ConfidentialTransfer_EnableConfidentialCredits: {"EnableConfidentialCredits", func() ConfidentialTransferInstructionData {
		return &ConfidentialTransferEnableConfidentialCreditsData{}
	}},
	ConfidentialTransfer_DisableConfidentialCredits: {"DisableConfidentialCredits", func() ConfidentialTransferInstructionData {
		return &ConfidentialTransferDisableConfidentialCreditsData{}
	}},
	ConfidentialTransfer_EnableNonConfidentialCredits: {"EnableNonConfidentialCredits", func() ConfidentialTransferInstructionData {
		return &ConfidentialTransferEnableNonConfidentialCreditsData{}
	}},
	ConfidentialTransfer_DisableNonConfidentialCredits: {"DisableNonConfidentialCredits", func() ConfidentialTransferInstructionData {
		return &ConfidentialTransferDisableNonConfidentialCreditsData{}
	}},
	ConfidentialTransfer_TransferWithFee: {"TransferWithFee", func() ConfidentialTransferInstructionData { return &ConfidentialTransferTransferWithFeeData{} }},
	ConfidentialTransfer_ConfigureAccountWithRegistry: {"ConfigureAccountWithRegistry", func() ConfidentialTransferInstructionData {
		return &ConfidentialTransferConfigureAccountWithRegistryData{}
	}},
}

// ConfidentialTransferExtension is the instruction wrapper for the ConfidentialTransfer extension (ID 27).
// This is a complex extension with many sub-instructions involving zero-knowledge proofs.
type ConfidentialTransferExtension struct {
	SubInstruction uint8
	// Raw data for the sub-instruction.
	RawData []byte

	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (obj *ConfidentialTransferExtension) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	obj.Accounts = ag_solanago.AccountMetaSlice(accounts)
	return nil
}

func (slice ConfidentialTransferExtension) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, slice.Accounts...)
	accounts = append(accounts, slice.Signers...)
	return
}

// DecodeSubInstructionData outputs the typed instruction data for SubInstruction.
func (obj ConfidentialTransferExtension) DecodeSubInstructionData() (ConfidentialTransferInstructionData, error) {
	if int(obj.SubInstruction) >= len(ctSubInstructions) {
		return nil, fmt.Errorf("token2022: unknown ConfidentialTransfer sub-instruction %d", obj.SubInstruction)
	}
	instructionData := ctSubInstructions[obj.SubInstruction].newData()
	if err := instructionData.UnmarshalBinary(obj.RawData); err != nil {
		return nil, err
	}
	return instructionData, nil
}

// subInstructionName is the name of SubInstruction, or "Unknown" for an ID the program does not define.
func (obj ConfidentialTransferExtension) subInstructionName() string {
	if int(obj.SubInstruction) >= len(ctSubInstructions) {
		return "Unknown"
	}
	return ctSubInstructions[obj.SubInstruction].name
}

func (inst ConfidentialTransferExtension) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   &inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ConfidentialTransferExtension),
	}}
}

func (inst ConfidentialTransferExtension) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfidentialTransferExtension) Validate() error {
	if len(inst.Accounts) == 0 {
		return errors.New("accounts is empty")
	}
	return nil
}

func (inst *ConfidentialTransferExtension) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfidentialTransfer." + inst.subInstructionName())).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if ctData, err := inst.DecodeSubInstructionData(); err == nil {
							paramsBranch.Child(ag_format.Param("Data", ctData))
						} else {
							// Fall back to the payload length for malformed instruction data.
							paramsBranch.Child(ag_format.Param("RawData (len)", len(inst.RawData)))
						}
					})
				})
		})
}

func (obj ConfidentialTransferExtension) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	err = encoder.WriteUint8(obj.SubInstruction)
	if err != nil {
		return err
	}
	if len(obj.RawData) > 0 {
		err = encoder.WriteBytes(obj.RawData, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ConfidentialTransferExtension) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	obj.SubInstruction, err = decoder.ReadUint8()
	if err != nil {
		return err
	}
	remaining := decoder.Remaining()
	if remaining > 0 {
		obj.RawData, err = decoder.ReadNBytes(remaining)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewConfidentialTransferInstruction creates a confidential transfer extension
// instruction from an raw sub-instruction payload.
//
// Prefer the typed NewConfidentialTransfer*Instruction builders.
func NewConfidentialTransferInstruction(
	subInstruction uint8,
	rawData []byte,
	accounts ...ag_solanago.AccountMeta,
) *ConfidentialTransferExtension {
	inst := &ConfidentialTransferExtension{
		SubInstruction: subInstruction,
		RawData:        rawData,
		Accounts:       make(ag_solanago.AccountMetaSlice, len(accounts)),
		Signers:        make(ag_solanago.AccountMetaSlice, 0),
	}
	for i := range accounts {
		inst.Accounts[i] = &accounts[i]
	}
	return inst
}
