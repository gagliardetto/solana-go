package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetBlocks returns a list of confirmed blocks between two slots.
// The result will be an array of u64 integers listing confirmed blocks
// between start_slot and either end_slot, if provided, or latest
// confirmed block, inclusive. Max range allowed is 500,000 slots.
//
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use `getConfirmedBlocks` for solana-core v1.6.
func (cl *Client) GetBlocks(
	ctx context.Context,
	startSlot uint64,
	endSlot *uint64, // optional
	commitment CommitmentType, // optional
) (out *BlocksResult, err error) {
	params := []interface{}{startSlot}
	if endSlot != nil {
		params = append(params, endSlot)
	}
	if commitment != "" {
		params = append(params,
			// TODO: provide commitment as string instead of object?
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getBlocks", params)

	return
}

type BlocksResult []bin.Uint64
