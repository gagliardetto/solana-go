package system

import (
	"fmt"
	"io"

	"github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
)

type SystemInstruction interface{}

func (si SystemInstruction) Pack(buf []byte, opt *struc.Options) (int, error) {
	// check the Instruction type
	var pack interface{}
	switch el := si.(type) {
	case *CreateAccount:
		buf[0] = 0x00
		pack = el
	case *Assign:
		buf[0] = 0x01
		pack = el
	case *Transfer:
		buf[0] = 0x02
		pack = el
	default:
		return 0, fmt.Errorf("unsupported SystemInstruction: %T", si)
	}
	// TODO: this won't work, right? we need to write in the right place, and stretch the `buf`
	// directly?
	written, err := struc.Pack(buf[1:], el, opt)
	written++
	if err != nil {
		return written, err
	}
	return written, nil
}
func (si SystemInstruction) Unpack(r io.Reader, length int, opt *struc.Options) error {
	// read first byte, then decode the rest based on the proper type
	return fmt.Errorf("not implemented")
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
