package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetTransactionResult struct {
	Slot bin.Uint64 `json:"slot"` // the slot this transaction was processed in
	// TODO: int64 as readable time
	BlockTime   bin.Int64          `json:"blockTime"` // estimated production time, as Unix timestamp (seconds since the Unix epoch) of when the transaction was processed. null if not available
	Transaction *ParsedTransaction `json:"transaction"`
	Meta        *TransactionMeta   `json:"meta,omitempty"`
}

type GetTransactionOpts struct {
	Encoding   EncodingType   `json:"encoding,omitempty"`
	Commitment CommitmentType `json:"commitment,omitempty"` // "processed" is not supported. If parameter not provided, the default is "finalized".
}

// GetTransaction returns transaction details for a confirmed transaction.
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use getConfirmedTransaction for solana-core v1.6
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
	err = cl.rpcClient.CallFor(&out, "getTransaction", params)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrNotFound
	}
	return
}
