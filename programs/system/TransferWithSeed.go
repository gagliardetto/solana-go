package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewTransferWithSeedInstruction(
	from solana.PublicKey,
	to solana.PublicKey,
	basePubKey solana.PublicKey,
	owner solana.PublicKey,
	seed string,
	lamports uint64,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_TransferWithSeed,

			Impl: &TransferWithSeed{
				Lamports: bin.Uint64(lamports),
				Seed:     seed,
				Owner:    owner,

				AccountMetaSlice: []*solana.AccountMeta{
					solana.NewAccountMeta(from, true, false),
					solana.NewAccountMeta(basePubKey, false, true),
					solana.NewAccountMeta(to, true, false),
				},
			},
		},
	}
}

type TransferWithSeed struct {
	Lamports bin.Uint64
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey

	// [0] = [WRITE] Funding account.
	// [1] = [SIGNER] Base for funding account.
	// [2] = [WRITE] Recipient account.
	solana.AccountMetaSlice `bin:"-"`
}
