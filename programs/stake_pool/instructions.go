// Copyright 2021 github.com/gagliardetto
// Copyright 2025 github.com/liquid-collective
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stakepool

import (
	"bytes"
	"fmt"

	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
)

var ProgramID = ag_solanago.SPLStakePoolProgramID

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "StakePool"

func init() {
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const (
	Instruction_Initialize                       uint8 = 0
	Instruction_AddValidatorToPool               uint8 = 1
	Instruction_RemoveValidatorFromPool          uint8 = 2
	Instruction_DecreaseValidatorStake           uint8 = 3
	Instruction_IncreaseValidatorStake           uint8 = 4
	Instruction_SetPreferredValidator            uint8 = 5
	Instruction_UpdateValidatorListBalance       uint8 = 6
	Instruction_UpdateStakePoolBalance           uint8 = 7
	Instruction_CleanupRemovedValidatorEntries   uint8 = 8
	Instruction_DepositStake                     uint8 = 9
	Instruction_WithdrawStake                    uint8 = 10
	Instruction_SetManager                       uint8 = 11
	Instruction_SetFee                           uint8 = 12
	Instruction_SetStaker                        uint8 = 13
	Instruction_DecreaseAdditionalValidatorStake uint8 = 14
	Instruction_IncreaseAdditionalValidatorStake uint8 = 15
	Instruction_Redelegate                       uint8 = 16
	Instruction_SetFundingAuthority              uint8 = 17
	Instruction_DepositSol                       uint8 = 18
	Instruction_WithdrawSol                      uint8 = 19
	Instruction_CreateTokenMetadata              uint8 = 20
	Instruction_UpdateTokenMetadata              uint8 = 21
	Instruction_DepositStakeWithSlippage         uint8 = 22
	Instruction_WithdrawStakeWithSlippage        uint8 = 23
	Instruction_DepositSolWithSlippage           uint8 = 24
	Instruction_WithdrawSolWithSlippage          uint8 = 25
)

func InstructionIDToName(id uint8) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_AddValidatorToPool:
		return "AddValidatorToPool"
	case Instruction_RemoveValidatorFromPool:
		return "RemoveValidatorFromPool"
	case Instruction_DecreaseValidatorStake:
		return "DecreaseValidatorStake"
	case Instruction_IncreaseValidatorStake:
		return "IncreaseValidatorStake"
	case Instruction_SetPreferredValidator:
		return "SetPreferredValidator"
	case Instruction_UpdateValidatorListBalance:
		return "UpdateValidatorListBalance"
	case Instruction_UpdateStakePoolBalance:
		return "UpdateStakePoolBalance"
	case Instruction_CleanupRemovedValidatorEntries:
		return "CleanupRemovedValidatorEntries"
	case Instruction_DepositStake:
		return "DepositStake"
	case Instruction_WithdrawStake:
		return "WithdrawStake"
	case Instruction_SetManager:
		return "SetManager"
	case Instruction_SetFee:
		return "SetFee"
	case Instruction_SetStaker:
		return "SetStaker"
	case Instruction_DecreaseAdditionalValidatorStake:
		return "DecreaseAdditionalValidatorStake"
	case Instruction_IncreaseAdditionalValidatorStake:
		return "IncreaseAdditionalValidatorStake"
	case Instruction_Redelegate:
		return "Redelegate"
	case Instruction_SetFundingAuthority:
		return "SetFundingAuthority"
	case Instruction_DepositSol:
		return "DepositSol"
	case Instruction_WithdrawSol:
		return "WithdrawSol"
	case Instruction_CreateTokenMetadata:
		return "CreateTokenMetadata"
	case Instruction_UpdateTokenMetadata:
		return "UpdateTokenMetadata"
	case Instruction_DepositStakeWithSlippage:
		return "DepositStakeWithSlippage"
	case Instruction_WithdrawStakeWithSlippage:
		return "WithdrawStakeWithSlippage"
	case Instruction_DepositSolWithSlippage:
		return "DepositSolWithSlippage"
	case Instruction_WithdrawSolWithSlippage:
		return "WithdrawSolWithSlippage"
	default:
		return "Unknown"
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.Uint8TypeIDEncoding,
	[]ag_binary.VariantType{

		{Name: "Initialize", Type: (*Initialize)(nil)},
		{Name: "AddValidatorToPool", Type: (*AddValidatorToPool)(nil)},
		{Name: "RemoveValidatorFromPool", Type: (*RemoveValidatorFromPool)(nil)},
		{Name: "DecreaseValidatorStake", Type: (*DecreaseValidatorStake)(nil)},
		{Name: "IncreaseValidatorStake", Type: (*IncreaseValidatorStake)(nil)},
		{Name: "SetPreferredValidator", Type: (*SetPreferredValidator)(nil)},
		{Name: "UpdateValidatorListBalance", Type: (*UpdateValidatorListBalance)(nil)},
		{Name: "UpdateStakePoolBalance", Type: (*UpdateStakePoolBalance)(nil)},
		{Name: "CleanupRemovedValidatorEntries", Type: (*CleanupRemovedValidatorEntries)(nil)},
		{Name: "DepositStake", Type: (*DepositStake)(nil)},
		{Name: "WithdrawStake", Type: (*WithdrawStake)(nil)},
		{Name: "SetManager", Type: (*SetManager)(nil)},
		{Name: "SetFee", Type: (*SetFee)(nil)},
		{Name: "SetStaker", Type: (*SetStaker)(nil)},
		{Name: "DecreaseAdditionalValidatorStake", Type: (*DecreaseAdditionalValidatorStake)(nil)},
		{Name: "IncreaseAdditionalValidatorStake", Type: (*IncreaseAdditionalValidatorStake)(nil)},
		{Name: "Redelegate", Type: (*Redelegate)(nil)},
		{Name: "SetFundingAuthority", Type: (*SetFundingAuthority)(nil)},
		{Name: "DepositSol", Type: (*DepositSol)(nil)},
		{Name: "WithdrawSol", Type: (*WithdrawSol)(nil)},
		{Name: "CreateTokenMetadata", Type: (*CreateTokenMetadata)(nil)},
		{Name: "UpdateTokenMetadata", Type: (*UpdateTokenMetadata)(nil)},
		{Name: "DepositStakeWithSlippage", Type: (*DepositStakeWithSlippage)(nil)},
		{Name: "WithdrawStakeWithSlippage", Type: (*WithdrawStakeWithSlippage)(nil)},
		{Name: "DepositSolWithSlippage", Type: (*DepositSolWithSlippage)(nil)},
		{Name: "WithdrawSolWithSlippage", Type: (*WithdrawSolWithSlippage)(nil)},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBinEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteUint8(inst.TypeID.Uint8())
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBinDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}

func (inst *Instruction) DecodeInstructionData() (instructionData interface{}, instructionType uint8, err error) {
	data, err := inst.Data()
	if err != nil {
		return nil, 0, err
	}

	if err := ag_binary.NewBinDecoder(data).Decode(&instructionType); err != nil {
		return nil, 0, err
	}

	switch instructionType {
	case Instruction_Initialize:
		instructionData = new(Initialize)
	case Instruction_AddValidatorToPool:
		instructionData = new(AddValidatorToPool)
	case Instruction_RemoveValidatorFromPool:
		instructionData = new(RemoveValidatorFromPool)
	case Instruction_DecreaseValidatorStake:
		instructionData = new(DecreaseValidatorStake)
	case Instruction_IncreaseValidatorStake:
		instructionData = new(IncreaseValidatorStake)
	case Instruction_SetPreferredValidator:
		instructionData = new(SetPreferredValidator)
	case Instruction_UpdateValidatorListBalance:
		instructionData = new(UpdateValidatorListBalance)
	case Instruction_UpdateStakePoolBalance:
		instructionData = new(UpdateStakePoolBalance)
	case Instruction_CleanupRemovedValidatorEntries:
		instructionData = new(CleanupRemovedValidatorEntries)
	case Instruction_DepositStake:
		instructionData = new(DepositStake)
	case Instruction_WithdrawStake:
		instructionData = new(WithdrawStake)
	case Instruction_SetManager:
		instructionData = new(SetManager)
	case Instruction_SetFee:
		instructionData = new(SetFee)
	case Instruction_SetStaker:
		instructionData = new(SetStaker)
	case Instruction_DecreaseAdditionalValidatorStake:
		instructionData = new(DecreaseAdditionalValidatorStake)
	case Instruction_IncreaseAdditionalValidatorStake:
		instructionData = new(IncreaseAdditionalValidatorStake)
	case Instruction_Redelegate:
		instructionData = new(Redelegate)
	case Instruction_SetFundingAuthority:
		instructionData = new(SetFundingAuthority)
	case Instruction_DepositSol:
		instructionData = new(DepositSol)
	case Instruction_WithdrawSol:
		instructionData = new(WithdrawSol)
	case Instruction_CreateTokenMetadata:
		instructionData = new(CreateTokenMetadata)
	case Instruction_UpdateTokenMetadata:
		instructionData = new(UpdateTokenMetadata)
	case Instruction_DepositStakeWithSlippage:
		instructionData = new(DepositStakeWithSlippage)
	case Instruction_WithdrawStakeWithSlippage:
		instructionData = new(WithdrawStakeWithSlippage)
	case Instruction_DepositSolWithSlippage:
		instructionData = new(DepositSolWithSlippage)
	case Instruction_WithdrawSolWithSlippage:
		instructionData = new(WithdrawSolWithSlippage)
	default:
		return nil, instructionType, fmt.Errorf("unknown instruction type: %v", instructionType)
	}

	if v, ok := instructionData.(ag_solanago.AccountsSettable); ok {
		_ = v.SetAccounts(inst.Accounts())
	}

	if err := ag_binary.NewBinDecoder(data[1:]).Decode(instructionData); err != nil {
		return nil, instructionType, err
	}

	return instructionData, instructionType, nil
}
