package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAdvanceNonceAccountInstruction(
	nonceAccount solana.PublicKey,
	nonceAuthority solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_AdvanceNonceAccount,

			Impl: &AdvanceNonceAccount{
				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(nonceAccount).WRITE(),
					solana.Meta(SysVarRecentBlockHashesPubkey),
					solana.Meta(nonceAuthority).SIGNER(),
				},
			},
		},
	}
}

type AdvanceNonceAccount struct {
	// [0] = [WRITE] Nonce account.
	// [1] = [] RecentBlockhashes sysvar.
	// [2] = [SIGNER] Nonce authority.
	solana.AccountMetaSlice `bin:"-"`
}
