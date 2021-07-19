package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type GetVoteAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	// (optional) Only return results for this validator vote address.
	VotePubkey solana.PublicKey `json:"votePubkey,omitempty"`
}

// GetVoteAccounts returns the account info and associated
// stake for all the voting accounts in the current bank.
func (cl *Client) GetVoteAccounts(
	ctx context.Context,
	opts *GetVoteAccountsOpts,
) (out *GetVoteAccountsResult, err error) {
	params := []interface{}{}
	if opts != nil {
		obj := M{}
		if opts.Commitment != "" {
			obj["commitment"] = string(opts.Commitment)
		}
		if !opts.VotePubkey.IsZero() {
			obj["votePubkey"] = opts.VotePubkey.String()
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	err = cl.rpcClient.CallFor(&out, "getVoteAccounts", params)
	return
}

type GetVoteAccountsResult struct {
	Current    []VoteAccountsResult `json:"current"`
	Delinquent []VoteAccountsResult `json:"delinquent"`
}

type VoteAccountsResult struct {
	// Vote account address.
	VotePubkey solana.PublicKey `json:"votePubkey,omitempty"`

	// Validator identity.
	NodePubkey solana.PublicKey `json:"nodePubkey,omitempty"`

	// The stake, in lamports, delegated to this vote account and active in this epoch.
	ActivatedStake uint64 `json:"activatedStake,omitempty"`

	// Whether the vote account is staked for this epoch.
	EpochVoteAccount bool `json:"epochVoteAccount,omitempty"`

	// Percentage (0-100) of rewards payout owed to the vote account.
	Commission uint8 `json:"commission,omitempty"`

	// Most recent slot voted on by this vote account.
	LastVote uint64 `json:"lastVote,omitempty"`

	RootSlot uint64 `json:"rootSlot,omitempty"` //

	// History of how many credits earned by the end of each epoch,
	// as an array of arrays containing: [epoch, credits, previousCredits]
	EpochCredits [][]int64 `json:"epochCredits,omitempty"`
}
