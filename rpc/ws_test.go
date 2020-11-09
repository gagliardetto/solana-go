package rpc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWSClient_ProgramSubscribe(t *testing.T) {

	c := NewWSClient("ws://api.mainnet-beta.solana.com:80/rpc")
	err := c.ProgramSubscribe("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o", "")
	require.NoError(t, err)
	time.Sleep(20 * time.Second)
}
