// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ReduceOrderWithFreeFunds is the `ReduceOrderWithFreeFunds` instruction.
type ReduceOrderWithFreeFunds struct {
	Params *ReduceOrderParams

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [WRITE, SIGNER] trader
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReduceOrderWithFreeFundsInstructionBuilder creates a new `ReduceOrderWithFreeFunds` instruction builder.
func NewReduceOrderWithFreeFundsInstructionBuilder() *ReduceOrderWithFreeFunds {
	nd := &ReduceOrderWithFreeFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *ReduceOrderWithFreeFunds) SetParams(params ReduceOrderParams) *ReduceOrderWithFreeFunds {
	inst.Params = &params
	return inst
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *ReduceOrderWithFreeFunds) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *ReduceOrderWithFreeFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *ReduceOrderWithFreeFunds) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *ReduceOrderWithFreeFunds) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *ReduceOrderWithFreeFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *ReduceOrderWithFreeFunds) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *ReduceOrderWithFreeFunds) SetMarketAccount(market ag_solanago.PublicKey) *ReduceOrderWithFreeFunds {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *ReduceOrderWithFreeFunds) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTraderAccount sets the "trader" account.
func (inst *ReduceOrderWithFreeFunds) SetTraderAccount(trader ag_solanago.PublicKey) *ReduceOrderWithFreeFunds {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(trader).WRITE().SIGNER()
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *ReduceOrderWithFreeFunds) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst ReduceOrderWithFreeFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ReduceOrderWithFreeFunds,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ReduceOrderWithFreeFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ReduceOrderWithFreeFunds) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PhoenixProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LogAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Trader is not set")
		}
	}
	return nil
}

func (inst *ReduceOrderWithFreeFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ReduceOrderWithFreeFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        trader", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj ReduceOrderWithFreeFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ReduceOrderWithFreeFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewReduceOrderWithFreeFundsInstruction declares a new ReduceOrderWithFreeFunds instruction with the provided parameters and accounts.
func NewReduceOrderWithFreeFundsInstruction(
	// Parameters:
	params ReduceOrderParams,
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	trader ag_solanago.PublicKey) *ReduceOrderWithFreeFunds {
	return NewReduceOrderWithFreeFundsInstructionBuilder().
		SetParams(params).
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetTraderAccount(trader)
}
