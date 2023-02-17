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

type SetComputeUnitLimit struct {
	UnitLimit *uint32
}

// NewSetComputeUnitLimitInstructionBuilder creates a new `SetComputeUnitLimit` instruction builder.
func NewSetComputeUnitLimitInstructionBuilder() *SetComputeUnitLimit {
	nd := &SetComputeUnitLimit{}
	return nd
}

// Unit limit
func (inst *SetComputeUnitLimit) SetUnitLimit(unitLimit uint32) *SetComputeUnitLimit {
	inst.UnitLimit = &unitLimit
	return inst
}

func (inst SetComputeUnitLimit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SetComputeUnitLimit),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetComputeUnitLimit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetComputeUnitLimit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UnitLimit == nil {
			return errors.New("UnitLimit parameter is not set")
		}
	}
	return nil
}

func (inst *SetComputeUnitLimit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetComputeUnitLimit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UnitLimit", *inst.UnitLimit))
					})
				})
		})
}

func (obj SetComputeUnitLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UnitLimit` param:
	err = encoder.Encode(obj.UnitLimit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetComputeUnitLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UnitLimit`:
	err = decoder.Decode(&obj.UnitLimit)
	if err != nil {
		return err
	}
	return nil
}

// NewSetComputeUnitLimitInstruction declares a new SetComputeUnitLimit instruction with the provided parameters and accounts.
func NewSetComputeUnitLimitInstruction(
	// Parameters:
	unitLimit uint32,
) *SetComputeUnitLimit {
	return NewSetComputeUnitLimitInstructionBuilder().SetUnitLimit(unitLimit)
}
