package ws

import (
	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type VoteResult struct {
	// The vote hash.
	Hash solana.Hash `json:"hash"`
	// The slots covered by the vote.
	Slots []bin.Uint64 `json:"slots"`
	// The timestamp of the vote.
	Timestamp *rpc.UnixTimeSeconds `json:"timestamp,omitempty"`
}

// VoteSubscribe (UNSTABLE, disabled by default) subscribes
// to receive notification anytime a new vote is observed in gossip.
// These votes are pre-consensus therefore there is
// no guarantee these votes will enter the ledger.
//
// This subscription is unstable and only available if the validator
// was started with the --rpc-pubsub-enable-vote-subscription flag.
// The format of this subscription may change in the future.
func (cl *Client) VoteSubscribe() (*VoteSubscription, error) {
	genSub, err := cl.subscribe(
		nil,
		nil,
		"voteSubscribe",
		"voteUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res VoteResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &VoteSubscription{
		sub: genSub,
	}, nil
}

type VoteSubscription struct {
	sub *Subscription
}

func (sw *VoteSubscription) Recv() (*VoteResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*VoteResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *VoteSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}
