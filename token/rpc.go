package token

import (
	"bytes"
	"context"
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/lunixbochs/struc"
)

func GetMint(ctx context.Context, cli *rpc.Client, mintPubKey solana.PublicKey) (*Mint, error) {
	accInfo, err := cli.GetAccountInfo(ctx, mintPubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base mint: %w", err)
	}

	accountData, err := accInfo.Value.DataToBytes()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve account data byte: %w", err)
	}

	var m Mint
	err = struc.Unpack(bytes.NewReader(accountData), &m)
	if err != nil {
		return nil, fmt.Errorf("unable to decode market: %w", err)
	}
	return &m, nil
}
