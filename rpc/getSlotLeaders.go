package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetSlotLeaders returns the slot leaders for a given slot range.
func (cl *Client) GetSlotLeaders(
	ctx context.Context,
	start uint64,
	limit uint64,
) (out []solana.PublicKey, err error) {
	params := []interface{}{start, limit}
	err = cl.rpcClient.CallForInto(ctx, &out, "getSlotLeaders", params)
	return
}
