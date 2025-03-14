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

type WithdrawStakeWithSlippage struct {
	PoolTokensIn   *uint64
	MinLamportsOut *uint64
	// [0] = [WRITE] stakePool
	// [1] = [WRITE] validatorList
	// [2] = [] withdrawAuthority
	// [3] = [WRITE] splitFrom
	// [4] = [WRITE] splitTo
	// [5] = [] beneficiary
	// [6] = [SIGNER] transferAuthority
	// [7] = [WRITE] burnFrom
	// [8] = [WRITE] managerFeeAccount
	// [9] = [WRITE] poolMint
	// [10] = []clock
	// [11] = []tokenProgram
	// [12] = []stakeProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewWithdrawStakeWithSlippageInstruction(
	// Parameters:
	poolTokensIn uint64,
	minLamportsOut uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	validatorList ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	splitFrom ag_solanago.PublicKey,
	splitTo ag_solanago.PublicKey,
	beneficiary ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	burnFrom ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
) *WithdrawStakeWithSlippage {
	return NewWithdrawStakeWithSlippageInstructionBuilder().
		SetPoolTokensIn(poolTokensIn).
		SetMinLamportsOut(minLamportsOut).
		SetStakePool(stakePool).
		SetValidatorList(validatorList).
		SetWithdrawAuthority(withdrawAuthority).
		SetSplitFrom(splitFrom).
		SetSplitTo(splitTo).
		SetBeneficiary(beneficiary).
		SetTransferAuthority(transferAuthority).
		SetBurnFrom(burnFrom).
		SetManagerFeeAccount(managerFeeAccount).
		SetPoolMint(poolMint).
		SetClock(clock).
		SetTokenProgram(tokenProgram).
		SetStakeProgram(stakeProgram)
}

func NewWithdrawStakeWithSlippageInstructionBuilder() *WithdrawStakeWithSlippage {
	return &WithdrawStakeWithSlippage{
		Accounts: make(ag_solanago.AccountMetaSlice, 13),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *WithdrawStakeWithSlippage) SetPoolTokensIn(in uint64) *WithdrawStakeWithSlippage {
	inst.PoolTokensIn = &in
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetMinLamportsOut(out uint64) *WithdrawStakeWithSlippage {
	inst.MinLamportsOut = &out
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetValidatorList(list ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[1] = ag_solanago.Meta(list).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[2] = ag_solanago.Meta(authority)
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetSplitFrom(from ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[3] = ag_solanago.Meta(from).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetSplitTo(to ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[4] = ag_solanago.Meta(to).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetBeneficiary(beneficiary ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[5] = ag_solanago.Meta(beneficiary)
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetTransferAuthority(authority ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[6] = ag_solanago.Meta(authority).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetBurnFrom(from ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[7] = ag_solanago.Meta(from).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetManagerFeeAccount(account ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[8] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetPoolMint(mint ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[9] = ag_solanago.Meta(mint).WRITE()
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetClock(clock ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[10] = ag_solanago.Meta(clock)
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetTokenProgram(program ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[11] = ag_solanago.Meta(program)
	return inst
}

func (inst *WithdrawStakeWithSlippage) SetStakeProgram(program ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	inst.Accounts[12] = ag_solanago.Meta(program)
	return inst
}

func (inst *WithdrawStakeWithSlippage) GetPoolTokensIn() *uint64 {
	return inst.PoolTokensIn
}

func (inst *WithdrawStakeWithSlippage) GetMinLamportsOut() *uint64 {
	return inst.MinLamportsOut
}

func (inst *WithdrawStakeWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetValidatorList() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetSplitFrom() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetSplitTo() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetBeneficiary() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetTransferAuthority() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetBurnFrom() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *WithdrawStakeWithSlippage) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *WithdrawStakeWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawStakeWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawStakeWithSlippage),
			Impl:   inst,
		},
	}
}

func (inst *WithdrawStakeWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawStakeWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.PoolTokensIn != nil {
							paramsBranch.Child(ag_format.Param("PoolTokensIn", *inst.PoolTokensIn))
						}
						if inst.MinLamportsOut != nil {
							paramsBranch.Child(ag_format.Param("MinLamportsOut", *inst.MinLamportsOut))
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

func (inst *WithdrawStakeWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.PoolTokensIn != nil {
		if err := encoder.Encode(inst.PoolTokensIn); err != nil {
			return err
		}
	}
	if inst.MinLamportsOut != nil {
		if err := encoder.Encode(inst.MinLamportsOut); err != nil {
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

func (inst *WithdrawStakeWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.PoolTokensIn != nil {
		if err := decoder.Decode(inst.PoolTokensIn); err != nil {
			return err
		}
	}
	if inst.MinLamportsOut != nil {
		if err := decoder.Decode(inst.MinLamportsOut); err != nil {
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

func (inst *WithdrawStakeWithSlippage) Validate() error {
	if inst.PoolTokensIn == nil {
		return errors.New("PoolTokensIn is not set")
	}
	if inst.MinLamportsOut == nil {
		return errors.New("MinLamportsOut is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.TransferAuthority should be a signer")
	}
	return nil
}
