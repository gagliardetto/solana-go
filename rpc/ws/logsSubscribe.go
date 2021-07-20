package ws

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type LogResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		// The transaction signature.
		Signature solana.Signature `json:"signature"`
		// Error if transaction failed, null if transaction succeeded.
		Err interface{} `json:"err"`
		// Array of log messages the transaction instructions output
		// during execution, null if simulation failed before the transaction
		// was able to execute (for example due to an invalid blockhash
		// or signature verification failure)
		Logs []string `json:"logs"`
	} `json:"value"`
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
) (*LogSubscription, error) {
	return cl.logsSubscribe(
		filter,
		commitment,
	)
}

// LogsSubscribe subscribes to all transactions that mention the provided Pubkey.
func (cl *Client) LogsSubscribeMentions(
	mentions solana.PublicKey,
	commitment rpc.CommitmentType,
) (*LogSubscription, error) {
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
) (*LogSubscription, error) {

	params := []interface{}{filter}
	conf := map[string]interface{}{}
	if commitment != "" {
		conf["commitment"] = commitment
	}

	genSub, err := cl.subscribe(
		params,
		conf,
		"logsSubscribe",
		"logsUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res LogResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &LogSubscription{
		sub: genSub,
	}, nil
}

type LogSubscription struct {
	sub *Subscription
}

func (sw *LogSubscription) Recv() (*LogResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*LogResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *LogSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}
