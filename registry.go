package solana

import "fmt"

type InstructionDecoder func(*CompiledInstruction) (interface{}, error)

var instructionDecoderRegistry = map[string]InstructionDecoder{}

func RegisterInstructionDecoder(programID PublicKey, decoder InstructionDecoder) {
	p := programID.String()
	if _, found := instructionDecoderRegistry[p]; found {
		panic(fmt.Sprintf("unable to re-register instruction decoder for program %q", p))
	}
	instructionDecoderRegistry[p] = decoder
}

func DecodeInstruction(programID PublicKey, inst *CompiledInstruction) (interface{}, error) {
	p := programID.String()

	decoder, found := instructionDecoderRegistry[p]
	if !found {
		return nil, fmt.Errorf("unknown programID, cannot find any instruction decoder %q", p)
	}

	return decoder(inst)
}
