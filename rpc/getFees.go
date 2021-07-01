package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

type GetFeesResult struct {
	RPCContext
	Value FeesResult `json:"value"`
}

type FeesResult struct {
	Blockhash            solana.Hash   `json:"blockhash"`            // a Hash as base-58 encoded string
	FeeCalculator        FeeCalculator `json:"feeCalculator"`        // FeeCalculator object, the fee schedule for this block hash
	LastValidSlot        bin.Uint64    `json:"lastValidSlot"`        // DEPRECATED - this value is inaccurate and should not be relied upon
	LastValidBlockHeight bin.Uint64    `json:"lastValidBlockHeight"` // last block height at which a blockhash will be valid
}

// GetFees returns a recent block hash from the ledger,
// a fee schedule that can be used to compute the cost
// of submitting a transaction using it, and the last
// slot in which the blockhash will be valid.
func (cl *Client) GetFees(
	ctx context.Context,
	commitment CommitmentType,
) (out *GetFeesResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallFor(&out, "getFees", params)
	return
}
