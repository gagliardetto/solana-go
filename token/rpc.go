package token

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

func KnownMints(network string) ([]*MintMeta, error) {
	box := rice.MustFindBox("mints-data").MustBytes(network + "-tokens.json")
	if box == nil {
		return nil, fmt.Errorf("unable to retrieve known markets")
	}

	dec := json.NewDecoder(bytes.NewReader(box))
	var markets []*MintMeta
	err := dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}

func GetMint(ctx context.Context, cli *rpc.Client, mintPubKey solana.PublicKey) (*Mint, error) {
	acctInfo, err := cli.GetAccountInfo(ctx, mintPubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base mint: %w", err)
	}

	return DecodeMint(acctInfo.Value.Data)
}
