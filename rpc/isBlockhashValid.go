package rpc

import (
	"context"
	"github.com/gagliardetto/solana-go"
)

// IsBlockhashValid returns the balance of the account of provided publicKey.
func (cl *Client) IsBlockhashValid(
	ctx context.Context,
	// Blockhash to be queried. Required.
	blockHash solana.Hash,

	// Commitment requirement. Optional.
	commitment CommitmentType,
) (out *IsValidBlockhashResult, err error) {
	params := []interface{}{blockHash}
	if commitment != "" {
		params = append(params, M{"commitment": string(commitment)})
	}

	err = cl.rpcClient.CallForInto(ctx, &out, "isBlockhashValid", params)
	return
}
