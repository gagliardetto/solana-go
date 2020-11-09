package rpc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWSClient_ProgramSubscribe(t *testing.T) {

	done := make(chan bool)
	c, err := NewWSClient("ws://api.mainnet-beta.solana.com:80/rpc")
	require.NoError(t, err)

	var closed bool
	err = c.ProgramSubscribe("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o", "", func(programResult result) {
		if closed {
			return
		}
		r := programResult.(*ProgramResult)
		fmt.Println("result", r)
		closed = true
		close(done)
	})
	require.NoError(t, err)

	select {
	case <-done:
	case <-time.After(2000 * time.Millisecond):
		t.Error("failed to run the giving time")
	}
}
