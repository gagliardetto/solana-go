package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetIdentity returns the identity pubkey for the current node.
func (cl *Client) GetIdentity(ctx context.Context) (out *GetIdentityResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getIdentity")
	return
}

type GetIdentityResult struct {
	// The identity pubkey of the current node.
	Identity solana.PublicKey `json:"identity"`
}
