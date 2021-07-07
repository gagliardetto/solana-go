package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type GetVoteAccountsResult struct {
	Current    []VoteAccountsResult `json:"current"`
	Delinquent []VoteAccountsResult `json:"delinquent"`
}
type VoteAccountsResult struct {
	VotePubkey       solana.PublicKey `json:"votePubkey,omitempty"`       // Vote account address
	NodePubkey       solana.PublicKey `json:"nodePubkey,omitempty"`       // Validator identity
	ActivatedStake   uint64           `json:"activatedStake,omitempty"`   // the stake, in lamports, delegated to this vote account and active in this epoch
	EpochVoteAccount bool             `json:"epochVoteAccount,omitempty"` // whether the vote account is staked for this epoch
	Commission       uint8            `json:"commission,omitempty"`       // percentage (0-100) of rewards payout owed to the vote account
	LastVote         uint64           `json:"lastVote,omitempty"`         // Most recent slot voted on by this vote account
	RootSlot         uint64           `json:"rootSlot,omitempty"`         //

	// History of how many credits earned by the end of each epoch,
	// as an array of arrays containing: [epoch, credits, previousCredits]
	EpochCredits [][]int64 `json:"epochCredits,omitempty"`
}

type GetVoteAccountsOpts struct {
	Commitment CommitmentType   `json:"commitment,omitempty"`
	VotePubkey solana.PublicKey `json:"votePubkey,omitempty"` // (optional) Only return results for this validator vote address
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
			obj["votePubkey"] = opts.VotePubkey
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	err = cl.rpcClient.CallFor(&out, "getVoteAccounts", params)
	return
}
