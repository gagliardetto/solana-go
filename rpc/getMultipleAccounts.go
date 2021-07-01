package rpc

import (
	"context"
	"errors"

	"github.com/dfuse-io/solana-go"
)

type GetMultipleAccountsResult struct {
	RPCContext
	Value []*Account `json:"value"`
}

// GetMultipleAccounts returns the account information for a list of Pubkeys.
func (cl *Client) GetMultipleAccounts(
	ctx context.Context,
	accounts []solana.PublicKey,
	encoding EncodingType,
	commitment CommitmentType,
	offset *uint,
	length *uint,
) (out *GetMultipleAccountsResult, err error) {
	obj := M{}

	if encoding != "" {
		obj["encoding"] = encoding
	}
	if commitment != "" {
		obj["commitment"] = commitment
	}
	if offset != nil && length != nil {
		obj["dataSlice"] = M{
			"offset": offset,
			"length": length,
		}
		if encoding == EncodingJSONParsed {
			return nil, errors.New("cannot use dataSlice with EncodingJSONParsed")
		}
	}

	params := []interface{}{accounts}
	if len(obj) > 0 {
		params = append(params, obj)
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
