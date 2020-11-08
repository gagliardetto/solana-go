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

//offset: 0,
//bytes: 0x736572756d2100000000000000
func TestClient_ProgramAccount(t *testing.T) {
	//c := NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	//programPubKey := solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")
	//accounts, err := c.GetProgramAccounts(context.Background(), programPubKey, &GetProgramAccountsOpts{
	//	Encoding: "base64",
	//	Filters:  []RPCFilter{
	//		&RPCFilter{
	//			Memcmp:   &RPCFilterMemcmp{
	//				Offset: 0,
	//				Bytes:  nil,
	//			},
	//		}
	//	},
	//})
	//require.NoError(t, err)
	//d, err := json.MarshalIndent(accInfo, "", " ")
	//require.NoError(t, err)
	//fmt.Println(string(d))
}
