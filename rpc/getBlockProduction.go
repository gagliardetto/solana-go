package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetBlockProductionResult struct {
	RPCContext
	Value BlockProductionResult `json:"value"`
}

type IdentityToSlotsBlocks map[string][2]int64

type BlockProductionResult struct {
	ByIdentity IdentityToSlotsBlocks `json:"byIdentity"` //  a dictionary of validator identities, as base-58 encoded strings. Value is a two element array containing the number of leader slots and the number of blocks produced.
	Range      SlotRangeResponse     `json:"range"`
}

type SlotRangeResponse struct {
	FirstSlot bin.Uint64 `json:"firstSlot"` // first slot of the block production information (inclusive)
	LastSlot  bin.Uint64 `json:"lastSlot"`  // last slot of block production information (inclusive)
}
type SlotRangeRequest struct {
	FirstSlot int     `json:"firstSlot"`          // first slot to return block production information for (inclusive)
	LastSlot  *int    `json:"lastSlot,omitempty"` // (optional) last slot to return block production information for (inclusive). If parameter not provided, defaults to the highest slot
	Identity  *string `json:"identity,omitempty"` // (optional) Only return results for this validator identity (base-58 encoded)
}

// GetBlockProduction returns recent block production information from the current or previous epoch.
func (cl *Client) GetBlockProduction(
	ctx context.Context,
) (out *GetBlockProductionResult, err error) {
	return cl.GetBlockProductionWithOpts(
		ctx,
		nil,
	)
}

type GetBlockProductionOpts struct {
	Commitment CommitmentType
	Range      *SlotRangeRequest // Slot range to return block production for. If parameter not provided, defaults to current epoch.
}

// GetBlockProduction returns recent block production information from the current or previous epoch.
func (cl *Client) GetBlockProductionWithOpts(
	ctx context.Context,
	opts *GetBlockProductionOpts,
) (out *GetBlockProductionResult, err error) {
	obj := M{}

	if opts != nil {
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if opts.Range != nil {
			rngObj := M{}
			rngObj["firstSlot"] = opts.Range.FirstSlot
			if opts.Range.LastSlot != nil {
				rngObj["lastSlot"] = opts.Range.LastSlot
			}
			if opts.Range.Identity != nil {
				rngObj["identity"] = opts.Range.Identity
			}
			obj["range"] = rngObj
		}
	}
	params := []interface{}{}
	if len(obj) != 0 {
		params = append(params, obj)
	}
	err = cl.rpcClient.CallFor(&out, "getBlockProduction", params)

	return
}
