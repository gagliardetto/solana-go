package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewCreateAccountWithSeedInstruction(
	base solana.PublicKey,
	seed string,
	lamports uint64,
	space uint64,
	owner solana.PublicKey,

	fundingAccount solana.PublicKey,
	newAccount solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_CreateAccountWithSeed,

			Impl: &CreateAccountWithSeed{

				Base:     base,
				Seed:     seed,
				Lamports: bin.Uint64(lamports),
				Space:    bin.Uint64(space),
				Owner:    owner,

				AccountMetaSlice: func() solana.AccountMetaSlice {
					res := solana.AccountMetaSlice{
						solana.Meta(fundingAccount).WRITE().SIGNER(),
						solana.Meta(newAccount).WRITE(),
					}

					if !base.Equals(fundingAccount) {
						res.Append(solana.Meta(base).SIGNER())
					}

					return res
				}(),
			},
		},
	}
}

type CreateAccountWithSeed struct {
	// Base public key.
	Base solana.PublicKey

	SeedSize int64 `bin:"sizeof=Seed"`
	// String of ASCII chars, no longer than solana.MAX_SEED_LEN
	Seed string

	// Number of lamports to transfer to the new account.
	Lamports bin.Uint64

	// Number of bytes of memory to allocate
	Space bin.Uint64

	// Owner program account address.
	Owner solana.PublicKey

	// [0] = [WRITE, SIGNER] Funding account.
	// [1] = [WRITE] Created account.
	// [2] = [SIGNER] Base account; the account matching the base Pubkey below must be provided as a signer,
	// 		 but may be the same as the funding account and provided as account 0
	solana.AccountMetaSlice `bin:"-"`
}
