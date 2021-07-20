package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

// GetStakeActivation returns epoch activation information for a stake account.
func (cl *Client) GetStakeActivation(
	ctx context.Context,
	// Pubkey of stake account to query
	account solana.PublicKey,

	commitment CommitmentType,

	// epoch for which to calculate activation details.
	// If parameter not provided, defaults to current epoch.
	epoch *uint64,
) (out *GetStakeActivationResult, err error) {
	params := []interface{}{account}
	{
		obj := M{}
		if commitment != "" {
			obj["commitment"] = commitment
		}
		if epoch != nil {
			obj["epoch"] = epoch
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getStakeActivation", params)
	return
}

type GetStakeActivationResult struct {
	// The stake account's activation state, one of: active, inactive, activating, deactivating.
	State ActivationStateType `json:"state"`

	// Stake active during the epoch.
	Active bin.Uint64 `json:"active"`

	// Stake inactive during the epoch.
	Inactive bin.Uint64 `json:"inactive"`
}

type ActivationStateType string

const (
	ActivationStateActive       ActivationStateType = "active"
	ActivationStateInactive     ActivationStateType = "inactive"
	ActivationStateActivating   ActivationStateType = "activating"
	ActivationStateDeactivating ActivationStateType = "deactivating"
)
