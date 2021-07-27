package rpc

import (
	"context"
)

// GetInflationRate returns the specific inflation values for the current epoch.
func (cl *Client) GetInflationRate(ctx context.Context) (out *GetInflationRateResult, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getInflationRate", nil)
	return
}

type GetInflationRateResult struct {
	// Total inflation.
	Total float64 `json:"total"`

	// Inflation allocated to validators.
	Validator float64 `json:"validator"`

	// Inflation allocated to the foundation.
	Foundation float64 `json:"foundation"`

	// Epoch for which these values are valid.
	Epoch float64 `json:"epoch"`
}
