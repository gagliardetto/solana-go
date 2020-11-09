package rpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWSClient_ProgramSubscribe(t *testing.T) {

	c, err := NewWSClient("ws://api.mainnet-beta.solana.com:80/rpc")
	require.NoError(t, err)

	stream, err := c.ProgramSubscribe("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o", "")
	require.NoError(t, err)

	select {
	case <-stream:
	case <-time.After(2000 * time.Millisecond):
		t.Error("failed to run the giving time")
	}
}
