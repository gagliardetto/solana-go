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

type DepositSolWithSlippage struct {
	LamportsIn   *uint64
	MinTokensOut *uint64
	// [0] = [WRITE] stakePool
	// [1] = [] withdrawAuthority
	// [2] = [WRITE] reserveStake
	// [3] = [SIGNER] depositFrom
	// [4] = [WRITE] mintTo
	// [5] = [WRITE] managerFeeAccount
	// [6] = [WRITE] referralFeeDest
	// [7] = [WRITE] poolMint
	// [8] = [] systemProgram
	// [9] = [] tokenProgram
	Accounts ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

func DepositSolWithSlippageInstruction(
	// Parameters:
	lamportsIn uint64,
	minTokensOut uint64,
	// Accounts:
	stakePool ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	reserveStake ag_solanago.PublicKey,
	depositFrom ag_solanago.PublicKey,
	mintTo ag_solanago.PublicKey,
	managerFeeAccount ag_solanago.PublicKey,
	referralFeeDest ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
) *DepositSolWithSlippage {
	return NewDepositSolWithSlippageBuilder().
		SetLamportsIn(lamportsIn).
		SetMinTokensOut(minTokensOut).
		SetStakePool(stakePool).
		SetWithdrawAuthority(withdrawAuthority).
		SetReserveStake(reserveStake).
		SetDepositFrom(depositFrom).
		SetMintTo(mintTo).
		SetManagerFeeAccount(managerFeeAccount).
		SetReferralFeeDest(referralFeeDest).
		SetPoolMint(poolMint).
		SetSystemProgram(systemProgram).
		SetTokenProgram(tokenProgram)
}

func NewDepositSolWithSlippageBuilder() *DepositSolWithSlippage {
	return &DepositSolWithSlippage{
		Accounts: make(ag_solanago.AccountMetaSlice, 10),
		Signers:  make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *DepositSolWithSlippage) SetLamportsIn(in uint64) *DepositSolWithSlippage {
	inst.LamportsIn = &in
	return inst
}

func (inst *DepositSolWithSlippage) SetMinTokensOut(out uint64) *DepositSolWithSlippage {
	inst.MinTokensOut = &out
	return inst
}

func (inst *DepositSolWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[1] = ag_solanago.Meta(authority)
	return inst
}

func (inst *DepositSolWithSlippage) SetReserveStake(stake ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[2] = ag_solanago.Meta(stake).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetDepositFrom(from ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[3] = ag_solanago.Meta(from).SIGNER()
	inst.Signers[0] = ag_solanago.Meta(from).SIGNER()
	return inst
}

func (inst *DepositSolWithSlippage) SetMintTo(to ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[4] = ag_solanago.Meta(to).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetManagerFeeAccount(account ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[5] = ag_solanago.Meta(account).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetReferralFeeDest(dest ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[6] = ag_solanago.Meta(dest).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetPoolMint(mint ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[7] = ag_solanago.Meta(mint).WRITE()
	return inst
}

func (inst *DepositSolWithSlippage) SetSystemProgram(program ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[8] = ag_solanago.Meta(program)
	return inst
}

func (inst *DepositSolWithSlippage) SetTokenProgram(program ag_solanago.PublicKey) *DepositSolWithSlippage {
	inst.Accounts[9] = ag_solanago.Meta(program)
	return inst
}

func (inst *DepositSolWithSlippage) GetLamportsIn() *uint64 {
	return inst.LamportsIn
}

func (inst *DepositSolWithSlippage) GetMinTokensOut() *uint64 {
	return inst.MinTokensOut
}

func (inst *DepositSolWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return inst.Accounts[0].PublicKey
}

func (inst *DepositSolWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return inst.Accounts[1].PublicKey
}

func (inst *DepositSolWithSlippage) GetReserveStake() ag_solanago.PublicKey {
	return inst.Accounts[2].PublicKey
}

func (inst *DepositSolWithSlippage) GetDepositFrom() ag_solanago.PublicKey {
	return inst.Accounts[3].PublicKey
}

func (inst *DepositSolWithSlippage) GetMintTo() ag_solanago.PublicKey {
	return inst.Accounts[4].PublicKey
}

func (inst *DepositSolWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return inst.Accounts[5].PublicKey
}

func (inst *DepositSolWithSlippage) GetReferralFeeDest() ag_solanago.PublicKey {
	return inst.Accounts[6].PublicKey
}

func (inst *DepositSolWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return inst.Accounts[7].PublicKey
}

func (inst *DepositSolWithSlippage) GetSystemProgram() ag_solanago.PublicKey {
	return inst.Accounts[8].PublicKey
}

func (inst *DepositSolWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return inst.Accounts[9].PublicKey
}

func (inst *DepositSolWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DepositSolWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DepositSolWithSlippage),
			Impl:   inst,
		},
	}
}

func (inst *DepositSolWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositSolWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if inst.LamportsIn != nil {
							paramsBranch.Child(ag_format.Param("LamportsIn", *inst.LamportsIn))
						}
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

func (inst *DepositSolWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if inst.LamportsIn != nil {
		if err := encoder.Encode(inst.LamportsIn); err != nil {
			return err
		}
	}
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

func (inst *DepositSolWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if inst.LamportsIn != nil {
		if err := decoder.Decode(inst.LamportsIn); err != nil {
			return err
		}
	}
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

func (inst *DepositSolWithSlippage) Validate() error {
	if inst.LamportsIn == nil {
		return errors.New("lamportsIn is not set")
	}
	if inst.MinTokensOut == nil {
		return errors.New("minTokensOut is not set")
	}
	for i, account := range inst.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(inst.Signers) == 0 || !inst.Signers[0].IsSigner {
		return errors.New("accounts.DepositFrom should be a signer")
	}
	return nil
}
