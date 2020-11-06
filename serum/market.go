package serum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/lunixbochs/struc"
	"github.com/stretchr/testify/require"

	"github.com/dfuse-io/solana-go"
)

type MarketMeta struct {
	Address               solana.PublicKey `json:"address"`
	Name                  string           `json:"name"`
	Deprecated            bool             `json:"deprecated"`
	ProgramID             solana.PublicKey `json:"programId"`
	BaseSPLTokenDecimals  uint64           `json:"base_spl_token_decimals"`
	QuoteSPLTokenDecimals uint64           `json:"quote_spl_token_decimals"`

	m *MarketV2
}

func KnownMarket() ([]*MarketMeta, error) {
	f, err := os.Open("serum/markets.json")
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve known markets: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var markets []*MarketMeta
	err = dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}

func FetchMarket() {

}

func getMarket(t *testing.T, marketAddr solana.PublicKey) *MarketV2 {
	c := rpc.NewClient("http://api.mainnet-beta.solana.com:80/rpc")
	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	accInfo, err := c.GetAccountInfo(context.Background(), pubKey)
	require.NoError(t, err)

	accountData, err := accInfo.Value.DataToBytes()
	require.NoError(t, err)

	var m MarketV2
	require.NoError(t, struc.Unpack(bytes.NewReader(accountData), &m))
	return &m
}
