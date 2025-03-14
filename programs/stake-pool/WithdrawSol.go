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

type WithdrawSol struct {
	Arg *uint64
	// [0] = [WRITE] stakePool
	// [1] = [] withdrawAuthority
	// [2] = [] transferAuthority
	// [3] = [WRITE] burnPoolTokens
	// [4] = [WRITE] reserveStakeAccount
	// [5] = [WRITE] withdrawAccount
	// [6] = [WRITE] feeTokenAccount
	// [7] = [WRITE] poolTokenMint
	// [8] = [] sysvarClock
	// [9] = [] sysvarStakeHistory
	// [10] = [] stakeProgram
	// [11] = [] tokenProgram
	// [12] = [SIGNER] solWithdrawAuthority
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func NewWithdrawSolInstruction(
	arg uint64,
	stakePool ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	burnPoolTokens ag_solanago.PublicKey,
	reserveStakeAccount ag_solanago.PublicKey,
	withdrawAccount ag_solanago.PublicKey,
	feeTokenAccount ag_solanago.PublicKey,
	poolTokenMint ag_solanago.PublicKey,
	sysvarClock ag_solanago.PublicKey,
	sysvarStakeHistory ag_solanago.PublicKey,
	stakeProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	solWithdrawAuthority ag_solanago.PublicKey,
) *WithdrawSol {
	return NewWithdrawSolBuilder().
		SetArg(arg).
		SetStakePool(stakePool).
		SetWithdrawAuthority(withdrawAuthority).
		SetTransferAuthority(transferAuthority).
		SetBurnPoolTokens(burnPoolTokens).
		SetReserveStakeAccount(reserveStakeAccount).
		SetWithdrawAccount(withdrawAccount).
		SetFeeTokenAccount(feeTokenAccount).
		SetPoolTokenMint(poolTokenMint).
		SetSysvarClock(sysvarClock).
		SetSysvarStakeHistory(sysvarStakeHistory).
		SetStakeProgram(stakeProgram).
		SetTokenProgram(tokenProgram).
		SetSolWithdrawAuthority(solWithdrawAuthority)
}

func NewWithdrawSolBuilder() *WithdrawSol {
	return &WithdrawSol{
		Accounts: make(ag_solanago.AccountMetaSlice, 13),
		Signers:  make(ag_solanago.AccountMetaSlice, 0),
	}
}

func (inst *WithdrawSol) SetArg(arg uint64) *WithdrawSol {
	inst.Arg = &arg
	return inst
}

func (inst *WithdrawSol) SetStakePool(stakePool ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[0] = ag_solanago.Meta(stakePool).WRITE()
	return inst
}

func (inst *WithdrawSol) SetWithdrawAuthority(withdrawAuthority ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[1] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

func (inst *WithdrawSol) SetTransferAuthority(transferAuthority ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[2] = ag_solanago.Meta(transferAuthority).SIGNER()
	return inst
}

func (inst *WithdrawSol) SetBurnPoolTokens(burnPoolTokens ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[3] = ag_solanago.Meta(burnPoolTokens).WRITE()
	return inst
}

func (inst *WithdrawSol) SetReserveStakeAccount(reserveStakeAccount ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[4] = ag_solanago.Meta(reserveStakeAccount).WRITE()
	return inst
}

func (inst *WithdrawSol) SetWithdrawAccount(withdrawAccount ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[5] = ag_solanago.Meta(withdrawAccount).WRITE()
	return inst
}

func (inst *WithdrawSol) SetFeeTokenAccount(feeTokenAccount ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[6] = ag_solanago.Meta(feeTokenAccount).WRITE()
	return inst
}

func (inst *WithdrawSol) SetPoolTokenMint(poolTokenMint ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[7] = ag_solanago.Meta(poolTokenMint).WRITE()
	return inst
}

func (inst *WithdrawSol) SetSysvarClock(sysvarClock ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[8] = ag_solanago.Meta(sysvarClock)
	return inst
}

func (inst *WithdrawSol) SetSysvarStakeHistory(sysvarStakeHistory ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[9] = ag_solanago.Meta(sysvarStakeHistory)
	return inst
}

func (inst *WithdrawSol) SetStakeProgram(stakeProgram ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[10] = ag_solanago.Meta(stakeProgram)
	return inst
}

func (inst *WithdrawSol) SetTokenProgram(tokenProgram ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

func (inst *WithdrawSol) SetSolWithdrawAuthority(solWithdrawAuthority ag_solanago.PublicKey) *WithdrawSol {
	inst.Accounts[12] = ag_solanago.Meta(solWithdrawAuthority).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(solWithdrawAuthority).SIGNER()
	return inst
}

func (inst *WithdrawSol) GetArg() *uint64 {
	return inst.Arg
}

func (inst *WithdrawSol) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *WithdrawSol) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *WithdrawSol) GetTransferAuthority() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *WithdrawSol) GetBurnPoolTokens() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *WithdrawSol) GetReserveStakeAccount() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *WithdrawSol) GetWithdrawAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *WithdrawSol) GetFeeTokenAccount() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *WithdrawSol) GetPoolTokenMint() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *WithdrawSol) GetSysvarClock() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *WithdrawSol) GetSysvarStakeHistory() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *WithdrawSol) GetStakeProgram() ag_solanago.PublicKey {
	return inst.Accounts[10].PublicKey
}

func (inst *WithdrawSol) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[11].PublicKey
}

func (inst *WithdrawSol) GetSolWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[12].PublicKey
}

func (inst *WithdrawSol) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawSol) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_WithdrawSol),
			Impl:   inst,
		},
	}
}

func (inst *WithdrawSol) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawSol")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.Arg != nil {
							paramsBranch.Child(ag_format.Param("Arg", *inst.Arg))
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

func (inst *WithdrawSol) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.Arg != nil {
		if err := encoder.Encode(inst.Arg); err != nil {
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

func (inst *WithdrawSol) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.Arg != nil {
		if err := decoder.Decode(inst.Arg); err != nil {
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

func (inst *WithdrawSol) Validate() error {
	if inst.Arg == nil {
		return errors.New("arg is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.solWithdrawAuthority should be a signer")
	}
	return nil
}
