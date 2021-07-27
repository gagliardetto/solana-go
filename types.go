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
	// A list of base-58 encoded signatures applied to the transaction.
	// The list is always of length `message.header.numRequiredSignatures` and not empty.
	// The signature at index `i` corresponds to the public key at index
	// `i` in `message.account_keys`. The first one is used as the transaction id.
	Signatures []Signature `json:"signatures"`

	// Defines the content of the transaction.
	Message Message `json:"message"`
}

func (t *Transaction) TouchAccount(account PublicKey) bool   { return t.Message.TouchAccount(account) }
func (t *Transaction) IsSigner(account PublicKey) bool       { return t.Message.IsSigner(account) }
func (t *Transaction) IsWritable(account PublicKey) bool     { return t.Message.IsWritable(account) }
func (t *Transaction) AccountMetaList() (out []*AccountMeta) { return t.Message.AccountMetaList() }
func (t *Transaction) ResolveProgramIDIndex(programIDIndex uint16) (PublicKey, error) {
	return t.Message.ResolveProgramIDIndex(programIDIndex)
}

type Message struct {
	// List of base-58 encoded public keys used by the transaction,
	// including by the instructions and for signatures.
	// The first `message.header.numRequiredSignatures` public keys must sign the transaction.
	AccountKeys []PublicKey `json:"accountKeys"`

	// Details the account types and signatures required by the transaction.
	Header MessageHeader `json:"header"`

	// A base-58 encoded hash of a recent block in the ledger used to
	// prevent transaction duplication and to give transactions lifetimes.
	RecentBlockhash Hash `json:"recentBlockhash"`

	// List of program instructions that will be executed in sequence
	// and committed in one atomic transaction if all succeed.
	Instructions []CompiledInstruction `json:"instructions"`
}

func (m *Message) AccountMetaList() (out []*AccountMeta) {
	for _, a := range m.AccountKeys {
		out = append(out, &AccountMeta{
			PublicKey:  a,
			IsSigner:   m.IsSigner(a),
			IsWritable: m.IsWritable(a),
		})
	}
	return out
}

func (m *Message) ResolveProgramIDIndex(programIDIndex uint16) (PublicKey, error) {
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
	// The total number of signatures required to make the transaction valid.
	// The signatures must match the first `numRequiredSignatures` of `message.account_keys`.
	NumRequiredSignatures uint8 `json:"numRequiredSignatures"`

	// The last numReadonlySignedAccounts of the signed keys are read-only accounts.
	// Programs may process multiple transactions that load read-only accounts within
	// a single PoH entry, but are not permitted to credit or debit lamports or modify
	// account data.
	// Transactions targeting the same read-write account are evaluated sequentially.
	NumReadonlySignedAccounts uint8 `json:"numReadonlySignedAccounts"`

	// The last `numReadonlyUnsignedAccounts` of the unsigned keys are read-only accounts.
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type CompiledInstruction struct {
	// Index into the message.accountKeys array indicating the program account that executes this instruction.
	ProgramIDIndex uint16 `json:"programIdIndex"`

	AccountCount bin.Varuint16 `json:"-" bin:"sizeof=Accounts"`
	DataLength   bin.Varuint16 `json:"-" bin:"sizeof=Data"`

	// List of ordered indices into the message.accountKeys array indicating which accounts to pass to the program.
	Accounts []uint16 `json:"accounts"`

	// The program input data encoded in a base-58 string.
	Data Base58 `json:"data"`
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
func MustTransactionFromData(in []byte) *Transaction {
	out, err := TransactionFromData(in)
	if err != nil {
		panic(err)
	}
	return out
}
