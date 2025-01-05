// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RemoveAllLiquidity is the `removeAllLiquidity` instruction.
type RemoveAllLiquidity struct {

	// [0] = [WRITE] position
	//
	// [1] = [WRITE] lbPair
	//
	// [2] = [WRITE] binArrayBitmapExtension
	//
	// [3] = [WRITE] userTokenX
	//
	// [4] = [WRITE] userTokenY
	//
	// [5] = [WRITE] reserveX
	//
	// [6] = [WRITE] reserveY
	//
	// [7] = [] tokenXMint
	//
	// [8] = [] tokenYMint
	//
	// [9] = [WRITE] binArrayLower
	//
	// [10] = [WRITE] binArrayUpper
	//
	// [11] = [SIGNER] sender
	//
	// [12] = [] tokenXProgram
	//
	// [13] = [] tokenYProgram
	//
	// [14] = [] eventAuthority
	//
	// [15] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRemoveAllLiquidityInstructionBuilder creates a new `RemoveAllLiquidity` instruction builder.
func NewRemoveAllLiquidityInstructionBuilder() *RemoveAllLiquidity {
	nd := &RemoveAllLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetPositionAccount sets the "position" account.
func (inst *RemoveAllLiquidity) SetPositionAccount(position ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *RemoveAllLiquidity) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *RemoveAllLiquidity) SetLbPairAccount(lbPair ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *RemoveAllLiquidity) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBinArrayBitmapExtensionAccount sets the "binArrayBitmapExtension" account.
func (inst *RemoveAllLiquidity) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "binArrayBitmapExtension" account.
func (inst *RemoveAllLiquidity) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserTokenXAccount sets the "userTokenX" account.
func (inst *RemoveAllLiquidity) SetUserTokenXAccount(userTokenX ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userTokenX).WRITE()
	return inst
}

// GetUserTokenXAccount gets the "userTokenX" account.
func (inst *RemoveAllLiquidity) GetUserTokenXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserTokenYAccount sets the "userTokenY" account.
func (inst *RemoveAllLiquidity) SetUserTokenYAccount(userTokenY ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userTokenY).WRITE()
	return inst
}

// GetUserTokenYAccount gets the "userTokenY" account.
func (inst *RemoveAllLiquidity) GetUserTokenYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReserveXAccount sets the "reserveX" account.
func (inst *RemoveAllLiquidity) SetReserveXAccount(reserveX ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserveX" account.
func (inst *RemoveAllLiquidity) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReserveYAccount sets the "reserveY" account.
func (inst *RemoveAllLiquidity) SetReserveYAccount(reserveY ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserveY" account.
func (inst *RemoveAllLiquidity) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenXMintAccount sets the "tokenXMint" account.
func (inst *RemoveAllLiquidity) SetTokenXMintAccount(tokenXMint ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenXMint)
	return inst
}

// GetTokenXMintAccount gets the "tokenXMint" account.
func (inst *RemoveAllLiquidity) GetTokenXMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenYMintAccount sets the "tokenYMint" account.
func (inst *RemoveAllLiquidity) SetTokenYMintAccount(tokenYMint ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenYMint)
	return inst
}

// GetTokenYMintAccount gets the "tokenYMint" account.
func (inst *RemoveAllLiquidity) GetTokenYMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetBinArrayLowerAccount sets the "binArrayLower" account.
func (inst *RemoveAllLiquidity) SetBinArrayLowerAccount(binArrayLower ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(binArrayLower).WRITE()
	return inst
}

// GetBinArrayLowerAccount gets the "binArrayLower" account.
func (inst *RemoveAllLiquidity) GetBinArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetBinArrayUpperAccount sets the "binArrayUpper" account.
func (inst *RemoveAllLiquidity) SetBinArrayUpperAccount(binArrayUpper ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(binArrayUpper).WRITE()
	return inst
}

// GetBinArrayUpperAccount gets the "binArrayUpper" account.
func (inst *RemoveAllLiquidity) GetBinArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSenderAccount sets the "sender" account.
func (inst *RemoveAllLiquidity) SetSenderAccount(sender ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(sender).SIGNER()
	return inst
}

// GetSenderAccount gets the "sender" account.
func (inst *RemoveAllLiquidity) GetSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenXProgramAccount sets the "tokenXProgram" account.
func (inst *RemoveAllLiquidity) SetTokenXProgramAccount(tokenXProgram ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenXProgram)
	return inst
}

// GetTokenXProgramAccount gets the "tokenXProgram" account.
func (inst *RemoveAllLiquidity) GetTokenXProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenYProgramAccount sets the "tokenYProgram" account.
func (inst *RemoveAllLiquidity) SetTokenYProgramAccount(tokenYProgram ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenYProgram)
	return inst
}

// GetTokenYProgramAccount gets the "tokenYProgram" account.
func (inst *RemoveAllLiquidity) GetTokenYProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *RemoveAllLiquidity) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *RemoveAllLiquidity) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetProgramAccount sets the "program" account.
func (inst *RemoveAllLiquidity) SetProgramAccount(program ag_solanago.PublicKey) *RemoveAllLiquidity {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *RemoveAllLiquidity) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst RemoveAllLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveAllLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveAllLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveAllLiquidity) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BinArrayBitmapExtension is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserTokenX is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserTokenY is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenXMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenYMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.BinArrayLower is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.BinArrayUpper is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Sender is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenXProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenYProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *RemoveAllLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveAllLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 lbPair", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("binArrayBitmapExtension", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             userTokenX", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             userTokenY", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               reserveX", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("               reserveY", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("             tokenXMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("             tokenYMint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          binArrayLower", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          binArrayUpper", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                 sender", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("          tokenXProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("          tokenYProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         eventAuthority", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj RemoveAllLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RemoveAllLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRemoveAllLiquidityInstruction declares a new RemoveAllLiquidity instruction with the provided parameters and accounts.
func NewRemoveAllLiquidityInstruction(
	// Accounts:
	position ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	userTokenX ag_solanago.PublicKey,
	userTokenY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	tokenXMint ag_solanago.PublicKey,
	tokenYMint ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	tokenXProgram ag_solanago.PublicKey,
	tokenYProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *RemoveAllLiquidity {
	return NewRemoveAllLiquidityInstructionBuilder().
		SetPositionAccount(position).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetUserTokenXAccount(userTokenX).
		SetUserTokenYAccount(userTokenY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetTokenXMintAccount(tokenXMint).
		SetTokenYMintAccount(tokenYMint).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetSenderAccount(sender).
		SetTokenXProgramAccount(tokenXProgram).
		SetTokenYProgramAccount(tokenYProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
