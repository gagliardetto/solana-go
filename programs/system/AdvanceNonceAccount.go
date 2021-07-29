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
					solana.NewAccountMeta(nonceAccount, true, false),
					solana.NewAccountMeta(SysVarRecentBlockHashesPubkey, false, false),
					solana.NewAccountMeta(nonceAuthority, true, false),
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
