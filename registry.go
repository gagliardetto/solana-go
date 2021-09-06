package solana

import (
	"fmt"
)

// InstructionDecoder receives the AccountMeta FOR THAT INSTRUCTION,
// and not the accounts of the *Message object. Resolve with
// CompiledInstruction.ResolveInstructionAccounts(message) beforehand.
type InstructionDecoder func(instructionAccounts []*AccountMeta, data []byte) (interface{}, error)

var InstructionDecoderRegistry = map[string]InstructionDecoder{}

func RegisterInstructionDecoder(programID PublicKey, decoder InstructionDecoder) {
	pid := programID.String()
	if _, found := InstructionDecoderRegistry[pid]; found {
		panic(fmt.Sprintf("unable to re-register instruction decoder for program %q", pid))
	}

	InstructionDecoderRegistry[pid] = decoder
}

func DecodeInstruction(programID PublicKey, accounts []*AccountMeta, data []byte) (interface{}, error) {
	pid := programID.String()

	decoder, found := InstructionDecoderRegistry[pid]
	if !found {
		return nil, fmt.Errorf("instruction decoder not found for %s", pid)
	}

	return decoder(accounts, data)
}
