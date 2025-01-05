// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_clmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Creates a pool for the given token pair and the initial price
//
// # Arguments
//
// * `ctx`- The context of accounts
// * `sqrt_price_x64` - the initial sqrt price (amount_token_1 / amount_token_0) of the pool as a Q64.64
//
type CreatePool struct {
	SqrtPriceX64 *ag_binary.Uint128
	OpenTime     *uint64

	// [0] = [WRITE, SIGNER] poolCreator
	// ··········· Address paying to create the pool. Can be anyone
	//
	// [1] = [] ammConfig
	// ··········· Which config the pool belongs to.
	//
	// [2] = [WRITE] poolState
	// ··········· Initialize an account to store the pool state
	//
	// [3] = [] tokenMint0
	// ··········· Token_0 mint, the key must be smaller then token_1 mint.
	//
	// [4] = [] tokenMint1
	// ··········· Token_1 mint
	//
	// [5] = [WRITE] tokenVault0
	// ··········· Token_0 vault for the pool
	//
	// [6] = [WRITE] tokenVault1
	// ··········· Token_1 vault for the pool
	//
	// [7] = [WRITE] observationState
	// ··········· Initialize an account to store oracle observations
	//
	// [8] = [WRITE] tickArrayBitmap
	// ··········· Initialize an account to store if a tick array is initialized.
	//
	// [9] = [] tokenProgram0
	// ··········· Spl token program or token program 2022
	//
	// [10] = [] tokenProgram1
	// ··········· Spl token program or token program 2022
	//
	// [11] = [] systemProgram
	// ··········· To create a new program account
	//
	// [12] = [] rent
	// ··········· Sysvar for program account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreatePoolInstructionBuilder creates a new `CreatePool` instruction builder.
func NewCreatePoolInstructionBuilder() *CreatePool {
	nd := &CreatePool{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetSqrtPriceX64 sets the "sqrtPriceX64" parameter.
func (inst *CreatePool) SetSqrtPriceX64(sqrtPriceX64 ag_binary.Uint128) *CreatePool {
	inst.SqrtPriceX64 = &sqrtPriceX64
	return inst
}

// SetOpenTime sets the "openTime" parameter.
func (inst *CreatePool) SetOpenTime(openTime uint64) *CreatePool {
	inst.OpenTime = &openTime
	return inst
}

// SetPoolCreatorAccount sets the "poolCreator" account.
// Address paying to create the pool. Can be anyone
func (inst *CreatePool) SetPoolCreatorAccount(poolCreator ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(poolCreator).WRITE().SIGNER()
	return inst
}

// GetPoolCreatorAccount gets the "poolCreator" account.
// Address paying to create the pool. Can be anyone
func (inst *CreatePool) GetPoolCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmConfigAccount sets the "ammConfig" account.
// Which config the pool belongs to.
func (inst *CreatePool) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ammConfig)
	return inst
}

// GetAmmConfigAccount gets the "ammConfig" account.
// Which config the pool belongs to.
func (inst *CreatePool) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolStateAccount sets the "poolState" account.
// Initialize an account to store the pool state
func (inst *CreatePool) SetPoolStateAccount(poolState ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
// Initialize an account to store the pool state
func (inst *CreatePool) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMint0Account sets the "tokenMint0" account.
// Token_0 mint, the key must be smaller then token_1 mint.
func (inst *CreatePool) SetTokenMint0Account(tokenMint0 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint0)
	return inst
}

// GetTokenMint0Account gets the "tokenMint0" account.
// Token_0 mint, the key must be smaller then token_1 mint.
func (inst *CreatePool) GetTokenMint0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenMint1Account sets the "tokenMint1" account.
// Token_1 mint
func (inst *CreatePool) SetTokenMint1Account(tokenMint1 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenMint1)
	return inst
}

// GetTokenMint1Account gets the "tokenMint1" account.
// Token_1 mint
func (inst *CreatePool) GetTokenMint1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenVault0Account sets the "tokenVault0" account.
// Token_0 vault for the pool
func (inst *CreatePool) SetTokenVault0Account(tokenVault0 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenVault0).WRITE()
	return inst
}

// GetTokenVault0Account gets the "tokenVault0" account.
// Token_0 vault for the pool
func (inst *CreatePool) GetTokenVault0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenVault1Account sets the "tokenVault1" account.
// Token_1 vault for the pool
func (inst *CreatePool) SetTokenVault1Account(tokenVault1 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenVault1).WRITE()
	return inst
}

// GetTokenVault1Account gets the "tokenVault1" account.
// Token_1 vault for the pool
func (inst *CreatePool) GetTokenVault1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetObservationStateAccount sets the "observationState" account.
// Initialize an account to store oracle observations
func (inst *CreatePool) SetObservationStateAccount(observationState ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(observationState).WRITE()
	return inst
}

// GetObservationStateAccount gets the "observationState" account.
// Initialize an account to store oracle observations
func (inst *CreatePool) GetObservationStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTickArrayBitmapAccount sets the "tickArrayBitmap" account.
// Initialize an account to store if a tick array is initialized.
func (inst *CreatePool) SetTickArrayBitmapAccount(tickArrayBitmap ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tickArrayBitmap).WRITE()
	return inst
}

// GetTickArrayBitmapAccount gets the "tickArrayBitmap" account.
// Initialize an account to store if a tick array is initialized.
func (inst *CreatePool) GetTickArrayBitmapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgram0Account sets the "tokenProgram0" account.
// Spl token program or token program 2022
func (inst *CreatePool) SetTokenProgram0Account(tokenProgram0 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram0)
	return inst
}

// GetTokenProgram0Account gets the "tokenProgram0" account.
// Spl token program or token program 2022
func (inst *CreatePool) GetTokenProgram0Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgram1Account sets the "tokenProgram1" account.
// Spl token program or token program 2022
func (inst *CreatePool) SetTokenProgram1Account(tokenProgram1 ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram1)
	return inst
}

// GetTokenProgram1Account gets the "tokenProgram1" account.
// Spl token program or token program 2022
func (inst *CreatePool) GetTokenProgram1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// To create a new program account
func (inst *CreatePool) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// To create a new program account
func (inst *CreatePool) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
// Sysvar for program account
func (inst *CreatePool) SetRentAccount(rent ag_solanago.PublicKey) *CreatePool {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Sysvar for program account
func (inst *CreatePool) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst CreatePool) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreatePool,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreatePool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreatePool) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SqrtPriceX64 == nil {
			return errors.New("SqrtPriceX64 parameter is not set")
		}
		if inst.OpenTime == nil {
			return errors.New("OpenTime parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PoolCreator is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMint0 is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenMint1 is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenVault0 is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenVault1 is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ObservationState is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TickArrayBitmap is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram0 is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram1 is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreatePool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreatePool")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("SqrtPriceX64", *inst.SqrtPriceX64))
						paramsBranch.Child(ag_format.Param("    OpenTime", *inst.OpenTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     poolCreator", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       ammConfig", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       poolState", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      tokenMint0", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      tokenMint1", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     tokenVault0", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     tokenVault1", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("observationState", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" tickArrayBitmap", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram0", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram1", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("            rent", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj CreatePool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SqrtPriceX64` param:
	err = encoder.Encode(obj.SqrtPriceX64)
	if err != nil {
		return err
	}
	// Serialize `OpenTime` param:
	err = encoder.Encode(obj.OpenTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreatePool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SqrtPriceX64`:
	err = decoder.Decode(&obj.SqrtPriceX64)
	if err != nil {
		return err
	}
	// Deserialize `OpenTime`:
	err = decoder.Decode(&obj.OpenTime)
	if err != nil {
		return err
	}
	return nil
}

// NewCreatePoolInstruction declares a new CreatePool instruction with the provided parameters and accounts.
func NewCreatePoolInstruction(
	// Parameters:
	sqrtPriceX64 ag_binary.Uint128,
	openTime uint64,
	// Accounts:
	poolCreator ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	tokenMint0 ag_solanago.PublicKey,
	tokenMint1 ag_solanago.PublicKey,
	tokenVault0 ag_solanago.PublicKey,
	tokenVault1 ag_solanago.PublicKey,
	observationState ag_solanago.PublicKey,
	tickArrayBitmap ag_solanago.PublicKey,
	tokenProgram0 ag_solanago.PublicKey,
	tokenProgram1 ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreatePool {
	return NewCreatePoolInstructionBuilder().
		SetSqrtPriceX64(sqrtPriceX64).
		SetOpenTime(openTime).
		SetPoolCreatorAccount(poolCreator).
		SetAmmConfigAccount(ammConfig).
		SetPoolStateAccount(poolState).
		SetTokenMint0Account(tokenMint0).
		SetTokenMint1Account(tokenMint1).
		SetTokenVault0Account(tokenVault0).
		SetTokenVault1Account(tokenVault1).
		SetObservationStateAccount(observationState).
		SetTickArrayBitmapAccount(tickArrayBitmap).
		SetTokenProgram0Account(tokenProgram0).
		SetTokenProgram1Account(tokenProgram1).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
