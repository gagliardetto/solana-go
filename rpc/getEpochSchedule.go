package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetEpochScheduleResult struct {
	SlotsPerEpoch            bin.Uint64 `json:"slotsPerEpoch"`            // the maximum number of slots in each epoch
	LeaderScheduleSlotOffset bin.Uint64 `json:"leaderScheduleSlotOffset"` // the number of slots before beginning of an epoch to calculate a leader schedule for that epoch
	Warmup                   bool       `json:"warmup"`                   // whether epochs start short and grow
	FirstNormalEpoch         bin.Uint64 `json:"firstNormalEpoch"`         // first normal-length epoch, log2(slotsPerEpoch) - log2(MINIMUM_SLOTS_PER_EPOCH)
	FirstNormalSlot          bin.Uint64 `json:"firstNormalSlot"`          // MINIMUM_SLOTS_PER_EPOCH * (2.pow(firstNormalEpoch) - 1)
}

// GetEpochSchedule returns epoch schedule information from this cluster's genesis config.
func (cl *Client) GetEpochSchedule(ctx context.Context) (out *GetEpochScheduleResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getEpochSchedule")
	return
}
