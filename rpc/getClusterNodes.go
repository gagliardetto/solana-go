package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetClusterNodesResult struct {
	Pubkey solana.PublicKey `json:"pubkey"` // Node public key, as base-58 encoded string
	Gossip *string          `json:"gossip"` // Gossip network address for the node
	TPU    *string          `json:"tpu"`    // TPU network address for the node

	// TODO: "" or nil ?
	RPC          *string   `json:"rpc"`          // JSON RPC network address for the node, or empty if the JSON RPC service is not enabled
	Version      *string   `json:"version"`      // The software version of the node, or empty if the version information is not available
	FeatureSet   bin.Int64 `json:"featureSet"`   // The unique identifier of the node's feature set
	ShredVersion bin.Int64 `json:"shredVersion"` // The shred version the node has been configured to use
}

// GetClusterNodes returns information about all the nodes participating in the cluster.
func (cl *Client) GetClusterNodes(ctx context.Context) (out []*GetClusterNodesResult, err error) {
	err = cl.rpcClient.CallFor(&out, "getClusterNodes")
	return
}
