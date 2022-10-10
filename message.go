// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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
	"encoding/base64"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go/text"
	"github.com/gagliardetto/treeout"
)

type MessageAddressTableLookupSlice []MessageAddressTableLookup

// NumLookups returns the number of accounts in all the MessageAddressTableLookupSlice
func (lookups MessageAddressTableLookupSlice) NumLookups() int {
	count := 0
	for _, lookup := range lookups {
		// TODO: check if this is correct.
		count += len(lookup.ReadonlyIndexes)
		count += len(lookup.WritableIndexes)
	}
	return count
}

type MessageAddressTableLookup struct {
	AccountKey      PublicKey // The account key of the address table.
	WritableIndexes []uint8
	ReadonlyIndexes []uint8
}

type MessageVersion int

const (
	MessageVersionLegacy MessageVersion = 0 // default
	MessageVersionV0     MessageVersion = 1 // v0
)

type Message struct {
	version MessageVersion
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

	// List of address table lookups used to load additional accounts for this transaction.
	addressTableLookups MessageAddressTableLookupSlice

	addressTables map[PublicKey][]PublicKey
}

func (mx *Message) SetAddressTables(tables map[PublicKey][]PublicKey) error {
	if mx.addressTables != nil {
		return fmt.Errorf("address tables already set")
	}
	mx.addressTables = tables
	return mx.resolveLookups(tables)
}

// GetTables returns the address tables used by this message.
// NOTE: you must have called `SetAddressTable` one or more times before being able to use this method.
func (mx *Message) GetTables() map[PublicKey][]PublicKey {
	return mx.addressTables
}

var _ bin.EncoderDecoder = &Message{}

// SetVersion sets the message version.
// This method forces the message to be encoded in the specified version.
// NOTE: if you set lookups, the version will default to V0.
func (m *Message) SetVersion(version MessageVersion) *Message {
	// check if the version is valid
	switch version {
	case MessageVersionV0, MessageVersionLegacy:
	default:
		panic(fmt.Errorf("invalid message version: %d", version))
	}
	m.version = version
	return m
}

// GetVersion returns the message version.
func (m *Message) GetVersion() MessageVersion {
	return m.version
}

func (mx *Message) SetAddressTableLookups(lookups []MessageAddressTableLookup) *Message {
	mx.addressTableLookups = lookups
	mx.version = MessageVersionV0
	return mx
}

func (mx *Message) AddAddressTableLookup(lookup MessageAddressTableLookup) *Message {
	mx.addressTableLookups = append(mx.addressTableLookups, lookup)
	mx.version = MessageVersionV0
	return mx
}

func (mx *Message) GetAddressTableLookups() []MessageAddressTableLookup {
	return mx.addressTableLookups
}

func (mx *Message) EncodeToTree(txTree treeout.Branches) {
	switch mx.version {
	case MessageVersionV0:
		txTree.Child("Version: v0")
	case MessageVersionLegacy:
		txTree.Child("Version: legacy")
	default:
		txTree.Child(text.Sf("Version (unknown): %v", mx.version))
	}
	txTree.Child(text.Sf("RecentBlockhash: %s", mx.RecentBlockhash))

	txTree.Child(fmt.Sprintf("AccountKeys[len=%v]", len(mx.AccountKeys))).ParentFunc(func(accountKeysBranch treeout.Branches) {
		for keyIndex, key := range mx.AccountKeys {
			isFromTable := mx.IsVersioned() && keyIndex >= len(mx.AccountKeys)-mx.addressTableLookups.NumLookups()
			if isFromTable {
				accountKeysBranch.Child(text.Sf("%s (from Address Table Lookup)", text.ColorizeBG(key.String())))
			} else {
				accountKeysBranch.Child(text.ColorizeBG(key.String()))
			}
		}
	})

	if mx.IsVersioned() {
		txTree.Child(fmt.Sprintf("AddressTableLookups[len=%v]", len(mx.addressTableLookups))).ParentFunc(func(lookupsBranch treeout.Branches) {
			for _, lookup := range mx.addressTableLookups {
				lookupsBranch.Child(text.Sf("%s", text.ColorizeBG(lookup.AccountKey.String()))).ParentFunc(func(lookupBranch treeout.Branches) {
					lookupBranch.Child(text.Sf("WritableIndexes: %v", lookup.WritableIndexes))
					lookupBranch.Child(text.Sf("ReadonlyIndexes: %v", lookup.ReadonlyIndexes))
				})
			}
		})
	}

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
	switch mx.version {
	case MessageVersionV0:
		return mx.MarshalV0()
	case MessageVersionLegacy:
		return mx.MarshalLegacy()
	default:
		return nil, fmt.Errorf("invalid message version: %d", mx.version)
	}
}

func (mx *Message) MarshalLegacy() ([]byte, error) {
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

func (mx *Message) MarshalV0() ([]byte, error) {
	buf := []byte{
		mx.Header.NumRequiredSignatures,
		mx.Header.NumReadonlySignedAccounts,
		mx.Header.NumReadonlyUnsignedAccounts,
	}
	{

		numHardcodedKeys := len(mx.AccountKeys) - mx.addressTableLookups.NumLookups()
		bin.EncodeCompactU16Length(&buf, numHardcodedKeys)
		for _, key := range mx.AccountKeys[:numHardcodedKeys] {
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
	}
	versionNum := byte(mx.version) // TODO: what number is this?
	if versionNum > 127 {
		return nil, fmt.Errorf("invalid message version: %d", mx.version)
	}
	buf = append([]byte{byte(versionNum + 127)}, buf...)

	if mx.addressTableLookups != nil && len(mx.addressTableLookups) > 0 {
		// wite length of address table lookups as u8
		buf = append(buf, byte(len(mx.addressTableLookups)))
		for _, lookup := range mx.addressTableLookups {
			// write account pubkey
			buf = append(buf, lookup.AccountKey[:]...)
			// write writable indexes
			bin.EncodeCompactU16Length(&buf, len(lookup.WritableIndexes))
			buf = append(buf, lookup.WritableIndexes...)
			// write readonly indexes
			bin.EncodeCompactU16Length(&buf, len(lookup.ReadonlyIndexes))
			buf = append(buf, lookup.ReadonlyIndexes...)
		}
	} else {
		buf = append(buf, 0)
	}
	return buf, nil
}

func (mx Message) MarshalWithEncoder(encoder *bin.Encoder) error {
	out, err := mx.MarshalBinary()
	if err != nil {
		return err
	}
	return encoder.WriteBytes(out, false)
}

func (mx Message) ToBase64() string {
	out, _ := mx.MarshalBinary()
	return base64.StdEncoding.EncodeToString(out)
}

func (mx *Message) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// peek first byte to determine if this is a legacy or v0 message
	versionNum, err := decoder.Peek(1)
	if err != nil {
		return err
	}
	// TODO: is this the right way to determine if this is a legacy or v0 message?
	if versionNum[0] < 127 {
		mx.version = MessageVersionLegacy
	} else {
		mx.version = MessageVersionV0
	}
	switch mx.version {
	case MessageVersionV0:
		return mx.UnmarshalV0(decoder)
	case MessageVersionLegacy:
		return mx.UnmarshalLegacy(decoder)
	default:
		return fmt.Errorf("invalid message version: %d", mx.version)
	}
}

func (mx *Message) UnmarshalBase64(b64 string) error {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}
	return mx.UnmarshalWithDecoder(bin.NewBinDecoder(b))
}

func (mx *Message) resolveLookups(tables map[PublicKey][]PublicKey) (err error) {
	// add accounts from the address table lookups
	for _, lookup := range mx.addressTableLookups {
		table, ok := tables[lookup.AccountKey]
		if !ok {
			return fmt.Errorf("address table lookup not found for account: %v", lookup.AccountKey)
		}
		for _, idx := range lookup.WritableIndexes {
			if int(idx) >= len(table) {
				return fmt.Errorf("address table lookup index out of range: %v", idx)
			}
			mx.AccountKeys = append(mx.AccountKeys, table[idx])
		}
		for _, idx := range lookup.ReadonlyIndexes {
			if int(idx) >= len(table) {
				return fmt.Errorf("address table lookup index out of range: %v", idx)
			}
			mx.AccountKeys = append(mx.AccountKeys, table[idx])
		}
	}
	return nil
}

func (mx *Message) UnmarshalV0(decoder *bin.Decoder) (err error) {
	version, err := decoder.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read message version: %w", err)
	}
	// TODO: check version
	mx.version = MessageVersion(version - 127)

	// The middle of the message is the same as the legacy message:
	err = mx.UnmarshalLegacy(decoder)
	if err != nil {
		return err
	}

	// Read address table lookups length:
	addressTableLookupsLen, err := decoder.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read address table lookups length: %w", err)
	}
	if addressTableLookupsLen > 0 {
		mx.addressTableLookups = make([]MessageAddressTableLookup, addressTableLookupsLen)
		for i := 0; i < int(addressTableLookupsLen); i++ {
			// read account pubkey
			_, err = decoder.Read(mx.addressTableLookups[i].AccountKey[:])
			if err != nil {
				return fmt.Errorf("failed to read account pubkey: %w", err)
			}

			// read writable indexes
			writableIndexesLen, err := decoder.ReadCompactU16Length()
			if err != nil {
				return fmt.Errorf("failed to read writable indexes length: %w", err)
			}
			mx.addressTableLookups[i].WritableIndexes = make([]byte, writableIndexesLen)
			_, err = decoder.Read(mx.addressTableLookups[i].WritableIndexes)
			if err != nil {
				return fmt.Errorf("failed to read writable indexes: %w", err)
			}

			// read readonly indexes
			readonlyIndexesLen, err := decoder.ReadCompactU16Length()
			if err != nil {
				return fmt.Errorf("failed to read readonly indexes length: %w", err)
			}
			mx.addressTableLookups[i].ReadonlyIndexes = make([]byte, readonlyIndexesLen)
			_, err = decoder.Read(mx.addressTableLookups[i].ReadonlyIndexes)
			if err != nil {
				return fmt.Errorf("failed to read readonly indexes: %w", err)
			}
		}
	}
	return nil
}

func (mx *Message) UnmarshalLegacy(decoder *bin.Decoder) (err error) {
	{
		mx.Header.NumRequiredSignatures, err = decoder.ReadUint8()
		if err != nil {
			return fmt.Errorf("unable to decode mx.Header.NumRequiredSignatures: %w", err)
		}
		mx.Header.NumReadonlySignedAccounts, err = decoder.ReadUint8()
		if err != nil {
			return fmt.Errorf("unable to decode mx.Header.NumReadonlySignedAccounts: %w", err)
		}
		mx.Header.NumReadonlyUnsignedAccounts, err = decoder.ReadUint8()
		if err != nil {
			return fmt.Errorf("unable to decode mx.Header.NumReadonlyUnsignedAccounts: %w", err)
		}
	}
	{
		numAccountKeys, err := decoder.ReadCompactU16()
		if err != nil {
			return fmt.Errorf("unable to decode numAccountKeys: %w", err)
		}
		mx.AccountKeys = make([]PublicKey, numAccountKeys)
		for i := 0; i < numAccountKeys; i++ {
			_, err := decoder.Read(mx.AccountKeys[i][:])
			if err != nil {
				return fmt.Errorf("unable to decode mx.AccountKeys[%d]: %w", i, err)
			}
		}
	}
	{
		_, err := decoder.Read(mx.RecentBlockhash[:])
		if err != nil {
			return fmt.Errorf("unable to decode mx.RecentBlockhash: %w", err)
		}
	}
	{
		numInstructions, err := decoder.ReadCompactU16()
		if err != nil {
			return fmt.Errorf("unable to decode numInstructions: %w", err)
		}
		mx.Instructions = make([]CompiledInstruction, numInstructions)
		for instructionIndex := 0; instructionIndex < numInstructions; instructionIndex++ {
			programIDIndex, err := decoder.ReadUint8()
			if err != nil {
				return fmt.Errorf("unable to decode mx.Instructions[%d].ProgramIDIndex: %w", instructionIndex, err)
			}
			mx.Instructions[instructionIndex].ProgramIDIndex = uint16(programIDIndex)

			{
				numAccounts, err := decoder.ReadCompactU16()
				if err != nil {
					return fmt.Errorf("unable to decode numAccounts for ix[%d]: %w", instructionIndex, err)
				}
				mx.Instructions[instructionIndex].Accounts = make([]uint16, numAccounts)
				for i := 0; i < numAccounts; i++ {
					accountIndex, err := decoder.ReadUint8()
					if err != nil {
						return fmt.Errorf("unable to decode accountIndex for ix[%d].Accounts[%d]: %w", instructionIndex, i, err)
					}
					mx.Instructions[instructionIndex].Accounts[i] = uint16(accountIndex)
				}
			}
			{
				dataLen, err := decoder.ReadCompactU16()
				if err != nil {
					return fmt.Errorf("unable to decode dataLen for ix[%d]: %w", instructionIndex, err)
				}
				dataBytes, err := decoder.ReadNBytes(dataLen)
				if err != nil {
					return fmt.Errorf("unable to decode dataBytes for ix[%d]: %w", instructionIndex, err)
				}
				mx.Instructions[instructionIndex].Data = (Base58)(dataBytes)
			}
		}
	}

	return nil
}

func (m Message) checkPreconditions() {
	// if this is versioned,
	// and there are > 0 lookups,
	// but the address table is empty,
	// then we can't build the account meta list:
	if m.IsVersioned() && m.addressTableLookups.NumLookups() > 0 && (m.addressTables == nil || len(m.addressTables) == 0) {
		panic("cannot build account meta list without address tables")
	}
}

func (m Message) AccountMetaList() AccountMetaSlice {
	m.checkPreconditions()
	out := make(AccountMetaSlice, len(m.AccountKeys))
	for i, a := range m.AccountKeys {
		out[i] = &AccountMeta{
			PublicKey:  a,
			IsSigner:   m.IsSigner(a),
			IsWritable: m.IsWritable(a),
		}
	}
	// TODO: include accounts from lookup tables
	return out
}

func (m Message) IsVersioned() bool {
	return m.version != MessageVersionLegacy
}

// Signers returns the pubkeys of all accounts that are signers.
func (m Message) Signers() PublicKeySlice {
	m.checkPreconditions()
	out := make(PublicKeySlice, 0, len(m.AccountKeys))
	for _, a := range m.AccountKeys {
		if m.IsSigner(a) {
			out = append(out, a)
		}
	}
	return out
}

// Writable returns the pubkeys of all accounts that are writable.
func (m Message) Writable() (out PublicKeySlice) {
	m.checkPreconditions()
	for _, a := range m.AccountKeys {
		if m.IsWritable(a) {
			out = append(out, a)
		}
	}
	return out
}

func (m Message) ResolveProgramIDIndex(programIDIndex uint16) (PublicKey, error) {
	m.checkPreconditions()
	if int(programIDIndex) < len(m.AccountKeys) {
		return m.AccountKeys[programIDIndex], nil
	}
	return PublicKey{}, fmt.Errorf("programID index not found %d", programIDIndex)
}

func (m Message) HasAccount(account PublicKey) bool {
	m.checkPreconditions()
	for _, a := range m.AccountKeys {
		if a.Equals(account) {
			return true
		}
	}
	return false
}

func (m Message) IsSigner(account PublicKey) bool {
	m.checkPreconditions()
	for idx, acc := range m.AccountKeys {
		if acc.Equals(account) {
			return idx < int(m.Header.NumRequiredSignatures)
		}
	}
	return false
}

func (m Message) IsWritable(account PublicKey) bool {
	m.checkPreconditions()
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

func (m Message) signerKeys() []PublicKey {
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
