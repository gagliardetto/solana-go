package solana

import (
	"fmt"
)

type InstructionDecoder func([]PublicKey, *CompiledInstruction) (interface{}, error)

var InstructionDecoderRegistry = map[string]InstructionDecoder{}

func RegisterInstructionDecoder(programID PublicKey, decoder InstructionDecoder) {
	p := programID.String()
	if _, found := InstructionDecoderRegistry[p]; found {
		panic(fmt.Sprintf("unable to re-register instruction decoder for program %q", p))
	}
	InstructionDecoderRegistry[p] = decoder
}

func DecodeInstruction(programID PublicKey, accounts []PublicKey, inst *CompiledInstruction) (interface{}, error) {
	p := programID.String()

	decoder, found := InstructionDecoderRegistry[p]
	if !found {
		return nil, fmt.Errorf("unknown programID, cannot find any instruction decoder %q", p)
	}

	return decoder(accounts, inst)
}
