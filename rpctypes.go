package solana

type ContactInfo struct {
	Pubkey  string `json:"pubkey"`
	Gossip  string `json:"gossip,omitempty"`
	TPU     string `json:"tpu,omitempty"`
	RPC     string `json:"rpc,omitempty"`
	Version string `json:"version,omitempty"`
}

type RPCContext struct {
	Context struct {
		Slot U64
	} `json:"context,omitempty"`
}
type GetBalanceRPCResult struct {
	RPCContext
	Value U64 `json:"value"`
}

type GetAccountInfoRPCResult struct {
	RPCContext
	Value *AccountInfoResult `json:"value"`
}
type AccountInfoResult struct {
	Executable bool   `json:"executable"`
	Owner      string `json:"owner"`
	Lamports   U64    `json:"lamports"`
	Data       Base58 `json:"data"`
	RentEpoch  U64    `json:"rentEpoch"`
}

type CommitmentType string

const (
	CommitmentMax          = CommitmentType("max")
	CommitmentRecent       = CommitmentType("recent")
	CommitmentRoot         = CommitmentType("root")
	CommitmentSingle       = CommitmentType("single")
	CommitmentSingleGossip = CommitmentType("singleGossip")
)
