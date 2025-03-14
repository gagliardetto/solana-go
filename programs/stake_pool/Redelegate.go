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

func (r *Redelegate) SetLamports(lamports uint64) *Redelegate {
	r.Lamports = &lamports
	return r
}

func (r *Redelegate) SetSourceTransientStakeSeed(sourceTransientStakeSeed uint64) *Redelegate {
	r.SourceTransientStakeSeed = &sourceTransientStakeSeed
	return r
}

func (r *Redelegate) SetEphemeralStakeSeed(ephemeralStakeSeed uint64) *Redelegate {
	r.EphemeralStakeSeed = &ephemeralStakeSeed
	return r
}

func (r *Redelegate) SetDestinationTransientStakeSeed(destinationTransientStakeSeed uint64) *Redelegate {
	r.DestinationTransientStakeSeed = &destinationTransientStakeSeed
	return r
}

func (r *Redelegate) SetStakePool(stakePool ag_solanago.PublicKey) *Redelegate {
	r.Accounts[0] = ag_solanago.Meta(stakePool)
	return r
}

func (r *Redelegate) SetStaker(staker ag_solanago.PublicKey) *Redelegate {
	r.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	r.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return r
}

func (r *Redelegate) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *Redelegate {
	r.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return r
}

func (r *Redelegate) SetValidatorList(validatorList ag_solanago.PublicKey) *Redelegate {
	r.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return r
}

func (r *Redelegate) SetSourceCanonicalStakeAccount(sourceCanonicalStakeAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[4] = ag_solanago.Meta(sourceCanonicalStakeAccount).WRITE()
	return r
}

func (r *Redelegate) SetSourceTransientStakeAccount(sourceTransientStakeAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[5] = ag_solanago.Meta(sourceTransientStakeAccount).WRITE()
	return r
}

func (r *Redelegate) SetUninitializedEphemeralStakeAccount(uninitializedEphemeralStakeAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[6] = ag_solanago.Meta(uninitializedEphemeralStakeAccount).WRITE()
	return r
}

func (r *Redelegate) SetDestinationTransientStakeAccount(destinationTransientStakeAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[7] = ag_solanago.Meta(destinationTransientStakeAccount).WRITE()
	return r
}

func (r *Redelegate) SetDestinationStakeAccount(destinationStakeAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[8] = ag_solanago.Meta(destinationStakeAccount)
	return r
}

func (r *Redelegate) SetDestinationValidatorVoteAccount(destinationValidatorVoteAccount ag_solanago.PublicKey) *Redelegate {
	r.Accounts[9] = ag_solanago.Meta(destinationValidatorVoteAccount)
	return r
}

func (r *Redelegate) SetClockSysvar(clockSysvar ag_solanago.PublicKey) *Redelegate {
	r.Accounts[10] = ag_solanago.Meta(clockSysvar)
	return r
}

func (r *Redelegate) SetStakeHistorySysvar(stakeHistorySysvar ag_solanago.PublicKey) *Redelegate {
	r.Accounts[11] = ag_solanago.Meta(stakeHistorySysvar)
	return r
}

func (r *Redelegate) SetStakeConfigSysvar(stakeConfigSysvar ag_solanago.PublicKey) *Redelegate {
	r.Accounts[12] = ag_solanago.Meta(stakeConfigSysvar)
	return r
}

func (r *Redelegate) SetSystemProgram(systemProgram ag_solanago.PublicKey) *Redelegate {
	r.Accounts[13] = ag_solanago.Meta(systemProgram)
	return r
}

func (r *Redelegate) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *Redelegate {
	r.Accounts[14] = ag_solanago.Meta(stakeProgram)
	return r
}

func (r *Redelegate) GetLamports() *uint64 {
	return r.Lamports
}

func (r *Redelegate) GetSourceTransientStakeSeed() *uint64 {
	return r.SourceTransientStakeSeed
}

func (r *Redelegate) GetEphemeralStakeSeed() *uint64 {
	return r.EphemeralStakeSeed
}

func (r *Redelegate) GetDestinationTransientStakeSeed() *uint64 {
	return r.DestinationTransientStakeSeed
}

func (r *Redelegate) GetStakePool() ag_solanago.PublicKey {
	return r.Accounts[0].PublicKey
}

func (r *Redelegate) GetStaker() ag_solanago.PublicKey {
	return r.Accounts[1].PublicKey
}

func (r *Redelegate) GetWithdrawAuthority() ag_solanago.PublicKey {
	return r.Accounts[2].PublicKey
}

func (r *Redelegate) GetValidatorList() ag_solanago.PublicKey {
	return r.Accounts[3].PublicKey
}

func (r *Redelegate) GetSourceCanonicalStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[4].PublicKey
}

func (r *Redelegate) GetSourceTransientStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[5].PublicKey
}

func (r *Redelegate) GetUninitializedEphemeralStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[6].PublicKey
}

func (r *Redelegate) GetDestinationTransientStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[7].PublicKey
}

func (r *Redelegate) GetDestinationStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[8].PublicKey
}

func (r *Redelegate) GetDestinationValidatorVoteAccount() ag_solanago.PublicKey {
	return r.Accounts[9].PublicKey
}

func (r *Redelegate) GetClockSysvar() ag_solanago.PublicKey {
	return r.Accounts[10].PublicKey
}

func (r *Redelegate) GetStakeHistorySysvar() ag_solanago.PublicKey {
	return r.Accounts[11].PublicKey
}

func (r *Redelegate) GetStakeConfigSysvar() ag_solanago.PublicKey {
	return r.Accounts[12].PublicKey
}

func (r *Redelegate) GetSystemProgram() ag_solanago.PublicKey {
	return r.Accounts[13].PublicKey
}

func (r *Redelegate) GetStakeProgram() ag_solanago.PublicKey {
	return r.Accounts[14].PublicKey
}

func (r *Redelegate) ValidateAndBuild() (*Instruction, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return r.Build(), nil
}

func (r *Redelegate) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_Redelegate),
			Impl:   r,
		},
	}
}

func (r *Redelegate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Redelegate")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if r.Lamports != nil {
							paramsBranch.Child(ag_format.Param("Lamports", *r.Lamports))
						}
						if r.SourceTransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("SourceTransientStakeSeed", *r.SourceTransientStakeSeed))
						}
						if r.EphemeralStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("EphemeralStakeSeed", *r.EphemeralStakeSeed))
						}
						if r.DestinationTransientStakeSeed != nil {
							paramsBranch.Child(ag_format.Param("DestinationTransientStakeSeed", *r.DestinationTransientStakeSeed))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range r.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(r.Signers)))
						for j, signer := range r.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (r *Redelegate) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if r.Lamports != nil {
		if err := encoder.Encode(r.Lamports); err != nil {
			return err
		}
	}
	if r.SourceTransientStakeSeed != nil {
		if err := encoder.Encode(r.SourceTransientStakeSeed); err != nil {
			return err
		}
	}
	if r.EphemeralStakeSeed != nil {
		if err := encoder.Encode(r.EphemeralStakeSeed); err != nil {
			return err
		}
	}
	if r.DestinationTransientStakeSeed != nil {
		if err := encoder.Encode(r.DestinationTransientStakeSeed); err != nil {
			return err
		}
	}
	for _, account := range r.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (r *Redelegate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if r.Lamports != nil {
		if err := decoder.Decode(r.Lamports); err != nil {
			return err
		}
	}
	if r.SourceTransientStakeSeed != nil {
		if err := decoder.Decode(r.SourceTransientStakeSeed); err != nil {
			return err
		}
	}
	if r.EphemeralStakeSeed != nil {
		if err := decoder.Decode(r.EphemeralStakeSeed); err != nil {
			return err
		}
	}
	if r.DestinationTransientStakeSeed != nil {
		if err := decoder.Decode(r.DestinationTransientStakeSeed); err != nil {
			return err
		}
	}
	for i := range r.Accounts {
		if err := decoder.Decode(r.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (r *Redelegate) Validate() error {
	if r.Lamports == nil {
		return errors.New("lamports is not set")
	}
	if r.SourceTransientStakeSeed == nil {
		return errors.New("SourceTransientStakeSeed is not set")
	}
	if r.EphemeralStakeSeed == nil {
		return errors.New("EphemeralStakeSeed is not set")
	}
	if r.DestinationTransientStakeSeed == nil {
		return errors.New("DestinationTransientStakeSeed is not set")
	}
	for i, account := range r.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(r.Signers) == 0 || !r.Signers[0].IsSigner {
		return errors.New("accounts.Staker should be a signer")
	}
	return nil
}
