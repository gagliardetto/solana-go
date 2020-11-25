package solana

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {

	a := NewAccount()
	privateKey := a.PrivateKey
	public := a.PublicKey()

	a2, err := AccountFromPrivateKeyBase58(privateKey.String())
	require.NoError(t, err)

	require.Equal(t, privateKey, a2.PrivateKey)
	require.Equal(t, public, a2.PublicKey())
}
