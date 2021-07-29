package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewTransferInstruction(
	lamports uint64,
	from solana.PublicKey,
	to solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_Transfer,

			Impl: &Transfer{
				Lamports: bin.Uint64(lamports),

				AccountMetaSlice: []*solana.AccountMeta{
					solana.NewAccountMeta(from, true, true),
					solana.NewAccountMeta(to, true, false),
				},
			},
		},
	}
}

type Transfer struct {
	Lamports bin.Uint64

	// [0] = [WRITE, SIGNER] Funding account.
	// [1] = [WRITE] Recipient account.
	solana.AccountMetaSlice `bin:"-"`
}
