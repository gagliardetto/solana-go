package system

import (
	"fmt"

	bin "github.com/dfuse-io/binary"
	solana "github.com/gagliardetto/solana-go"
)

func NewAdvanceNonceAccountInstruction(
	nonceAccount solana.PublicKey,
	nonceAuthority solana.PublicKey,
) *Instruction {
	return NewAdvanceNonceAccountBuilder().
		WithNonceAccount(nonceAccount).
		WithNonceAuthority(nonceAuthority).
		Build()
}

// Consumes a stored nonce, replacing it with a successor.
type AdvanceNonceAccount struct {
	// [0] = [WRITE] Nonce account.
	// [1] = [] RecentBlockhashes sysvar.
	// [2] = [SIGNER] Nonce authority.
	solana.AccountMetaSlice `bin:"-"`
}

// NewAdvanceNonceAccountBuilder initializes a new AdvanceNonceAccount builder.
func NewAdvanceNonceAccountBuilder() *AdvanceNonceAccount {
	nb := &AdvanceNonceAccount{
		AccountMetaSlice: make(solana.AccountMetaSlice, 3),
	}
	nb.AccountMetaSlice[1] = solana.Meta(SysVarRecentBlockHashesPubkey)
	return nb
}

func (ins *AdvanceNonceAccount) WithNonceAccount(nonceAccount solana.PublicKey) *AdvanceNonceAccount {
	ins.AccountMetaSlice[0] = solana.Meta(nonceAccount).WRITE()
	return ins
}

func (ins *AdvanceNonceAccount) GetNonceAccount() *solana.PublicKey {
	ac := ins.AccountMetaSlice[0]
	if ac == nil {
		return nil
	}
	return &ac.PublicKey
}

func (ins *AdvanceNonceAccount) WithNonceAuthority(nonceAuthority solana.PublicKey) *AdvanceNonceAccount {
	ins.AccountMetaSlice[2] = solana.Meta(nonceAuthority).SIGNER()
	return ins
}

func (ins *AdvanceNonceAccount) GetNonceAuthority() *solana.PublicKey {
	ac := ins.AccountMetaSlice[2]
	if ac == nil {
		return nil
	}
	return &ac.PublicKey
}

func (ins *AdvanceNonceAccount) Validate() error {
	for accIndex, acc := range ins.AccountMetaSlice {
		if acc == nil {
			return fmt.Errorf("ins.AccountMetaSlice[%v] is nil", accIndex)
		}
	}
	return nil
}

func (ins *AdvanceNonceAccount) Build() *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{

			TypeID: Instruction_AdvanceNonceAccount,

			Impl: ins,
		},
	}
}
