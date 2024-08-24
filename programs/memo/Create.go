// Copyright 2021 github.com/gagliardetto
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

package memo

import (
	"errors"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

type Create struct {
	// The memo message
	Message []byte

	// [0] = [SIGNER] Signer
	// ··········· The account that will pay for the transaction
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewMemoInstructionBuilder creates a new `Memo` instruction builder.
func NewMemoInstructionBuilder() *Create {
	nd := &Create{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetMessage sets the memo message
func (inst *Create) SetMessage(message []byte) *Create {
	inst.Message = message
	return inst
}

// SetSigner sets the signer account
func (inst *Create) SetSigner(signer ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

func (inst *Create) GetSigner() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

func (inst Create) Build() *MemoInstruction {

	return &MemoInstruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.NoTypeIDDefaultID,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Create) ValidateAndBuild() (*MemoInstruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Create) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if len(inst.Message) == 0 {
			return errors.New("Message not set")
		}
	}

	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}
func (inst *Create) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program("Memo", ag_solanago.MemoProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Create")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("Signer", inst.AccountMetaSlice[0]))
					})
				})
		})
}

func (inst Create) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	// Serialize `Message` param:
	{
		err := encoder.Encode(inst.Message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *Create) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	// Deserialize `Message` param:
	{
		err := decoder.Decode(&inst.Message)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewMemoInstruction declares a new Memo instruction with the provided parameters and accounts.
func NewMemoInstruction(
	// Parameters:
	message []byte,
	// Accounts:
	signer ag_solanago.PublicKey) *Create {
	return NewMemoInstructionBuilder().
		SetMessage(message).
		SetSigner(signer)
}
