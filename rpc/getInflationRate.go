package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetInflationRateResult struct {
	Total      bin.JSONFloat64 `json:"total"`      // total inflation
	Validator  bin.JSONFloat64 `json:"validator"`  // inflation allocated to validators
	Foundation bin.JSONFloat64 `json:"foundation"` // inflation allocated to the foundation
	Epoch      bin.JSONFloat64 `json:"epoch"`      // epoch for which these values are valid
}

// GetInflationRate returns the specific inflation values for the current epoch.
func (cl *Client) GetInflationRate(ctx context.Context) (out *GetInflationRateResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getInflationRate")
	return
}
