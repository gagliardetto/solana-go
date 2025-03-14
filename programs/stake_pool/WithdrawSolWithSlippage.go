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

func (w *WithdrawSolWithSlippage) SetTokensIn(in uint64) *WithdrawSolWithSlippage {
	w.TokensIn = &in
	return w
}

func (w *WithdrawSolWithSlippage) SetMinLamportsOut(out uint64) *WithdrawSolWithSlippage {
	w.MinLamportsOut = &out
	return w
}

func (w *WithdrawSolWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[1] = ag_solanago.Meta(authority)
	return w
}

func (w *WithdrawSolWithSlippage) SetTransferAuthority(authority ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[2] = ag_solanago.Meta(authority).SIGNER()
	w.Signers[0] = ag_solanago.Meta(authority).SIGNER()
	return w
}

func (w *WithdrawSolWithSlippage) SetBurnFrom(burnFrom ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[3] = ag_solanago.Meta(burnFrom).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetReserveStake(reserveStake ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[4] = ag_solanago.Meta(reserveStake).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetWithdrawTo(withdrawTo ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[5] = ag_solanago.Meta(withdrawTo).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetManagerFeeAccount(managerFeeAccount ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[6] = ag_solanago.Meta(managerFeeAccount).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetPoolMint(poolMint ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[7] = ag_solanago.Meta(poolMint).WRITE()
	return w
}

func (w *WithdrawSolWithSlippage) SetClock(clock ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[8] = ag_solanago.Meta(clock)
	return w
}

func (w *WithdrawSolWithSlippage) SetStakeHistory(stakeHistory ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[9] = ag_solanago.Meta(stakeHistory)
	return w
}

func (w *WithdrawSolWithSlippage) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[10] = ag_solanago.Meta(stakeProgram)
	return w
}

func (w *WithdrawSolWithSlippage) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *WithdrawSolWithSlippage {
	w.Accounts[11] = ag_solanago.Meta(tokenProgram)
	return w
}

func (w *WithdrawSolWithSlippage) GetTokensIn() *uint64 {
	return w.TokensIn
}

func (w *WithdrawSolWithSlippage) GetMinLamportsOut() *uint64 {
	return w.MinLamportsOut
}

func (w *WithdrawSolWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return w.Accounts[0].PublicKey
}

func (w *WithdrawSolWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return w.Accounts[1].PublicKey
}

func (w *WithdrawSolWithSlippage) GetTransferAuthority() ag_solanago.PublicKey {
	return w.Accounts[2].PublicKey
}

func (w *WithdrawSolWithSlippage) GetBurnFrom() ag_solanago.PublicKey {
	return w.Accounts[3].PublicKey
}

func (w *WithdrawSolWithSlippage) GetReserveStake() ag_solanago.PublicKey {
	return w.Accounts[4].PublicKey
}

func (w *WithdrawSolWithSlippage) GetWithdrawTo() ag_solanago.PublicKey {
	return w.Accounts[5].PublicKey
}

func (w *WithdrawSolWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return w.Accounts[6].PublicKey
}

func (w *WithdrawSolWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return w.Accounts[7].PublicKey
}

func (w *WithdrawSolWithSlippage) GetClock() ag_solanago.PublicKey {
	return w.Accounts[8].PublicKey
}

func (w *WithdrawSolWithSlippage) GetStakeHistory() ag_solanago.PublicKey {
	return w.Accounts[9].PublicKey
}

func (w *WithdrawSolWithSlippage) GetStakeProgram() ag_solanago.PublicKey {
	return w.Accounts[10].PublicKey
}

func (w *WithdrawSolWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return w.Accounts[11].PublicKey
}

func (w *WithdrawSolWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := w.Validate(); err != nil {
		return nil, err
	}
	return w.Build(), nil
}

func (w *WithdrawSolWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawSolWithSlippage),
			Impl:   w,
		},
	}
}

func (w *WithdrawSolWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawSolWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if w.TokensIn != nil {
							paramsBranch.Child(ag_format.Param("TokensIn", *w.TokensIn))
						}
						if w.MinLamportsOut != nil {
							paramsBranch.Child(ag_format.Param("MinLamportsOut", *w.MinLamportsOut))
						}
					})
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						for i, account := range w.Accounts {
							accountsBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", i), account))
						}
						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(w.Signers)))
						for j, signer := range w.Signers {
							signersBranch.Child(ag_format.Meta(fmt.Sprintf("[%v]", j), signer))
						}
					})
				})
		})
}

func (w *WithdrawSolWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if w.TokensIn != nil {
		if err := encoder.Encode(w.TokensIn); err != nil {
			return err
		}
	}
	if w.MinLamportsOut != nil {
		if err := encoder.Encode(w.MinLamportsOut); err != nil {
			return err
		}
	}
	for _, account := range w.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (w *WithdrawSolWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if w.TokensIn != nil {
		if err := decoder.Decode(w.TokensIn); err != nil {
			return err
		}
	}
	if w.MinLamportsOut != nil {
		if err := decoder.Decode(w.MinLamportsOut); err != nil {
			return err
		}
	}
	for i := range w.Accounts {
		if err := decoder.Decode(w.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (w *WithdrawSolWithSlippage) Validate() error {
	if w.TokensIn == nil {
		return errors.New("tokensIn is not set")
	}
	if w.MinLamportsOut == nil {
		return errors.New("minLamportsOut is not set")
	}
	for i, account := range w.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(w.Signers) == 0 || !w.Signers[0].IsSigner {
		return errors.New("accounts.TransferAuthority should be a signer")
	}
	return nil
}
