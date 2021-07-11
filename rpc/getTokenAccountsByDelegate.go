package rpc

import (
	"context"
	"errors"

	"github.com/gagliardetto/solana-go"
)

type GetTokenAccountsResult struct {
	RPCContext
	Value []*Account `json:"value"`
}

type GetTokenAccountsConfig struct {
	Mint solana.PublicKey `json:"mint"` // Pubkey of the specific token Mint to limit accounts to
	// OR:
	ProgramId solana.PublicKey `json:"programId"` // Pubkey of the Token program ID that owns the accounts
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
		obj := M{}
		if !conf.Mint.IsZero() {
			obj["mint"] = conf.Mint
		}
		if !conf.ProgramId.IsZero() {
			obj["programId"] = conf.ProgramId
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	{
		obj := M{}
		if opts != nil {
			if opts.Commitment != "" {
				obj["commitment"] = opts.Commitment
			}
			if opts.Encoding != "" {
				// TODO: remove option?
				obj["encoding"] = opts.Encoding
			}
			if opts.DataSlice != nil {
				obj["dataSlice"] = M{
					"offset": opts.DataSlice.Offset,
					"length": opts.DataSlice.Length,
				}
			}
			if len(obj) > 0 {
				params = append(params, obj)
			}
		}
	}

	err = cl.rpcClient.CallFor(&out, "getTokenAccountsByDelegate", params)
	return
}
