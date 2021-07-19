package rpc

import (
	"context"
)

// GetBlockHeight returns the current block height of the node.
func (cl *Client) GetBlockHeight(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out uint64, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallFor(&out, "getBlockHeight", params)
	return
}
