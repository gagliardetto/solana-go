package rpc

import (
	"context"
)

// GetHealth returns the current health of the node.
// If one or more --trusted-validator arguments are provided
// to solana-validator, "ok" is returned when the node has within
// HEALTH_CHECK_SLOT_DISTANCE slots of the highest trusted validator,
// otherwise an error is returned. "ok" is always returned if no
// trusted validators are provided.
func (cl *Client) GetHealth(ctx context.Context) (out string, err error) {
	err = cl.rpcClient.CallFor(&out, "getHealth")
	return
}

const HealthOk = "ok"
