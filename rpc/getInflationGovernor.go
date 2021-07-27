package rpc

import (
	"context"
)

// GetInflationGovernor returns the current inflation governor.
func (cl *Client) GetInflationGovernor(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out *GetInflationGovernorResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getInflationGovernor", params)
	return
}

type GetInflationGovernorResult struct {
	// The initial inflation percentage from time 0.
	Initial float64 `json:"initial"`

	// Terminal inflation percentage.
	Terminal float64 `json:"terminal"`

	// Rate per year at which inflation is lowered. Rate reduction is derived using the target slot time in genesis config.
	Taper float64 `json:"taper"`

	// Percentage of total inflation allocated to the foundation.
	Foundation float64 `json:"foundation"`

	// Duration of foundation pool inflation in years.
	FoundationTerm float64 `json:"foundationTerm"`
}
