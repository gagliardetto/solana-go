package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetInflationRate returns the specific inflation values for the current epoch.
func (cl *Client) GetInflationRate(ctx context.Context) (out *GetInflationRateResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getInflationRate")
	return
}

type GetInflationRateResult struct {
	// Total inflation.
	Total bin.JSONFloat64 `json:"total"`

	// Inflation allocated to validators.
	Validator bin.JSONFloat64 `json:"validator"`

	// Inflation allocated to the foundation.
	Foundation bin.JSONFloat64 `json:"foundation"`

	// Epoch for which these values are valid.
	Epoch bin.JSONFloat64 `json:"epoch"`
}
