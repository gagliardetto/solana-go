package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewWithdrawNonceAccountInstruction(
	lamports uint64,

	nonceAccount solana.PublicKey,
	recipientAccount solana.PublicKey,
	nonceAuthority solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_WithdrawNonceAccount,

			Impl: &WithdrawNonceAccount{
				Lamports: bin.Uint64(lamports),
				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(nonceAccount).WRITE(),
					solana.Meta(recipientAccount).WRITE(),
					solana.Meta(SysVarRecentBlockHashesPubkey),
					solana.Meta(SysVarRentPubkey),
					solana.Meta(nonceAuthority).SIGNER(),
				},
			},
		},
	}
}

type WithdrawNonceAccount struct {
	// The u64 parameter is the lamports to withdraw, which must leave the account balance above the rent exempt reserve or at zero.
	Lamports bin.Uint64

	// [0] = [WRITE] Nonce account.
	// [1] = [WRITE] Recipient account.
	// [2] = [] RecentBlockhashes sysvar.
	// [3] = [] Rent sysvar.
	// [4] = [SIGNER] Nonce authority.
	solana.AccountMetaSlice `bin:"-"`
}
