package system

import (
	"io"

	solana "github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
)

var systemInstructionDef = solana.NewVariantDefinition([]solana.VariantType{
	{"CreateAccount", (*CreateAccount)(nil)},
	{"Assign", (*Assign)(nil)},
	{"Transfer", (*Transfer)(nil)},
})

type SystemInstruction struct {
	solana.BaseVariant
}

func NewInstruction(impl interface{}) *SystemInstruction {
	return &SystemInstruction{
		solana.BaseVariant{
			Type: systemInstructionDef.IDForType(impl),
			Impl: impl,
		},
	}
}

func (si *SystemInstruction) Unpack(r io.Reader, length int, opt *struc.Options) error {
	return si.BaseVariant.Unpack(systemInstructionDef, r, length, opt)
}

type CreateAccount struct {
	// prefixed with byte 0x00
	Lamports solana.U64
	Space    solana.U64
	Owner    solana.PublicKey
}

type Assign struct {
	// prefixed with byte 0x01
	Owner solana.PublicKey
}

type Transfer struct {
	// Prefixed with byte 0x02
	Lamports solana.U64
}

type CreateAccountWithSeed struct {
	// Prefixed with byte 0x03
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Lamports solana.U64
	Space    solana.U64
	Owner    solana.PublicKey
}

type AdvanceNonceAccount struct {
	// Prefix with 0x04
}

type WithdrawNonceAccount struct {
	// Prefix with 0x05
	Lamports solana.U64
}

type InitializeNonceAccount struct {
	// Prefix with 0x06
	AuthorizedAccount solana.PublicKey
}

type AuthorizeNonceAccount struct {
	// Prefix with 0x07
	AuthorizeAccount solana.PublicKey
}

type Allocate struct {
	// Prefix with 0x08
	Space solana.U64
}

type AllocateWithSeed struct {
	// Prefixed with byte 0x09
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Space    solana.U64
	Owner    solana.PublicKey
}

type AssignWithSeed struct {
	// Prefixed with byte 0x0a
	Base     solana.PublicKey
	SeedSize int `struc:"sizeof=Seed"`
	Seed     string
	Owner    solana.PublicKey
}
