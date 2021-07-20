package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
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
	AbsoluteSlot bin.Uint64 `json:"absoluteSlot"`

	// The current block height.
	BlockHeight bin.Uint64 `json:"blockHeight"`

	// The current epoch.
	Epoch bin.Uint64 `json:"epoch"`

	// The current slot relative to the start of the current epoch.
	SlotIndex bin.Uint64 `json:"slotIndex"`

	// The number of slots in this epoch.
	SlotsInEpoch bin.Uint64 `json:"slotsInEpoch"`

	TransactionCount bin.Uint64 `json:"transactionCount"`
}
