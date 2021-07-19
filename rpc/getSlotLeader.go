package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetSlotLeader returns the current slot leader.
func (cl *Client) GetSlotLeader(
	ctx context.Context,
	commitment CommitmentType, // optional
) (out solana.PublicKey, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params, M{"commitment": commitment})
	}
	err = cl.rpcClient.CallFor(&out, "getSlotLeader", params)
	return
}
