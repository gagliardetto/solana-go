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
					solana.Meta(from).WRITE().SIGNER(),
					solana.Meta(to).WRITE(),
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
