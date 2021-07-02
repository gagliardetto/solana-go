package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetTransactionCount returns the current Transaction count from the ledger.
func (cl *Client) GetTransactionCount(
	ctx context.Context,
	commitment CommitmentType,
) (out bin.Uint64, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getTransactionCount", params)
	return
}
