package solana

import (
	"bytes"
	"fmt"
	"sort"

	bin "github.com/dfuse-io/binary"

	"go.uber.org/zap"
)

type TransactionInstruction interface {
	Accounts() []*AccountMeta // returns the list of accounts the instructions requires
	ProgramID() PublicKey     // the programID the instruction acts on
	Data() ([]byte, error)    // the binary encoded instructions
}

type Options struct {
	payer *PublicKey
}

func TransactionWithInstructions(instructions []TransactionInstruction, opt *Options) (*Transaction, error) {
	if len(instructions) == 0 {
		return nil, fmt.Errorf("requires at-least one instruction to create a transaction")
	}

	if opt == nil {
		opt = &Options{}
	}

	feePayer := opt.payer
	if feePayer == nil {
		for _, act := range instructions[0].Accounts() {
			if act.IsSigner {
				feePayer = &act.PublicKey
				break
			}
		}
	}
	if feePayer == nil {
		return nil, fmt.Errorf("cannot determine fee payer. You can ether pass the fee payer vai the 'TransactionWithInstructions' option parameter or it fallback to the first instruction's first signer")
	}

	programIDs := map[string]bool{}
	accounts := []*AccountMeta{}
	for _, instruction := range instructions {
		for _, key := range instruction.Accounts() {
			accounts = append(accounts, key)
		}
		programIDs[instruction.ProgramID().String()] = true
	}

	// Add programID to the account list
	for programId, _ := range programIDs {
		accounts = append(accounts, &AccountMeta{
			PublicKey:  MustPublicKeyFromBase58(programId),
			IsSigner:   false,
			IsWritable: false,
		})
	}

	// Sort. Prioritizing first by signer, then by writable
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].less(accounts[j])
	})

	uniqAccountsMap := map[string]uint64{}
	uniqAccounts := []*AccountMeta{}
	for _, acc := range accounts {
		if index, found := uniqAccountsMap[acc.PublicKey.String()]; found {
			uniqAccounts[index].IsWritable = uniqAccounts[index].IsWritable || acc.IsWritable
			continue
		}
		uniqAccounts = append(uniqAccounts, acc)
		uniqAccountsMap[acc.PublicKey.String()] = uint64(len(uniqAccounts) - 1)
	}

	zlog.Debug("unique account sorted", zap.Int("account_count", len(uniqAccounts)))
	// Move fee payer to the front
	feePayerIndex := -1
	for idx, acc := range uniqAccounts {
		if acc.PublicKey.Equals(*feePayer) {
			feePayerIndex = idx
		}
	}
	zlog.Debug("current fee payer index", zap.Int("fee_payer_index", feePayerIndex))

	accountCount := len(uniqAccounts)
	if feePayerIndex < 0 {
		// fee payer is not part of accounts we want to add it
		accountCount++
	}
	finalAccounts := make([]*AccountMeta, accountCount)

	itr := 1
	for idx, uniqAccount := range uniqAccounts {
		if idx == feePayerIndex {
			uniqAccount.IsSigner = true
			uniqAccount.IsWritable = true
			finalAccounts[0] = uniqAccount
			continue
		}
		finalAccounts[itr] = uniqAccount
		itr++
	}

	message := Message{
		AccountKeys:     nil,
		RecentBlockhash: PublicKey{},
		Instructions:    nil,
	}
	accountKeyIndex := map[string]uint8{}
	for idx, acc := range finalAccounts {
		message.AccountKeys = append(message.AccountKeys, acc.PublicKey)
		accountKeyIndex[acc.PublicKey.String()] = uint8(idx)
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

	for trxIdx, instruction := range instructions {
		accounts = instruction.Accounts()
		accountIndex := make([]uint8, len(accounts))
		for idx, acc := range accounts {
			accountIndex[idx] = accountKeyIndex[acc.PublicKey.String()]
		}
		data, err := instruction.Data()
		if err != nil {
			return nil, fmt.Errorf("unable to encode instructions [%d]: %w", trxIdx, err)
		}
		message.Instructions = append(message.Instructions, CompiledInstruction{
			ProgramIDIndex: accountKeyIndex[instruction.ProgramID().String()],
			AccountCount:   bin.Varuint16(uint16(len(accountIndex))),
			Accounts:       accountIndex,
			DataLength:     bin.Varuint16(uint16(len(data))),
			Data:           data,
		})
	}

	return &Transaction{
		Message: message,
	}, nil

}

type privateKeyGetter func(key PublicKey) *PrivateKey

func (t *Transaction) Sign(getter privateKeyGetter) (out []Signature, err error) {
	buf := new(bytes.Buffer)
	if err = bin.NewEncoder(buf).Encode(t.Message); err != nil {
		return nil, fmt.Errorf("unable to encode message for signing: %w", err)
	}
	messageCnt := buf.Bytes()

	signerKeys := t.Message.signerKeys()

	for _, key := range signerKeys {
		privateKey := getter(key)
		if privateKey == nil {
			return nil, fmt.Errorf("signer key %q not found. Ensure all the signer keys are in the vault", key.String())
		}

		s, err := privateKey.Sign(messageCnt)
		if err != nil {
			return nil, fmt.Errorf("failed to signed with key %q: %w", key.String(), err)
		}

		t.Signatures = append(t.Signatures, s)
	}
	return t.Signatures, nil
}
