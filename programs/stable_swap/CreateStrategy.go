// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateStrategy is the `create_strategy` instruction.
type CreateStrategy struct {
	AmpMinFactor    *uint16
	AmpMaxFactor    *uint16
	RampMinStep     *uint16
	RampMaxStep     *uint16
	RampMinDuration *uint32
	RampMaxDuration *uint32

	// ····· owner_only: [0] = [WRITE] pool
	//
	// [1] = [WRITE] strategy
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateStrategyInstructionBuilder creates a new `CreateStrategy` instruction builder.
func NewCreateStrategyInstructionBuilder() *CreateStrategy {
	nd := &CreateStrategy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetAmpMinFactor sets the "amp_min_factor" parameter.
func (inst *CreateStrategy) SetAmpMinFactor(amp_min_factor uint16) *CreateStrategy {
	inst.AmpMinFactor = &amp_min_factor
	return inst
}

// SetAmpMaxFactor sets the "amp_max_factor" parameter.
func (inst *CreateStrategy) SetAmpMaxFactor(amp_max_factor uint16) *CreateStrategy {
	inst.AmpMaxFactor = &amp_max_factor
	return inst
}

// SetRampMinStep sets the "ramp_min_step" parameter.
func (inst *CreateStrategy) SetRampMinStep(ramp_min_step uint16) *CreateStrategy {
	inst.RampMinStep = &ramp_min_step
	return inst
}

// SetRampMaxStep sets the "ramp_max_step" parameter.
func (inst *CreateStrategy) SetRampMaxStep(ramp_max_step uint16) *CreateStrategy {
	inst.RampMaxStep = &ramp_max_step
	return inst
}

// SetRampMinDuration sets the "ramp_min_duration" parameter.
func (inst *CreateStrategy) SetRampMinDuration(ramp_min_duration uint32) *CreateStrategy {
	inst.RampMinDuration = &ramp_min_duration
	return inst
}

// SetRampMaxDuration sets the "ramp_max_duration" parameter.
func (inst *CreateStrategy) SetRampMaxDuration(ramp_max_duration uint32) *CreateStrategy {
	inst.RampMaxDuration = &ramp_max_duration
	return inst
}

type CreateStrategyOwnerOnlyAccountsBuilder struct {
	ag_solanago.AccountMetaSlice `bin:"-"`
}

func NewCreateStrategyOwnerOnlyAccountsBuilder() *CreateStrategyOwnerOnlyAccountsBuilder {
	return &CreateStrategyOwnerOnlyAccountsBuilder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *CreateStrategy) SetOwnerOnlyAccountsFromBuilder(createStrategyOwnerOnlyAccountsBuilder *CreateStrategyOwnerOnlyAccountsBuilder) *CreateStrategy {
	inst.AccountMetaSlice[1] = createStrategyOwnerOnlyAccountsBuilder.GetPoolAccount()
	return inst
}

// SetPoolAccount sets the "pool" account.
func (inst *CreateStrategyOwnerOnlyAccountsBuilder) SetPoolAccount(pool ag_solanago.PublicKey) *CreateStrategyOwnerOnlyAccountsBuilder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *CreateStrategyOwnerOnlyAccountsBuilder) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStrategyAccount sets the "strategy" account.
func (inst *CreateStrategy) SetStrategyAccount(strategy ag_solanago.PublicKey) *CreateStrategy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(strategy).WRITE()
	return inst
}

// GetStrategyAccount gets the "strategy" account.
func (inst *CreateStrategy) GetStrategyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst CreateStrategy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateStrategy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateStrategy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmpMinFactor == nil {
			return errors.New("AmpMinFactor parameter is not set")
		}
		if inst.AmpMaxFactor == nil {
			return errors.New("AmpMaxFactor parameter is not set")
		}
		if inst.RampMinStep == nil {
			return errors.New("RampMinStep parameter is not set")
		}
		if inst.RampMaxStep == nil {
			return errors.New("RampMaxStep parameter is not set")
		}
		if inst.RampMinDuration == nil {
			return errors.New("RampMinDuration parameter is not set")
		}
		if inst.RampMaxDuration == nil {
			return errors.New("RampMaxDuration parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.OwnerOnlyPool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Strategy is not set")
		}
	}
	return nil
}

func (inst *CreateStrategy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     AmpMinFactor", *inst.AmpMinFactor))
						paramsBranch.Child(ag_format.Param("     AmpMaxFactor", *inst.AmpMaxFactor))
						paramsBranch.Child(ag_format.Param("      RampMinStep", *inst.RampMinStep))
						paramsBranch.Child(ag_format.Param("      RampMaxStep", *inst.RampMaxStep))
						paramsBranch.Child(ag_format.Param("  RampMinDuration", *inst.RampMinDuration))
						paramsBranch.Child(ag_format.Param("  RampMaxDuration", *inst.RampMaxDuration))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("owner_only/pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       strategy", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj CreateStrategy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmpMinFactor` param:
	err = encoder.Encode(obj.AmpMinFactor)
	if err != nil {
		return err
	}
	// Serialize `AmpMaxFactor` param:
	err = encoder.Encode(obj.AmpMaxFactor)
	if err != nil {
		return err
	}
	// Serialize `RampMinStep` param:
	err = encoder.Encode(obj.RampMinStep)
	if err != nil {
		return err
	}
	// Serialize `RampMaxStep` param:
	err = encoder.Encode(obj.RampMaxStep)
	if err != nil {
		return err
	}
	// Serialize `RampMinDuration` param:
	err = encoder.Encode(obj.RampMinDuration)
	if err != nil {
		return err
	}
	// Serialize `RampMaxDuration` param:
	err = encoder.Encode(obj.RampMaxDuration)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateStrategy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmpMinFactor`:
	err = decoder.Decode(&obj.AmpMinFactor)
	if err != nil {
		return err
	}
	// Deserialize `AmpMaxFactor`:
	err = decoder.Decode(&obj.AmpMaxFactor)
	if err != nil {
		return err
	}
	// Deserialize `RampMinStep`:
	err = decoder.Decode(&obj.RampMinStep)
	if err != nil {
		return err
	}
	// Deserialize `RampMaxStep`:
	err = decoder.Decode(&obj.RampMaxStep)
	if err != nil {
		return err
	}
	// Deserialize `RampMinDuration`:
	err = decoder.Decode(&obj.RampMinDuration)
	if err != nil {
		return err
	}
	// Deserialize `RampMaxDuration`:
	err = decoder.Decode(&obj.RampMaxDuration)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateStrategyInstruction declares a new CreateStrategy instruction with the provided parameters and accounts.
func NewCreateStrategyInstruction(
	// Parameters:
	amp_min_factor uint16,
	amp_max_factor uint16,
	ramp_min_step uint16,
	ramp_max_step uint16,
	ramp_min_duration uint32,
	ramp_max_duration uint32,
	// Accounts:
	ownerOnlyPool ag_solanago.PublicKey,
	strategy ag_solanago.PublicKey) *CreateStrategy {
	return NewCreateStrategyInstructionBuilder().
		SetAmpMinFactor(amp_min_factor).
		SetAmpMaxFactor(amp_max_factor).
		SetRampMinStep(ramp_min_step).
		SetRampMaxStep(ramp_max_step).
		SetRampMinDuration(ramp_min_duration).
		SetRampMaxDuration(ramp_max_duration).
		SetOwnerOnlyAccountsFromBuilder(
			NewCreateStrategyOwnerOnlyAccountsBuilder().
				SetPoolAccount(ownerOnlyPool),
		).
		SetStrategyAccount(strategy)
}
