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

// Swap is the `Swap` instruction.
type Swap struct {
	OrderPacket OrderPacket

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [SIGNER] trader
	//
	// [4] = [WRITE] baseAccount
	//
	// [5] = [WRITE] quoteAccount
	//
	// [6] = [WRITE] baseVault
	//
	// [7] = [WRITE] quoteVault
	//
	// [8] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSwapInstructionBuilder creates a new `Swap` instruction builder.
func NewSwapInstructionBuilder() *Swap {
	nd := &Swap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetOrderPacket sets the "orderPacket" parameter.
func (inst *Swap) SetOrderPacket(orderPacket OrderPacket) *Swap {
	inst.OrderPacket = orderPacket
	return inst
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *Swap) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *Swap) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *Swap) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *Swap) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *Swap) SetMarketAccount(market ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *Swap) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTraderAccount sets the "trader" account.
func (inst *Swap) SetTraderAccount(trader ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(trader).SIGNER()
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *Swap) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBaseAccountAccount sets the "baseAccount" account.
func (inst *Swap) SetBaseAccountAccount(baseAccount ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(baseAccount).WRITE()
	return inst
}

// GetBaseAccountAccount gets the "baseAccount" account.
func (inst *Swap) GetBaseAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetQuoteAccountAccount sets the "quoteAccount" account.
func (inst *Swap) SetQuoteAccountAccount(quoteAccount ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(quoteAccount).WRITE()
	return inst
}

// GetQuoteAccountAccount gets the "quoteAccount" account.
func (inst *Swap) GetQuoteAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBaseVaultAccount sets the "baseVault" account.
func (inst *Swap) SetBaseVaultAccount(baseVault ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(baseVault).WRITE()
	return inst
}

// GetBaseVaultAccount gets the "baseVault" account.
func (inst *Swap) GetBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetQuoteVaultAccount sets the "quoteVault" account.
func (inst *Swap) SetQuoteVaultAccount(quoteVault ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(quoteVault).WRITE()
	return inst
}

// GetQuoteVaultAccount gets the "quoteVault" account.
func (inst *Swap) GetQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Swap) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Swap) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst Swap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Swap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Swap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Swap) Validate() error {
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
			return errors.New("accounts.BaseAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.QuoteAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.BaseVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.QuoteVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *Swap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Swap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("OrderPacket", inst.OrderPacket))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        trader", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          base", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         quote", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     baseVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("    quoteVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj Swap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *Swap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewSwapInstruction declares a new Swap instruction with the provided parameters and accounts.
func NewSwapInstruction(
	// Parameters:
	orderPacket OrderPacket,
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	trader ag_solanago.PublicKey,
	baseAccount ag_solanago.PublicKey,
	quoteAccount ag_solanago.PublicKey,
	baseVault ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *Swap {
	return NewSwapInstructionBuilder().
		SetOrderPacket(orderPacket).
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetTraderAccount(trader).
		SetBaseAccountAccount(baseAccount).
		SetQuoteAccountAccount(quoteAccount).
		SetBaseVaultAccount(baseVault).
		SetQuoteVaultAccount(quoteVault).
		SetTokenProgramAccount(tokenProgram)
}
