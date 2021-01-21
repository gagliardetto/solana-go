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

func (t *Transaction) TouchAccount(account PublicKey) bool   { return t.Message.TouchAccount(account) }
func (t *Transaction) IsSigner(account PublicKey) bool       { return t.Message.IsSigner(account) }
func (t *Transaction) IsWritable(account PublicKey) bool     { return t.Message.IsWritable(account) }
func (t *Transaction) AccountMetaList() (out []*AccountMeta) { return t.Message.AccountMetaList() }
func (t *Transaction) ResolveProgramIDIndex(programIDIndex uint8) (PublicKey, error) {
	return t.Message.ResolveProgramIDIndex(programIDIndex)
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []PublicKey   `json:"accountKeys"`
	RecentBlockhash PublicKey/* TODO: change to Hash */ `json:"recentBlockhash"`
	Instructions    []CompiledInstruction `json:"instructions"`
}

func (m *Message) AccountMetaList() (out []*AccountMeta) {
	return m.AccountMetaList()
	for _, a := range m.AccountKeys {
		out = append(out, &AccountMeta{
			PublicKey:  a,
			IsSigner:   m.IsSigner(a),
			IsWritable: m.IsWritable(a),
		})
	}
	return out
}

func (m *Message) ResolveProgramIDIndex(programIDIndex uint8) (PublicKey, error) {
	if int(programIDIndex) < len(m.AccountKeys) {
		return m.AccountKeys[programIDIndex], nil
	}
	return PublicKey{}, fmt.Errorf("programID index not found %d", programIDIndex)
}

func (m *Message) TouchAccount(account PublicKey) bool {
	for _, a := range m.AccountKeys {
		if a.Equals(account) {
			return true
		}
	}
	return false
}

func (m *Message) IsSigner(account PublicKey) bool {
	for idx, acc := range m.AccountKeys {
		if acc.Equals(account) {
			return idx < int(m.Header.NumRequiredSignatures)
		}
	}
	return false
}

func (m *Message) IsWritable(account PublicKey) bool {
	index := 0
	found := false
	for idx, acc := range m.AccountKeys {
		if acc.Equals(account) {
			found = true
			index = idx
		}
	}
	if !found {
		return false
	}
	h := m.Header
	return (index < int(h.NumRequiredSignatures-h.NumReadonlySignedAccounts)) ||
		((index >= int(h.NumRequiredSignatures)) && (index < len(m.AccountKeys)-int(h.NumReadonlyUnsignedAccounts)))
}

func (m *Message) signerKeys() []PublicKey {
	return m.AccountKeys[0:m.Header.NumRequiredSignatures]
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type CompiledInstruction struct {
	ProgramIDIndex uint8         `json:"programIdIndex"`
	AccountCount   bin.Varuint16 `json:"-" bin:"sizeof=Accounts"`
	Accounts       []uint8       `json:"accounts"`
	DataLength     bin.Varuint16 `json:"-" bin:"sizeof=Data"`
	Data           Base58        `json:"data"`
}

func (ci *CompiledInstruction) ResolveInstructionAccounts(message *Message) (out []*AccountMeta) {
	metas := message.AccountMetaList()
	for _, acct := range ci.Accounts {
		out = append(out, metas[acct])
	}
	return
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
