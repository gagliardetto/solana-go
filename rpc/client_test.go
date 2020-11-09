package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dfuse-io/solana-go"
	"github.com/stretchr/testify/require"
)

func TestClient_GetAccountInfo(t *testing.T) {
	c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	accInfo, err := c.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)
	d, err := json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))
	pubKey = solana.MustPublicKeyFromBase58("EXnGBBSamqzd3uxEdRLUiYzjJkTwQyorAaFXdfteuGXe")
	accInfo, err = c.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)
	d, err = json.MarshalIndent(accInfo, "", " ")
	require.NoError(t, err)
	fmt.Println(string(d))

}
