package solana

import (
	"bytes"

	"github.com/lunixbochs/struc"
)

type Transaction struct {
	Signatures []Signature `json:"signatures"`
	Message    Message     `json:"message"`
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []PublicKey   `json:"accountKeys"`
	RecentBlockhash PublicKey/* TODO: change to Hash */ `json:"recentBlockhash"`
	Instructions    []CompiledInstruction `json:"instructions"`
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyunsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type CompiledInstruction struct {
	ProgramIDIndex uint8    `json:"programIdIndex"`
	AccountsCount  ShortVec `json:"-" struc:"sizeof=Accounts"`
	Accounts       []uint8  `json:"accounts"`
	DataLength     ShortVec `json:"-" struc:"sizeof=Data"`
	Data           Base58   `json:"data"`
}

func TransactionFromData(in []byte) (*Transaction, error) {
	var out Transaction
	err := struc.Unpack(bytes.NewReader(in), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
