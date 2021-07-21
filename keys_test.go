package solana

import (
	"encoding/hex"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublicKeyFromBytes(t *testing.T) {
	tests := []struct {
		name     string
		inHex    string
		expected PublicKey
	}{
		{
			"empty",
			"",
			MustPublicKeyFromBase58("11111111111111111111111111111111"),
		},
		{
			"smaller than required",
			"010203040506",
			MustPublicKeyFromBase58("4wBqpZM9k69W87zdYXT2bRtLViWqTiJV3i2Kn9q7S6j"),
		},
		{
			"equal to 32 bytes",
			"0102030405060102030405060102030405060102030405060102030405060101",
			MustPublicKeyFromBase58("4wBqpZM9msxygzsdeLPq6Zw3LoiAxJk3GjtKPpqkcsi"),
		},
		{
			"longer than required",
			"0102030405060102030405060102030405060102030405060102030405060101FFFFFFFFFF",
			MustPublicKeyFromBase58("4wBqpZM9msxygzsdeLPq6Zw3LoiAxJk3GjtKPpqkcsi"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bytes, err := hex.DecodeString(test.inHex)
			require.NoError(t, err)

			actual := PublicKeyFromBytes(bytes)
			assert.Equal(t, test.expected, actual, "%s != %s", test.expected, actual)
		})
	}
}

func TestPublicKeyFromBase58(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		expected    PublicKey
		expectedErr error
	}{
		{
			"hand crafted",
			"SerumkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
			MustPublicKeyFromBase58("SerumkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
			nil,
		},
		{
			"hand crafted error",
			"SerkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
			zeroPublicKey,
			errors.New("invalid length, expected 32, got 30"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := PublicKeyFromBase58(test.in)
			if test.expectedErr == nil {
				require.NoError(t, err)
				assert.Equal(t, test.expected, actual)
			} else {
				assert.Equal(t, test.expectedErr, err)
			}
		})
	}
}

func TestPrivateKeyFromSolanaKeygenFile(t *testing.T) {
	tests := []struct {
		inFile      string
		expected    PrivateKey
		expectedPub PublicKey
		expectedErr error
	}{
		{
			"testdata/standard.solana-keygen.json",
			MustPrivateKeyFromBase58("66cDvko73yAf8LYvFMM3r8vF5vJtkk7JKMgEKwkmBC86oHdq41C7i1a2vS3zE1yCcdLLk6VUatUb32ZzVjSBXtRs"),
			MustPublicKeyFromBase58("F8UvVsKnzWyp2nF8aDcqvQ2GVcRpqT91WDsAtvBKCMt9"),
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.inFile, func(t *testing.T) {
			actual, err := PrivateKeyFromSolanaKeygenFile(test.inFile)
			if test.expectedErr == nil {
				require.NoError(t, err)
				assert.Equal(t, test.expected, actual)
				assert.Equal(t, test.expectedPub, actual.PublicKey(), "%s != %s", test.expectedPub, actual.PublicKey())

			} else {
				assert.Equal(t, test.expectedErr, err)
			}
		})
	}
}

func TestPublicKey_MarshalText(t *testing.T) {
	keyString := "4wBqpZM9k69W87zdYXT2bRtLViWqTiJV3i2Kn9q7S6j"
	keyParsed := MustPublicKeyFromBase58(keyString)

	var key PublicKey
	err := key.UnmarshalText([]byte(keyString))
	require.NoError(t, err)

	assert.True(t, keyParsed.Equals(key))

	keyText, err := key.MarshalText()
	require.NoError(t, err)
	assert.Equal(t, []byte(keyString), keyText)

	type IdentityToSlotsBlocks map[PublicKey][2]int64

	var payload IdentityToSlotsBlocks
	data := `{"` + keyString + `":[3,4]}`
	err = json.Unmarshal([]byte(data), &payload)
	require.NoError(t, err)

	assert.Equal(t,
		IdentityToSlotsBlocks{
			keyParsed: [2]int64{3, 4},
		},
		payload,
	)
}
