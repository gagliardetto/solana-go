package tokenregistry

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/dfuse-io/solana-go/text"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

func init() {
	solana.RegisterInstructionDecoder(ProgramID(), registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, rawInstruction *solana.CompiledInstruction) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, rawInstruction)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, compiledInstruction *solana.CompiledInstruction) (*Instruction, error) {
	var inst Instruction
	if err := bin.NewDecoder(compiledInstruction.Data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountSettable); ok {
		err := v.SetAccounts(accounts, compiledInstruction.Accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return &inst, nil
}

type Instruction struct {
	bin.BaseVariant
}

var InstructionDefVariant = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"register_token", (*RegisterToken)(nil)},
})

func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

func (i *Instruction) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionDefVariant)
}

func (i *Instruction) MarshalBinary(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(i.TypeID, binary.LittleEndian)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(i.Impl)
}

type RegisterTokenAccounts struct {
	MintMeta *solana.AccountMeta `text:"linear,notype"`
	Owner    *solana.AccountMeta `text:"linear,notype"`
	Token    *solana.AccountMeta `text:"linear,notype"`
}

type RegisterToken struct {
	Logo     Logo
	Name     Name
	Symbol   Symbol
	Accounts *RegisterTokenAccounts `bin:"-"`
}

func NewRegisterToken(logo Logo, name Name, symbol Symbol, programIdIndex uint8, mintMetaIdx uint8, ownerIdx uint8, tokenIdx uint8) (*solana.CompiledInstruction, error) {
	instruction := &Instruction{
		BaseVariant: bin.BaseVariant{
			TypeID: 0,
			Impl: &RegisterToken{
				Logo:   logo,
				Name:   name,
				Symbol: symbol,
			},
		},
	}

	buf := new(bytes.Buffer)
	if err := bin.NewEncoder(buf).Encode(instruction); err != nil {
		return nil, fmt.Errorf("new register token: encode: %w", err)
	}
	data := buf.Bytes()

	/// 0. `[writable]` The register data's account to initialize
	/// 1. `[signer]` The registry's owner
	/// 2. `[]` The mint address to link with this registration

	return &solana.CompiledInstruction{
		ProgramIDIndex: programIdIndex,
		AccountCount:   3,
		Accounts:       []uint8{mintMetaIdx, ownerIdx, tokenIdx},
		DataLength:     bin.Varuint16(len(data)),
		Data:           data,
	}, nil

}

func (i *RegisterToken) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 9 {
		return fmt.Errorf("insuficient account")
	}
	i.Accounts = &RegisterTokenAccounts{
		MintMeta: accounts[instructionActIdx[0]],
		Owner:    accounts[instructionActIdx[1]],
		Token:    accounts[instructionActIdx[2]],
	}

	return nil
}
