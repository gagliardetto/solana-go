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

package serum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
	"github.com/dfuse-io/solana-go/rpc/ws"
	"go.uber.org/zap"
)

//go:generate rice embed-go

// TODO: hit the chain and
func KnownMarket() ([]*MarketMeta, error) {
	box := rice.MustFindBox("data").MustBytes("markets.json")
	if box == nil {
		return nil, fmt.Errorf("unable to retrieve known markets")
	}

	dec := json.NewDecoder(bytes.NewReader(box))
	var markets []*MarketMeta
	err := dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}

func FetchOpenOrders(ctx context.Context, rpcCli *rpc.Client, openOrdersAddr solana.PublicKey) (*OpenOrdersMeta, error) {
	acctInfo, err := rpcCli.GetAccountInfo(ctx, openOrdersAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get open orders account:%w", err)
	}

	openOrdersMeta := &OpenOrdersMeta{}

	if err := openOrdersMeta.OpenOrders.Decode(acctInfo.Value.Data); err != nil {
		return nil, fmt.Errorf("decoding market v2: %w", err)
	}

	return openOrdersMeta, nil
}

func FetchMarket(ctx context.Context, rpcCli *rpc.Client, marketAddr solana.PublicKey) (*MarketMeta, error) {
	acctInfo, err := rpcCli.GetAccountInfo(ctx, marketAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to get market account:%w", err)
	}

	meta := &MarketMeta{
		Address: marketAddr,
	}

	dataLen := len(acctInfo.Value.Data)
	switch dataLen {
	// case 380:
	// 	// if err := meta.MarketV1.Decode(acctInfo.Value.Data); err != nil {
	// 	// 	return nil, fmt.Errorf("decoding market v1: %w", err)
	// 	// }
	// 	return nil, fmt.Errorf("Unsupported market version, w/ data length of 380")

	case 388:
		if err := meta.MarketV2.Decode(acctInfo.Value.Data); err != nil {
			return nil, fmt.Errorf("decoding market v2: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported market data length: %d", dataLen)
	}

	if err := rpcCli.GetAccountDataIn(ctx, meta.MarketV2.QuoteMint, &meta.QuoteMint); err != nil {
		return nil, fmt.Errorf("getting quote mint: %w", err)
	}

	if err := rpcCli.GetAccountDataIn(ctx, meta.MarketV2.BaseMint, &meta.BaseMint); err != nil {
		return nil, fmt.Errorf("getting base token: %w", err)
	}

	return meta, nil
}

func StreamOpenOrders(client *ws.Client) error {
	sub, err := client.ProgramSubscribe(DEXProgramIDV2, rpc.CommitmentSingleGossip)
	if err != nil {
		return fmt.Errorf("unable to subscribe to programID %q: %w", DEXProgramIDV2, err)
	}
	count := 0
	for {
		d, err := sub.Recv()
		if err != nil {
			return fmt.Errorf("received error from programID subscription: %w", err)
		}
		res := d.(*ws.ProgramResult)

		var f *AccountFlag
		err = bin.NewDecoder(res.Value.Account.Data).Decode(&f)
		if err != nil {
			fmt.Println("***********************************", err)
			zlog.Debug("unable to decoce account flag for account... skipping",
				zap.Stringer("account_address", res.Value.PubKey),
			)
			continue
		}
		fmt.Printf("%d - %s\n", count, f.String())
		count++
	}
}
