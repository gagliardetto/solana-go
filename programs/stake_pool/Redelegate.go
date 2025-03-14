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

type Redelegate struct {
	Lamports                      *uint64
	SourceTransientStakeSeed      *uint64
	EphemeralStakeSeed            *uint64
	DestinationTransientStakeSeed *uint64
	// [0] = [] stakePool
	// [1] = [SIGNER] staker
	// [2] = [] withdrawAuthority
	// [3] = [WRITE] validatorList
	// [4] = [WRITE] sourceCanonicalStakeAccount
	// [5] = [WRITE] sourceTransientStakeAccount
	// [6] = [WRITE] uninitializedEphemeralStakeAccount
	// [7] = [WRITE] destinationTransientStakeAccount
	// [8] = [WRITE] destinationStakeAccount
	// [9] = [] destinationValidatorVoteAccount
	// [10] = [] clockSysvar
	// [11] = [] stakeHistorySysvar
	// [12] = [] stakeConfigSysvar
	// [13] = [] systemProgram
	// [14] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewRedelegateInstruction(
	// Parameters:
	lamports uint64,
	sourceTransientStakeSeed uint64,
	ephemeralStakeSeed uint64,
	destinationTransientStakeSeed uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	sourceCanonicalStakeAccount ag_solanago.PublicKey,
	sourceTransientStakeAccount ag_solanago.PublicKey,
	uninitializedEphemeralStakeAccount ag_solanago.PublicKey,
	destinationTransientStakeAccount ag_solanago.PublicKey,
	destinationStakeAccount ag_solanago.PublicKey,
	destinationValidatorVoteAccount ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	stakeHistorySysvar ag_solanago.PublicKey,
	stakeConfigSysvar ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,

) *Redelegate {
	return NewRedelegateBuilder().
		SetLamports(lamports).
		SetSourceTransientStakeSeed(sourceTransientStakeSeed).
		SetEphemeralStakeSeed(ephemeralStakeSeed).
		SetDestinationTransientStakeSeed(destinationTransientStakeSeed).
		SetStakePool(stakePool).
		SetStaker(staker).
		SetValidatorList(validatorList).
		SetSourceCanonicalStakeAccount(sourceCanonicalStakeAccount).
		SetSourceTransientStakeAccount(sourceTransientStakeAccount).
		SetUninitializedEphemeralStakeAccount(uninitializedEphemeralStakeAccount).
		SetDestinationTransientStakeAccount(destinationTransientStakeAccount).
		SetDestinationStakeAccount(destinationStakeAccount).
		SetDestinationValidatorVoteAccount(destinationValidatorVoteAccount).
		SetClockSysvar(clockSysvar).
		SetStakeHistorySysvar(stakeHistorySysvar).
		SetStakeConfigSysvar(stakeConfigSysvar).
		SetSystemProgram(systemProgram).
		SetStakeProgram(stakeProgram)
}

func NewRedelegateBuilder() *Redelegate {
	return &Redelegate{
		Accounts: make(ag_solanago.AccountMetaSlice, 15),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *Redelegate) SetLamports(lamports uint64) *Redelegate {
	inst.Lamports = &lamports
	return inst
}

func (inst *Redelegate) SetSourceTransientStakeSeed(sourceTransientStakeSeed uint64) *Redelegate {
	inst.SourceTransientStakeSeed = &sourceTransientStakeSeed
	return inst
}

func (inst *Redelegate) SetEphemeralStakeSeed(ephemeralStakeSeed uint64) *Redelegate {
	inst.EphemeralStakeSeed = &ephemeralStakeSeed
	return inst
}

func (inst *Redelegate) SetDestinationTransientStakeSeed(destinationTransientStakeSeed uint64) *Redelegate {
	inst.DestinationTransientStakeSeed = &destinationTransientStakeSeed
	return inst
}

func (inst *Redelegate) SetStakePool(stakePool ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[0] = ag_solanago.Meta(stakePool)
	return inst
}

func (inst *Redelegate) SetStaker(staker ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return inst
}

func (inst *Redelegate) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *Redelegate) SetValidatorList(validatorList ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return inst
}

func (inst *Redelegate) SetSourceCanonicalStakeAccount(sourceCanonicalStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[4] = ag_solanago.Meta(sourceCanonicalStakeAccount).WRITE()
	return inst
}

func (inst *Redelegate) SetSourceTransientStakeAccount(sourceTransientStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[5] = ag_solanago.Meta(sourceTransientStakeAccount).WRITE()
	return inst
}

func (inst *Redelegate) SetUninitializedEphemeralStakeAccount(uninitializedEphemeralStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[6] = ag_solanago.Meta(uninitializedEphemeralStakeAccount).WRITE()
	return inst
}

func (inst *Redelegate) SetDestinationTransientStakeAccount(destinationTransientStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[7] = ag_solanago.Meta(destinationTransientStakeAccount).WRITE()
	return inst
}

func (inst *Redelegate) SetDestinationStakeAccount(destinationStakeAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[8] = ag_solanago.Meta(destinationStakeAccount)
	return inst
}

func (inst *Redelegate) SetDestinationValidatorVoteAccount(destinationValidatorVoteAccount ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[9] = ag_solanago.Meta(destinationValidatorVoteAccount)
	return inst
}

func (inst *Redelegate) SetClockSysvar(clockSysvar ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[10] = ag_solanago.Meta(clockSysvar)
	return inst
}

func (inst *Redelegate) SetStakeHistorySysvar(stakeHistorySysvar ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[11] = ag_solanago.Meta(stakeHistorySysvar)
	return inst
}

func (inst *Redelegate) SetStakeConfigSysvar(stakeConfigSysvar ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[12] = ag_solanago.Meta(stakeConfigSysvar)
	return inst
}

func (inst *Redelegate) SetSystemProgram(systemProgram ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[13] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *Redelegate) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *Redelegate {
	inst.Accounts[14] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *Redelegate) GetLamports() *uint64 {
	return inst.Lamports
}

func (inst *Redelegate) GetSourceTransientStakeSeed() *uint64 {
	return inst.SourceTransientStakeSeed
}

func (inst *Redelegate) GetEphemeralStakeSeed() *uint64 {
	return inst.EphemeralStakeSeed
}

func (inst *Redelegate) GetDestinationTransientStakeSeed() *uint64 {
	return inst.DestinationTransientStakeSeed
}

func (inst *Redelegate) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *Redelegate) GetStaker() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *Redelegate) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *Redelegate) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *Redelegate) GetSourceCanonicalStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *Redelegate) GetSourceTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *Redelegate) GetUninitializedEphemeralStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *Redelegate) GetDestinationTransientStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *Redelegate) GetDestinationStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *Redelegate) GetDestinationValidatorVoteAccount() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *Redelegate) GetClockSysvar() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *Redelegate) GetStakeHistorySysvar() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *Redelegate) GetStakeConfigSysvar() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *Redelegate) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[13].PublicKey
}

func (inst *Redelegate) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[14].PublicKey
}

func (inst *Redelegate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Redelegate) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_Redelegate),
			Impl:   inst,
		},
	}
}

func (inst *Redelegate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Redelegate")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Lamports != nil {
							paramsBranch.Child(ag_format.Param("Lamports", *inst.Lamports))
						}
						if inst.SourceTransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("SourceTransientStakeSeed", *inst.SourceTransientStakeSeed))
						}
						if inst.EphemeralStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("EphemeralStakeSeed", *inst.EphemeralStakeSeed))
						}
						if inst.DestinationTransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("DestinationTransientStakeSeed", *inst.DestinationTransientStakeSeed))
						}
					})
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

func (inst *Redelegate) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Lamports != nil {
		if err := encoder.Encode(inst.Lamports); err != nil {
			return err
		}
	}
	if inst.SourceTransientStakeSeed != nil {
		if err := encoder.Encode(inst.SourceTransientStakeSeed); err != nil {
			return err
		}
	}
	if inst.EphemeralStakeSeed != nil {
		if err := encoder.Encode(inst.EphemeralStakeSeed); err != nil {
			return err
		}
	}
	if inst.DestinationTransientStakeSeed != nil {
		if err := encoder.Encode(inst.DestinationTransientStakeSeed); err != nil {
			return err
		}
	}
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *Redelegate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Lamports != nil {
		if err := decoder.Decode(inst.Lamports); err != nil {
			return err
		}
	}
	if inst.SourceTransientStakeSeed != nil {
		if err := decoder.Decode(inst.SourceTransientStakeSeed); err != nil {
			return err
		}
	}
	if inst.EphemeralStakeSeed != nil {
		if err := decoder.Decode(inst.EphemeralStakeSeed); err != nil {
			return err
		}
	}
	if inst.DestinationTransientStakeSeed != nil {
		if err := decoder.Decode(inst.DestinationTransientStakeSeed); err != nil {
			return err
		}
	}
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *Redelegate) Validate() error {
	if inst.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if inst.SourceTransientStakeSeed == nil {
		return errors.New("SourceTransientStakeSeed is not set")
	}
	if inst.EphemeralStakeSeed == nil {
		return errors.New("EphemeralStakeSeed is not set")
	}
	if inst.DestinationTransientStakeSeed == nil {
		return errors.New("DestinationTransientStakeSeed is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
