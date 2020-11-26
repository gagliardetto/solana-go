package tokenregistry

import (
	"context"
	"fmt"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
)

func GetTokenRegistryEntry(ctx context.Context, rpcCli *rpc.Client, mintAddress solana.PublicKey) (*TokenMeta, error) {
	resp, err := rpcCli.GetProgramAccounts(
		ctx,
		ProgramID(),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					Memcmp: &rpc.RPCFilterMemcmp{
						Offset: 5,
						Bytes:  mintAddress[:], // hackey to convert [32]byte to []byte
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("resp empty... cannot find account")
	}

	for _, keyedAcct := range resp {
		acct := keyedAcct.Account
		t, err := DecodeTokenMeta(acct.Data)
		if err != nil {
			return nil, fmt.Errorf("unable to decode token meta %q: %w", acct.Owner.String(), err)
		}
		return t, nil
	}
	return nil, rpc.ErrNotFound
}

func GetEntries(ctx context.Context, rpcCli *rpc.Client) (out []*TokenMeta, err error) {
	resp, err := rpcCli.GetProgramAccounts(
		ctx,
		ProgramID(),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{
				{
					DataSize: TOKEN_META_SIZE,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("resp empty... cannot find accounts")
	}

	for _, keyedAcct := range resp {
		acct := keyedAcct.Account
		t, err := DecodeTokenMeta(acct.Data)
		if err != nil {
			return nil, fmt.Errorf("unable to decode token meta %q: %w", acct.Owner.String(), err)
		}
		out = append(out, t)
	}
	return out, nil
}
