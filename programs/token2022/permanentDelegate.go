package token2022

import (
	"bytes"
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

func CreateInitializePermanentDelegateInstruction(
	mint solana.PublicKey,
	permanentDelegate *solana.PublicKey,
) solana.Instruction {
	programID := ProgramID

	delegate := solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	if permanentDelegate != nil {
		delegate = *permanentDelegate
	}
	pointerData := createInitializePermanentDelegateInstructionData{
		Instruction: InitializePermanentDelegate,
		Delegate:    delegate,
	}

	ix := &instruction{
		programID: programID,
		accounts: []*solana.AccountMeta{
			solana.Meta(mint).WRITE(),
		},
		data: pointerData.encode(),
	}

	return ix
}

type createInitializePermanentDelegateInstructionData struct {
	Instruction TokenInstruction
	Delegate    solana.PublicKey
}

func (data *createInitializePermanentDelegateInstructionData) encode() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, data.Instruction)
	buf.Write(data.Delegate.Bytes())
	return buf.Bytes()
}
