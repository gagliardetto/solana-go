package rpc

import (
	"context"

	"github.com/dfuse-io/solana-go"
)

type GetIdentityResult struct {
	Identity solana.PublicKey `json:"identity"` // the identity pubkey of the current node
}

// GetIdentity returns the identity pubkey for the current node.
func (cl *Client) GetIdentity(ctx context.Context) (out *GetIdentityResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getIdentity")
	return
}
