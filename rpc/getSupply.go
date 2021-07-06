package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetSupplyResult struct {
	RPCContext
	Value *SupplyResult `json:"value"`
}

type SupplyResult struct {
	Total                  bin.Uint64         `json:"total"`                  // Total supply in lamports
	Circulating            bin.Uint64         `json:"circulating"`            // Circulating supply in lamports
	NonCirculating         bin.Uint64         `json:"nonCirculating"`         // Non-circulating supply in lamports
	NonCirculatingAccounts []solana.PublicKey `json:"nonCirculatingAccounts"` // an array of account addresses of non-circulating accounts, as strings
}

// GetSupply returns information about the current supply.
func (cl *Client) GetSupply(
	ctx context.Context,
	commitment CommitmentType,
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
