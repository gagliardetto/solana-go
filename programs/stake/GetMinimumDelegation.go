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

package stake

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_format "github.com/gagliardetto/solana-go/text/format"
)

// GetMinimumDelegation instruction
type GetMinimumDelegation struct {
}

func (inst GetMinimumDelegation) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst GetMinimumDelegation) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_GetMinimumDelegation, ag_binary.LE),
	}}
}

func (inst *GetMinimumDelegation) Validate() error {
	return nil
}

func (inst *GetMinimumDelegation) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("GetMinimumDelegation")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

				})
		})
}

func (inst GetMinimumDelegation) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (inst *GetMinimumDelegation) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewGetMinimumDelegationInstruction declares a new GetMinimumDelegation instruction with the provided accounts.
func NewGetMinimumDelegationInstruction() *GetMinimumDelegation {
	return &GetMinimumDelegation{}
}
