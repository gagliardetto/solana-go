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

type WithdrawSolWithSlippage struct {
	TokensIn       *uint64
	MinLamportsOut *uint64
	// [0] = [WRITE] stakePool
	// [1] = [] withdrawAuthority
	// [2] = [SIGNER] transferAuthority
	// [3] = [WRITE] burnFrom
	// [4] = [WRITE] reserveStake
	// [5] = [WRITE] withdrawTo
	// [6] = [WRITE] managerFeeAccount
	// [7] = [WRITE] poolMint
	// [8] = [] clock
	// [9] = [] stakeHistory
	// [10] = [] stakeProgram
	// [11] = [] tokenProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewWithdrawSolWithSlippageInstruction(
	// Parameters:
	tokensIn uint64,
	minLamportsOut uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	burnFrom ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	withdrawTo ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	stakeHistory ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *WithdrawSolWithSlippage {
	return NewWithdrawSolWithSlippageInstructionBuilder().
		SetTokensIn(tokensIn).
		SetMinLamportsOut(minLamportsOut).
		SetStakePool(stakePool).
		SetWithdrawAuthority(withdrawAuthority).
		SetTransferAuthority(transferAuthority).
		SetBurnFrom(burnFrom).
		SetReserveStake(reserveStake).
		SetWithdrawTo(withdrawTo).
		SetManagerFeeAccount(managerFeeAccount).
		SetPoolMint(poolMint).
		SetClock(clock).
		SetStakeHistory(stakeHistory).
		SetStakeProgram(stakeProgram).
		SetTokenProgram(tokenProgram)

}

func NewWithdrawSolWithSlippageInstructionBuilder() *WithdrawSolWithSlippage {
	return &WithdrawSolWithSlippage{
		Accounts: make(ag_solanago.AccountMetaSlice, 12),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *WithdrawSolWithSlippage) SetTokensIn(in uint64) *WithdrawSolWithSlippage {
	inst.TokensIn = &in
	return inst
}

func (inst *WithdrawSolWithSlippage) SetMinLamportsOut(out uint64) *WithdrawSolWithSlippage {
	inst.MinLamportsOut = &out
	return inst
}

func (inst *WithdrawSolWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[1] = ag_solanago.Meta(authority)
	return inst
}

func (inst *WithdrawSolWithSlippage) SetTransferAuthority(authority ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[2] = ag_solanago.Meta(authority).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetBurnFrom(burnFrom ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[3] = ag_solanago.Meta(burnFrom).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetReserveStake(reserveStake ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetWithdrawTo(withdrawTo ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[5] = ag_solanago.Meta(withdrawTo).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetManagerFeeAccount(managerFeeAccount ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[6] = ag_solanago.Meta(managerFeeAccount).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetPoolMint(poolMint ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[7] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

func (inst *WithdrawSolWithSlippage) SetClock(clock ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[8] = ag_solanago.Meta(clock)
	return inst
}

func (inst *WithdrawSolWithSlippage) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[9] = ag_solanago.Meta(stakeHistory)
	return inst
}

func (inst *WithdrawSolWithSlippage) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[10] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *WithdrawSolWithSlippage) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	inst.Accounts[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

func (inst *WithdrawSolWithSlippage) GetTokensIn() *uint64 {
	return inst.TokensIn
}

func (inst *WithdrawSolWithSlippage) GetMinLamportsOut() *uint64 {
	return inst.MinLamportsOut
}

func (inst *WithdrawSolWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetTransferAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetBurnFrom() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetWithdrawTo() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetClock() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *WithdrawSolWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *WithdrawSolWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawSolWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawSolWithSlippage),
			Impl:   inst,
		},
	}
}

func (inst *WithdrawSolWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawSolWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.TokensIn != nil {
							paramsBranch.Child(ag_format.Param("TokensIn", *inst.TokensIn))
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

func (inst *WithdrawSolWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.TokensIn != nil {
		if err := encoder.Encode(inst.TokensIn); err != nil {
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

func (inst *WithdrawSolWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.TokensIn != nil {
		if err := decoder.Decode(inst.TokensIn); err != nil {
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

func (inst *WithdrawSolWithSlippage) Validate() error {
	if inst.TokensIn == nil {
		return errors.New("tokensIn is not set")
	}
	if inst.MinLamportsOut == nil {
		return errors.New("minLamportsOut is not set")
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
