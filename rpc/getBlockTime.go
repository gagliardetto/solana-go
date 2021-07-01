package rpc

import (
	"context"
)

// Estimated production time, as Unix timestamp (seconds since the Unix epoch)
type GetBlockTimeResult int64

// GetBlockTime returns the estimated production time of a block.
// Each validator reports their UTC time to the ledger on a regular
// interval by intermittently adding a timestamp to a Vote for a
// particular block. A requested block's time is calculated from
// the stake-weighted mean of the Vote timestamps in a set of
// recent blocks recorded on the ledger.
//
// The result will be an int64 estimated production time,
// as Unix timestamp (seconds since the Unix epoch),
// or nil if the timestamp is not available for this block.
func (cl *Client) GetBlockTime(
	ctx context.Context,
	block uint64, // block, identified by Slot
) (out *GetBlockTimeResult, err error) {
	params := []interface{}{block}
	err = cl.rpcClient.CallFor(&out, "getBlockTime", params...)
	return
}
