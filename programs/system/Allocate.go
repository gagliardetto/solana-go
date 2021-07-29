package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAllocateInstruction(
	space uint64,
	accountPubKey solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_Allocate,

			Impl: &Allocate{
				Space: bin.Uint64(space),
				AccountMetaSlice: []*solana.AccountMeta{
					solana.NewAccountMeta(accountPubKey, true, true),
				},
			},
		},
	}
}

type Allocate struct {
	// Number of bytes of memory to allocate.
	Space bin.Uint64

	// [0] = [WRITE, SIGNER] New account.
	solana.AccountMetaSlice `bin:"-"`
}
