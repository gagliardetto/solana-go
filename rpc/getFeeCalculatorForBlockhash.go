package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetFeeCalculatorForBlockhash returns the fee calculator
// associated with the query blockhash, or null if the blockhash has expired.
func (cl *Client) GetFeeCalculatorForBlockhash(
	ctx context.Context,
	hash solana.Hash, // query blockhash
	commitment CommitmentType, // optional
) (out *GetFeeCalculatorForBlockhashResult, err error) {
	params := []interface{}{hash}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getFeeCalculatorForBlockhash", params)
	return
}

type GetFeeCalculatorForBlockhashResult struct {
	RPCContext

	// Value will be nil if the query blockhash has expired.
	Value *FeeCalculatorForBlockhashResult `json:"value"`
}

type FeeCalculatorForBlockhashResult struct {
	// Object describing the cluster fee rate at the queried blockhash
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}
