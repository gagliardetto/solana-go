package rpc

import (
	"context"
)

// GetMaxRetransmitSlot returns the max slot seen from retransmit stage.
func (cl *Client) GetMaxRetransmitSlot(ctx context.Context) (out uint64, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getMaxRetransmitSlot", nil)
	return
}
