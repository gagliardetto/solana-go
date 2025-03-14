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

func (d *DepositSolWithSlippage) SetLamportsIn(in uint64) *DepositSolWithSlippage {
	d.LamportsIn = &in
	return d
}

func (d *DepositSolWithSlippage) SetMinTokensOut(out uint64) *DepositSolWithSlippage {
	d.MinTokensOut = &out
	return d
}

func (d *DepositSolWithSlippage) SetStakePool(pool ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[0] = ag_solanago.Meta(pool).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetWithdrawAuthority(authority ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[1] = ag_solanago.Meta(authority)
	return d
}

func (d *DepositSolWithSlippage) SetReserveStake(stake ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[2] = ag_solanago.Meta(stake).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetDepositFrom(from ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[3] = ag_solanago.Meta(from).SIGNER()
	d.Signers[0] = ag_solanago.Meta(from).SIGNER()
	return d
}

func (d *DepositSolWithSlippage) SetMintTo(to ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[4] = ag_solanago.Meta(to).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetManagerFeeAccount(account ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[5] = ag_solanago.Meta(account).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetReferralFeeDest(dest ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[6] = ag_solanago.Meta(dest).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetPoolMint(mint ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[7] = ag_solanago.Meta(mint).WRITE()
	return d
}

func (d *DepositSolWithSlippage) SetSystemProgram(program ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[8] = ag_solanago.Meta(program)
	return d
}

func (d *DepositSolWithSlippage) SetTokenProgram(program ag_solanago.PublicKey) *DepositSolWithSlippage {
	d.Accounts[9] = ag_solanago.Meta(program)
	return d
}

func (d *DepositSolWithSlippage) GetLamportsIn() *uint64 {
	return d.LamportsIn
}

func (d *DepositSolWithSlippage) GetMinTokensOut() *uint64 {
	return d.MinTokensOut
}

func (d *DepositSolWithSlippage) GetStakePool() ag_solanago.PublicKey {
	return d.Accounts[0].PublicKey
}

func (d *DepositSolWithSlippage) GetWithdrawAuthority() ag_solanago.PublicKey {
	return d.Accounts[1].PublicKey
}

func (d *DepositSolWithSlippage) GetReserveStake() ag_solanago.PublicKey {
	return d.Accounts[2].PublicKey
}

func (d *DepositSolWithSlippage) GetDepositFrom() ag_solanago.PublicKey {
	return d.Accounts[3].PublicKey
}

func (d *DepositSolWithSlippage) GetMintTo() ag_solanago.PublicKey {
	return d.Accounts[4].PublicKey
}

func (d *DepositSolWithSlippage) GetManagerFeeAccount() ag_solanago.PublicKey {
	return d.Accounts[5].PublicKey
}

func (d *DepositSolWithSlippage) GetReferralFeeDest() ag_solanago.PublicKey {
	return d.Accounts[6].PublicKey
}

func (d *DepositSolWithSlippage) GetPoolMint() ag_solanago.PublicKey {
	return d.Accounts[7].PublicKey
}

func (d *DepositSolWithSlippage) GetSystemProgram() ag_solanago.PublicKey {
	return d.Accounts[8].PublicKey
}

func (d *DepositSolWithSlippage) GetTokenProgram() ag_solanago.PublicKey {
	return d.Accounts[9].PublicKey
}

func (d *DepositSolWithSlippage) ValidateAndBuild() (*Instruction, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d.Build(), nil
}

func (d *DepositSolWithSlippage) Build() *Instruction {
	return &Instruction{
		BaseVariant: ag_binary.BaseVariant{
			TypeID: ag_binary.TypeIDFromUint8(Instruction_DepositSolWithSlippage),
			Impl:   d,
		},
	}
}

func (d *DepositSolWithSlippage) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DepositSolWithSlippage")).
				ParentFunc(func(instructionBranch ag_treeout.Branches) {
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						if d.LamportsIn != nil {
							paramsBranch.Child(ag_format.Param("LamportsIn", *d.LamportsIn))
						}
						if d.MinTokensOut != nil {
							paramsBranch.Child(ag_format.Param("MinTokensOut", *d.MinTokensOut))
						}
					})
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

func (d *DepositSolWithSlippage) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if d.LamportsIn != nil {
		if err := encoder.Encode(d.LamportsIn); err != nil {
			return err
		}
	}
	if d.MinTokensOut != nil {
		if err := encoder.Encode(d.MinTokensOut); err != nil {
			return err
		}
	}
	for _, account := range d.Accounts {
		if err := encoder.Encode(account); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepositSolWithSlippage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if d.LamportsIn != nil {
		if err := decoder.Decode(d.LamportsIn); err != nil {
			return err
		}
	}
	if d.MinTokensOut != nil {
		if err := decoder.Decode(d.MinTokensOut); err != nil {
			return err
		}
	}
	for i := range d.Accounts {
		if err := decoder.Decode(d.Accounts[i]); err != nil {
			return err
		}
	}
	return nil
}

func (d *DepositSolWithSlippage) Validate() error {
	if d.LamportsIn == nil {
		return errors.New("lamportsIn is not set")
	}
	if d.MinTokensOut == nil {
		return errors.New("minTokensOut is not set")
	}
	for i, account := range d.Accounts {
		if account == nil {
			return fmt.Errorf("accounts[%v] is not set", i)
		}
	}
	if len(d.Signers) == 0 || !d.Signers[0].IsSigner {
		return errors.New("accounts.DepositFrom should be a signer")
	}
	return nil
}
