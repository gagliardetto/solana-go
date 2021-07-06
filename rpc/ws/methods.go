package ws

import (
	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// AccountSubscribe subscribes to an account to receive notifications when the lamports or data for a given account public key changes.
func (cl *Client) AccountSubscribe(
	account solana.PublicKey,
	commitment rpc.CommitmentType,
) (*Subscription, error) {
	return cl.AccountSubscribeWithOpts(
		account,
		commitment,
		"",
	)
}

// AccountSubscribe subscribes to an account to receive notifications when the lamports or data for a given account public key changes.
func (cl *Client) AccountSubscribeWithOpts(
	account solana.PublicKey,
	commitment rpc.CommitmentType,
	encoding rpc.EncodingType,
) (*Subscription, error) {

	params := []interface{}{account.String()}
	conf := map[string]interface{}{
		"encoding": "base64",
	}
	if commitment != "" {
		conf["commitment"] = commitment
	}
	if encoding != "" {
		conf["encoding"] = encoding
	}

	return cl.subscribe(
		params,
		conf,
		"accountSubscribe",
		"accountUnsubscribe",
		AccountResult{},
	)
}

type LogsSubscribeFilterType string

const (
	LogsSubscribeFilterAll          LogsSubscribeFilterType = "all"          // subscribe to all transactions except for simple vote transactions
	LogsSubscribeFilterAllWithVotes LogsSubscribeFilterType = "allWithVotes" // subscribe to all transactions including simple vote transactions
)

// LogsSubscribe subscribes to transaction logging.
func (cl *Client) LogsSubscribe(
	filter LogsSubscribeFilterType,
	commitment rpc.CommitmentType,
) (*Subscription, error) {
	return cl.logsSubscribe(
		filter,
		commitment,
	)
}

// LogsSubscribe subscribes to all transactions that mention the provided Pubkey.
func (cl *Client) LogsSubscribeMentions(
	mentions solana.PublicKey,
	commitment rpc.CommitmentType,
) (*Subscription, error) {
	return cl.logsSubscribe(
		rpc.M{
			"mentions": []string{mentions.String()},
		},
		commitment,
	)
}

// LogsSubscribe subscribes to transaction logging.
func (cl *Client) logsSubscribe(
	filter interface{},
	commitment rpc.CommitmentType,
) (*Subscription, error) {

	params := []interface{}{filter}
	conf := map[string]interface{}{}
	if commitment != "" {
		conf["commitment"] = commitment
	}

	return cl.subscribe(
		params,
		conf,
		"logsSubscribe",
		"logsUnsubscribe",
		LogResult{},
	)
}

// ProgramSubscribe subscribes to a program to receive notifications
// when the lamports or data for a given account owned by the program changes.
func (cl *Client) ProgramSubscribe(
	programID solana.PublicKey,
	commitment rpc.CommitmentType,
) (*Subscription, error) {
	return cl.ProgramSubscribeWithOpts(
		programID,
		commitment,
		"",
		nil,
	)
}

// ProgramSubscribe subscribes to a program to receive notifications
// when the lamports or data for a given account owned by the program changes.
func (cl *Client) ProgramSubscribeWithOpts(
	programID solana.PublicKey,
	commitment rpc.CommitmentType,
	encoding rpc.EncodingType,
	filters []rpc.RPCFilter,
) (*Subscription, error) {

	params := []interface{}{programID.String()}
	conf := map[string]interface{}{
		"encoding": "base64",
	}
	if commitment != "" {
		conf["commitment"] = commitment
	}
	if encoding != "" {
		conf["encoding"] = encoding
	}
	if filters != nil && len(filters) > 0 {
		conf["filters"] = filters
	}

	return cl.subscribe(
		params,
		conf,
		"programSubscribe",
		"programUnsubscribe",
		ProgramResult{},
	)
}

// SignatureSubscribe subscribes to a transaction signature to receive
// notification when the transaction is confirmed On signatureNotification,
// the subscription is automatically cancelled
func (cl *Client) SignatureSubscribe(
	signature solana.Signature,
	commitment rpc.CommitmentType,
) (*Subscription, error) {

	params := []interface{}{signature.String()}
	conf := map[string]interface{}{}
	if commitment != "" {
		conf["commitment"] = commitment
	}

	return cl.subscribe(
		params,
		conf,
		"signatureSubscribe",
		"signatureUnsubscribe",
		SignatureResult{},
	)
}

// SlotSubscribe subscribes to receive notification anytime a slot is processed by the validator.
func (cl *Client) SlotSubscribe() (*Subscription, error) {
	return cl.subscribe(
		nil,
		nil,
		"slotSubscribe",
		"slotUnsubscribe",
		SlotResult{},
	)
}

// SignatureSubscribe subscribes to receive notification
// anytime a new root is set by the validator.
func (cl *Client) RootSubscribe() (*Subscription, error) {
	return cl.subscribe(
		nil,
		nil,
		"rootSubscribe",
		"rootUnsubscribe",
		RootResult(bin.Uint64(1)),
	)
}

// VoteSubscribe (UNSTABLE) subscribes to receive notification anytime
// a new vote is observed in gossip.
// These votes are pre-consensus therefore there is
// no guarantee these votes will enter the ledger.
//
// This subscription is unstable and only available if the validator
// was started with the --rpc-pubsub-enable-vote-subscription flag.
// The format of this subscription may change in the future.
func (cl *Client) VoteSubscribe() (*Subscription, error) {
	return cl.subscribe(
		nil,
		nil,
		"voteSubscribe",
		"voteUnsubscribe",
		VoteResult{},
	)
}
