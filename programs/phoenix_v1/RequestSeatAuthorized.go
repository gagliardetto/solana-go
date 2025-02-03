// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RequestSeatAuthorized is the `RequestSeatAuthorized` instruction.
type RequestSeatAuthorized struct {

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [SIGNER] marketAuthority
	//
	// [4] = [WRITE, SIGNER] payer
	//
	// [5] = [] trader
	//
	// [6] = [WRITE] seat
	//
	// [7] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRequestSeatAuthorizedInstructionBuilder creates a new `RequestSeatAuthorized` instruction builder.
func NewRequestSeatAuthorizedInstructionBuilder() *RequestSeatAuthorized {
	nd := &RequestSeatAuthorized{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *RequestSeatAuthorized) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *RequestSeatAuthorized) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *RequestSeatAuthorized) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *RequestSeatAuthorized) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *RequestSeatAuthorized) SetMarketAccount(market ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *RequestSeatAuthorized) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMarketAuthorityAccount sets the "marketAuthority" account.
func (inst *RequestSeatAuthorized) SetMarketAuthorityAccount(marketAuthority ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(marketAuthority).SIGNER()
	return inst
}

// GetMarketAuthorityAccount gets the "marketAuthority" account.
func (inst *RequestSeatAuthorized) GetMarketAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *RequestSeatAuthorized) SetPayerAccount(payer ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *RequestSeatAuthorized) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTraderAccount sets the "trader" account.
func (inst *RequestSeatAuthorized) SetTraderAccount(trader ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(trader)
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *RequestSeatAuthorized) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSeatAccount sets the "seat" account.
func (inst *RequestSeatAuthorized) SetSeatAccount(seat ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(seat).WRITE()
	return inst
}

// GetSeatAccount gets the "seat" account.
func (inst *RequestSeatAuthorized) GetSeatAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RequestSeatAuthorized) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RequestSeatAuthorized {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RequestSeatAuthorized) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst RequestSeatAuthorized) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RequestSeatAuthorized),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RequestSeatAuthorized) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RequestSeatAuthorized) Validate() error {
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
			return errors.New("accounts.MarketAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Trader is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Seat is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RequestSeatAuthorized) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RequestSeatAuthorized")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("marketAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         trader", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           seat", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj RequestSeatAuthorized) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RequestSeatAuthorized) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRequestSeatAuthorizedInstruction declares a new RequestSeatAuthorized instruction with the provided parameters and accounts.
func NewRequestSeatAuthorizedInstruction(
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	marketAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	trader ag_solanago.PublicKey,
	seat ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RequestSeatAuthorized {
	return NewRequestSeatAuthorizedInstructionBuilder().
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetMarketAuthorityAccount(marketAuthority).
		SetPayerAccount(payer).
		SetTraderAccount(trader).
		SetSeatAccount(seat).
		SetSystemProgramAccount(systemProgram)
}
