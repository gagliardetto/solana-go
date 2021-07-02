package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetVersionResult struct {
	SolanaCore string    `json:"solana-core"` // software version of solana-core
	FeatureSet bin.Int64 `json:"feature-set"` // unique identifier of the current software's feature set
}

// GetVersion returns the current solana versions running on the node.
func (cl *Client) GetVersion(ctx context.Context) (out *GetVersionResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getVersion")
	return
}
