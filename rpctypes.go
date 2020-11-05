package solana

// type ContactInfo struct {
// 	Pubkey  string `json:"pubkey"`
// 	Gossip  string `json:"gossip,omitempty"`
// 	TPU     string `json:"tpu,omitempty"`
// 	RPC     string `json:"rpc,omitempty"`
// 	Version string `json:"version,omitempty"`
// }

type RPCContext struct {
	Context struct {
		Slot U64
	} `json:"context,omitempty"`
}

///

type GetBalanceResult struct {
	RPCContext
	Value U64 `json:"value"`
}

///

type GetSlotResult U64

///

type GetRecentBlockhashResult struct {
	RPCContext
	Value BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     PublicKey     `json:"blockhash"` /* make this a `Hash` type, which is a copy of the PublicKey` type */
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature U64 `json:"lamportsPerSignature"`
}

///

type GetConfirmedBlockResult struct {
	Blockhash         PublicKey             `json:"blockhash"`
	PreviousBlockhash PublicKey             `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        U64                   `json:"parentSlot"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []BlockReward         `json:"rewards"`
	BlockTime         U64                   `json:"blockTime,omitempty"`
}

type BlockReward struct {
	Pubkey   PublicKey `json:"pubkey"`
	Lamports U64       `json:"lamports"`
}

type TransactionWithMeta struct {
	Transaction *Transaction     `json:"transaction"`
	Meta        *TransactionMeta `json:"meta,omitempty"`
}

type TransactionMeta struct {
	Err          interface{} `json:"err"`
	Fee          U64         `json:"fee"`
	PreBalances  []U64       `json:"preBalances"`
	PostBalances []U64       `json:"postBalances"`
}

///

type GetAccountInfoResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	Lamports   U64       `json:"lamports"`
	Data       []byte    `json:"data"`
	Owner      PublicKey `json:"owner"`
	Executable bool      `json:"executable"`
	RentEpoch  U64       `json:"rentEpoch"`
}

type KeyedAccount struct {
	Pubkey  PublicKey `json:"pubkey"`
	Account *Account  `json:"account"`
}

///

type GetProgramAccountsResult []*KeyedAccount

type GetProgramAccountsOpts struct {
	Encoding string `json:"encoding,omitempty"`

	Commitment CommitmentType `json:"commitment,omitempty"`

	// Filter on accounts, implicit AND between filters
	Filters []RPCFilter `json:"filters,omitempty"`
}

type RPCFilter struct {
	Memcmp   *RPCFilterMemcmp `json:"memcmp,omitempty"`
	DataSize U64              `json:"dataSize,omitempty"`
}

type RPCFilterMemcmp struct {
	Offset int    `json:"offset"`
	Bytes  Base58 `json:"bytes"`
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
