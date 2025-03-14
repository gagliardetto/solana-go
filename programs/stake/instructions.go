// Copyright 2021 github.com/gagliardetto
// Copyright 2024 github.com/cordialsys
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

package stake

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/treeout"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
)

var ProgramID solana.PublicKey = solana.StakeProgramID

func SetProgramID(pubkey solana.PublicKey) {
	ProgramID = pubkey
	solana.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Stake"

func init() {
	solana.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const (
	// Initializes a new stake account
	Instruction_Initialize uint32 = iota
	// Authorize a key to manage stake or withdrawal
	Instruction_Authorize
	// Delegate a stake account to a validator vote account
	Instruction_DelegateStake
	// Split an active stake account into a new stake account
	Instruction_Split
	// Withdraw unstaked lamports from the stake account
	Instruction_Withdraw
	// Deactivates the stake in the account
	Instruction_Deactivate
	// Sets the lockup for the stake account
	Instruction_SetLockup
	// Merges two stake accounts
	Instruction_Merge
	// Authorize a key to manage stake or withdrawal with seed
	Instruction_AuthorizeWithSeed
	// Initializes a new stake account with checked authorities
	Instruction_InitializeChecked
	// Authorize a key to manage stake or withdrawal with checked authorities
	Instruction_AuthorizeChecked
	// Authorize a key to manage stake or withdrawal with checked authorities and seed
	Instruction_AuthorizeCheckedWithSeed
	// Sets the lockup for the stake account with checked authorities
	Instruction_SetLockupChecked
	// Gets the minimum delegation for the stake account
	Instruction_GetMinimumDelegation
	// Deactivates delinquent stake accounts
	Instruction_DeactivateDelinquent
	// Moves stake from one account to another
	Instruction_MoveStake
	// Moves lamports from one account to another
	Instruction_MoveLamports
)

func InstructionIDToName(id uint32) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_Authorize:
		return "Authorize"
	case Instruction_DelegateStake:
		return "DelegateStake"
	case Instruction_Split:
		return "Split"
	case Instruction_Withdraw:
		return "Withdraw"
	case Instruction_Deactivate:
		return "Deactivate"
	case Instruction_SetLockup:
		return "SetLockup"
	case Instruction_Merge:
		return "Merge"
	case Instruction_AuthorizeWithSeed:
		return "AuthorizeWithSeed"
	case Instruction_InitializeChecked:
		return "InitializeChecked"
	case Instruction_AuthorizeChecked:
		return "AuthorizeChecked"
	case Instruction_AuthorizeCheckedWithSeed:
		return "AuthorizeCheckedWithSeed"
	case Instruction_SetLockupChecked:
		return "SetLockupChecked"
	case Instruction_GetMinimumDelegation:
		return "GetMinimumDelegation"
	case Instruction_DeactivateDelinquent:
		return "DeactivateDelinquent"
	case Instruction_MoveStake:
		return "MoveStake"
	case Instruction_MoveLamports:
		return "MoveLamports"
	default:
		return ""
	}
}

type Instruction struct {
	bin.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent treeout.Branches) {
	if enToTree, ok := inst.Impl.(text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(spew.Sdump(inst))
	}
}

var InstructionImplDef = bin.NewVariantDefinition(
	bin.Uint32TypeIDEncoding,
	[]bin.VariantType{
		{
			"Initialize", (*Initialize)(nil),
		},
		{
			"Authorize", (*Authorize)(nil),
		},
		{
			"DelegateStake", (*DelegateStake)(nil),
		},
		{
			"Split", (*Split)(nil),
		},
		{
			"Withdraw", (*Withdraw)(nil),
		},
		{
			"Deactivate", (*Deactivate)(nil),
		},
		{
			"SetLockup", (*SetLockup)(nil),
		},
		{
			"Merge", (*Merge)(nil),
		},
		{
			"AuthorizeWithSeed", (*AuthorizeWithSeed)(nil),
		},
		{
			"InitializeChecked", (*InitializeChecked)(nil),
		},
		{
			"AuthorizeChecked", (*AuthorizeChecked)(nil),
		},
		{
			"AuthorizeCheckedWithSeed", (*AuthorizeCheckedWithSeed)(nil),
		},
		{
			"SetLockupChecked", (*SetLockupChecked)(nil),
		},
		{
			"GetMinimumDelegation", (*GetMinimumDelegation)(nil),
		},
		{
			"DeactivateDelinquent", (*DeactivateDelinquent)(nil),
		},
		{
			"MoveStake", (*MoveStake)(nil),
		},
		{
			"MoveLamports", (*MoveLamports)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() solana.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*solana.AccountMeta) {
	return inst.Impl.(solana.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewBinEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst Instruction) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(inst.TypeID.Uint32(), binary.LittleEndian)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := bin.NewBinDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(solana.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
