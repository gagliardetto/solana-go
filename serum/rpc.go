package serum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/token"
	"github.com/lunixbochs/struc"
)

//go:generate rice embed-go

// TODO: hit the chain and
func KnownMarket() ([]*MarketMeta, error) {
	box := rice.MustFindBox("data").MustBytes("markets.json")
	if box == nil {
		return nil, fmt.Errorf("unable to retrieve known markets")
	}

	dec := json.NewDecoder(bytes.NewReader(box))
	var markets []*MarketMeta
	err := dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}

func FetchMarket(ctx context.Context, rpcCli *rpc.Client, marketAddr solana.PublicKey) (*MarketMeta, error) {
	accInfo, err := rpcCli.GetAccountInfo(ctx, marketAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get market account:%w", err)
	}

	accountData, err := accInfo.Value.DataToBytes()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve account data byte: %w", err)
	}

	var m MarketV2
	err = struc.Unpack(bytes.NewReader(accountData), &m)
	if err != nil {
		return nil, fmt.Errorf("unable to decode market: %w", err)
	}

	baseMint, err := token.GetMint(ctx, rpcCli, m.BaseMint)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base token: %w", err)
	}

	quoteMint, err := token.GetMint(ctx, rpcCli, m.QuoteMint)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base token: %w", err)
	}

	return &MarketMeta{
		Address:   marketAddr,
		MarketV2:  &m,
		QuoteMint: quoteMint,
		BaseMint:  baseMint,
	}, nil

}
