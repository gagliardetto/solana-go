package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

// GetFees returns a recent block hash from the ledger,
// a fee schedule that can be used to compute the cost
// of submitting a transaction using it, and the last
// slot in which the blockhash will be valid.
func (cl *Client) GetFees(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out *GetFeesResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getFees", params)
	return
}

type GetFeesResult struct {
	RPCContext
	Value *FeesResult `json:"value"`
}

type FeesResult struct {
	// A Hash.
	Blockhash solana.Hash `json:"blockhash"`

	// FeeCalculator object, the fee schedule for this block hash.
	FeeCalculator FeeCalculator `json:"feeCalculator"`

	// DEPRECATED - this value is inaccurate and should not be relied upon.
	LastValidSlot bin.Uint64 `json:"lastValidSlot"`

	// Last block height at which a blockhash will be valid.
	LastValidBlockHeight bin.Uint64 `json:"lastValidBlockHeight"`
}
