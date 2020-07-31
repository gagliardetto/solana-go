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

type GetBalanceRPCResult struct {
	RPCContext
	Value U64 `json:"value"`
}

///

type GetAccountInfoRPCResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	Lamports   U64       `json:"lamports"`
	Data       Base58    `json:"data"`
	Owner      PublicKey `json:"owner"`
	Executable bool      `json:"executable"`
	RentEpoch  U64       `json:"rentEpoch"`
}

type KeyedAccount struct {
	Pubkey  PublicKey `json:"pubkey"`
	Account *Account  `json:"account"`
}

///

type GetProgramAccountsRPCResult []*KeyedAccount

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
