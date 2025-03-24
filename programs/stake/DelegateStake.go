// Copyright 2024 github.com/cordialsys
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stake

import (
	"fmt"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type DelegateStake struct {
	// [0] = [WRITE SIGNER] StakeAccount
	// ··········· Stake account getting initialized
	//
	// [1] = [] Vote Account
	// ··········· The validator vote account being delegated to
	//
	// [2] = [] Clock Sysvar
	// ··········· The Clock Sysvar Account
	//
	// [3] = [] Stake History Sysvar
	// ··········· The Stake History Sysvar Account
	//
	// [4] = [] Stake Config Account
	// ··········· The Stake Config Account
	//
	// [5] = [WRITE SIGNER] Stake Authoriy
	// ··········· The Stake Authority
	//
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func (inst *DelegateStake) Validate() error {
	// Check whether all accounts are set:
	for accIndex, acc := range inst.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is not set", accIndex)
		}
	}
	return nil
}
func (inst *DelegateStake) SetStakeAccount(stakeAccount ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(stakeAccount).WRITE().SIGNER()
	return inst
}
func (inst *DelegateStake) SetVoteAccount(voteAcc ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(voteAcc)
	return inst
}
func (inst *DelegateStake) SetClockSysvar(clockSysVarAcc ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(clockSysVarAcc)
	return inst
}
func (inst *DelegateStake) SetStakeHistorySysvar(stakeHistorySysVarAcc ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stakeHistorySysVarAcc)
	return inst
}
func (inst *DelegateStake) SetConfigAccount(stakeConfigAcc ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stakeConfigAcc)
	return inst
}
func (inst *DelegateStake) SetStakeAuthority(stakeAuthority ag_solanago.PublicKey) *DelegateStake {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(stakeAuthority).WRITE().SIGNER()
	return inst
}
func (inst *DelegateStake) GetStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}
func (inst *DelegateStake) GetVoteAccount() *ag_solanago.AccountMeta { return inst.AccountMetaSlice[1] }
func (inst *DelegateStake) GetClockSysvar() *ag_solanago.AccountMeta { return inst.AccountMetaSlice[2] }
func (inst *DelegateStake) GetStakeHistorySysvar() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}
func (inst *DelegateStake) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}
func (inst *DelegateStake) GetStakeAuthority() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

func (inst DelegateStake) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint32(Instruction_DelegateStake, ag_binary.LE),
	}}
}

func (inst *DelegateStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DelegateStake")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("StakeAccount", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("VoteAccount", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("ClockSysvar", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("StakeHistorySysvar", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("StakeConfigAccount", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("StakeAuthoriy", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

// NewDelegateStakeInstructionBuilder creates a new `DelegateStake` instruction builder.
func NewDelegateStakeInstructionBuilder() *DelegateStake {
	nd := &DelegateStake{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// NewDelegateStakeInstruction declares a new DelegateStake instruction with the provided parameters and accounts.
func NewDelegateStakeInstruction(
	// Accounts:
	validatorVoteAccount ag_solanago.PublicKey,
	stakeAuthority ag_solanago.PublicKey,
	stakeAccount ag_solanago.PublicKey,
) *DelegateStake {
	return NewDelegateStakeInstructionBuilder().
		SetStakeAccount(stakeAccount).
		SetVoteAccount(validatorVoteAccount).
		SetClockSysvar(ag_solanago.SysVarClockPubkey).
		SetStakeHistorySysvar(ag_solanago.SysVarStakeHistoryPubkey).
		SetConfigAccount(ag_solanago.SysVarStakeConfigPubkey).
		SetStakeAuthority(stakeAuthority)
}
