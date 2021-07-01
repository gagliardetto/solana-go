package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetFeeRateGovernorResult struct {
	RPCContext
	Value FeeRateGovernorResult `json:"value"`
}
type FeeRateGovernorResult struct {
	FeeRateGovernor FeeRateGovernor `json:"feeRateGovernor"`
}
type FeeRateGovernor struct {
	BurnPercent                uint8      `json:"burnPercent"`                // Percentage of fees collected to be destroyed
	MaxLamportsPerSignature    bin.Uint64 `json:"maxLamportsPerSignature"`    // Largest value lamportsPerSignature can attain for the next slot
	MinLamportsPerSignature    bin.Uint64 `json:"minLamportsPerSignature"`    // Smallest value lamportsPerSignature can attain for the next slot
	TargetLamportsPerSignature bin.Uint64 `json:"targetLamportsPerSignature"` // Desired fee rate for the cluster
	TargetSignaturesPerSlot    bin.Uint64 `json:"targetSignaturesPerSlot"`    // Desired signature rate for the cluster
}

// GetFeeRateGovernor returns the fee rate governor information from the root bank.
func (cl *Client) GetFeeRateGovernor(ctx context.Context) (out *GetFeeRateGovernorResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getFeeRateGovernor")
	return
}
