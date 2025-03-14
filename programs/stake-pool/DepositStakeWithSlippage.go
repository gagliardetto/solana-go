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

type DepositStakeWithSlippage struct {
	MinTokensOut *uint64
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

func NewDepositStakeWithSlippageInstruction(
	// Parameters:
	minTokensOut uint64,
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
) *DepositStakeWithSlippage {
	return NewDepositStakeWithSlippageBuilder().
		SetMinTokensOut(minTokensOut).
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

func NewDepositStakeWithSlippageBuilder() *DepositStakeWithSlippage {
	return &DepositStakeWithSlippage{
		Accounts: make(ag_solanago.AccountMetaSlice, 15),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *DepositStakeWithSlippage) SetMinTokensOut(out uint64) *DepositStakeWithSlippage {
	inst.MinTokensOut = &out
	return inst
}

func (inst *DepositStakeWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetValidatorList(list ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[1] = ag_solanago.Meta(list).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetStakeDepositAuthority(authority ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[2] = ag_solanago.Meta(authority).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

func (inst *DepositStakeWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[3] = ag_solanago.Meta(authority)
	return inst
}

func (inst *DepositStakeWithSlippage) SetStakeDepositing(depositing ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[4] = ag_solanago.Meta(depositing).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetValidatorStakeAccount(account ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[5] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetReserveStake(stake ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[6] = ag_solanago.Meta(stake).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetMintTo(to ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[7] = ag_solanago.Meta(to).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetManagerFeeAccount(account ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[8] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetReferralFeeDest(dest ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[9] = ag_solanago.Meta(dest).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetPoolMint(mint ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[10] = ag_solanago.Meta(mint).WRITE()
	return inst
}

func (inst *DepositStakeWithSlippage) SetClock(clock ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[11] = ag_solanago.Meta(clock)
	return inst
}

func (inst *DepositStakeWithSlippage) SetStakeHistory(history ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[12] = ag_solanago.Meta(history)
	return inst
}

func (inst *DepositStakeWithSlippage) SetTokenProgram(program ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[13] = ag_solanago.Meta(program)
	return inst
}

func (inst *DepositStakeWithSlippage) SetStakeProgram(program ag_solanago.PublicKey) *DepositStakeWithSlippage {
	inst.Accounts[14] = ag_solanago.Meta(program)
	return inst
}

func (inst *DepositStakeWithSlippage) GetMinTokensOut() *uint64 {
	return inst.MinTokensOut
}

func (inst *DepositStakeWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *DepositStakeWithSlippage) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *DepositStakeWithSlippage) GetStakeDepositAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *DepositStakeWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *DepositStakeWithSlippage) GetStakeDepositing() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *DepositStakeWithSlippage) GetValidatorStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *DepositStakeWithSlippage) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *DepositStakeWithSlippage) GetMintTo() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *DepositStakeWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *DepositStakeWithSlippage) GetReferralFeeDest() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *DepositStakeWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *DepositStakeWithSlippage) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *DepositStakeWithSlippage) GetStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *DepositStakeWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[13].PublicKey
}

func (inst *DepositStakeWithSlippage) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[14].PublicKey
}

func (inst *DepositStakeWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositStakeWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DepositStakeWithSlippage),
			Impl:   inst,
		},
	}
}

func (inst *DepositStakeWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositStakeWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.MinTokensOut != nil {
							paramsBranch.Child(ag_format.Param("MinTokensOut", *inst.MinTokensOut))
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

func (inst *DepositStakeWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.MinTokensOut != nil {
		if err := encoder.Encode(inst.MinTokensOut); err != nil {
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

func (inst *DepositStakeWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.MinTokensOut != nil {
		if err := decoder.Decode(inst.MinTokensOut); err != nil {
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

func (inst *DepositStakeWithSlippage) Validate() error {
	if inst.MinTokensOut == nil {
		return errors.New("minTokensOut is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.StakeDepositAuthority should be a signer")
	}
	return nil
}
