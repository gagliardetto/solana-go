package rpc

import (
	"context"

	"github.com/dfuse-io/solana-go"
)

type GetFeeCalculatorForBlockhashResult struct {
	RPCContext
	Value FeeCalculatorForBlockhashResult `json:"value"`
}

type FeeCalculatorForBlockhashResult struct {
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

// GetFeeCalculatorForBlockhash returns the fee calculator
// associated with the query blockhash, or null if the blockhash has expired.
func (cl *Client) GetFeeCalculatorForBlockhash(
	ctx context.Context,
	hash solana.Hash, // query blockhash as a Base58 encoded string
	commitment CommitmentType,
) (out *GetFeeCalculatorForBlockhashResult, err error) {
	params := []interface{}{hash}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallFor(&out, "getFeeCalculatorForBlockhash", params)
	return
}
