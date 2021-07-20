package rpc

import (
	"context"
)

// MinimumLedgerSlot returns the lowest slot that the node
// has information about in its ledger. This value may increase
// over time if the node is configured to purge older ledger data.
func (cl *Client) MinimumLedgerSlot(ctx context.Context) (out uint64, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "minimumLedgerSlot", nil)
	return
}
