package rpc

import (
	"context"
)

// GetBlocksWithLimit returns a list of confirmed blocks starting at the given slot.
// The result field will be an array of u64 integers listing
// confirmed blocks starting at startSlot for up to limit blocks, inclusive.
//
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use getConfirmedBlocksWithLimit for solana-core v1.6
func (cl *Client) GetBlocksWithLimit(
	ctx context.Context,
	startSlot uint64,
	limit uint64,
	commitment CommitmentType, // optional; "processed" is not supported. If parameter not provided, the default is "finalized".
) (out *BlocksResult, err error) {
	params := []interface{}{startSlot, limit}
	if commitment != "" {
		params = append(params,
			// TODO: provide commitment as string instead of object?
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getBlocksWithLimit", params)
	return
}
