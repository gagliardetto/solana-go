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

package token

import (
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type GetAccountDataSize struct {
	// [0] = [WRITE] mint
	// ··········· The mint.
	Mint *ag_solanago.AccountMeta `bin:"-" borsh_skip:"true"`
}

func (inst *GetAccountDataSize) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	if len(accounts) != 1 {
		return fmt.Errorf("expected 1 account, got %v", len(accounts))
	}
	inst.Mint = accounts[0]
	return nil
}

func (inst GetAccountDataSize) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Mint)
	return
}

func (inst GetAccountDataSize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (inst *GetAccountDataSize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

func NewGetAccountDataSizeInstructionBuilder() *GetAccountDataSize {
	return &GetAccountDataSize{}
}

func (inst *GetAccountDataSize) SetMint(mint ag_solanago.PublicKey) *GetAccountDataSize {
	inst.Mint = ag_solanago.Meta(mint).WRITE()
	return inst
}

func (inst *GetAccountDataSize) GetMint() *ag_solanago.AccountMeta {
	return inst.Mint
}

func (inst GetAccountDataSize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_GetAccountDataSize),
	}}
}

func (inst GetAccountDataSize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *GetAccountDataSize) Validate() error {
	if inst.Mint == nil {
		return errors.New("accounts.Mint is not set")
	}
	return nil
}

func (inst *GetAccountDataSize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("GetAccountDataSize")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("mint", inst.Mint))
					})
				})
		})
}
