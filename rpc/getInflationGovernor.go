package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
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
	err = cl.rpcClient.CallFor(&out, "getInflationGovernor", params)
	return
}

type GetInflationGovernorResult struct {
	// The initial inflation percentage from time 0.
	Initial bin.JSONFloat64 `json:"initial"`

	// Terminal inflation percentage.
	Terminal bin.JSONFloat64 `json:"terminal"`

	// Rate per year at which inflation is lowered. Rate reduction is derived using the target slot time in genesis config.
	Taper bin.JSONFloat64 `json:"taper"`

	// Percentage of total inflation allocated to the foundation.
	Foundation bin.JSONFloat64 `json:"foundation"`

	// Duration of foundation pool inflation in years.
	FoundationTerm bin.JSONFloat64 `json:"foundationTerm"`
}
