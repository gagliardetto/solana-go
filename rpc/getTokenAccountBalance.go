package rpc

import (
	"context"

	"github.com/dfuse-io/solana-go"
)

type GetTokenAccountBalanceResult struct {
	RPCContext
	Value *UiTokenAmount `json:"value"`
}

// GetTokenAccountBalance returns the token balance of an SPL Token account.
func (cl *Client) GetTokenAccountBalance(
	ctx context.Context,
	account solana.PublicKey,
	commitment CommitmentType,
) (out *GetTokenAccountBalanceResult, err error) {
	params := []interface{}{account}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getTokenAccountBalance", params)
	return
}
