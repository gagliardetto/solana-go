// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelUpTo is the `CancelUpTo` instruction.
type CancelUpTo struct {
	Params *CancelUpToParams

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

// NewCancelUpToInstructionBuilder creates a new `CancelUpTo` instruction builder.
func NewCancelUpToInstructionBuilder() *CancelUpTo {
	nd := &CancelUpTo{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *CancelUpTo) SetParams(params CancelUpToParams) *CancelUpTo {
	inst.Params = &params
	return inst
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *CancelUpTo) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *CancelUpTo) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *CancelUpTo) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *CancelUpTo) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *CancelUpTo) SetMarketAccount(market ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CancelUpTo) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTraderAccount sets the "trader" account.
func (inst *CancelUpTo) SetTraderAccount(trader ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(trader).SIGNER()
	return inst
}

// GetTraderAccount gets the "trader" account.
func (inst *CancelUpTo) GetTraderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBaseAccountAccount sets the "baseAccount" account.
func (inst *CancelUpTo) SetBaseAccountAccount(baseAccount ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(baseAccount).WRITE()
	return inst
}

// GetBaseAccountAccount gets the "baseAccount" account.
func (inst *CancelUpTo) GetBaseAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetQuoteAccountAccount sets the "quoteAccount" account.
func (inst *CancelUpTo) SetQuoteAccountAccount(quoteAccount ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(quoteAccount).WRITE()
	return inst
}

// GetQuoteAccountAccount gets the "quoteAccount" account.
func (inst *CancelUpTo) GetQuoteAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBaseVaultAccount sets the "baseVault" account.
func (inst *CancelUpTo) SetBaseVaultAccount(baseVault ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(baseVault).WRITE()
	return inst
}

// GetBaseVaultAccount gets the "baseVault" account.
func (inst *CancelUpTo) GetBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetQuoteVaultAccount sets the "quoteVault" account.
func (inst *CancelUpTo) SetQuoteVaultAccount(quoteVault ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(quoteVault).WRITE()
	return inst
}

// GetQuoteVaultAccount gets the "quoteVault" account.
func (inst *CancelUpTo) GetQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CancelUpTo) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CancelUpTo {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CancelUpTo) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst CancelUpTo) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelUpTo,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelUpTo) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelUpTo) Validate() error {
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

func (inst *CancelUpTo) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelUpTo")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
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

func (obj CancelUpTo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelUpTo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelUpToInstruction declares a new CancelUpTo instruction with the provided parameters and accounts.
func NewCancelUpToInstruction(
	// Parameters:
	params CancelUpToParams,
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	trader ag_solanago.PublicKey,
	baseAccount ag_solanago.PublicKey,
	quoteAccount ag_solanago.PublicKey,
	baseVault ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CancelUpTo {
	return NewCancelUpToInstructionBuilder().
		SetParams(params).
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
