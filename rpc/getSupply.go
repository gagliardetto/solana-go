package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetSupply returns information about the current supply.
func (cl *Client) GetSupply(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out *GetSupplyResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getSupply", params)
	return
}

type GetSupplyResult struct {
	RPCContext
	Value *SupplyResult `json:"value"`
}

type SupplyResult struct {
	// TODO: use bin.Uint64 ???

	// Total supply in lamports
	Total uint64 `json:"total"`

	// Circulating supply in lamports.
	Circulating uint64 `json:"circulating"`

	// Non-circulating supply in lamports.
	NonCirculating uint64 `json:"nonCirculating"`

	// An array of account addresses of non-circulating accounts.
	NonCirculatingAccounts []solana.PublicKey `json:"nonCirculatingAccounts"`
}
