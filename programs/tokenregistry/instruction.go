package tokenregistry

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/gagliardetto/solana-go/text"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

func init() {
	solana.RegisterInstructionDecoder(ProgramID(), registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	var inst Instruction
	if err := bin.NewBinDecoder(data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return &inst, nil
}

func NewRegisterTokenInstruction(logo Logo, name Name, symbol Symbol, website Website, tokenMetaKey, ownerKey, tokenKey solana.PublicKey) *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{
			TypeID: 0,
			Impl: &RegisterToken{
				Logo:    logo,
				Name:    name,
				Website: website,
				Symbol:  symbol,
				Accounts: &RegisterTokenAccounts{
					TokenMeta: &solana.AccountMeta{tokenMetaKey, false, true},
					Owner:     &solana.AccountMeta{ownerKey, true, false},
					Token:     &solana.AccountMeta{tokenKey, false, false},
				},
			},
		},
	}
}

type Instruction struct {
	bin.BaseVariant
}

var _ bin.EncoderDecoder = &Instruction{}

func (i *Instruction) Accounts() (out []*solana.AccountMeta) {
	switch i.TypeID {
	case 0:
		accounts := i.Impl.(*RegisterToken).Accounts
		out = []*solana.AccountMeta{accounts.TokenMeta, accounts.Owner, accounts.Token}
	}
	return
}

func (i *Instruction) ProgramID() solana.PublicKey {
	return ProgramID()
}

func (i *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewBinEncoder(buf).Encode(i); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

var InstructionDefVariant = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"register_token", (*RegisterToken)(nil)},
})

func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

func (i *Instruction) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionDefVariant)
}

func (i *Instruction) MarshalWithEncoder(encoder *bin.Encoder) error {
	err := encoder.WriteUint32(i.TypeID, binary.LittleEndian)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(i.Impl)
}

type RegisterTokenAccounts struct {
	TokenMeta *solana.AccountMeta `text:"linear,notype"`
	Owner     *solana.AccountMeta `text:"linear,notype"`
	Token     *solana.AccountMeta `text:"linear,notype"`
}

type RegisterToken struct {
	Logo     Logo
	Name     Name
	Website  Website
	Symbol   Symbol
	Accounts *RegisterTokenAccounts `bin:"-"`
}

func (i *RegisterToken) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insufficient account")
	}
	i.Accounts = &RegisterTokenAccounts{
		TokenMeta: accounts[0],
		Owner:     accounts[1],
		Token:     accounts[2],
	}

	return nil
}
