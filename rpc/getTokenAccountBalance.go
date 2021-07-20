package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetTokenAccountBalance returns the token balance of an SPL Token account.
func (cl *Client) GetTokenAccountBalance(
	ctx context.Context,
	account solana.PublicKey,
	commitment CommitmentType, // optional
) (out *GetTokenAccountBalanceResult, err error) {
	params := []interface{}{account}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getTokenAccountBalance", params)
	return
}

type GetTokenAccountBalanceResult struct {
	RPCContext
	Value *UiTokenAmount `json:"value"`
}
