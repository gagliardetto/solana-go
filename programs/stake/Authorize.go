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

package stake

import (
	"errors"
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type Authorize struct {
	NewAuthority  *ag_solanago.PublicKey
	AuthorityType StakeAuthorize

	// [0] = [WRITE] StakeAccount
	// ··········· The stake account to authorize.
	//
	// [1] = [] ClockSysvar
	// ··········· Clock sysvar account.
	//
	// [2] = [] StakeOrWithdrawAuthority
	// ··········· Stake or withdraw authority.
	//
	// [3...] = [SIGNER] Signers
	// ··········· M signer accounts.
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *Authorize) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.Accounts, inst.Signers = ag_solanago.AccountMetaSlice(accounts).SplitFrom(3)
	return nil
}

func (inst Authorize) GetAccounts() (accounts []*ag_solanago.AccountMeta) {
	accounts = append(accounts, inst.Accounts...)
	accounts = append(accounts, inst.Signers...)
	return
}

func (inst *Authorize) Validate() error {
	// Check whether all (required) parameters are set:
	if inst.NewAuthority == nil {
		return errors.New("NewAuthority parameter is not set")
	}

	// Check whether all (required) accounts are set:
	if inst.Accounts[0] == nil {
		return errors.New("accounts.StakeAccount is not set")
	}
	if inst.Accounts[1] == nil {
		return errors.New("accounts.ClockSysvar is not set")
	}
	if inst.Accounts[2] == nil {
		return errors.New("accounts.StakeOrWithdrawAuthority is not set")
	}
	if !inst.Accounts[2].IsSigner && len(inst.Signers) == 0 {
		return fmt.Errorf("accounts.Signers is not set")
	}
	if len(inst.Signers) > MAX_SIGNERS {
		return fmt.Errorf("too many signers; got %v, but max is 11", len(inst.Signers))
	}
	return nil
}

func (inst *Authorize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Authorize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewAuthority", *inst.NewAuthority))
						paramsBranch.Child(ag_format.Param("AuthorityType", inst.AuthorityType))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("StakeAccount", inst.Accounts[0]))
						accountsBranch.Child(ag_format.Meta("ClockSysvar", inst.Accounts[1]))
						accountsBranch.Child(ag_format.Meta("StakeOrWithdrawAuthority", inst.Accounts[2]))

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(inst.Signers)))
						for i, v := range inst.Signers {
							if len(inst.Signers) > 9 && i < 10 {
								signersBranch.Child(ag_format.Meta(fmt.Sprintf(" [%v]", i), v))
							} else {
								signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), v))
							}
						}
					})
				})
		})
}

func (inst Authorize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAuthority` param:
	err = encoder.Encode(inst.NewAuthority)
	if err != nil {
		return err
	}
	// Serialize `AuthorityType` param:
	err = encoder.Encode(inst.AuthorityType)
	if err != nil {
		return err
	}
	return nil
}

func (inst *Authorize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAuthority`:
	err = decoder.Decode(&inst.NewAuthority)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityType`:
	err = decoder.Decode(&inst.AuthorityType)
	if err != nil {
		return err
	}
	return nil
}

func (inst Authorize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_Authorize, ag_binary.LE),
	}}
}

func NewAuthorizeInstruction(
	newAuthority ag_solanago.PublicKey,
	authorityType StakeAuthorize,
	stakeAccount ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	stakeOrWithdrawAuthority ag_solanago.PublicKey,
	multisigSigners []ag_solanago.PublicKey,
) *Authorize {
	return &Authorize{
		NewAuthority:  &newAuthority,
		AuthorityType: authorityType,
		Accounts: ag_solanago.AccountMetaSlice{
			ag_solanago.Meta(stakeAccount).WRITE(),
			ag_solanago.Meta(clockSysvar),
			ag_solanago.Meta(stakeOrWithdrawAuthority),
		},
		Signers: func(keys []ag_solanago.PublicKey) ag_solanago.AccountMetaSlice {
			metas := make(ag_solanago.AccountMetaSlice, len(keys))
			for i, key := range keys {
				metas[i] = ag_solanago.Meta(key).SIGNER()
			}
			return metas
		}(multisigSigners),
	}
}
