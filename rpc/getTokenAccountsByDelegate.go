package rpc

import (
	"context"
	"errors"

	"github.com/gagliardetto/solana-go"
)

type GetTokenAccountsConfig struct {
	// Pubkey of the specific token Mint to limit accounts to.
	Mint solana.PublicKey `json:"mint"`

	// OR:

	// Pubkey of the Token program ID that owns the accounts.
	ProgramId solana.PublicKey `json:"programId"`
}

type GetTokenAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	Encoding solana.EncodingType `json:"encoding,omitempty"`

	DataSlice *DataSlice `json:"dataSlice,omitempty"`
}

// GetTokenAccountsByDelegate returns all SPL Token accounts by approved Delegate.
func (cl *Client) GetTokenAccountsByDelegate(
	ctx context.Context,
	account solana.PublicKey, // Pubkey of account delegate to query
	conf *GetTokenAccountsConfig,
	opts *GetTokenAccountsOpts,
) (out *GetTokenAccountsResult, err error) {
	params := []interface{}{account}
	if conf == nil {
		return nil, errors.New("conf is nil")
	}
	if !conf.Mint.IsZero() && !conf.ProgramId.IsZero() {
		return nil, errors.New("conf.Mint and conf.ProgramId are both set; must be just one of them")
	}

	{
		confObj := M{}
		if !conf.Mint.IsZero() {
			confObj["mint"] = conf.Mint
		}
		if !conf.ProgramId.IsZero() {
			confObj["programId"] = conf.ProgramId
		}
		if len(confObj) > 0 {
			params = append(params, confObj)
		}
	}
	{
		optsObj := M{}
		if opts != nil {
			if opts.Commitment != "" {
				optsObj["commitment"] = opts.Commitment
			}
			if opts.Encoding != "" {
				optsObj["encoding"] = opts.Encoding
			}
			if opts.DataSlice != nil {
				optsObj["dataSlice"] = M{
					"offset": opts.DataSlice.Offset,
					"length": opts.DataSlice.Length,
				}
			}
			if len(optsObj) > 0 {
				params = append(params, optsObj)
			}
		}
	}

	err = cl.rpcClient.CallForInto(ctx, &out, "getTokenAccountsByDelegate", params)
	return
}

type GetTokenAccountsResult struct {
	RPCContext
	Value []*Account `json:"value"`
}
