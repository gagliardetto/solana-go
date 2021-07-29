package system

import (
	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAssignInstruction(
	// The account that is being assigned.
	accountPubkey solana.PublicKey,
	// The program that is becoming the new owner of the account.
	assignToProgramID solana.PublicKey,
) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_Assign,

			Impl: &Assign{

				NewOwner: assignToProgramID,

				AccountMetaSlice: []*solana.AccountMeta{
					solana.Meta(accountPubkey).WRITE().SIGNER(),
				},
			},
		},
	}
}

type Assign struct {
	// Owner program account.
	NewOwner solana.PublicKey

	// [0] = [WRITE, SIGNER] Assigned account public key.
	solana.AccountMetaSlice `bin:"-"`
}
