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

package vote

import (
	"fmt"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
	"github.com/gagliardetto/treeout"
)

type Vote struct {
	Slots     []uint64
	Hash      solana.Hash
	Timestamp *int64

	// [0] = [WRITE] VoteAccount
	// ··········· Vote account to vote with
	//
	// [1] = [] SysVarSlotHashes
	// ··········· Slot hashes sysvar
	//
	// [2] = [] SysVarClock
	// ··········· Clock sysvar
	//
	// [3] = [SIGNER] VoteAuthority
	// ··········· Vote authority
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (v *Vote) UnmarshalWithDecoder(dec *bin.Decoder) error {
	v.Slots = nil
	var numSlots uint64
	if err := dec.Decode(&numSlots); err != nil {
		return err
	}
	for i := uint64(0); i < numSlots; i++ {
		var slot uint64
		if err := dec.Decode(&slot); err != nil {
			return err
		}
		v.Slots = append(v.Slots, slot)
	}
	if err := dec.Decode(&v.Hash); err != nil {
		return err
	}
	var timestampVariant uint8
	if err := dec.Decode(&timestampVariant); err != nil {
		return err
	}
	switch timestampVariant {
	case 0:
		break
	case 1:
		var ts int64
		if err := dec.Decode(&ts); err != nil {
			return err
		}
		v.Timestamp = &ts
	default:
		return fmt.Errorf("invalid vote timestamp variant %#08x", timestampVariant)
	}
	return nil
}

func (inst *Vote) Validate() error {
	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}

func (inst *Vote) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Vote")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("Slots", fmt.Sprintf("%v", inst.Slots)))
						paramsBranch.Child(format.Param("Hash", inst.Hash.String()))
						var ts time.Time
						if inst.Timestamp != nil {
							ts = time.Unix(*inst.Timestamp, 0).UTC()
						}
						paramsBranch.Child(format.Param("Timestamp", ts))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("Vote Account      ", inst.AccountMetaSlice[0]))
						accountsBranch.Child(format.Meta("Slot Hashes Sysvar", inst.AccountMetaSlice[1]))
						accountsBranch.Child(format.Meta("Clock Sysvar      ", inst.AccountMetaSlice[2]))
						accountsBranch.Child(format.Meta("Vote Authority    ", inst.AccountMetaSlice[3]))
					})
				})
		})
}
