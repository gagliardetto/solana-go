package serum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
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
	acctInfo, err := rpcCli.GetAccountInfo(ctx, marketAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get market account:%w", err)
	}

	meta := &MarketMeta{
		Address: marketAddr,
	}

	dataLen := len(acctInfo.Value.Data)
	switch dataLen {
	// case 380:
	// 	// if err := meta.MarketV1.Decode(acctInfo.Value.Data); err != nil {
	// 	// 	return nil, fmt.Errorf("decoding market v1: %w", err)
	// 	// }
	// 	return nil, fmt.Errorf("Unsupported market version, w/ data length of 380")

	case 388:
		if err := meta.MarketV2.Decode(acctInfo.Value.Data); err != nil {
			return nil, fmt.Errorf("decoding market v2: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported market data length: %d", dataLen)
	}

	if err := rpcCli.GetAccountDataIn(ctx, meta.MarketV2.QuoteMint, &meta.QuoteMint); err != nil {
		return nil, fmt.Errorf("getting quote mint: %w", err)
	}

	if err := rpcCli.GetAccountDataIn(ctx, meta.MarketV2.BaseMint, &meta.BaseMint); err != nil {
		return nil, fmt.Errorf("getting base token: %w", err)
	}

	return meta, nil
}
