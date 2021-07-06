package rpc

import (
	"context"
	"errors"

	"github.com/gagliardetto/solana-go"
)

// GetTokenAccountsByOwner returns all SPL Token accounts by token owner.
func (cl *Client) GetTokenAccountsByOwner(
	ctx context.Context,
	owner solana.PublicKey,
	conf *GetTokenAccountsConfig,
	opts *GetTokenAccountsOpts,
) (out *GetTokenAccountsResult, err error) {
	params := []interface{}{owner}
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

	err = cl.rpcClient.CallFor(&out, "getTokenAccountsByOwner", params)
	return
}
