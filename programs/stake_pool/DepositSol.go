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

type DepositSol struct {
	// [0] = [WRITE] stakePool
	// [1] = [] withdrawAuthority
	// [2] = [WRITE] reserveStake
	// [3] = [SIGNER] fundingAccount
	// [4] = [WRITE] destinationAccount
	// [5] = [WRITE] managerFeeAccount
	// [6] = [WRITE] referralFeeAccount
	// [7] = [WRITE] poolMint
	// [8] = [] systemProgram
	// [9] = [] tokenProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewDepositSolInstruction(
	stakePool ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	fundingAccount ag_solanago.PublicKey,
	destinationAccount ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	referralFeeAccount ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *DepositSol {
	return NewDepositSolBuilder().
		SetStakePool(stakePool).
		SetWithdrawAuthority(withdrawAuthority).
		SetReserveStake(reserveStake).
		SetFundingAccount(fundingAccount).
		SetDestinationAccount(destinationAccount).
		SetManagerFeeAccount(managerFeeAccount).
		SetReferralFeeAccount(referralFeeAccount).
		SetPoolMint(poolMint).
		SetSystemProgram(systemProgram).
		SetTokenProgram(tokenProgram)
}

func NewDepositSolBuilder() *DepositSol {
	return &DepositSol{
		Accounts: make(ag_solanago.AccountMetaSlice, 10),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (inst *DepositSol) SetStakePool(stakePool ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return inst
}

func (inst *DepositSol) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[1] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *DepositSol) SetReserveStake(reserveStake ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[2] = ag_solanago.Meta(reserveStake).WRITE()
	return inst
}

func (inst *DepositSol) SetFundingAccount(fundingAccount ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[3] = ag_solanago.Meta(fundingAccount).SIGNER()
	inst.Signers = append(inst.Signers, ag_solanago.Meta(fundingAccount).SIGNER())
	return inst
}

func (inst *DepositSol) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[4] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

func (inst *DepositSol) SetManagerFeeAccount(managerFeeAccount ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[5] = ag_solanago.Meta(managerFeeAccount).WRITE()
	return inst
}

func (inst *DepositSol) SetReferralFeeAccount(referralFeeAccount ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[6] = ag_solanago.Meta(referralFeeAccount).WRITE()
	return inst
}

func (inst *DepositSol) SetPoolMint(poolMint ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[7] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

func (inst *DepositSol) SetSystemProgram(systemProgram ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[8] = ag_solanago.Meta(systemProgram)
	return inst
}

func (inst *DepositSol) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *DepositSol {
	inst.Accounts[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

func (inst *DepositSol) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *DepositSol) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *DepositSol) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *DepositSol) GetFundingAccount() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *DepositSol) GetDestinationAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *DepositSol) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *DepositSol) GetReferralFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *DepositSol) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *DepositSol) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *DepositSol) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *DepositSol) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositSol) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DepositSol),
			Impl:   inst,
		},
	}
}

func (inst *DepositSol) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositSol")).
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

func (inst *DepositSol) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	for _, account := range inst.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (inst *DepositSol) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	for i := range inst.Accounts {
		if err := decoder.Decode(inst.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (inst *DepositSol) Validate() error {
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.FundingAccount should be a signer")
	}
	return nil
}
