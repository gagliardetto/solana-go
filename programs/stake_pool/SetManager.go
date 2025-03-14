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
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type SetManager struct {
	// [0] = [WRITE] stakePool
	// [1] = [SIGNER] manager
	// [2] = [SIGNER] newManager
	// [3] = [] newManagerFeeAccount
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewSetManagerInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	manager ag_solanago.PublicKey,
	newManager ag_solanago.PublicKey,
	newManagerFeeAccount ag_solanago.PublicKey,
) *SetManager {
	return NewSetManagerInstructionBuilder().
		SetStakePool(stakePool).
		SetManager(manager).
		SetNewManager(newManager).
		SetNewManagerFeeAccount(newManagerFeeAccount)
}

func NewSetManagerInstructionBuilder() *SetManager {
	return &SetManager{
		Accounts: make(ag_solanago.AccountMetaSlice, 4),
		Signers:  make(ag_solanago.AccountMetaSlice, 2),
	}
}

func (inst *SetManager) SetStakePool(pool ag_solanago.PublicKey) *SetManager {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *SetManager) SetManager(manager ag_solanago.PublicKey) *SetManager {
	inst.Accounts[1] = ag_solanago.Meta(manager).SIGNER()
	return inst
}

func (inst *SetManager) SetNewManager(newManager ag_solanago.PublicKey) *SetManager {
	inst.Accounts[2] = ag_solanago.Meta(newManager).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(newManager).SIGNER()
	return inst
}

func (inst *SetManager) SetNewManagerFeeAccount(newManagerFeeAccount ag_solanago.PublicKey) *SetManager {
	inst.Accounts[3] = ag_solanago.Meta(newManagerFeeAccount)
	return inst
}

func (inst *SetManager) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *SetManager) GetManager() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *SetManager) GetNewManager() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *SetManager) GetNewManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *SetManager) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetManager) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_SetManager),
			Impl:   inst,
		},
	}
}

func (inst *SetManager) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetManager")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range inst.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(inst.Signers)))
						for j, signer := range inst.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (inst *SetManager) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *SetManager) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *SetManager) Validate() error {
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.Manager should be a signer")
	}
	return nil
}
