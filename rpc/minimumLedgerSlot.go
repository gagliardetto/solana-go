package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// MinimumLedgerSlot returns the lowest slot that the node
// has information about in its ledger. This value may increase
// over time if the node is configured to purge older ledger data.
func (cl *Client) MinimumLedgerSlot(ctx context.Context) (out bin.Uint64, err error) {
	err = cl.rpcClient.CallFor(&out, "minimumLedgerSlot")
	return
}
