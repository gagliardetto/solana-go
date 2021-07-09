// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package rpc

import (
	"context"
	"errors"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

// GetAccountInfo returns all information associated with the account of provided publicKey.
func (cl *Client) GetAccountInfo(ctx context.Context, account solana.PublicKey) (out *GetAccountInfoResult, err error) {
	return cl.GetAccountInfoWithOpts(
		ctx,
		account,
		&GetAccountInfoOpts{
			Encoding:   EncodingBase64,
			Commitment: "",
			Offset:     nil,
			Length:     nil,
		},
	)
}

// GetAccountInfo populates the provided `inVar` parameter with all information associated with the account of provided publicKey.
func (cl *Client) GetAccountDataIn(ctx context.Context, account solana.PublicKey, inVar interface{}) (err error) {
	resp, err := cl.GetAccountInfo(ctx, account)
	if err != nil {
		return err
	}

	return bin.NewDecoder(resp.Value.Data).Decode(inVar)
}

type GetAccountInfoOpts struct {
	Encoding   EncodingType
	Commitment CommitmentType
	Offset     *uint64
	Length     *uint64
}

// GetAccountInfoWithOpts returns all information associated with the account of provided publicKey.
// You can limit the returned account data with the offset and length parameters.
// You can specify the encoding of the returned data with the encoding parameter.
func (cl *Client) GetAccountInfoWithOpts(
	ctx context.Context,
	account solana.PublicKey,
	opts *GetAccountInfoOpts,
) (out *GetAccountInfoResult, err error) {

	obj := M{}

	if opts != nil {
		if opts.Encoding != "" {
			obj["encoding"] = opts.Encoding
		}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if opts.Offset != nil && opts.Length != nil {
			obj["dataSlice"] = M{
				"offset": opts.Offset,
				"length": opts.Length,
			}
			if opts.Encoding == EncodingJSONParsed {
				return nil, errors.New("cannot use dataSlice with EncodingJSONParsed")
			}
		}
	}

	params := []interface{}{account}
	if len(obj) > 0 {
		params = append(params, obj)
	}

	err = cl.rpcClient.CallFor(&out, "getAccountInfo", params)
	if err != nil {
		return nil, err
	}

	if out.Value == nil {
		return nil, ErrNotFound
	}

	return out, nil
}
