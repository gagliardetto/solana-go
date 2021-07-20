package rpc

import (
	"context"
)

// GetSnapshotSlot returns the highest slot that the node has a snapshot for.
func (cl *Client) GetSnapshotSlot(ctx context.Context) (out uint64, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getSnapshotSlot", nil)
	return
}
