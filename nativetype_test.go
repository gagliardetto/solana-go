package solana

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMustPublicKeyFromBase58(t *testing.T) {
	require.Panics(t, func() {
		MustPublicKeyFromBase58("toto")
	})
}
