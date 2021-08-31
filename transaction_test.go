package solana

import (
	"encoding/base64"
	"testing"

	bin "github.com/dfuse-io/binary"
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
			AccountCount:   2,
			Accounts:       []uint16{0, 01},
			DataLength:     2,
			Data:           []byte{0xaa, 0xbb},
		},
		{
			ProgramIDIndex: 6,
			AccountCount:   4,
			Accounts:       []uint16{4, 3, 1, 2},
			DataLength:     2,
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
				AccountCount:   2,
				DataLength:     12,
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
