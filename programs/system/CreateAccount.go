package system

import (
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
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_CreateAccount,

			Impl: &CreateAccount{

				Lamports: bin.Uint64(lamports),
				Space:    bin.Uint64(space),
				Owner:    owner,

				AccountMetaSlice: []*solana.AccountMeta{
					solana.NewAccountMeta(fundingAccount, true, true),
					solana.NewAccountMeta(newAccount, true, true),
				},
			},
		},
	}
}

type CreateAccount struct {
	Lamports bin.Uint64
	Space    bin.Uint64
	Owner    solana.PublicKey

	// [0] = [WRITE, SIGNER] Funding account.
	// [1] = [WRITE, SIGNER] New account.
	solana.AccountMetaSlice `bin:"-"`
}
