package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetSlotLeaders returns the slot leaders for a given slot range.
func (cl *Client) GetSlotLeaders(
	ctx context.Context,
	start int,
	limit int,
) (out []solana.PublicKey, err error) {
	params := []interface{}{start, limit}
	err = cl.rpcClient.CallFor(&out, "getSlotLeaders", params)
	return
}
