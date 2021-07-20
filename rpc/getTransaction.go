package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetTransactionOpts struct {
	Encoding solana.EncodingType `json:"encoding,omitempty"`

	// Desired commitment. "processed" is not supported. If parameter not provided, the default is "finalized".
	Commitment CommitmentType `json:"commitment,omitempty"`
}

// GetTransaction returns transaction details for a confirmed transaction.
//
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use `getConfirmedTransaction` for solana-core v1.6
func (cl *Client) GetTransaction(
	ctx context.Context,
	txSig solana.Signature, // transaction signature
	opts *GetTransactionOpts,
) (out *GetTransactionResult, err error) {
	params := []interface{}{txSig}
	if opts != nil {
		obj := M{}
		if opts.Encoding != "" {
			obj["encoding"] = opts.Encoding
		}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getTransaction", params)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrNotFound
	}
	return
}

type GetTransactionResult struct {
	// The slot this transaction was processed in.
	Slot bin.Uint64 `json:"slot"`

	// Estimated production time, as Unix timestamp (seconds since the Unix epoch)
	// of when the transaction was processed.
	// Nil if not available.
	BlockTime *UnixTimeSeconds `json:"blockTime"`

	Transaction *ParsedTransaction `json:"transaction"`
	Meta        *TransactionMeta   `json:"meta,omitempty"`
}
