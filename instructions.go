package solana

import (
	"bytes"

	"github.com/lunixbochs/struc"
)

type AccountMeta struct {
	PublicKey  PublicKey
	IsSigner   bool
	IsWritable bool
}

type Instruction struct {
	ProgramID PublicKey
	Accounts  []AccountMeta
	Data      Base58
}

func NewInstruction(programID PublicKey, accountMetas []AccountMeta, instruction interface{}) (*Instruction, error) {
	buf := &bytes.Buffer{}
	err := struc.Pack(buf, instruction)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		ProgramID: programID,
		Accounts:  accountMetas,
		Data:      Base58(buf.Bytes()),
	}, nil
}
