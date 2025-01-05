// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// remove a strategy by advance payment
type RemoveStrategy2 struct {
	MaxAdminPayAmount *uint64

	// [0] = [WRITE] vault
	// ··········· Vault account
	//
	// [1] = [WRITE] strategy
	// ··········· Strategy account
	//
	// [2] = [] strategyProgram
	//
	// [3] = [WRITE] collateralVault
	// ··········· Collateral vault account
	//
	// [4] = [WRITE] reserve
	//
	// [5] = [WRITE] tokenVault
	// ··········· token_vault
	//
	// [6] = [WRITE] tokenAdminAdvancePayment
	// ··········· token_advance_payemnt
	// ··········· the owner of token_advance_payment must be admin
	//
	// [7] = [WRITE] tokenVaultAdvancePayment
	// ··········· token_vault_advance_payment
	// ··········· the account must be different from token_vault
	// ··········· the owner of token_advance_payment must be vault
	//
	// [8] = [WRITE] feeVault
	// ··········· fee_vault
	//
	// [9] = [WRITE] lpMint
	// ··········· lp_mint
	//
	// [10] = [] tokenProgram
	// ··········· token_program
	//
	// [11] = [SIGNER] admin
	// ··········· admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRemoveStrategy2InstructionBuilder creates a new `RemoveStrategy2` instruction builder.
func NewRemoveStrategy2InstructionBuilder() *RemoveStrategy2 {
	nd := &RemoveStrategy2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetMaxAdminPayAmount sets the "maxAdminPayAmount" parameter.
func (inst *RemoveStrategy2) SetMaxAdminPayAmount(maxAdminPayAmount uint64) *RemoveStrategy2 {
	inst.MaxAdminPayAmount = &maxAdminPayAmount
	return inst
}

// SetVaultAccount sets the "vault" account.
// Vault account
func (inst *RemoveStrategy2) SetVaultAccount(vault ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
// Vault account
func (inst *RemoveStrategy2) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStrategyAccount sets the "strategy" account.
// Strategy account
func (inst *RemoveStrategy2) SetStrategyAccount(strategy ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(strategy).WRITE()
	return inst
}

// GetStrategyAccount gets the "strategy" account.
// Strategy account
func (inst *RemoveStrategy2) GetStrategyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStrategyProgramAccount sets the "strategyProgram" account.
func (inst *RemoveStrategy2) SetStrategyProgramAccount(strategyProgram ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(strategyProgram)
	return inst
}

// GetStrategyProgramAccount gets the "strategyProgram" account.
func (inst *RemoveStrategy2) GetStrategyProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollateralVaultAccount sets the "collateralVault" account.
// Collateral vault account
func (inst *RemoveStrategy2) SetCollateralVaultAccount(collateralVault ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collateralVault).WRITE()
	return inst
}

// GetCollateralVaultAccount gets the "collateralVault" account.
// Collateral vault account
func (inst *RemoveStrategy2) GetCollateralVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReserveAccount sets the "reserve" account.
func (inst *RemoveStrategy2) SetReserveAccount(reserve ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(reserve).WRITE()
	return inst
}

// GetReserveAccount gets the "reserve" account.
func (inst *RemoveStrategy2) GetReserveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenVaultAccount sets the "tokenVault" account.
// token_vault
func (inst *RemoveStrategy2) SetTokenVaultAccount(tokenVault ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenVault).WRITE()
	return inst
}

// GetTokenVaultAccount gets the "tokenVault" account.
// token_vault
func (inst *RemoveStrategy2) GetTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenAdminAdvancePaymentAccount sets the "tokenAdminAdvancePayment" account.
// token_advance_payemnt
// the owner of token_advance_payment must be admin
func (inst *RemoveStrategy2) SetTokenAdminAdvancePaymentAccount(tokenAdminAdvancePayment ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenAdminAdvancePayment).WRITE()
	return inst
}

// GetTokenAdminAdvancePaymentAccount gets the "tokenAdminAdvancePayment" account.
// token_advance_payemnt
// the owner of token_advance_payment must be admin
func (inst *RemoveStrategy2) GetTokenAdminAdvancePaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenVaultAdvancePaymentAccount sets the "tokenVaultAdvancePayment" account.
// token_vault_advance_payment
// the account must be different from token_vault
// the owner of token_advance_payment must be vault
func (inst *RemoveStrategy2) SetTokenVaultAdvancePaymentAccount(tokenVaultAdvancePayment ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenVaultAdvancePayment).WRITE()
	return inst
}

// GetTokenVaultAdvancePaymentAccount gets the "tokenVaultAdvancePayment" account.
// token_vault_advance_payment
// the account must be different from token_vault
// the owner of token_advance_payment must be vault
func (inst *RemoveStrategy2) GetTokenVaultAdvancePaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFeeVaultAccount sets the "feeVault" account.
// fee_vault
func (inst *RemoveStrategy2) SetFeeVaultAccount(feeVault ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(feeVault).WRITE()
	return inst
}

// GetFeeVaultAccount gets the "feeVault" account.
// fee_vault
func (inst *RemoveStrategy2) GetFeeVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetLpMintAccount sets the "lpMint" account.
// lp_mint
func (inst *RemoveStrategy2) SetLpMintAccount(lpMint ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(lpMint).WRITE()
	return inst
}

// GetLpMintAccount gets the "lpMint" account.
// lp_mint
func (inst *RemoveStrategy2) GetLpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// token_program
func (inst *RemoveStrategy2) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// token_program
func (inst *RemoveStrategy2) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAdminAccount sets the "admin" account.
// admin
func (inst *RemoveStrategy2) SetAdminAccount(admin ag_solanago.PublicKey) *RemoveStrategy2 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
// admin
func (inst *RemoveStrategy2) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst RemoveStrategy2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveStrategy2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveStrategy2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveStrategy2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MaxAdminPayAmount == nil {
			return errors.New("MaxAdminPayAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Strategy is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StrategyProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CollateralVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Reserve is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenVault is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenAdminAdvancePayment is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenVaultAdvancePayment is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FeeVault is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.LpMint is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *RemoveStrategy2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveStrategy2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MaxAdminPayAmount", *inst.MaxAdminPayAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   vault", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                strategy", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         strategyProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         collateralVault", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 reserve", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              tokenVault", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("tokenAdminAdvancePayment", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("tokenVaultAdvancePayment", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                feeVault", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                  lpMint", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj RemoveStrategy2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MaxAdminPayAmount` param:
	err = encoder.Encode(obj.MaxAdminPayAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RemoveStrategy2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MaxAdminPayAmount`:
	err = decoder.Decode(&obj.MaxAdminPayAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewRemoveStrategy2Instruction declares a new RemoveStrategy2 instruction with the provided parameters and accounts.
func NewRemoveStrategy2Instruction(
	// Parameters:
	maxAdminPayAmount uint64,
	// Accounts:
	vault ag_solanago.PublicKey,
	strategy ag_solanago.PublicKey,
	strategyProgram ag_solanago.PublicKey,
	collateralVault ag_solanago.PublicKey,
	reserve ag_solanago.PublicKey,
	tokenVault ag_solanago.PublicKey,
	tokenAdminAdvancePayment ag_solanago.PublicKey,
	tokenVaultAdvancePayment ag_solanago.PublicKey,
	feeVault ag_solanago.PublicKey,
	lpMint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *RemoveStrategy2 {
	return NewRemoveStrategy2InstructionBuilder().
		SetMaxAdminPayAmount(maxAdminPayAmount).
		SetVaultAccount(vault).
		SetStrategyAccount(strategy).
		SetStrategyProgramAccount(strategyProgram).
		SetCollateralVaultAccount(collateralVault).
		SetReserveAccount(reserve).
		SetTokenVaultAccount(tokenVault).
		SetTokenAdminAdvancePaymentAccount(tokenAdminAdvancePayment).
		SetTokenVaultAdvancePaymentAccount(tokenVaultAdvancePayment).
		SetFeeVaultAccount(feeVault).
		SetLpMintAccount(lpMint).
		SetTokenProgramAccount(tokenProgram).
		SetAdminAccount(admin)
}
