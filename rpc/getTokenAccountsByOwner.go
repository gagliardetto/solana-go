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
				if opts.Encoding == solana.EncodingJSONParsed {
					return nil, errors.New("cannot use dataSlice with EncodingJSONParsed")
				}
			}
			if len(optsObj) > 0 {
				params = append(params, optsObj)
			}
		}
	}

	err = cl.rpcClient.CallForInto(ctx, &out, "getTokenAccountsByOwner", params)
	return
}
