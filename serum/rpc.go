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

type SerumClient struct {
	rpc.Client
}

func NewSerumClient(endpoint string) *SerumClient {
	endpoint = "http://api.mainnet-beta.solana.com:80/rpc"
	return &SerumClient{
		*rpc.NewClient(endpoint),
	}
}

func (s *SerumClient) FetchMarket(ctx context.Context, marketAddr solana.PublicKey) (*MarketMeta, error) {
	accInfo, err := s.GetAccountInfo(ctx, marketAddr)
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

	baseMint, err := s.getToken(ctx, m.BaseMint)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base token: %w", err)
	}

	cnt, _ = json.MarshalIndent(baseMint, "", "  ")
	fmt.Println(string(cnt))

	quoteMint, err := s.getToken(ctx, m.QuoteMint)
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

func (s *SerumClient) getToken(ctx context.Context, mintPubKey solana.PublicKey) (*token.Mint, error) {
	accInfo, err := s.GetAccountInfo(ctx, mintPubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base mint: %w", err)
	}

	fmt.Println(accInfo, mintPubKey)

	accountData, err := accInfo.Value.DataToBytes()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve account data byte: %w", err)
	}

	var m token.Mint
	err = struc.Unpack(bytes.NewReader(accountData), &m)
	if err != nil {
		return nil, fmt.Errorf("unable to decode market: %w", err)
	}
	return &m, nil
}
