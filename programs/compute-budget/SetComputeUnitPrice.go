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

package computebudget

import (
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

type SetComputeUnitPrice struct {
	UnitPrice *uint32
}

// NewSetComputeUnitPriceInstructionBuilder creates a new `SetComputeUnitPrice` instruction builder.
func NewSetComputeUnitPriceInstructionBuilder() *SetComputeUnitPrice {
	nd := &SetComputeUnitPrice{}
	return nd
}

func (inst *SetComputeUnitPrice) SetUnitPrice(unitPrice uint32) *SetComputeUnitPrice {
	inst.UnitPrice = &unitPrice
	return inst
}

func (inst SetComputeUnitPrice) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SetComputeUnitPrice),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetComputeUnitPrice) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetComputeUnitPrice) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UnitPrice == nil {
			return errors.New("UnitPrice parameter is not set")
		}
	}
	return nil
}

func (inst *SetComputeUnitPrice) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetComputeUnitPrice")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UnitPrice", *inst.UnitPrice))
					})
				})
		})
}

func (obj SetComputeUnitPrice) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UnitPrice` param:
	err = encoder.Encode(obj.UnitPrice)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetComputeUnitPrice) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UnitPrice`:
	err = decoder.Decode(&obj.UnitPrice)
	if err != nil {
		return err
	}
	return nil
}

// NewSetComputeUnitPriceInstruction declares a new SetComputeUnitPrice instruction with the provided parameters and accounts.
func NewSetComputeUnitPriceInstruction(
	// Parameters:
	unitPrice uint32,
) *SetComputeUnitPrice {
	return NewSetComputeUnitPriceInstructionBuilder().SetUnitPrice(unitPrice)
}
