package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewInitializeNonceAccountInstruction(
	authPubKey solana.PublicKey,
	noncePubKey solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_InitializeNonceAccount,

			Impl: &InitializeNonceAccount{
				AuthPubKey: authPubKey,
				AccountMetaSlice: []*solana.AccountMeta{
					solana.NewAccountMeta(noncePubKey, true, false),
					solana.NewAccountMeta(SysVarRecentBlockHashesPubkey, false, false),
					solana.NewAccountMeta(SysVarRentPubkey, false, false),
				},
			},
		},
	}
}

type InitializeNonceAccount struct {
	// The Pubkey parameter specifies the entity
	// authorized to execute nonce instruction on the account.
	// No signatures are required to execute this instruction,
	// enabling derived nonce account addresses
	AuthPubKey solana.PublicKey

	// [0] = [WRITE] Nonce account.
	// [1] = [] RecentBlockhashes sysvar.
	// [2] = [] Rent sysvar.
	solana.AccountMetaSlice `bin:"-"`
}
