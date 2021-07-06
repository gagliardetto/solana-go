package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetInflationRewardResult struct {
	Epoch         bin.Uint64 `json:"epoch"`         // epoch for which reward occured
	EffectiveSlot bin.Uint64 `json:"effectiveSlot"` // the slot in which the rewards are effective
	Amount        bin.Uint64 `json:"amount"`        // reward amount in lamports
	PostBalance   bin.Uint64 `json:"postBalance"`   // post balance of the account in lamports
}

type GetInflationRewardOpts struct {
	Commitment CommitmentType
	Epoch      *bin.Uint64
}

// GetInflationReward returns the inflation reward for a list of addresses for an epoch.
func (cl *Client) GetInflationReward(
	ctx context.Context,
	addresses []solana.PublicKey,
	opts *GetInflationRewardOpts,
) (out []*GetInflationRewardResult, err error) {
	params := []interface{}{addresses}
	if opts != nil {
		obj := M{}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if opts.Epoch != nil {
			obj["epoch"] = opts.Epoch
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	// TODO: check
	err = cl.rpcClient.CallFor(&out, "getInflationReward", params)
	return
}
