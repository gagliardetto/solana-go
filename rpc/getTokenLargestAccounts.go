package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetTokenLargestAccounts returns the 20 largest accounts of a particular SPL Token type.
func (cl *Client) GetTokenLargestAccounts(
	ctx context.Context,
	tokenMint solana.PublicKey, // Pubkey of token Mint to query
	commitment CommitmentType, // optional
) (out *GetTokenLargestAccountsResult, err error) {
	params := []interface{}{tokenMint}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getTokenLargestAccounts", params)
	return
}

type GetTokenLargestAccountsResult struct {
	RPCContext
	Value []*TokenLargestAccountsResult `json:"value"`
}
type TokenLargestAccountsResult struct {
	Address string `json:"address"` // the address of the token account
	UiTokenAmount
}
