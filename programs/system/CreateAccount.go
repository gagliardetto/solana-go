package system

import (
	"fmt"

	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewCreateAccountInstruction(
	lamports uint64,
	space uint64,
	owner solana.PublicKey,

	fundingAccount solana.PublicKey,
	newAccount solana.PublicKey,
) *Instruction {
	return NewCreateAccountBuilder().
		WithLamports(lamports).
		WithSpace(space).
		WithOwner(owner).
		WithFundingAccount(fundingAccount).
		WithNewAccount(newAccount).
		Build()
}

// Create a new account.
type CreateAccount struct {
	// Number of lamports to transfer to the new account.
	Lamports bin.Uint64
	// Number of bytes of memory to allocate.
	Space bin.Uint64
	// Address of program that will own the new account.
	Owner solana.PublicKey

	// [0] = [WRITE, SIGNER] Funding account.
	// [1] = [WRITE, SIGNER] New account.
	solana.AccountMetaSlice `bin:"-"`
}

// NewCreateAccountBuilder initializes a new CreateAccount builder.
func NewCreateAccountBuilder() *CreateAccount {
	return &CreateAccount{
		AccountMetaSlice: make(solana.AccountMetaSlice, 2),
	}
}

// WithLamports sets the number of lamports to transfer to the new account.
func (ins *CreateAccount) WithLamports(lamports uint64) *CreateAccount {
	ins.Lamports = bin.Uint64(lamports)
	return ins
}

// WithSpace sets the number of bytes of memory to allocate.
func (ins *CreateAccount) WithSpace(space uint64) *CreateAccount {
	ins.Space = bin.Uint64(space)
	return ins
}

// WithOwner sets the address of program that will own the new account.
func (ins *CreateAccount) WithOwner(owner solana.PublicKey) *CreateAccount {
	ins.Owner = owner
	return ins
}

// WithFundingAccount sets the account that will fund the new account.
func (ins *CreateAccount) WithFundingAccount(fundingAccount solana.PublicKey) *CreateAccount {
	ins.AccountMetaSlice[0] = solana.Meta(fundingAccount).WRITE().SIGNER()
	return ins
}

// GetFundingAccount gets the account that will fund the new account.
func (ins *CreateAccount) GetFundingAccount() *solana.PublicKey {
	ac := ins.AccountMetaSlice[0]
	if ac == nil {
		return nil
	}
	return &ac.PublicKey
}

// WithNewAccount sets the new account that will be created.
func (ins *CreateAccount) WithNewAccount(newAccount solana.PublicKey) *CreateAccount {
	ins.AccountMetaSlice[1] = solana.Meta(newAccount).WRITE().SIGNER()
	return ins
}

// GetNewAccount gets the new account.
func (ins *CreateAccount) GetNewAccount() *solana.PublicKey {
	ac := ins.AccountMetaSlice[1]
	if ac == nil {
		return nil
	}
	return &ac.PublicKey
}

func (ins *CreateAccount) Validate() error {
	for accIndex, acc := range ins.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is nil", accIndex)
		}
	}
	return nil
}

func (ins *CreateAccount) Build() *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_CreateAccount,

			Impl: ins,
		},
	}
}
