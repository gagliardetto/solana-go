package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

// GetSlotLeaders returns the slot leaders for a given slot range.
func (cl *Client) GetSlotLeaders(
	ctx context.Context,
	start bin.Uint64,
	limit bin.Uint64,
) (out []solana.PublicKey, err error) {
	params := []interface{}{start, limit}
	err = cl.rpcClient.CallFor(&out, "getSlotLeaders", params)
	return
}
