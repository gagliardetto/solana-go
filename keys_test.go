package solana

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPublicKeyFromBase58(t *testing.T) {
	emptyKey := PublicKey{}

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
			emptyKey,
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
