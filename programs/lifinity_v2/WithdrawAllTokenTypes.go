// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lifinity_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawAllTokenTypes is the `withdrawAllTokenTypes` instruction.
type WithdrawAllTokenTypes struct {
	PoolTokenAmount     *uint64
	MinimumTokenAAmount *uint64
	MinimumTokenBAmount *uint64

	// [0] = [WRITE] amm
	//
	// [1] = [] authority
	//
	// [2] = [SIGNER] userTransferAuthorityInfo
	//
	// [3] = [WRITE] sourceInfo
	//
	// [4] = [WRITE] tokenA
	//
	// [5] = [WRITE] tokenB
	//
	// [6] = [WRITE] poolMint
	//
	// [7] = [WRITE] destTokenAInfo
	//
	// [8] = [WRITE] destTokenBInfo
	//
	// [9] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawAllTokenTypesInstructionBuilder creates a new `WithdrawAllTokenTypes` instruction builder.
func NewWithdrawAllTokenTypesInstructionBuilder() *WithdrawAllTokenTypes {
	nd := &WithdrawAllTokenTypes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetPoolTokenAmount sets the "poolTokenAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetPoolTokenAmount(poolTokenAmount uint64) *WithdrawAllTokenTypes {
	inst.PoolTokenAmount = &poolTokenAmount
	return inst
}

// SetMinimumTokenAAmount sets the "minimumTokenAAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetMinimumTokenAAmount(minimumTokenAAmount uint64) *WithdrawAllTokenTypes {
	inst.MinimumTokenAAmount = &minimumTokenAAmount
	return inst
}

// SetMinimumTokenBAmount sets the "minimumTokenBAmount" parameter.
func (inst *WithdrawAllTokenTypes) SetMinimumTokenBAmount(minimumTokenBAmount uint64) *WithdrawAllTokenTypes {
	inst.MinimumTokenBAmount = &minimumTokenBAmount
	return inst
}

// SetAmmAccount sets the "amm" account.
func (inst *WithdrawAllTokenTypes) SetAmmAccount(amm ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *WithdrawAllTokenTypes) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawAllTokenTypes) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawAllTokenTypes) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserTransferAuthorityInfoAccount sets the "userTransferAuthorityInfo" account.
func (inst *WithdrawAllTokenTypes) SetUserTransferAuthorityInfoAccount(userTransferAuthorityInfo ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userTransferAuthorityInfo).SIGNER()
	return inst
}

// GetUserTransferAuthorityInfoAccount gets the "userTransferAuthorityInfo" account.
func (inst *WithdrawAllTokenTypes) GetUserTransferAuthorityInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSourceInfoAccount sets the "sourceInfo" account.
func (inst *WithdrawAllTokenTypes) SetSourceInfoAccount(sourceInfo ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(sourceInfo).WRITE()
	return inst
}

// GetSourceInfoAccount gets the "sourceInfo" account.
func (inst *WithdrawAllTokenTypes) GetSourceInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenAAccount sets the "tokenA" account.
func (inst *WithdrawAllTokenTypes) SetTokenAAccount(tokenA ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenA).WRITE()
	return inst
}

// GetTokenAAccount gets the "tokenA" account.
func (inst *WithdrawAllTokenTypes) GetTokenAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenBAccount sets the "tokenB" account.
func (inst *WithdrawAllTokenTypes) SetTokenBAccount(tokenB ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenB).WRITE()
	return inst
}

// GetTokenBAccount gets the "tokenB" account.
func (inst *WithdrawAllTokenTypes) GetTokenBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolMintAccount sets the "poolMint" account.
func (inst *WithdrawAllTokenTypes) SetPoolMintAccount(poolMint ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(poolMint).WRITE()
	return inst
}

// GetPoolMintAccount gets the "poolMint" account.
func (inst *WithdrawAllTokenTypes) GetPoolMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetDestTokenAInfoAccount sets the "destTokenAInfo" account.
func (inst *WithdrawAllTokenTypes) SetDestTokenAInfoAccount(destTokenAInfo ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(destTokenAInfo).WRITE()
	return inst
}

// GetDestTokenAInfoAccount gets the "destTokenAInfo" account.
func (inst *WithdrawAllTokenTypes) GetDestTokenAInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetDestTokenBInfoAccount sets the "destTokenBInfo" account.
func (inst *WithdrawAllTokenTypes) SetDestTokenBInfoAccount(destTokenBInfo ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(destTokenBInfo).WRITE()
	return inst
}

// GetDestTokenBInfoAccount gets the "destTokenBInfo" account.
func (inst *WithdrawAllTokenTypes) GetDestTokenBInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *WithdrawAllTokenTypes) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *WithdrawAllTokenTypes) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst WithdrawAllTokenTypes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawAllTokenTypes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawAllTokenTypes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawAllTokenTypes) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PoolTokenAmount == nil {
			return errors.New("PoolTokenAmount parameter is not set")
		}
		if inst.MinimumTokenAAmount == nil {
			return errors.New("MinimumTokenAAmount parameter is not set")
		}
		if inst.MinimumTokenBAmount == nil {
			return errors.New("MinimumTokenBAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserTransferAuthorityInfo is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SourceInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenA is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenB is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PoolMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.DestTokenAInfo is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.DestTokenBInfo is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawAllTokenTypes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawAllTokenTypes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    PoolTokenAmount", *inst.PoolTokenAmount))
						paramsBranch.Child(ag_format.Param("MinimumTokenAAmount", *inst.MinimumTokenAAmount))
						paramsBranch.Child(ag_format.Param("MinimumTokenBAmount", *inst.MinimumTokenBAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                      amm", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("userTransferAuthorityInfo", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("               sourceInfo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                   tokenA", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   tokenB", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                 poolMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           destTokenAInfo", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           destTokenBInfo", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("             tokenProgram", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj WithdrawAllTokenTypes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PoolTokenAmount` param:
	err = encoder.Encode(obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `MinimumTokenAAmount` param:
	err = encoder.Encode(obj.MinimumTokenAAmount)
	if err != nil {
		return err
	}
	// Serialize `MinimumTokenBAmount` param:
	err = encoder.Encode(obj.MinimumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawAllTokenTypes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PoolTokenAmount`:
	err = decoder.Decode(&obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `MinimumTokenAAmount`:
	err = decoder.Decode(&obj.MinimumTokenAAmount)
	if err != nil {
		return err
	}
	// Deserialize `MinimumTokenBAmount`:
	err = decoder.Decode(&obj.MinimumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawAllTokenTypesInstruction declares a new WithdrawAllTokenTypes instruction with the provided parameters and accounts.
func NewWithdrawAllTokenTypesInstruction(
	// Parameters:
	poolTokenAmount uint64,
	minimumTokenAAmount uint64,
	minimumTokenBAmount uint64,
	// Accounts:
	amm ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	userTransferAuthorityInfo ag_solanago.PublicKey,
	sourceInfo ag_solanago.PublicKey,
	tokenA ag_solanago.PublicKey,
	tokenB ag_solanago.PublicKey,
	poolMint ag_solanago.PublicKey,
	destTokenAInfo ag_solanago.PublicKey,
	destTokenBInfo ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *WithdrawAllTokenTypes {
	return NewWithdrawAllTokenTypesInstructionBuilder().
		SetPoolTokenAmount(poolTokenAmount).
		SetMinimumTokenAAmount(minimumTokenAAmount).
		SetMinimumTokenBAmount(minimumTokenBAmount).
		SetAmmAccount(amm).
		SetAuthorityAccount(authority).
		SetUserTransferAuthorityInfoAccount(userTransferAuthorityInfo).
		SetSourceInfoAccount(sourceInfo).
		SetTokenAAccount(tokenA).
		SetTokenBAccount(tokenB).
		SetPoolMintAccount(poolMint).
		SetDestTokenAInfoAccount(destTokenAInfo).
		SetDestTokenBInfoAccount(destTokenBInfo).
		SetTokenProgramAccount(tokenProgram)
}
