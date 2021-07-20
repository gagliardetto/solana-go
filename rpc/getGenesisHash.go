package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetGenesisHash returns the genesis hash.
func (cl *Client) GetGenesisHash(ctx context.Context) (out solana.Hash, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getGenesisHash", nil)
	return
}
