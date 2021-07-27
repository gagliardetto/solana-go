package rpc

import (
	"context"
)

// GetBlockCommitment returns commitment for particular block.
func (cl *Client) GetBlockCommitment(
	ctx context.Context,
	block uint64, // block, identified by Slot
) (out *GetBlockCommitmentResult, err error) {
	params := []interface{}{block}
	err = cl.rpcClient.CallForInto(ctx, &out, "getBlockCommitment", params)
	return
}

type GetBlockCommitmentResult struct {
	// nil if Unknown block, or array of u64 integers
	// logging the amount of cluster stake in lamports
	// that has voted on the block at each depth from 0 to `MAX_LOCKOUT_HISTORY` + 1
	Commitment []uint64 `json:"commitment"`

	// Total active stake, in lamports, of the current epoch.
	TotalStake uint64 `json:"totalStake"`
}
