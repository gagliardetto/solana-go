package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type BlocksResult []bin.Uint64

// GetBlocks returns a list of confirmed blocks between two slots.
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use getConfirmedBlocks for solana-core v1.6.
// The result will be an array of u64 integers listing confirmed blocks
// between start_slot and either end_slot, if provided, or latest
// confirmed block, inclusive. Max range allowed is 500,000 slots.
func (cl *Client) GetBlocks(
	ctx context.Context,
	startSlot int,
	endSlot *int,
	commitment CommitmentType,
) (out *BlocksResult, err error) {
	params := []interface{}{startSlot}
	if endSlot != nil {
		params = append(params, endSlot)
	}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getBlocks", params)

	return
}
