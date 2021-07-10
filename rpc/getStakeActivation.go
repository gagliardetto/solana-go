package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetStakeActivationResult struct {
	State    ActivationStateType `json:"state"`    // the stake account's activation state, one of: active, inactive, activating, deactivating
	Active   bin.Uint64          `json:"active"`   // stake active during the epoch
	Inactive bin.Uint64          `json:"inactive"` // stake inactive during the epoch
}

type ActivationStateType string

const (
	ActivationStateActive       ActivationStateType = "active"
	ActivationStateInactive     ActivationStateType = "inactive"
	ActivationStateActivating   ActivationStateType = "activating"
	ActivationStateDeactivating ActivationStateType = "deactivating"
)

// GetStakeActivation returns epoch activation information for a stake account.
func (cl *Client) GetStakeActivation(
	ctx context.Context,
	account solana.PublicKey, // Pubkey of stake account to query
	commitment CommitmentType,
	epoch *uint64, // epoch for which to calculate activation details. If parameter not provided, defaults to current epoch.
) (out *GetStakeActivationResult, err error) {
	params := []interface{}{account}
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
	err = cl.rpcClient.CallFor(&out, "getStakeActivation", params)
	return
}
