package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetVersion returns the current solana versions running on the node.
func (cl *Client) GetVersion(ctx context.Context) (out *GetVersionResult, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getVersion", nil)
	return
}

type GetVersionResult struct {
	// Software version of solana-core.
	SolanaCore string `json:"solana-core"`

	// Unique identifier of the current software's feature set.
	FeatureSet bin.Int64 `json:"feature-set"`
}
