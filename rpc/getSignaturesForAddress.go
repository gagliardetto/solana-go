package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type GetSignaturesForAddressOpts struct {
	Limit      *int             `json:"limit,omitempty"`      // (optional) maximum transaction signatures to return (between 1 and 1,000, default: 1,000).
	Before     solana.Signature `json:"before,omitempty"`     // (optional) start searching backwards from this transaction signature. If not provided the search starts from the top of the highest max confirmed block.
	Until      solana.Signature `json:"until,omitempty"`      // (optional) search until this transaction signature, if found before limit reached.
	Commitment CommitmentType   `json:"commitment,omitempty"` // (optional) Commitment; "processed" is not supported. If parameter not provided, the default is "finalized".
}

// GetSignaturesForAddress returns confirmed signatures for transactions
// involving an address backwards in time from the provided signature
// or most recent confirmed block.
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use getConfirmedSignaturesForAddress2 for solana-core v1.6
func (cl *Client) GetSignaturesForAddress(
	ctx context.Context,
	account solana.PublicKey,
) (out []*TransactionSignature, err error) {
	return cl.GetSignaturesForAddressWithOpts(
		ctx,
		account,
		nil,
	)
}

// GetSignaturesForAddressWithOpts returns confirmed signatures for transactions
// involving an address backwards in time from the provided signature
// or most recent confirmed block.
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use getConfirmedSignaturesForAddress2 for solana-core v1.6
func (cl *Client) GetSignaturesForAddressWithOpts(
	ctx context.Context,
	account solana.PublicKey,
	// TODO: adopt opts style for all funcs that have lots of options.
	opts *GetSignaturesForAddressOpts,
) (out []*TransactionSignature, err error) {
	params := []interface{}{account}
	if opts != nil {
		obj := M{}
		if opts.Limit != nil {
			obj["limit"] = opts.Limit
		}
		if !opts.Before.IsZero() {
			obj["before"] = opts.Before
		}
		if !opts.Until.IsZero() {
			obj["until"] = opts.Until
		}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		params = append(params, obj)
	}

	err = cl.rpcClient.CallFor(&out, "getSignaturesForAddress", params)
	return
}
