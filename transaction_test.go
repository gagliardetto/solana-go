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
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/magiconair/properties/assert"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/require"
)

type testTransactionInstructions struct {
	accounts  []*AccountMeta
	data      []byte
	programID PublicKey
}

func (t *testTransactionInstructions) Accounts() []*AccountMeta {
	return t.accounts
}

func (t *testTransactionInstructions) ProgramID() PublicKey {
	return t.programID
}

func (t *testTransactionInstructions) Data() ([]byte, error) {
	return t.data, nil
}

func TestNewTransaction(t *testing.T) {
	debugNewTransaction = true

	instructions := []Instruction{
		&testTransactionInstructions{
			accounts: []*AccountMeta{
				{PublicKey: MustPublicKeyFromBase58("A9QnpgfhCkmiBSjgBuWk76Wo3HxzxvDopUq9x6UUMmjn"), IsSigner: true, IsWritable: false},
				{PublicKey: MustPublicKeyFromBase58("9hFtYBYmBJCVguRYs9pBTWKYAFoKfjYR7zBPpEkVsmD"), IsSigner: true, IsWritable: true},
			},
			data:      []byte{0xaa, 0xbb},
			programID: MustPublicKeyFromBase58("11111111111111111111111111111111"),
		},
		&testTransactionInstructions{
			accounts: []*AccountMeta{
				{PublicKey: MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111"), IsSigner: false, IsWritable: false},
				{PublicKey: MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111"), IsSigner: false, IsWritable: true},
				{PublicKey: MustPublicKeyFromBase58("9hFtYBYmBJCVguRYs9pBTWKYAFoKfjYR7zBPpEkVsmD"), IsSigner: false, IsWritable: true},
				{PublicKey: MustPublicKeyFromBase58("6FzXPEhCJoBx7Zw3SN9qhekHemd6E2b8kVguitmVAngW"), IsSigner: true, IsWritable: false},
			},
			data:      []byte{0xcc, 0xdd},
			programID: MustPublicKeyFromBase58("Vote111111111111111111111111111111111111111"),
		},
	}

	blockhash, err := HashFromBase58("A9QnpgfhCkmiBSjgBuWk76Wo3HxzxvDopUq9x6UUMmjn")
	require.NoError(t, err)

	trx, err := NewTransaction(instructions, blockhash)
	require.NoError(t, err)

	assert.Equal(t, trx.Message.Header, MessageHeader{
		NumRequiredSignatures:       3,
		NumReadonlySignedAccounts:   1,
		NumReadonlyUnsignedAccounts: 3,
	})

	assert.Equal(t, trx.Message.RecentBlockhash, blockhash)

	assert.Equal(t, trx.Message.AccountKeys, []PublicKey{
		MustPublicKeyFromBase58("A9QnpgfhCkmiBSjgBuWk76Wo3HxzxvDopUq9x6UUMmjn"),
		MustPublicKeyFromBase58("9hFtYBYmBJCVguRYs9pBTWKYAFoKfjYR7zBPpEkVsmD"),
		MustPublicKeyFromBase58("6FzXPEhCJoBx7Zw3SN9qhekHemd6E2b8kVguitmVAngW"),
		MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111"),
		MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111"),
		MustPublicKeyFromBase58("11111111111111111111111111111111"),
		MustPublicKeyFromBase58("Vote111111111111111111111111111111111111111"),
	})

	assert.Equal(t, trx.Message.Instructions, []CompiledInstruction{
		{
			ProgramIDIndex: 5,
			Accounts:       []uint16{0, 01},
			Data:           []byte{0xaa, 0xbb},
		},
		{
			ProgramIDIndex: 6,
			Accounts:       []uint16{4, 3, 1, 2},
			Data:           []byte{0xcc, 0xdd},
		},
	})
}

func TestTransactionDecode(t *testing.T) {
	encoded := "AfjEs3XhTc3hrxEvlnMPkm/cocvAUbFNbCl00qKnrFue6J53AhEqIFmcJJlJW3EDP5RmcMz+cNTTcZHW/WJYwAcBAAEDO8hh4VddzfcO5jbCt95jryl6y8ff65UcgukHNLWH+UQGgxCGGpgyfQVQV02EQYqm4QwzUt2qf9f1gVLM7rI4hwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6ANIF55zOZWROWRkeh+lExxZBnKFqbvIxZDLE7EijjoBAgIAAQwCAAAAOTAAAAAAAAA="
	data, err := base64.StdEncoding.DecodeString(encoded)
	require.NoError(t, err)

	tx, err := TransactionFromDecoder(bin.NewBinDecoder(data))
	require.NoError(t, err)
	require.NotNil(t, tx)

	require.Len(t, tx.Signatures, 1)
	require.Equal(t,
		MustSignatureFromBase58("5yUSwqQqeZLEEYKxnG4JC4XhaaBpV3RS4nQbK8bQTyjLX5btVq9A1Ja5nuJzV7Z3Zq8G6EVKFvN4DKUL6PSAxmTk"),
		tx.Signatures[0],
	)

	require.Equal(t,
		[]PublicKey{
			MustPublicKeyFromBase58("52NGrUqh6tSGhr59ajGxsH3VnAaoRdSdTbAaV9G3UW35"),
			MustPublicKeyFromBase58("SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"),
			MustPublicKeyFromBase58("11111111111111111111111111111111"),
		},
		tx.Message.AccountKeys,
	)

	require.Equal(t,
		MessageHeader{
			NumRequiredSignatures:       1,
			NumReadonlySignedAccounts:   0,
			NumReadonlyUnsignedAccounts: 1,
		},
		tx.Message.Header,
	)

	require.Equal(t,
		MustHashFromBase58("GcgVK9buRA7YepZh3zXuS399GJAESCisLnLDBCmR5Aoj"),
		tx.Message.RecentBlockhash,
	)

	decodedData, err := base58.Decode("3Bxs4ART6LMJ13T5")
	require.NoError(t, err)
	require.Equal(t,
		[]CompiledInstruction{
			{
				ProgramIDIndex: 2,
				Accounts: []uint16{
					0,
					1,
				},
				Data: Base58(decodedData),
			},
		},
		tx.Message.Instructions,
	)
}

func TestTransactionVerifySignatures(t *testing.T) {
	type testCase struct {
		Transaction string
	}

	testCases := []testCase{
		{
			Transaction: "AVBFwRrn4wroV9+NVQfgg/GbjFtQFodLnNI5oTpDMQiQ4HfZNyFzcFamHSSFW4p5wc3efeEKvykbmk8jzf2LCQwBAAIGjYddInd/DSl2KJCP18GhEDlaJyPKVrgBGGsr3TF6jSYPgr3AdITNKr2UQVQ5I+Wh5StQv/a5XdLr6VN4Y21My1M/Y1FNK5wQLKJa1LYfN/HAudufFVtc0fRPR6AMUJ9UrkRI7sjY/PnpcXLF7A7SBvJrWu+o8+7QIaD8sL9aXkGFDy1uAqR6+CTQmradxC1wyyjL+iSft+5XudJWwSdi7wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAi+i1vCST+HNO0DEchpEJImMHhZ1BReuf7poRqmXpeA8CBAUBAgMCAgcAAwAAAAEABQIAAAwCAAAA6w0AAAAAAAA=",
		},
		{
			Transaction: "AWwhMTxKhl9yZOlidY0u3gYmy+J/6V3kFSXU7GgK5zwN+SwljR2dOlHgKtUDRX8uee2HtfeyL3t4lB3n749L4QQBAAIEFg+6wTr33dgF0xcKPeDGvZcSah4CwNJZ0Khu+CHW5cehpkZfTC6/JEwx2AvJXCc0WjQk5CjC3vM+ztnpDT9wGwan1RcYx3TJKFZjmGkdXraLXrijm0ttXHNVWyEAAAAA3OXr4eScO58RTLVUTFCpnsDWktY/Vnla4Cmsg9nqi+Jr/+AAgahV8wmBK4mnz9WwJSryq8x2Ic0asytADGhLZAEDAwABAigCAAAABwAAAAEAAAAAAAAAz+dyuQIAAAAIn18BAAAAAPsVKAcAAAAA",
		},
		{
			Transaction: "ARZsk8+AvvT9onUT8FU1VRaiC8Sp+FKveOwhdPoigWHA+MGNcIOqbow6mwSILEYvvyOB/fi3UQ/xKQCjEtxBRgIBAAIFKIX92BRrkgEfrLEXAvXtw7OgPPhHU+62C8DB5QPoMgNSbKXgdub0sr7Yp3Nvdrsp6SDoJ4gdoyRad2AV+Japj0dRtYW4OxE78FvRZTeqHFy2My/m12/afGIPS8iUnMGlBqfVFxjHdMkoVmOYaR1etoteuKObS21cc1VbIQAAAAC/jt8clGtWu0PSX5i4e2vlERcwCmEmGvn5+U7telqAiK4hdAN78GteFjqtJrxLXxpVNKsu1lfdcFPXa/Kcg4e5AQQEAQADAicmMiQQAiGujz0xoTQSQCgAMPOroDk5F0hQ/BgzEkBBvVKWIY41EkA=",
		},
		{
			Transaction: "Ad7TPpYTvSpO//KNA5YTZVojVwz4NlH4gH9ktl+rTObJcgo8QkqmHK4t6DQr9dD58B/A/5/N7v9K+0j6y1TVCAsBAAMFA9maY4S727Z/lOSb08nHehVFsC32kTKMMPjPJp111bKM0Fl1Dg04vV2x9nL2TCqSHmjT8xg6wUAzjZa1+6YCBQan1RcZLwqvxvJl4/t3zHragsUp0L47E24tAFUgAAAABqfVFxjHdMkoVmOYaR1etoteuKObS21cc1VbIQAAAAAHYUgdNXR0u3xNdiTr072z2DVec9EQQ/wNo1OAAAAAAJDQfslK1yQFkGqDXWu6cthRNuYGlajYMOmtoSJB6hmPAQQEAQIDAE0CAAAAAwAAAAAAAAD5FSgHAAAAAPoVKAcAAAAA+xUoBwAAAADECMJOPX7e7fOF5Hrq9xhdch2Uqhg8vQOYyZM/6V983gHQ0gNiAAAAAA==",
		},
		{
			Transaction: "Ak8jvC3ch5hq3lhOHPkACoFepIUON2zEN4KRcw4lDS6GBsQfnSdzNGPETm/yi0hPKk75/i2VXFj0FLUWnGR64ADyUbqnirFjFtaSNgcGi02+Tm7siT4CPpcaTq0jxfYQK/h9FdxXXPnLry74J+RE8yji/BtJ/Cjxbx+TIHigeIYJAgEBBByE1Y6EqCJKsr7iEupU6lsBHtBdtI4SK3yWMCFA0iEKeFPgnGmtp+1SIX1Ak+sN65iBaR7v4Iim5m1OEuFQTgi9N57UnhNpCNuUePaTt7HJaFBmyeZB3deXeKWVudpY3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWVECK/n3a7QR6OKWYR4DuAVjS6FXgZj82W0dJpSIPnEBAwQAAgEDDAIAAABAQg8AAAAAAA==",
		},
	}

	for _, tc := range testCases {
		txBin, err := base64.StdEncoding.DecodeString(tc.Transaction)
		require.NoError(t, err)
		tx, err := TransactionFromDecoder(bin.NewBinDecoder(txBin))
		require.NoError(t, err)
		require.NoError(t, tx.VerifySignatures())
		require.Equal(t, len(tx.Signatures), len(tx.Message.Signers()))
	}
}

func BenchmarkTransactionFromDecoder(b *testing.B) {
	txString := "Ak8jvC3ch5hq3lhOHPkACoFepIUON2zEN4KRcw4lDS6GBsQfnSdzNGPETm/yi0hPKk75/i2VXFj0FLUWnGR64ADyUbqnirFjFtaSNgcGi02+Tm7siT4CPpcaTq0jxfYQK/h9FdxXXPnLry74J+RE8yji/BtJ/Cjxbx+TIHigeIYJAgEBBByE1Y6EqCJKsr7iEupU6lsBHtBdtI4SK3yWMCFA0iEKeFPgnGmtp+1SIX1Ak+sN65iBaR7v4Iim5m1OEuFQTgi9N57UnhNpCNuUePaTt7HJaFBmyeZB3deXeKWVudpY3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWVECK/n3a7QR6OKWYR4DuAVjS6FXgZj82W0dJpSIPnEBAwQAAgEDDAIAAABAQg8AAAAAAA=="
	txBin, err := base64.StdEncoding.DecodeString(txString)
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := TransactionFromDecoder(bin.NewBinDecoder(txBin))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTransactionVerifySignatures(b *testing.B) {
	txString := "Ak8jvC3ch5hq3lhOHPkACoFepIUON2zEN4KRcw4lDS6GBsQfnSdzNGPETm/yi0hPKk75/i2VXFj0FLUWnGR64ADyUbqnirFjFtaSNgcGi02+Tm7siT4CPpcaTq0jxfYQK/h9FdxXXPnLry74J+RE8yji/BtJ/Cjxbx+TIHigeIYJAgEBBByE1Y6EqCJKsr7iEupU6lsBHtBdtI4SK3yWMCFA0iEKeFPgnGmtp+1SIX1Ak+sN65iBaR7v4Iim5m1OEuFQTgi9N57UnhNpCNuUePaTt7HJaFBmyeZB3deXeKWVudpY3gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWVECK/n3a7QR6OKWYR4DuAVjS6FXgZj82W0dJpSIPnEBAwQAAgEDDAIAAABAQg8AAAAAAA=="
	txBin, err := base64.StdEncoding.DecodeString(txString)
	if err != nil {
		b.Error(err)
	}

	tx, err := TransactionFromDecoder(bin.NewBinDecoder(txBin))
	if err != nil {
		b.Error(err)
	}
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tx.VerifySignatures()
	}
}
