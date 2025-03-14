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

type DepositStake struct {
	// [0] = [WRITE] stakePool
	// [1] = [WRITE] validatorList
	// [2] = [SIGNER] stakeDepositAuthority
	// [3] = [] withdrawAuthority
	// [4] = [WRITE] stakeDepositing
	// [5] = [WRITE] validatorStakeAccount
	// [6] = [WRITE] reserveStake
	// [7] = [WRITE] mintTo
	// [8] = [WRITE] managerFeeAccount
	// [9] = [WRITE] referralFeeDest
	// [10] = [WRITE] poolMint
	// [11] = [] clock
	// [12] = [] stakeHistory
	// [13] = [] tokenProgram
	// [14] = [] stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewDepositStakeInstruction(
	// Accounts:
	stakePool ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	stakeDepositAuthority ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	stakeDepositing ag_solanago.PublicKey,
	validatorStakeAccount ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	mintTo ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	referralFeeDest ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *DepositStake {
	return NewDepositStakeBuilder().
		SetStakePool(stakePool).
		SetValidatorList(validatorList).
		SetStakeDepositAuthority(stakeDepositAuthority).
		SetWithdrawAuthority(withdrawAuthority).
		SetStakeDepositing(stakeDepositing).
		SetValidatorStakeAccount(validatorStakeAccount).
		SetReserveStake(reserveStake).
		SetMintTo(mintTo).
		SetManagerFeeAccount(managerFeeAccount).
		SetReferralFeeDest(referralFeeDest).
		SetPoolMint(poolMint).
		SetClock(clock).
		SetStakeHistory(stakeHistory).
		SetTokenProgram(tokenProgram).
		SetStakeProgram(stakeProgram)
}

func NewDepositStakeBuilder() *DepositStake {
	return &DepositStake{
		Accounts: make(ag_solanago.AccountMetaSlice, 15),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (d *DepositStake) SetStakePool(stakePool ag_solanago.PublicKey) *DepositStake {
	d.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return d
}

func (d *DepositStake) SetValidatorList(validatorList ag_solanago.PublicKey) *DepositStake {
	d.Accounts[1] = ag_solanago.Meta(validatorList).WRITE()
	return d
}

func (d *DepositStake) SetStakeDepositAuthority(stakeDepositAuthority ag_solanago.PublicKey) *DepositStake {
	d.Accounts[2] = ag_solanago.Meta(stakeDepositAuthority).SIGNER()
	d.Signers[0] = ag_solanago.Meta(stakeDepositAuthority).SIGNER()
	return d
}

func (d *DepositStake) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *DepositStake {
	d.Accounts[3] = ag_solanago.Meta(withdrawAuthority)
	return d
}

func (d *DepositStake) SetStakeDepositing(stakeDepositing ag_solanago.PublicKey) *DepositStake {
	d.Accounts[4] = ag_solanago.Meta(stakeDepositing).WRITE()
	return d
}

func (d *DepositStake) SetValidatorStakeAccount(validatorStakeAccount ag_solanago.PublicKey) *DepositStake {
	d.Accounts[5] = ag_solanago.Meta(validatorStakeAccount).WRITE()
	return d
}

func (d *DepositStake) SetReserveStake(reserveStake ag_solanago.PublicKey) *DepositStake {
	d.Accounts[6] = ag_solanago.Meta(reserveStake).WRITE()
	return d
}

func (d *DepositStake) SetMintTo(mintTo ag_solanago.PublicKey) *DepositStake {
	d.Accounts[7] = ag_solanago.Meta(mintTo).WRITE()
	return d
}

func (d *DepositStake) SetManagerFeeAccount(managerFeeAccount ag_solanago.PublicKey) *DepositStake {
	d.Accounts[8] = ag_solanago.Meta(managerFeeAccount).WRITE()
	return d
}

func (d *DepositStake) SetReferralFeeDest(referralFeeDest ag_solanago.PublicKey) *DepositStake {
	d.Accounts[9] = ag_solanago.Meta(referralFeeDest).WRITE()
	return d
}

func (d *DepositStake) SetPoolMint(poolMint ag_solanago.PublicKey) *DepositStake {
	d.Accounts[10] = ag_solanago.Meta(poolMint).WRITE()
	return d
}

func (d *DepositStake) SetClock(clock ag_solanago.PublicKey) *DepositStake {
	d.Accounts[11] = ag_solanago.Meta(clock)
	return d
}

func (d *DepositStake) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *DepositStake {
	d.Accounts[12] = ag_solanago.Meta(stakeHistory)
	return d
}

func (d *DepositStake) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *DepositStake {
	d.Accounts[13] = ag_solanago.Meta(tokenProgram)
	return d
}

func (d *DepositStake) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *DepositStake {
	d.Accounts[14] = ag_solanago.Meta(stakeProgram)
	return d
}

func (d *DepositStake) GetStakePool() ag_solanago.PublicKey {
	return d.Accounts[0].PublicKey
}

func (d *DepositStake) GetValidatorList() ag_solanago.PublicKey {
	return d.Accounts[1].PublicKey
}

func (d *DepositStake) GetStakeDepositAuthority() ag_solanago.PublicKey {
	return d.Accounts[2].PublicKey
}

func (d *DepositStake) GetWithdrawAuthority() ag_solanago.PublicKey {
	return d.Accounts[3].PublicKey
}

func (d *DepositStake) GetStakeDepositing() ag_solanago.PublicKey {
	return d.Accounts[4].PublicKey
}

func (d *DepositStake) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return d.Accounts[5].PublicKey
}

func (d *DepositStake) GetReserveStake() ag_solanago.PublicKey {
	return d.Accounts[6].PublicKey
}

func (d *DepositStake) GetMintTo() ag_solanago.PublicKey {
	return d.Accounts[7].PublicKey
}

func (d *DepositStake) GetManagerFeeAccount() ag_solanago.PublicKey {
	return d.Accounts[8].PublicKey
}

func (d *DepositStake) GetReferralFeeDest() ag_solanago.PublicKey {
	return d.Accounts[9].PublicKey
}

func (d *DepositStake) GetPoolMint() ag_solanago.PublicKey {
	return d.Accounts[10].PublicKey
}

func (d *DepositStake) GetClock() ag_solanago.PublicKey {
	return d.Accounts[11].PublicKey
}

func (d *DepositStake) GetStakeHistory() ag_solanago.PublicKey {
	return d.Accounts[12].PublicKey
}

func (d *DepositStake) GetTokenProgram() ag_solanago.PublicKey {
	return d.Accounts[13].PublicKey
}

func (d *DepositStake) GetStakeProgram() ag_solanago.PublicKey {
	return d.Accounts[14].PublicKey
}

func (d *DepositStake) ValidateAndBuild() (*Instruction, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d.Build(), nil
}

func (d *DepositStake) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DepositStake),
			Impl:   d,
		},
	}
}

func (d *DepositStake) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositStake")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range d.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(d.Signers)))
						for j, signer := range d.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (d *DepositStake) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range d.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepositStake) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range d.Accounts {
		if err := decoder.Decode(d.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepositStake) Validate() error {
	for i, account := range d.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(d.Signers) == 0 || !d.Signers[0].IsSigner {
		return errors.New("accounts.StakeDepositAuthority should be a signer")
	}
	return nil
}
