package system

import (
	"bytes"
	"fmt"
	"io"

	solana "github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
)

type SystemInstruction struct {
	Type    solana.Varuint16
	Variant interface{}
}

func (si *SystemInstruction) Unpack(r io.Reader, length int, opt *struc.Options) (err error) {
	fmt.Println("CALLING OUR SI UNPACK", length)

	if err = struc.Unpack(r, &si.Type); err != nil {
		return
	}

	var el interface{}
	switch si.Type {
	case 0:
		el = &CreateAccount{}
	case 1:
		el = &Assign{}
	case 2:
		el = &Transfer{}
	default:
		return fmt.Errorf("unsupported System Instruction variant %d", si.Type)
	}
	si.Variant = el

	return struc.Unpack(r, el)
}

func (si SystemInstruction) Pack(p []byte, opt *struc.Options) (written int, err error) {
	fmt.Println("CALLING OUR SI PACK", len(p), cap(p))

	buf := &bytes.Buffer{}
	w := &solana.ByteCountWriter{Writer: buf}

	switch si.Variant.(type) {
	case *CreateAccount:
		si.Type = 0
	case *Assign:
		si.Type = 1
	case *Transfer:
		si.Type = 2
	default:
		return 0, fmt.Errorf("unsupported variant %T", si.Variant)
	}

	err = struc.Pack(w, si.Type)
	if err != nil {
		return 0, fmt.Errorf("pack type: %w", err)
	}

	err = struc.Pack(w, si.Variant)
	if err != nil {
		return 0, fmt.Errorf("pack impl: %w", err)
	}

	fmt.Println("byte count", w.ByteCount)

	copy(p, buf.Bytes())

	return w.ByteCount, nil
}

func (si SystemInstruction) Size(opt *struc.Options) int {
	s1, err := struc.Sizeof(si.Type)
	if err != nil {
		panic(err)
	}
	s2, err := struc.Sizeof(si.Variant)
	if err != nil {
		panic(err)
	}
	return s1 + s2
}

func (si SystemInstruction) String() string { return fmt.Sprintf("variant %d, %T", si.Type, si.Variant) }

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
