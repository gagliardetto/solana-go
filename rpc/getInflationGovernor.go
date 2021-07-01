package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetInflationGovernorResult struct {
	Initial        bin.JSONFloat64 `json:"initial"`        // the initial inflation percentage from time 0
	Terminal       bin.JSONFloat64 `json:"terminal"`       // terminal inflation percentage
	Taper          bin.JSONFloat64 `json:"taper"`          // rate per year at which inflation is lowered. Rate reduction is derived using the target slot time in genesis config
	Foundation     bin.JSONFloat64 `json:"foundation"`     // percentage of total inflation allocated to the foundation
	FoundationTerm bin.JSONFloat64 `json:"foundationTerm"` // duration of foundation pool inflation in years
}

// GetInflationGovernor returns the current inflation governor.
func (cl *Client) GetInflationGovernor(
	ctx context.Context,
	commitment CommitmentType,
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
