package rpc

import (
	"context"
	"errors"

	"github.com/gagliardetto/solana-go"
)

type GetMultipleAccountsResult struct {
	RPCContext
	Value []*Account `json:"value"`
}

// GetMultipleAccounts returns the account information for a list of Pubkeys.
func (cl *Client) GetMultipleAccounts(
	ctx context.Context,
	accounts ...solana.PublicKey, // An array of Pubkeys to query
) (out *GetMultipleAccountsResult, err error) {
	return cl.GetMultipleAccountsWithOpts(
		ctx,
		accounts,
		nil,
	)
}

// GetMultipleAccountsWithOpts returns the account information for a list of Pubkeys.
func (cl *Client) GetMultipleAccountsWithOpts(
	ctx context.Context,
	accounts []solana.PublicKey,
	opts *GetAccountInfoOpts,
) (out *GetMultipleAccountsResult, err error) {
	params := []interface{}{accounts}

	if opts != nil {
		obj := M{}
		if opts.Encoding != "" {
			obj["encoding"] = opts.Encoding
		}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if opts.DataSlice != nil {
			obj["dataSlice"] = M{
				"offset": opts.DataSlice.Offset,
				"length": opts.DataSlice.Length,
			}
			if opts.Encoding == solana.EncodingJSONParsed {
				return nil, errors.New("cannot use dataSlice with EncodingJSONParsed")
			}
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}

	err = cl.rpcClient.CallFor(&out, "getMultipleAccounts", params)
	if err != nil {
		return nil, err
	}
	if out.Value == nil {
		return nil, ErrNotFound
	}
	return
}
