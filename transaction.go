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
	"bytes"
	"encoding/base64"
	"fmt"
	"sort"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/treeout"
	"go.uber.org/zap"

	"github.com/gagliardetto/solana-go/text"
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

// UnmarshalBase64 decodes a base64 encoded transaction.
func (tx *Transaction) UnmarshalBase64(b64 string) error {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}
	return tx.UnmarshalWithDecoder(bin.NewBinDecoder(b))
}

var _ bin.EncoderDecoder = &Transaction{}

func (t *Transaction) HasAccount(account PublicKey) (bool, error) {
	return t.Message.HasAccount(account)
}

func (t *Transaction) IsSigner(account PublicKey) bool {
	return t.Message.IsSigner(account)
}

func (t *Transaction) IsWritable(account PublicKey) (bool, error) {
	return t.Message.IsWritable(account)
}

func (t *Transaction) AccountMetaList() ([]*AccountMeta, error) {
	return t.Message.AccountMetaList()
}

func (t *Transaction) ResolveProgramIDIndex(programIDIndex uint16) (PublicKey, error) {
	return t.Message.ResolveProgramIDIndex(programIDIndex)
}

func (t *Transaction) GetAccountIndex(account PublicKey) (uint16, error) {
	return t.Message.GetAccountIndex(account)
}

// TransactionFromDecoder decodes a transaction from a decoder.
func TransactionFromDecoder(decoder *bin.Decoder) (*Transaction, error) {
	var out *Transaction
	err := decoder.Decode(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MustTransactionFromDecoder decodes a transaction from a decoder.
// Panics on error.
func MustTransactionFromDecoder(decoder *bin.Decoder) *Transaction {
	out, err := TransactionFromDecoder(decoder)
	if err != nil {
		panic(err)
	}
	return out
}

type CompiledInstruction struct {
	// Index into the message.accountKeys array indicating the program account that executes this instruction.
	// NOTE: it is actually a uint8, but using a uint16 because uint8 is treated as a byte everywhere,
	// and that can be an issue.
	ProgramIDIndex uint16 `json:"programIdIndex"`

	// List of ordered indices into the message.accountKeys array indicating which accounts to pass to the program.
	// NOTE: it is actually a []uint8, but using a uint16 because []uint8 is treated as a []byte everywhere,
	// and that can be an issue.
	Accounts []uint16 `json:"accounts"`

	// The program input data encoded in a base-58 string.
	Data Base58 `json:"data"`
}

func (ci *CompiledInstruction) ResolveInstructionAccounts(message *Message) ([]*AccountMeta, error) {
	out := make([]*AccountMeta, len(ci.Accounts))
	metas, err := message.AccountMetaList()
	if err != nil {
		return nil, err
	}
	for i, acct := range ci.Accounts {
		out[i] = metas[acct]
	}

	return out, nil
}

type Instruction interface {
	ProgramID() PublicKey     // the programID the instruction acts on
	Accounts() []*AccountMeta // returns the list of accounts the instructions requires
	Data() ([]byte, error)    // the binary encoded instructions
}

type TransactionOption interface {
	apply(opts *transactionOptions)
}

type transactionOptions struct {
	payer         PublicKey
	addressTables map[PublicKey]PublicKeySlice // [tablePubkey]addresses
}

type transactionOptionFunc func(opts *transactionOptions)

func (f transactionOptionFunc) apply(opts *transactionOptions) {
	f(opts)
}

func TransactionPayer(payer PublicKey) TransactionOption {
	return transactionOptionFunc(func(opts *transactionOptions) { opts.payer = payer })
}

func TransactionAddressTables(tables map[PublicKey]PublicKeySlice) TransactionOption {
	return transactionOptionFunc(func(opts *transactionOptions) { opts.addressTables = tables })
}

var debugNewTransaction = false

type TransactionBuilder struct {
	instructions    []Instruction
	recentBlockHash Hash
	opts            []TransactionOption
}

// NewTransactionBuilder creates a new instruction builder.
func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{}
}

// AddInstruction adds the provided instruction to the builder.
func (builder *TransactionBuilder) AddInstruction(instruction Instruction) *TransactionBuilder {
	builder.instructions = append(builder.instructions, instruction)
	return builder
}

// SetRecentBlockHash sets the recent blockhash for the instruction builder.
func (builder *TransactionBuilder) SetRecentBlockHash(recentBlockHash Hash) *TransactionBuilder {
	builder.recentBlockHash = recentBlockHash
	return builder
}

// WithOpt adds a TransactionOption.
func (builder *TransactionBuilder) WithOpt(opt TransactionOption) *TransactionBuilder {
	builder.opts = append(builder.opts, opt)
	return builder
}

// Set transaction fee payer.
// If not set, defaults to first signer account of the first instruction.
func (builder *TransactionBuilder) SetFeePayer(feePayer PublicKey) *TransactionBuilder {
	builder.opts = append(builder.opts, TransactionPayer(feePayer))
	return builder
}

// Build builds and returns a *Transaction.
func (builder *TransactionBuilder) Build() (*Transaction, error) {
	return NewTransaction(
		builder.instructions,
		builder.recentBlockHash,
		builder.opts...,
	)
}

type addressTablePubkeyWithIndex struct {
	addressTable PublicKey
	index        uint8
}

func NewTransaction(instructions []Instruction, recentBlockHash Hash, opts ...TransactionOption) (*Transaction, error) {
	if len(instructions) == 0 {
		return nil, fmt.Errorf("requires at-least one instruction to create a transaction")
	}

	options := transactionOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	feePayer := options.payer
	if feePayer.IsZero() {
		found := false
		for _, act := range instructions[0].Accounts() {
			if act.IsSigner {
				feePayer = act.PublicKey
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("cannot determine fee payer. You can ether pass the fee payer via the 'TransactionWithInstructions' option parameter or it falls back to the first instruction's first signer")
		}
	}

	addressLookupKeysMap := make(map[PublicKey]addressTablePubkeyWithIndex) // all accounts from tables as map
	for addressTablePubKey, addressTable := range options.addressTables {
		if len(addressTable) > 256 {
			return nil, fmt.Errorf("max lookup table index exceeded for %s table", addressTablePubKey)
		}

		for i, address := range addressTable {
			_, ok := addressLookupKeysMap[address]
			if ok {
				continue
			}

			addressLookupKeysMap[address] = addressTablePubkeyWithIndex{
				addressTable: addressTablePubKey,
				index:        uint8(i),
			}
		}
	}

	programIDs := make(PublicKeySlice, 0)
	accounts := []*AccountMeta{}
	for _, instruction := range instructions {
		accounts = append(accounts, instruction.Accounts()...)
		programIDs.UniqueAppend(instruction.ProgramID())
	}

	// for IsInvoked check
	programIDsMap := make(map[PublicKey]struct{}, len(programIDs))
	// Add programID to the account list
	for _, programID := range programIDs {
		accounts = append(accounts, &AccountMeta{
			PublicKey:  programID,
			IsSigner:   false,
			IsWritable: false,
		})

		programIDsMap[programID] = struct{}{}
	}

	// Sort. Prioritizing first by signer, then by writable
	sort.SliceStable(accounts, func(i, j int) bool {
		return accounts[i].less(accounts[j])
	})

	uniqAccountsMap := map[PublicKey]uint64{}
	uniqAccounts := []*AccountMeta{}
	for _, acc := range accounts {
		if index, found := uniqAccountsMap[acc.PublicKey]; found {
			uniqAccounts[index].IsWritable = uniqAccounts[index].IsWritable || acc.IsWritable
			continue
		}
		uniqAccounts = append(uniqAccounts, acc)
		uniqAccountsMap[acc.PublicKey] = uint64(len(uniqAccounts) - 1)
	}

	if debugNewTransaction {
		zlog.Debug("unique account sorted", zap.Int("account_count", len(uniqAccounts)))
	}
	// Move fee payer to the front
	feePayerIndex := -1
	for idx, acc := range uniqAccounts {
		if acc.PublicKey.Equals(feePayer) {
			feePayerIndex = idx
		}
	}
	if debugNewTransaction {
		zlog.Debug("current fee payer index", zap.Int("fee_payer_index", feePayerIndex))
	}

	accountCount := len(uniqAccounts)
	if feePayerIndex < 0 {
		// fee payer is not part of accounts we want to add it
		accountCount++
	}
	allKeys := make([]*AccountMeta, accountCount)

	itr := 1
	for idx, uniqAccount := range uniqAccounts {
		if idx == feePayerIndex {
			uniqAccount.IsSigner = true
			uniqAccount.IsWritable = true
			allKeys[0] = uniqAccount
			continue
		}
		allKeys[itr] = uniqAccount
		itr++
	}

	if feePayerIndex < 0 {
		// fee payer is not part of accounts we want to add it
		feePayerAccount := &AccountMeta{
			PublicKey:  feePayer,
			IsSigner:   true,
			IsWritable: true,
		}
		allKeys[0] = feePayerAccount
	}

	message := Message{
		RecentBlockhash: recentBlockHash,
	}
	lookupsMap := make(map[PublicKey]struct { // extended MessageAddressTableLookup
		AccountKey      PublicKey // The account key of the address table.
		WritableIndexes []uint8
		Writable        []PublicKey
		ReadonlyIndexes []uint8
		Readonly        []PublicKey
	})
	for idx, acc := range allKeys {

		if debugNewTransaction {
			zlog.Debug("transaction account",
				zap.Int("account_index", idx),
				zap.Stringer("account_pub_key", acc.PublicKey),
			)
		}

		addressLookupKeyEntry, isPresentedInTables := addressLookupKeysMap[acc.PublicKey]
		_, isInvoked := programIDsMap[acc.PublicKey]
		// skip fee payer
		if isPresentedInTables && idx != 0 && !acc.IsSigner && !isInvoked {
			lookup := lookupsMap[addressLookupKeyEntry.addressTable]
			if acc.IsWritable {
				lookup.WritableIndexes = append(lookup.WritableIndexes, addressLookupKeyEntry.index)
				lookup.Writable = append(lookup.Writable, acc.PublicKey)
			} else {
				lookup.ReadonlyIndexes = append(lookup.ReadonlyIndexes, addressLookupKeyEntry.index)
				lookup.Readonly = append(lookup.Readonly, acc.PublicKey)
			}

			lookupsMap[addressLookupKeyEntry.addressTable] = lookup
			continue // prevent changing message.Header properties
		}

		message.AccountKeys = append(message.AccountKeys, acc.PublicKey)

		if acc.IsSigner {
			message.Header.NumRequiredSignatures++
			if !acc.IsWritable {
				message.Header.NumReadonlySignedAccounts++
			}
			continue
		}

		if !acc.IsWritable {
			message.Header.NumReadonlyUnsignedAccounts++
		}
	}

	var lookupsWritableKeys []PublicKey
	var lookupsReadOnlyKeys []PublicKey
	if len(lookupsMap) > 0 {
		lookups := make([]MessageAddressTableLookup, 0, len(lookupsMap))

		for tablePubKey, l := range lookupsMap {
			lookupsWritableKeys = append(lookupsWritableKeys, l.Writable...)
			lookupsReadOnlyKeys = append(lookupsReadOnlyKeys, l.Readonly...)

			lookups = append(lookups, MessageAddressTableLookup{
				AccountKey:      tablePubKey,
				WritableIndexes: l.WritableIndexes,
				ReadonlyIndexes: l.ReadonlyIndexes,
			})
		}

		// prevent error created in ResolveLookups
		err := message.SetAddressTables(options.addressTables)
		if err != nil {
			return nil, fmt.Errorf("SetAddressTables: %s", err)
		}
		message.SetAddressTableLookups(lookups)
	}

	var idx uint16
	accountKeyIndex := make(map[string]uint16, len(message.AccountKeys)+len(lookupsWritableKeys)+len(lookupsReadOnlyKeys))
	for _, acc := range message.AccountKeys {
		accountKeyIndex[acc.String()] = idx
		idx++
	}
	for _, acc := range lookupsWritableKeys {
		accountKeyIndex[acc.String()] = idx
		idx++
	}
	for _, acc := range lookupsReadOnlyKeys {
		accountKeyIndex[acc.String()] = idx
		idx++
	}

	if debugNewTransaction {
		zlog.Debug("message header compiled",
			zap.Uint8("num_required_signatures", message.Header.NumRequiredSignatures),
			zap.Uint8("num_readonly_signed_accounts", message.Header.NumReadonlySignedAccounts),
			zap.Uint8("num_readonly_unsigned_accounts", message.Header.NumReadonlyUnsignedAccounts),
		)
	}

	for txIdx, instruction := range instructions {
		accounts = instruction.Accounts()
		accountIndex := make([]uint16, len(accounts))
		for idx, acc := range accounts {
			accountIndex[idx] = accountKeyIndex[acc.PublicKey.String()]
		}
		data, err := instruction.Data()
		if err != nil {
			return nil, fmt.Errorf("unable to encode instructions [%d]: %w", txIdx, err)
		}
		message.Instructions = append(message.Instructions, CompiledInstruction{
			ProgramIDIndex: accountKeyIndex[instruction.ProgramID().String()],
			Accounts:       accountIndex,
			Data:           data,
		})
	}

	return &Transaction{
		Message: message,
	}, nil
}

type privateKeyGetter func(key PublicKey) *PrivateKey

func (tx *Transaction) MarshalBinary() ([]byte, error) {
	messageContent, err := tx.Message.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode tx.Message to binary: %w", err)
	}

	var signatureCount []byte
	bin.EncodeCompactU16Length(&signatureCount, len(tx.Signatures))
	output := make([]byte, 0, len(signatureCount)+len(signatureCount)*64+len(messageContent))
	output = append(output, signatureCount...)
	for _, sig := range tx.Signatures {
		output = append(output, sig[:]...)
	}
	output = append(output, messageContent...)

	return output, nil
}

func (tx Transaction) MarshalWithEncoder(encoder *bin.Encoder) error {
	out, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	return encoder.WriteBytes(out, false)
}

func (tx *Transaction) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	{
		numSignatures, err := decoder.ReadCompactU16()
		if err != nil {
			return fmt.Errorf("unable to read numSignatures: %w", err)
		}
		if numSignatures > decoder.Remaining()/64 {
			return fmt.Errorf("numSignatures %d is too large for remaining bytes %d", numSignatures, decoder.Remaining())
		}

		tx.Signatures = make([]Signature, numSignatures)
		for i := 0; i < numSignatures; i++ {
			_, err := decoder.Read(tx.Signatures[i][:])
			if err != nil {
				return fmt.Errorf("unable to read tx.Signatures[%d]: %w", i, err)
			}
		}
	}
	{
		err := tx.Message.UnmarshalWithDecoder(decoder)
		if err != nil {
			return fmt.Errorf("unable to decode tx.Message: %w", err)
		}
	}
	return nil
}

func (tx *Transaction) PartialSign(getter privateKeyGetter) (out []Signature, err error) {
	messageContent, err := tx.Message.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to encode message for signing: %w", err)
	}
	signerKeys := tx.Message.signerKeys()

	signedSignatures := []Signature{}
	for _, key := range signerKeys {
		privateKey := getter(key)
		if privateKey != nil {
			s, err := privateKey.Sign(messageContent)
			if err != nil {
				return nil, fmt.Errorf("failed to signed with key %q: %w", key.String(), err)
			}
			signedSignatures = append(signedSignatures, s)
		}
	}
	tx.Signatures = append(tx.Signatures, signedSignatures...)
	return tx.Signatures, nil
}

func (tx *Transaction) Sign(getter privateKeyGetter) (out []Signature, err error) {
	signerKeys := tx.Message.signerKeys()
	for _, key := range signerKeys {
		if getter(key) == nil {
			return nil, fmt.Errorf("signer key %q not found. Ensure all the signer keys are in the vault", key.String())
		}
	}
	return tx.PartialSign(getter)
}

func (tx *Transaction) EncodeTree(encoder *text.TreeEncoder) (int, error) {
	tx.EncodeToTree(encoder)
	return encoder.WriteString(encoder.Tree.String())
}

// String returns a human-readable string representation of the transaction data.
// To disable colors, set "github.com/gagliardetto/solana-go/text".DisableColors = true
func (tx *Transaction) String() string {
	buf := new(bytes.Buffer)
	_, err := tx.EncodeTree(text.NewTreeEncoder(buf, ""))
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (tx Transaction) ToBase64() (string, error) {
	out, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(out), nil
}

func (tx Transaction) MustToBase64() string {
	out, err := tx.ToBase64()
	if err != nil {
		panic(err)
	}
	return out
}

func (tx *Transaction) EncodeToTree(parent treeout.Branches) {
	parent.ParentFunc(func(txTree treeout.Branches) {
		txTree.Child(fmt.Sprintf("Signatures[len=%v]", len(tx.Signatures))).ParentFunc(func(signaturesBranch treeout.Branches) {
			for _, sig := range tx.Signatures {
				signaturesBranch.Child(sig.String())
			}
		})

		txTree.Child("Message").ParentFunc(func(messageBranch treeout.Branches) {
			tx.Message.EncodeToTree(messageBranch)
		})
	})

	parent.Child(fmt.Sprintf("Instructions[len=%v]", len(tx.Message.Instructions))).ParentFunc(func(message treeout.Branches) {
		for _, inst := range tx.Message.Instructions {

			progKey, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
			if err == nil {
				accounts, err := inst.ResolveInstructionAccounts(&tx.Message)
				if err != nil {
					message.Child(fmt.Sprintf(text.RedBG("cannot ResolveInstructionAccounts: %s"), err))
					return
				}
				decodedInstruction, err := DecodeInstruction(progKey, accounts, inst.Data)
				if err == nil {
					if enToTree, ok := decodedInstruction.(text.EncodableToTree); ok {
						enToTree.EncodeToTree(message)
					} else {
						message.Child(spew.Sdump(decodedInstruction))
					}
				} else {
					// TODO: log error?
					message.Child(fmt.Sprintf(text.RedBG("cannot decode instruction for %s program: %s"), progKey, err)).
						Child(text.IndigoBG("Program") + ": " + text.Bold("<unknown>") + " " + text.ColorizeBG(progKey.String())).
						//
						ParentFunc(func(programBranch treeout.Branches) {
							programBranch.Child(text.Purple(text.Bold("Instruction")) + ": " + text.Bold("<unknown>")).
								//
								ParentFunc(func(instructionBranch treeout.Branches) {
									// Data of the instruction call:
									instructionBranch.Child(text.Sf("data[len=%v bytes]", len(inst.Data))).ParentFunc(func(paramsBranch treeout.Branches) {
										paramsBranch.Child(bin.FormatByteSlice(inst.Data))
									})

									// Accounts of the instruction call:
									instructionBranch.Child(text.Sf("accounts[len=%v]", len(accounts))).ParentFunc(func(accountsBranch treeout.Branches) {
										for i := range accounts {
											accountsBranch.Child(formatMeta(text.Sf("accounts[%v]", i), accounts[i]))
										}
									})
								})
						})
				}
			} else {
				message.Child(fmt.Sprintf(text.RedBG("cannot ResolveProgramIDIndex: %s"), err))
			}
		}
	})
}

func formatMeta(name string, meta *AccountMeta) string {
	if meta == nil {
		return text.Shakespeare(name) + ": " + "<nil>"
	}
	out := text.Shakespeare(name) + ": " + text.ColorizeBG(meta.PublicKey.String())
	out += " ["
	if meta.IsWritable {
		out += "WRITE"
	}
	if meta.IsSigner {
		if meta.IsWritable {
			out += ", "
		}
		out += "SIGN"
	}
	out += "] "
	return out
}

// VerifySignatures verifies all the signatures in the transaction
// against the pubkeys of the signers.
func (tx *Transaction) VerifySignatures() error {
	msg, err := tx.Message.MarshalBinary()
	if err != nil {
		return err
	}

	signers := tx.Message.Signers()

	if len(signers) != len(tx.Signatures) {
		return fmt.Errorf(
			"got %v signers, but %v signatures",
			len(signers),
			len(tx.Signatures),
		)
	}

	for i, sig := range tx.Signatures {
		if !sig.Verify(signers[i], msg) {
			return fmt.Errorf("invalid signature by %s", signers[i].String())
		}
	}

	return nil
}
