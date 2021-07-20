package rpc

import (
	"context"
)

// GetFirstAvailableBlock returns the slot of the lowest confirmed block that has not been purged from the ledger.
func (cl *Client) GetFirstAvailableBlock(ctx context.Context) (out uint64, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getFirstAvailableBlock", nil)
	return
}
