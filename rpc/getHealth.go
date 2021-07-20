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
//
// - If the node is healthy: "ok"
// - If the node is unhealthy, a JSON RPC error response is returned.
//   The specifics of the error response are UNSTABLE and may change in the future.
func (cl *Client) GetHealth(ctx context.Context) (out string, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getHealth", nil)
	return
}

const HealthOk = "ok"
