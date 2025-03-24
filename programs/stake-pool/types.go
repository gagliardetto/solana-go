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
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

// Fee rate as a ratio, minted on `UpdateStakePoolBalance` as a proportion of
// the rewards
type Fee struct {
	// denominator of the fee ratio
	Denominator uint64
	// numerator of the fee ratio
	Numerator uint64
}

type FutureEpochFee interface {
	isFutureEpochFee()
}

type None struct{}

func (None) isFutureEpochFee() {}

type One struct {
	Fee Fee
}

func (One) isFutureEpochFee() {}

type Two struct {
	Fee Fee
}

func (Two) isFutureEpochFee() {}

// StakeStatus represents the status of the stake account in the validator list, for accounting.
type StakeStatus uint8

const (
	// Stake account is active, there may be a transient stake as well
	StakeStatusActive StakeStatus = iota
	// Only transient stake account exists, when a transient stake is deactivating during validator removal
	StakeStatusDeactivatingTransient
	// No more validator stake accounts exist, entry ready for removal during `UpdateStakePoolBalance`
	StakeStatusReadyForRemoval
	// Only the validator stake account is deactivating, no transient stake account exists
	StakeStatusDeactivatingValidator
	// Both the transient and validator stake account are deactivating, when a validator is removed with a transient stake active
	StakeStatusDeactivatingAll
)

func (s StakeStatus) String() string {
	switch s {
	case StakeStatusActive:
		return "Active"
	case StakeStatusDeactivatingTransient:
		return "DeactivatingTransient"
	case StakeStatusReadyForRemoval:
		return "ReadyForRemoval"
	case StakeStatusDeactivatingValidator:
		return "DeactivatingValidator"
	case StakeStatusDeactivatingAll:
		return "DeactivatingAll"
	default:
		return "Unknown"
	}
}

type FeeType interface {
	isFeeType()
}

type SolReferral struct {
	Fee uint8
}

func (SolReferral) isFeeType() {}

type StakeReferral struct {
	Fee uint8
}

func (StakeReferral) isFeeType() {}

type Epoch struct {
	Fee Fee
}

func (Epoch) isFeeType() {}

type StakeWithdrawal struct {
	Fee Fee
}

func (StakeWithdrawal) isFeeType() {}

type SolDeposit struct {
	Fee Fee
}

func (SolDeposit) isFeeType() {}

type StakeDeposit struct {
	Fee Fee
}

func (StakeDeposit) isFeeType() {}

type SolWithdrawal struct {
	Fee Fee
}

func (SolWithdrawal) isFeeType() {}

type FundingType uint8

const (
	// Sets the stake deposit authority
	FundingTypeStakeDeposit FundingType = iota
	// Sets the SOL deposit authority
	FundingTypeSolDeposit
	// Sets the SOL withdraw authority
	FundingTypeSolWithdraw
)

type AdditionalValidatorStakeArgs struct {
	Lamports           uint64
	TransientStakeSeed uint64
	EphemeralStakeSeed uint64
}

// PreferredValidatorType represents the type of preferred validator.
type PreferredValidatorType uint8

const (
	// Preferred validator is a validator in the validator list
	PreferredValidatorTypeValidator PreferredValidatorType = iota
	// Preferred validator is a validator not in the validator list
	PreferredValidatorTypeUntrustedValidator
)

type UpdateValidatorListBalanceArgs struct {
	StartIndex uint32
	NoMerge    bool
}

func (obj *UpdateValidatorListBalanceArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NoMerge` param:
	err = encoder.Encode(obj.StartIndex)
	if err != nil {
		return err
	}
	// Serialize `StartIndex` param:
	return encoder.Encode(obj.NoMerge)
}

func (obj *UpdateValidatorListBalanceArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NoMerge`:
	err = decoder.Decode(&obj.StartIndex)
	if err != nil {
		return err
	}
	// Deserialize `StartIndex`:
	return decoder.Decode(&obj.NoMerge)
}

// ValidatorStakeInfo represents information about a validator in the pool.
type ValidatorStakeInfo struct {
	// Amount of lamports on the validator stake account, including rent
	// Note that if `last_update_epoch` does not match the current epoch then
	// this field may not be accurate
	ActiveStakeLamports uint64

	// Amount of transient stake delegated to this validator
	// Note that if `last_update_epoch` does not match the current epoch then
	// this field may not be accurate
	TransientStakeLamports uint64

	// Last epoch the active and transient stake lamports fields were updated
	LastUpdateEpoch uint64

	// Transient account seed suffix start, used to derive the transient stake account address
	TransientSeedSuffixStart uint64

	// Transient account seed suffix end, used to derive the transient stake account address
	TransientSeedSuffixEnd uint64

	// Status of the validator stake account
	Status StakeStatus

	// Validator vote account address
	VoteAccountAddress ag_solanago.PublicKey
}
