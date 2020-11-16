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
	"fmt"

	bin "github.com/dfuse-io/binary"
)

type Transaction struct {
	Signatures []Signature `json:"signatures"`
	Message    Message     `json:"message"`
}

func (t *Transaction) TouchAccount(account PublicKey) bool {
	for _, a := range t.Message.AccountKeys {
		if a.Equals(account) {
			return true
		}
	}
	return false
}

func (t *Transaction) IsSigner(account PublicKey) bool {
	for idx, acc := range t.Message.AccountKeys {
		if acc.Equals(account) {
			return idx < int(t.Message.Header.NumRequiredSignatures)
		}
	}
	return false
}

func (t *Transaction) IsWritable(account PublicKey) bool {
	index := 0
	found := false
	for idx, acc := range t.Message.AccountKeys {
		if acc.Equals(account) {
			found = true
			index = idx
		}
	}
	if !found {
		return false
	}
	h := t.Message.Header
	return (index < int(h.NumRequiredSignatures-h.NumReadonlySignedAccounts)) ||
		((index >= int(h.NumRequiredSignatures)) && (index < len(t.Message.AccountKeys)-int(h.NumReadonlyunsignedAccounts)))
}

func (t *Transaction) ResolveProgramIDIndex(programIDIndex uint8) (PublicKey, error) {
	if int(programIDIndex) < len(t.Message.AccountKeys) {
		return t.Message.AccountKeys[programIDIndex], nil
	}
	return PublicKey{}, fmt.Errorf("programID index not found %d", programIDIndex)
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
	ProgramIDIndex uint8         `json:"programIdIndex"`
	AccountCount   bin.Varuint16 `json:"-" bin:"sizeof=Accounts"`
	Accounts       []uint8       `json:"accounts"`
	DataLength     bin.Varuint16 `json:"-" bin:"sizeof=Data"`
	Data           Base58        `json:"data"`
}

func TransactionFromData(in []byte) (*Transaction, error) {
	var out *Transaction
	decoder := bin.NewDecoder(in)
	err := decoder.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
