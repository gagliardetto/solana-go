package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetEpochInfoResult struct {
	AbsoluteSlot     bin.Uint64 `json:"absoluteSlot"` // the current slot
	BlockHeight      bin.Uint64 `json:"blockHeight"`  // the current block height
	Epoch            bin.Uint64 `json:"epoch"`        // the current epoch
	SlotIndex        bin.Uint64 `json:"slotIndex"`    // the current slot relative to the start of the current epoch
	SlotsInEpoch     bin.Uint64 `json:"slotsInEpoch"` // the number of slots in this epoch
	TransactionCount bin.Uint64 `json:"transactionCount"`
}

// GetEpochInfo returns information about the current epoch.
func (cl *Client) GetEpochInfo(
	ctx context.Context,
	commitment CommitmentType,
) (out *GetEpochInfoResult, err error) {
	var params []interface{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	// TODO: why `params)` and not `params...)` ??? (getting `-32602:`params` should be an array`)
	err = cl.rpcClient.CallFor(&out, "getEpochInfo", params)
	return
}
