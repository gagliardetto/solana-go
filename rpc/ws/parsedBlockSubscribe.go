// Copyright 2022 github.com/gagliardetto
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

package ws

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type ParsedBlockResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Slot  uint64                    `json:"slot"`
		Err   interface{}               `json:"err,omitempty"`
		Block *rpc.GetParsedBlockResult `json:"block,omitempty"`
	} `json:"value"`
}

// NOTE: Unstable, disabled by default
//
// Subscribe to receive notification anytime a new block is Confirmed or Finalized.
//
// **This subscription is unstable and only available if the validator was started
// with the `--rpc-pubsub-enable-block-subscription` flag. The format of this
// subscription may change in the future**
func (cl *Client) ParsedBlockSubscribe(
	filter BlockSubscribeFilter,
	opts *BlockSubscribeOpts,
) (*ParsedBlockSubscription, error) {
	var params []interface{}
	if filter != nil {
		switch v := filter.(type) {
		case BlockSubscribeFilterAll:
			params = append(params, "all")
		case *BlockSubscribeFilterMentionsAccountOrProgram:
			params = append(params, rpc.M{"mentionsAccountOrProgram": v.Pubkey})
		}
	}
	if opts != nil {
		obj := make(rpc.M)
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		obj["encoding"] = solana.EncodingJSONParsed
		if opts.TransactionDetails != "" {
			obj["transactionDetails"] = opts.TransactionDetails
		}
		if opts.Rewards != nil {
			obj["rewards"] = opts.Rewards
		}
		if opts.MaxSupportedTransactionVersion != nil {
			obj["maxSupportedTransactionVersion"] = *opts.MaxSupportedTransactionVersion
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	genSub, err := cl.subscribe(
		params,
		nil,
		"blockSubscribe",
		"blockUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res ParsedBlockResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &ParsedBlockSubscription{
		sub: genSub,
	}, nil
}

type ParsedBlockSubscription struct {
	sub *Subscription
}

func (sw *ParsedBlockSubscription) Recv(ctx context.Context) (*ParsedBlockResult, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case d := <-sw.sub.stream:
		return d.(*ParsedBlockResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *ParsedBlockSubscription) Err() <-chan error {
	return sw.sub.err
}

func (sw *ParsedBlockSubscription) Response() <-chan *ParsedBlockResult {
	typedChan := make(chan *ParsedBlockResult, 1)
	go func(ch chan *ParsedBlockResult) {
		// TODO: will this subscription yield more than one result?
		d, ok := <-sw.sub.stream
		if !ok {
			return
		}
		ch <- d.(*ParsedBlockResult)
	}(typedChan)
	return typedChan
}

func (sw *ParsedBlockSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}
