package rpc

import (
	"context"
)

// GetEpochInfo returns information about the current epoch.
func (cl *Client) GetEpochInfo(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out *GetEpochInfoResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getEpochInfo", params)
	return
}

type GetEpochInfoResult struct {
	// The current slot.
	AbsoluteSlot uint64 `json:"absoluteSlot"`

	// The current block height.
	BlockHeight uint64 `json:"blockHeight"`

	// The current epoch.
	Epoch uint64 `json:"epoch"`

	// The current slot relative to the start of the current epoch.
	SlotIndex uint64 `json:"slotIndex"`

	// The number of slots in this epoch.
	SlotsInEpoch uint64 `json:"slotsInEpoch"`

	TransactionCount uint64 `json:"transactionCount"`
}
