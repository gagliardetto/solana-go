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

func (w *WithdrawStakeWithSlippage) SetPoolTokensIn(in uint64) *WithdrawStakeWithSlippage {
	w.PoolTokensIn = &in
	return w
}

func (w *WithdrawStakeWithSlippage) SetMinLamportsOut(out uint64) *WithdrawStakeWithSlippage {
	w.MinLamportsOut = &out
	return w
}

func (w *WithdrawStakeWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetValidatorList(list ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[1] = ag_solanago.Meta(list).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[2] = ag_solanago.Meta(authority)
	return w
}

func (w *WithdrawStakeWithSlippage) SetSplitFrom(from ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[3] = ag_solanago.Meta(from).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetSplitTo(to ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[4] = ag_solanago.Meta(to).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetBeneficiary(beneficiary ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[5] = ag_solanago.Meta(beneficiary)
	return w
}

func (w *WithdrawStakeWithSlippage) SetTransferAuthority(authority ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[6] = ag_solanago.Meta(authority).SIGNER()
	w.Signers[0] = ag_solanago.Meta(authority).SIGNER()
	return w
}

func (w *WithdrawStakeWithSlippage) SetBurnFrom(from ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[7] = ag_solanago.Meta(from).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetManagerFeeAccount(account ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[8] = ag_solanago.Meta(account).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetPoolMint(mint ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[9] = ag_solanago.Meta(mint).WRITE()
	return w
}

func (w *WithdrawStakeWithSlippage) SetClock(clock ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[10] = ag_solanago.Meta(clock)
	return w
}

func (w *WithdrawStakeWithSlippage) SetTokenProgram(program ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[11] = ag_solanago.Meta(program)
	return w
}

func (w *WithdrawStakeWithSlippage) SetStakeProgram(program ag_solanago.PublicKey) *WithdrawStakeWithSlippage {
	w.Accounts[12] = ag_solanago.Meta(program)
	return w
}

func (w *WithdrawStakeWithSlippage) GetPoolTokensIn() *uint64 {
	return w.PoolTokensIn
}

func (w *WithdrawStakeWithSlippage) GetMinLamportsOut() *uint64 {
	return w.MinLamportsOut
}

func (w *WithdrawStakeWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return w.Accounts[0].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetValidatorList() ag_solanago.PublicKey {
	return w.Accounts[1].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return w.Accounts[2].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetSplitFrom() ag_solanago.PublicKey {
	return w.Accounts[3].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetSplitTo() ag_solanago.PublicKey {
	return w.Accounts[4].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetBeneficiary() ag_solanago.PublicKey {
	return w.Accounts[5].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetTransferAuthority() ag_solanago.PublicKey {
	return w.Accounts[6].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetBurnFrom() ag_solanago.PublicKey {
	return w.Accounts[7].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return w.Accounts[8].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return w.Accounts[9].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetClock() ag_solanago.PublicKey {
	return w.Accounts[10].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return w.Accounts[11].PublicKey
}

func (w *WithdrawStakeWithSlippage) GetStakeProgram() ag_solanago.PublicKey {
	return w.Accounts[12].PublicKey
}

func (w *WithdrawStakeWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := w.Validate(); err != nil {
		return nil, err
	}
	return w.Build(), nil
}

func (w *WithdrawStakeWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawStakeWithSlippage),
			Impl:   w,
		},
	}
}

func (w *WithdrawStakeWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawStakeWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if w.PoolTokensIn != nil {
							paramsBranch.Child(ag_format.Param("PoolTokensIn", *w.PoolTokensIn))
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

func (w *WithdrawStakeWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if w.PoolTokensIn != nil {
		if err := encoder.Encode(w.PoolTokensIn); err != nil {
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

func (w *WithdrawStakeWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if w.PoolTokensIn != nil {
		if err := decoder.Decode(w.PoolTokensIn); err != nil {
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

func (w *WithdrawStakeWithSlippage) Validate() error {
	if w.PoolTokensIn == nil {
		return errors.New("PoolTokensIn is not set")
	}
	if w.MinLamportsOut == nil {
		return errors.New("MinLamportsOut is not set")
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
