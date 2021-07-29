package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAssignWithSeedInstruction(
	accountPubKey solana.PublicKey,
	basePubKey solana.PublicKey,
	seed string,
	owner solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_AssignWithSeed,

			Impl: &AssignWithSeed{
				Base:  basePubKey,
				Seed:  seed,
				Owner: owner,

				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(accountPubKey).WRITE(),
					solana.Meta(basePubKey).SIGNER(),
				},
			},
		},
	}
}

type AssignWithSeed struct {
	Base     solana.PublicKey
	SeedSize int `bin:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey

	// [0] = [WRITE] Assigned account.
	// [1] = [SIGNER] Base account.
	solana.AccountMetaSlice `bin:"-"`
}
