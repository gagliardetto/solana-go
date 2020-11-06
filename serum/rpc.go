package serum

import (
	"bytes"
	"context"
	"fmt"

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
	err = struc.Unpack(bytes.NewReader(accountData), *m)
	if err != nil {
		return nil, fmt.Errorf("unable to decode market: %w", err)
	}

	return &MarketMeta{
		Address: marketAddr,
		m:       m,
	}, nil

}
