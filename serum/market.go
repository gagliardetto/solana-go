package serum

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/token"
	"github.com/dfuse-io/solana-go/rpc"
)

type MarketMeta struct {
	Address    solana.PublicKey `json:"address"`
	Name       string           `json:"name"`
	Deprecated bool             `json:"deprecated"`
	QuoteMint  *token.Mint
	BaseMint   *token.Mint

	MarketV2 *MarketV2
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
