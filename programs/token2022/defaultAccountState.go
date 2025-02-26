package token2022

import (
	"bytes"
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
)

// Instruction to initialize the DefaultAccountState Extension
func CreateInitializeDefaultAccountStateInstruction(
	mint solana.PublicKey, // Mint Account address
	accountState token.AccountState, // Default AccountState
) solana.Instruction {
	programID := ProgramID

	pointerData := createInitializeDefaultAccountStateInstructionData{
		Instruction:                    DefaultAccountStateExtension,
		DefaultAccountStateInstruction: Initialize,
		AccountState:                   accountState,
	}

	ix := &createInitializeDefaultAccountStateInstruction{
		programID: programID,
		accounts: []*solana.AccountMeta{
			solana.Meta(mint).WRITE(),
		},
		data: pointerData.encode(),
	}

	return ix
}

type createInitializeDefaultAccountStateInstructionData struct {
	Instruction                    TokenInstruction
	DefaultAccountStateInstruction MetadataPointerInstruction
	AccountState                   token.AccountState
}

func (data *createInitializeDefaultAccountStateInstructionData) encode() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data.Instruction)
	binary.Write(&buf, binary.LittleEndian, data.DefaultAccountStateInstruction)
	buf.Write([]byte{byte(data.AccountState)})
	return buf.Bytes()
}

type createInitializeDefaultAccountStateInstruction struct {
	programID solana.PublicKey
	accounts  []*solana.AccountMeta
	data      []byte
}

func (inst *createInitializeDefaultAccountStateInstruction) ProgramID() solana.PublicKey {
	return inst.programID
}

func (inst *createInitializeDefaultAccountStateInstruction) Accounts() (out []*solana.AccountMeta) {
	return inst.accounts
}

func (inst *createInitializeDefaultAccountStateInstruction) Data() ([]byte, error) {
	return inst.data, nil
}
