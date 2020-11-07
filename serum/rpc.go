package serum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dfuse-io/solana-go/token"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/lunixbochs/struc"
)

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

	fmt.Println("market addr", marketAddr)
	cnt, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(cnt))

	baseMint, err := token.GetMint(ctx, rpcCli, m.BaseMint)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base token: %w", err)
	}

	cnt, _ = json.MarshalIndent(baseMint, "", "  ")
	fmt.Println(string(cnt))

	quoteMint, err := token.GetMint(ctx, rpcCli, m.QuoteMint)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base token: %w", err)
	}

	cnt, _ = json.MarshalIndent(quoteMint, "", "  ")
	fmt.Println(string(cnt))

	return &MarketMeta{
		Address:   marketAddr,
		MarketV2:  &m,
		QuoteMint: quoteMint,
		BaseMint:  baseMint,
	}, nil

}
