package rpc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dfuse-io/solana-go"

	"github.com/stretchr/testify/require"
)

func TestClient_GetConfirmedBlock(t *testing.T) {

	//rpcClient := NewRPCClient("api.mainnet-beta.solana.com:443")
	c := NewClient("http://api.mainnet-beta.solana.com/rpc")
	err := c.GetConfirmedBlock(46243868)
	require.NoError(t, err)
}

func TestClient_GetAccountInfo(t *testing.T) {

	//rpcClient := NewRPCClient("api.mainnet-beta.solana.com:443")
	c := NewClient("http://api.mainnet-beta.solana.com/rpc")
	accInfo, err := c.GetAccountInfo(solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932"))
	require.NoError(t, err)
	d, err := json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))
}
