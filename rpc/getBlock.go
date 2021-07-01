package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

// GetBlock returns identity and transaction information about a confirmed block in the ledger.
// NEW: This method is only available in solana-core v1.7 or newer. Please use getConfirmedBlock for solana-core v1.6
func (cl *Client) GetBlock(
	ctx context.Context,
	slot bin.Uint64,
) (out *GetBlockResult, err error) {
	return cl.GetBlockWithOpts(
		ctx,
		slot,
		"",
		nil,
		"",
	)
}

type TransactionDetailsType string

const (
	TransactionDetailsFull       TransactionDetailsType = "full"
	TransactionDetailsSignatures TransactionDetailsType = "signatures"
	TransactionDetailsNone       TransactionDetailsType = "none"
)

func (cl *Client) GetBlockWithOpts(
	ctx context.Context,
	slot bin.Uint64,
	// encoding EncodingType,
	transactionDetails TransactionDetailsType, // level of transaction detail to return. If parameter not provided, the default detail level is "full".
	rewards *bool, // whether to populate the rewards array. If parameter not provided, the default includes rewards.
	commitment CommitmentType, // "processed" is not supported. If parameter not provided, the default is "finalized".
) (out *GetBlockResult, err error) {
	obj := M{
		"encoding": EncodingJSON,
		// transactionDetails: <string>
		// rewards: bool
		// commitment
	}

	if transactionDetails != "" {
		obj["transactionDetails"] = transactionDetails
	}
	if rewards != nil {
		obj["rewards"] = rewards
	}
	if commitment != "" {
		obj["commitment"] = commitment
	}

	params := []interface{}{slot, obj}

	err = cl.rpcClient.CallFor(&out, "getBlock", params...)
	return
}

type GetBlockResult struct {
	// The blockhash of this block, as base-58 encoded string.
	Blockhash solana.Hash `json:"blockhash"`

	// The blockhash of this block's parent, as base-58 encoded string;
	// if the parent block is not available due to ledger cleanup,
	// this field will return "11111111111111111111111111111111".
	PreviousBlockhash solana.Hash `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable

	// The slot index of this block's parent.
	ParentSlot bin.Uint64 `json:"parentSlot"`

	// Present if "full" transaction details are requested.
	Transactions []TransactionWithMeta `json:"transactions"`

	// Present if "signatures" are requested for transaction details;
	// an array of signatures strings, corresponding to the transaction order in the block.
	Signatures []solana.Signature `json:"signatures"`

	// Present if rewards are requested.
	Rewards []BlockReward `json:"rewards"`

	// estimated production time, as Unix timestamp (seconds since the Unix epoch). null if not available
	BlockTime *bin.Int64 `json:"blockTime"`
	// the number of blocks beneath this block
	BlockHeight *bin.Uint64 `json:"blockHeight"`
}
