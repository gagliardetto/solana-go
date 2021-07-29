package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAuthorizeNonceAccountInstruction(
	authorizePubKey solana.PublicKey,

	noncePubKey solana.PublicKey,
	nonceAuthority solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_AuthorizeNonceAccount,

			Impl: &AuthorizeNonceAccount{
				PubKey: authorizePubKey,
				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(noncePubKey).WRITE(),
					solana.Meta(nonceAuthority).SIGNER(),
				},
			},
		},
	}
}

type AuthorizeNonceAccount struct {
	// The Pubkey parameter identifies the entity to authorize.
	PubKey solana.PublicKey

	// [0] = [WRITE] Nonce account.
	// [1] = [SIGNER] Nonce authority.
	solana.AccountMetaSlice `bin:"-"`
}
