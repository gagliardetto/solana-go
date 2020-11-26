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

package token

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
)

//go:generate rice embed-go

func FetchMints(ctx context.Context, rpcCli *rpc.Client) (out []*Mint, err error) {
	resp, err := rpcCli.GetProgramAccounts(
		ctx,
		TOKEN_PROGRAM_ID,
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					DataSize: MINT_SIZE,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("resp empty... program account not found")
	}

	for _, keyedAcct := range resp {
		acct := keyedAcct.Account

		m, err := DecodeMint(acct.Data)
		if err != nil {
			return nil, fmt.Errorf("unable to decode mint %q: %w", acct.Owner.String(), err)
		}
		out = append(out, m)

	}
	return
}

func KnownMints(network string) ([]*MintMeta, error) {
	box := rice.MustFindBox("mints-data").MustBytes(network + "-tokens.json")
	if box == nil {
		return nil, fmt.Errorf("unable to retrieve known markets")
	}

	dec := json.NewDecoder(bytes.NewReader(box))
	var markets []*MintMeta
	err := dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}

func GetMint(ctx context.Context, cli *rpc.Client, mintPubKey solana.PublicKey) (*Mint, error) {
	acctInfo, err := cli.GetAccountInfo(ctx, mintPubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve base mint: %w", err)
	}

	return DecodeMint(acctInfo.Value.Data)
}
