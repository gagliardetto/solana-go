// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelAllOrdersWithFreeFunds is the `CancelAllOrdersWithFreeFunds` instruction.
type CancelAllOrdersWithFreeFunds struct {

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [SIGNER] trader
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelAllOrdersWithFreeFundsInstructionBuilder creates a new `CancelAllOrdersWithFreeFunds` instruction builder.
func NewCancelAllOrdersWithFreeFundsInstructionBuilder() *CancelAllOrdersWithFreeFunds {
	nd := &CancelAllOrdersWithFreeFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *CancelAllOrdersWithFreeFunds) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *CancelAllOrdersWithFreeFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *CancelAllOrdersWithFreeFunds) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *CancelAllOrdersWithFreeFunds) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *CancelAllOrdersWithFreeFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *CancelAllOrdersWithFreeFunds) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *CancelAllOrdersWithFreeFunds) SetMarketAccount(market ag_solanago.PublicKey) *CancelAllOrdersWithFreeFunds {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CancelAllOrdersWithFreeFunds) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTraderAccount sets the "trader" account.
func (inst *CancelAllOrdersWithFreeFunds) SetTraderAccount(trader ag_solanago.PublicKey) *CancelAllOrdersWithFreeFunds {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(trader).SIGNER()
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *CancelAllOrdersWithFreeFunds) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CancelAllOrdersWithFreeFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CancelAllOrdersWithFreeFunds),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelAllOrdersWithFreeFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelAllOrdersWithFreeFunds) Validate() error {
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

func (inst *CancelAllOrdersWithFreeFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelAllOrdersWithFreeFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

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

func (obj CancelAllOrdersWithFreeFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CancelAllOrdersWithFreeFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCancelAllOrdersWithFreeFundsInstruction declares a new CancelAllOrdersWithFreeFunds instruction with the provided parameters and accounts.
func NewCancelAllOrdersWithFreeFundsInstruction(
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	trader ag_solanago.PublicKey) *CancelAllOrdersWithFreeFunds {
	return NewCancelAllOrdersWithFreeFundsInstructionBuilder().
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetTraderAccount(trader)
}
