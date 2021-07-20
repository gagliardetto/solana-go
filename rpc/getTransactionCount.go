package rpc

import (
	"context"
)

// GetTransactionCount returns the current Transaction count from the ledger.
func (cl *Client) GetTransactionCount(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out uint64, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getTransactionCount", params)
	return
}
