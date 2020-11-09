// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	ProgramIDIndex uint8     `json:"programIdIndex"`
	AccountsCount  Varuint16 `json:"-" struc:"sizeof=Accounts"`
	Accounts       []uint8   `json:"accounts"`
	DataLength     Varuint16 `json:"-" struc:"sizeof=Data"`
	Data           Base58    `json:"data"`
}

func TransactionFromData(in []byte) (*Transaction, error) {
	var out Transaction
	err := struc.Unpack(bytes.NewReader(in), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
