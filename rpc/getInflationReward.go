package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetInflationRewardOpts struct {
	Commitment CommitmentType

	// An epoch for which the reward occurs.
	// If omitted, the previous epoch will be used.
	Epoch *uint64
}

// GetInflationReward returns the inflation reward for a list of addresses for an epoch.
func (cl *Client) GetInflationReward(
	ctx context.Context,

	// An array of addresses to query.
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

type GetInflationRewardResult struct {
	// Epoch for which reward occured.
	Epoch bin.Uint64 `json:"epoch"`

	// The slot in which the rewards are effective.
	EffectiveSlot bin.Uint64 `json:"effectiveSlot"`

	// Reward amount in lamports.
	Amount bin.Uint64 `json:"amount"`

	// Post balance of the account in lamports.
	PostBalance bin.Uint64 `json:"postBalance"`

	// Vote account commission when the reward was credited.
	Commission *uint8 `json:"commission,omitempty"`
}
