package rpc

import (
	"encoding/base64"
	"fmt"

	"github.com/dfuse-io/solana-go"
)

// type ContactInfo struct {
// 	Pubkey  string `json:"pubkey"`
// 	Gossip  string `json:"gossip,omitempty"`
// 	TPU     string `json:"tpu,omitempty"`
// 	RPC     string `json:"rpc,omitempty"`
// 	Version string `json:"version,omitempty"`
// }

type RPCContext struct {
	Context struct {
		Slot solana.U64
	} `json:"context,omitempty"`
}

///

type GetBalanceResult struct {
	RPCContext
	Value solana.U64 `json:"value"`
}

///

type GetSlotResult solana.U64

///

type GetRecentBlockhashResult struct {
	RPCContext
	Value BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     solana.PublicKey `json:"blockhash"` /* make this a `Hash` type, which is a copy of the PublicKey` type */
	FeeCalculator FeeCalculator    `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature solana.U64 `json:"lamportsPerSignature"`
}

///

type GetConfirmedBlockResult struct {
	Blockhash         solana.PublicKey      `json:"blockhash"`
	PreviousBlockhash solana.PublicKey      `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        solana.U64            `json:"parentSlot"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []BlockReward         `json:"rewards"`
	BlockTime         solana.U64            `json:"blockTime,omitempty"`
}

type BlockReward struct {
	Pubkey   solana.PublicKey `json:"pubkey"`
	Lamports solana.U64       `json:"lamports"`
}

type TransactionWithMeta struct {
	Transaction *solana.Transaction `json:"transaction"`
	Meta        *TransactionMeta    `json:"meta,omitempty"`
}

type TransactionMeta struct {
	Err          interface{}  `json:"err"`
	Fee          solana.U64   `json:"fee"`
	PreBalances  []solana.U64 `json:"preBalances"`
	PostBalances []solana.U64 `json:"postBalances"`
}

///

type GetAccountInfoResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	Lamports   solana.U64       `json:"lamports"`
	Data       []string         `json:"data"`
	Owner      solana.PublicKey `json:"owner"`
	Executable bool             `json:"executable"`
	RentEpoch  solana.U64       `json:"rentEpoch"`
}

func (a *Account) MustDataToBytes() []byte {
	d, err := a.DataToBytes()
	if err != nil {
		panic(fmt.Sprintf("failed to base64 decode: %s", err))
	}
	return d
}

func (a *Account) DataToBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(a.Data[0])
}

type KeyedAccount struct {
	Pubkey  solana.PublicKey `json:"pubkey"`
	Account *Account         `json:"account"`
}
type GetProgramAccountsResult []*KeyedAccount

type GetProgramAccountsOpts struct {
	Encoding string `json:"encoding,omitempty"`

	Commitment CommitmentType `json:"commitment,omitempty"`

	// Filter on accounts, implicit AND between filters
	Filters []RPCFilter `json:"filters,omitempty"`
}

type RPCFilter struct {
	Memcmp   *RPCFilterMemcmp `json:"memcmp,omitempty"`
	DataSize solana.U64       `json:"dataSize,omitempty"`
}

type RPCFilterMemcmp struct {
	Offset int           `json:"offset"`
	Bytes  solana.Base58 `json:"bytes"`
}

///

type CommitmentType string

const (
	CommitmentMax          = CommitmentType("max")
	CommitmentRecent       = CommitmentType("recent")
	CommitmentRoot         = CommitmentType("root")
	CommitmentSingle       = CommitmentType("single")
	CommitmentSingleGossip = CommitmentType("singleGossip")
)
