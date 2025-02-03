// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SwapWithFreeFunds is the `SwapWithFreeFunds` instruction.
type SwapWithFreeFunds struct {
	OrderPacket OrderPacket

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [SIGNER] trader
	//
	// [4] = [] seat
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSwapWithFreeFundsInstructionBuilder creates a new `SwapWithFreeFunds` instruction builder.
func NewSwapWithFreeFundsInstructionBuilder() *SwapWithFreeFunds {
	nd := &SwapWithFreeFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetOrderPacket sets the "orderPacket" parameter.
func (inst *SwapWithFreeFunds) SetOrderPacket(orderPacket OrderPacket) *SwapWithFreeFunds {
	inst.OrderPacket = orderPacket
	return inst
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *SwapWithFreeFunds) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *SwapWithFreeFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *SwapWithFreeFunds) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *SwapWithFreeFunds) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *SwapWithFreeFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *SwapWithFreeFunds) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *SwapWithFreeFunds) SetMarketAccount(market ag_solanago.PublicKey) *SwapWithFreeFunds {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *SwapWithFreeFunds) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTraderAccount sets the "trader" account.
func (inst *SwapWithFreeFunds) SetTraderAccount(trader ag_solanago.PublicKey) *SwapWithFreeFunds {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(trader).SIGNER()
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *SwapWithFreeFunds) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSeatAccount sets the "seat" account.
func (inst *SwapWithFreeFunds) SetSeatAccount(seat ag_solanago.PublicKey) *SwapWithFreeFunds {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(seat)
	return inst
}

// GetSeatAccount gets the "seat" account.
func (inst *SwapWithFreeFunds) GetSeatAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst SwapWithFreeFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SwapWithFreeFunds),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SwapWithFreeFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SwapWithFreeFunds) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OrderPacket == nil {
			return errors.New("OrderPacket parameter is not set")
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
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Seat is not set")
		}
	}
	return nil
}

func (inst *SwapWithFreeFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SwapWithFreeFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("OrderPacket", inst.OrderPacket))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        trader", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          seat", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj SwapWithFreeFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrderPacket` param:
	{
		tmp := orderPacketContainer{}
		switch realvalue := obj.OrderPacket.(type) {
		case *OrderPacketPostOnly:
			tmp.Enum = 0
			tmp.PostOnly = *realvalue
		case *OrderPacketLimit:
			tmp.Enum = 1
			tmp.Limit = *realvalue
		case *OrderPacketImmediateOrCancel:
			tmp.Enum = 2
			tmp.ImmediateOrCancel = *realvalue
		}
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (obj *SwapWithFreeFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrderPacket`:
	{
		tmp := new(orderPacketContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.OrderPacket = &tmp.PostOnly
		case 1:
			obj.OrderPacket = &tmp.Limit
		case 2:
			obj.OrderPacket = &tmp.ImmediateOrCancel
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	return nil
}

// NewSwapWithFreeFundsInstruction declares a new SwapWithFreeFunds instruction with the provided parameters and accounts.
func NewSwapWithFreeFundsInstruction(
	// Parameters:
	orderPacket OrderPacket,
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	trader ag_solanago.PublicKey,
	seat ag_solanago.PublicKey) *SwapWithFreeFunds {
	return NewSwapWithFreeFundsInstructionBuilder().
		SetOrderPacket(orderPacket).
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetTraderAccount(trader).
		SetSeatAccount(seat)
}
