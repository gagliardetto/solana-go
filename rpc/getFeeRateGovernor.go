package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetFeeRateGovernor returns the fee rate governor information from the root bank.
func (cl *Client) GetFeeRateGovernor(ctx context.Context) (out *GetFeeRateGovernorResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getFeeRateGovernor")
	return
}

type GetFeeRateGovernorResult struct {
	RPCContext
	Value FeeRateGovernorResult `json:"value"`
}
type FeeRateGovernorResult struct {
	FeeRateGovernor FeeRateGovernor `json:"feeRateGovernor"`
}
type FeeRateGovernor struct {
	// Percentage of fees collected to be destroyed.
	BurnPercent uint8 `json:"burnPercent"`

	// Largest value lamportsPerSignature can attain for the next slot.
	MaxLamportsPerSignature bin.Uint64 `json:"maxLamportsPerSignature"`

	// Smallest value lamportsPerSignature can attain for the next slot.
	MinLamportsPerSignature bin.Uint64 `json:"minLamportsPerSignature"`

	// Desired fee rate for the cluster.
	TargetLamportsPerSignature bin.Uint64 `json:"targetLamportsPerSignature"`

	// Desired signature rate for the cluster.
	TargetSignaturesPerSlot bin.Uint64 `json:"targetSignaturesPerSlot"`
}
