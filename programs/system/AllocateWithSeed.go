package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAllocateWithSeedInstruction(
	accountPubKey solana.PublicKey,
	basePubKey solana.PublicKey,
	seed string,
	space uint64,
	owner solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_AllocateWithSeed,

			Impl: &AllocateWithSeed{
				Base:  basePubKey,
				Seed:  seed,
				Space: bin.Uint64(space),
				Owner: owner,

				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(accountPubKey).WRITE(),
					solana.Meta(basePubKey).SIGNER(),
				},
			},
		},
	}
}

type AllocateWithSeed struct {
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Space    bin.Uint64
	Owner    solana.PublicKey

	// [0] = [WRITE] Allocated account.
	// [1] = [SIGNER] Base account.
	solana.AccountMetaSlice `bin:"-"`
}
