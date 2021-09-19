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
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go/text"
	"github.com/gagliardetto/treeout"
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

var _ bin.EncoderDecoder = &Transaction{}

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

var _ bin.EncoderDecoder = &Message{}

func (mx *Message) EncodeToTree(txTree treeout.Branches) {
	txTree.Child(text.Sf("RecentBlockhash: %s", mx.RecentBlockhash))

	txTree.Child(fmt.Sprintf("AccountKeys[len=%v]", len(mx.AccountKeys))).ParentFunc(func(accountKeysBranch treeout.Branches) {
		for _, key := range mx.AccountKeys {
			accountKeysBranch.Child(text.ColorizeBG(key.String()))
		}
	})

	txTree.Child("Header").ParentFunc(func(message treeout.Branches) {
		mx.Header.EncodeToTree(message)
	})
}

func (header *MessageHeader) EncodeToTree(mxBranch treeout.Branches) {
	mxBranch.Child(text.Sf("NumRequiredSignatures: %v", header.NumRequiredSignatures))
	mxBranch.Child(text.Sf("NumReadonlySignedAccounts: %v", header.NumReadonlySignedAccounts))
	mxBranch.Child(text.Sf("NumReadonlyUnsignedAccounts: %v", header.NumReadonlyUnsignedAccounts))
}

func (mx *Message) MarshalBinary() ([]byte, error) {
	buf := []byte{
		mx.Header.NumRequiredSignatures,
		mx.Header.NumReadonlySignedAccounts,
		mx.Header.NumReadonlyUnsignedAccounts,
	}

	bin.EncodeCompactU16Length(&buf, len(mx.AccountKeys))
	for _, key := range mx.AccountKeys {
		buf = append(buf, key[:]...)
	}

	buf = append(buf, mx.RecentBlockhash[:]...)

	bin.EncodeCompactU16Length(&buf, len(mx.Instructions))
	for _, instruction := range mx.Instructions {
		buf = append(buf, byte(instruction.ProgramIDIndex))
		bin.EncodeCompactU16Length(&buf, len(instruction.Accounts))
		for _, accountIdx := range instruction.Accounts {
			buf = append(buf, byte(accountIdx))
		}

		bin.EncodeCompactU16Length(&buf, len(instruction.Data))
		buf = append(buf, instruction.Data...)
	}
	return buf, nil
}

func (mx *Message) MarshalWithEncoder(encoder *bin.Encoder) error {
	out, err := mx.MarshalBinary()
	if err != nil {
		return err
	}
	return encoder.WriteBytes(out, false)
}

func (mx *Message) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	{
		mx.Header.NumRequiredSignatures, err = decoder.ReadUint8()
		if err != nil {
			return err
		}
		mx.Header.NumReadonlySignedAccounts, err = decoder.ReadUint8()
		if err != nil {
			return err
		}
		mx.Header.NumReadonlyUnsignedAccounts, err = decoder.ReadUint8()
		if err != nil {
			return err
		}
	}
	{
		numAccountKeys, err := bin.DecodeCompactU16LengthFromByteReader(decoder)
		if err != nil {
			return err
		}
		for i := 0; i < numAccountKeys; i++ {
			pubkeyBytes, err := decoder.ReadNBytes(32)
			if err != nil {
				return err
			}
			var sig PublicKey
			copy(sig[:], pubkeyBytes)
			mx.AccountKeys = append(mx.AccountKeys, sig)
		}
	}
	{
		recentBlockhashBytes, err := decoder.ReadNBytes(32)
		if err != nil {
			return err
		}
		var recentBlockhash Hash
		copy(recentBlockhash[:], recentBlockhashBytes)
		mx.RecentBlockhash = recentBlockhash
	}
	{
		numInstructions, err := bin.DecodeCompactU16LengthFromByteReader(decoder)
		if err != nil {
			return err
		}
		for i := 0; i < numInstructions; i++ {
			programIDIndex, err := decoder.ReadUint8()
			if err != nil {
				return err
			}
			var compInst CompiledInstruction
			compInst.ProgramIDIndex = uint16(programIDIndex)

			{
				numAccounts, err := bin.DecodeCompactU16LengthFromByteReader(decoder)
				if err != nil {
					return err
				}
				compInst.AccountCount = bin.Varuint16(numAccounts)
				for i := 0; i < numAccounts; i++ {
					accountIndex, err := decoder.ReadUint8()
					if err != nil {
						return err
					}
					compInst.Accounts = append(compInst.Accounts, uint16(accountIndex))
				}
			}
			{
				dataLen, err := bin.DecodeCompactU16LengthFromByteReader(decoder)
				if err != nil {
					return err
				}
				dataBytes, err := decoder.ReadNBytes(dataLen)
				if err != nil {
					return err
				}
				compInst.DataLength = bin.Varuint16(dataLen)
				compInst.Data = Base58(dataBytes)
			}
			mx.Instructions = append(mx.Instructions, compInst)
		}
	}

	return nil
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
	// NOTE: it is actually a uint8, but using a uint16 because uint8 is treated as a byte everywhere,
	// and that can be an issue.
	ProgramIDIndex uint16 `json:"programIdIndex"`

	AccountCount bin.Varuint16 `json:"-" bin:"sizeof=Accounts"`
	DataLength   bin.Varuint16 `json:"-" bin:"sizeof=Data"`

	// List of ordered indices into the message.accountKeys array indicating which accounts to pass to the program.
	// NOTE: it is actually a []uint8, but using a uint16 because []uint8 is treated as a []byte everywhere,
	// and that can be an issue.
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

func TransactionFromDecoder(decoder *bin.Decoder) (*Transaction, error) {
	var out *Transaction
	err := decoder.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func MustTransactionFromDecoder(decoder *bin.Decoder) *Transaction {
	out, err := TransactionFromDecoder(decoder)
	if err != nil {
		panic(err)
	}
	return out
}

// Unix timestamp (seconds since the Unix epoch)
type UnixTimeSeconds int64

func (res UnixTimeSeconds) Time() time.Time {
	return time.Unix(int64(res), 0)
}

func (res UnixTimeSeconds) String() string {
	return time.Unix(int64(res), 0).String()
}
