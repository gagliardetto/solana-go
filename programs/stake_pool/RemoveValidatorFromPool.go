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

type RemoveValidatorFromPool struct {
	// [0]  = [WRITE] stakePool
	// [1]  = [SIGNER] staker
	// [2]  = [] withdrawAuthority
	// [3]  = [WRITE] validatorList
	// [4]  = [WRITE] validatorStakeAccount
	// [5]  = [WRITE] transientStakeAccount
	// [6]  = [] clock
	// [7]  = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewRemoveValidatorFromPoolInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	staker ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	transientStakeAccount ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *RemoveValidatorFromPool {
	return NewRemoveValidatorFromPoolInstructionBuilder().
		SetStakePool(stakePool).
		SetStaker(staker).
		SetWithdrawAuthority(withdrawAuthority).
		SetValidatorList(validatorList).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetTransientStakeAccount(transientStakeAccount).
		SetClock(clock).
		SetStakeProgram(stakeProgram)
}

func NewRemoveValidatorFromPoolInstructionBuilder() *RemoveValidatorFromPool {
	return &RemoveValidatorFromPool{
		Accounts: make(ag_solanago.AccountMetaSlice, 8),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (r *RemoveValidatorFromPool) SetStakePool(pool ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return r
}

func (r *RemoveValidatorFromPool) SetStaker(staker ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[1] = ag_solanago.Meta(staker).SIGNER()
	r.Signers[0] = ag_solanago.Meta(staker).SIGNER()
	return r
}

func (r *RemoveValidatorFromPool) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[2] = ag_solanago.Meta(withdrawAuthority)
	return r
}

func (r *RemoveValidatorFromPool) SetValidatorList(validatorList ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[3] = ag_solanago.Meta(validatorList).WRITE()
	return r
}

func (r *RemoveValidatorFromPool) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[4] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return r
}

func (r *RemoveValidatorFromPool) SetTransientStakeAccount(transientStakeAccount ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[5] = ag_solanago.Meta(transientStakeAccount).WRITE()
	return r
}

func (r *RemoveValidatorFromPool) SetClock(clock ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[6] = ag_solanago.Meta(clock)
	return r
}

func (r *RemoveValidatorFromPool) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *RemoveValidatorFromPool {
	r.Accounts[7] = ag_solanago.Meta(stakeProgram)
	return r
}

func (r *RemoveValidatorFromPool) GetStakePool() ag_solanago.PublicKey {
	return r.Accounts[0].PublicKey
}

func (r *RemoveValidatorFromPool) GetStaker() ag_solanago.PublicKey {
	return r.Accounts[1].PublicKey
}

func (r *RemoveValidatorFromPool) GetWithdrawAuthority() ag_solanago.PublicKey {
	return r.Accounts[2].PublicKey
}

func (r *RemoveValidatorFromPool) GetValidatorList() ag_solanago.PublicKey {
	return r.Accounts[3].PublicKey
}

func (r *RemoveValidatorFromPool) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[4].PublicKey
}

func (r *RemoveValidatorFromPool) GetTransientStakeAccount() ag_solanago.PublicKey {
	return r.Accounts[5].PublicKey
}

func (r *RemoveValidatorFromPool) GetClock() ag_solanago.PublicKey {
	return r.Accounts[6].PublicKey
}

func (r *RemoveValidatorFromPool) GetStakeProgram() ag_solanago.PublicKey {
	return r.Accounts[7].PublicKey
}

func (r *RemoveValidatorFromPool) ValidateAndBuild() (*Instruction, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return r.Build(), nil
}

func (r *RemoveValidatorFromPool) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_RemoveValidatorFromPool),
			Impl:   r,
		},
	}
}

func (r *RemoveValidatorFromPool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveValidatorFromPool")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
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

func (r *RemoveValidatorFromPool) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range r.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (r *RemoveValidatorFromPool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range r.Accounts {
		if err := decoder.Decode(r.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (r *RemoveValidatorFromPool) Validate() error {
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
