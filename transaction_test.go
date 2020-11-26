package solana

import (
	"testing"

	"github.com/magiconair/properties/assert"
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

func TestTransactionWithInstructions(t *testing.T) {
	instructions := []TransactionInstruction{
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

	trx, err := TransactionWithInstructions(instructions, nil)
	require.NoError(t, err)

	assert.Equal(t, trx.Message.Header, MessageHeader{
		NumRequiredSignatures:       3,
		NumReadonlySignedAccounts:   1,
		NumReadonlyUnsignedAccounts: 3,
	})

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
			Accounts:       []uint8{0, 01},
			DataLength:     2,
			Data:           []byte{0xaa, 0xbb},
		},
		{
			ProgramIDIndex: 6,
			AccountCount:   4,
			Accounts:       []uint8{4, 3, 1, 2},
			DataLength:     2,
			Data:           []byte{0xcc, 0xdd},
		},
	})
}
