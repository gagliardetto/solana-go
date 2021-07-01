package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

type GetLargestAccountsResult struct {
	RPCContext
	Value []LargestAccountsResult `json:"value"`
}

type LargestAccountsResult struct {
	Address  solana.PublicKey `json:"address"`  // address of the account
	Lamports bin.Uint64       `json:"lamports"` // number of lamports in the account
}

type LargestAccountsFilterType string

const (
	LargestAccountsFilterCirculating    LargestAccountsFilterType = "circulating"
	LargestAccountsFilterNonCirculating LargestAccountsFilterType = "nonCirculating"
)

// GetLargestAccounts returns the 20 largest accounts,
// by lamport balance (results may be cached up to two hours).
func (cl *Client) GetLargestAccounts(
	ctx context.Context,
	commitment CommitmentType,
	filter LargestAccountsFilterType, // filter results by account type; currently supported: circulating|nonCirculating
) (out *GetLargestAccountsResult, err error) {
	params := []interface{}{}
	obj := M{}
	if commitment != "" {
		obj["commitment"] = commitment
	}
	if filter != "" {
		obj["filter"] = filter
	}
	if len(obj) > 0 {
		params = append(params, obj)
	}
	err = cl.rpcClient.CallFor(&out, "getLargestAccounts", params)
	return
}
