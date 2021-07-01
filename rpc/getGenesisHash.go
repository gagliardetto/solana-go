package rpc

import (
	"context"

	"github.com/dfuse-io/solana-go"
)

type GetGenesisHashResult struct{}

// GetGenesisHash returns the genesis hash.
func (cl *Client) GetGenesisHash(ctx context.Context) (out solana.Hash, err error) {
	err = cl.rpcClient.CallFor(&out, "getGenesisHash")
	return
}
