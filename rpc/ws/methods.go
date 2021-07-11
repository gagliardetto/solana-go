package ws

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// AccountSubscribe subscribes to an account to receive notifications when the lamports or data for a given account public key changes.
func (cl *Client) AccountSubscribe(
	account solana.PublicKey,
	commitment rpc.CommitmentType,
) (*AccountSubscription, error) {
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
	encoding solana.EncodingType,
) (*AccountSubscription, error) {

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

	genSub, err := cl.subscribe(
		params,
		conf,
		"accountSubscribe",
		"accountUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res AccountResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &AccountSubscription{
		sub: genSub,
	}, nil
}

type AccountSubscription struct {
	sub *Subscription
}

func (sw *AccountSubscription) Recv() (*AccountResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*AccountResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *AccountSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
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

// ProgramSubscribe subscribes to a program to receive notifications
// when the lamports or data for a given account owned by the program changes.
func (cl *Client) ProgramSubscribe(
	programID solana.PublicKey,
	commitment rpc.CommitmentType,
) (*ProgramSubscription, error) {
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
	encoding solana.EncodingType,
	filters []rpc.RPCFilter,
) (*ProgramSubscription, error) {

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

	genSub, err := cl.subscribe(
		params,
		conf,
		"programSubscribe",
		"programUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res ProgramResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &ProgramSubscription{
		sub: genSub,
	}, nil
}

type ProgramSubscription struct {
	sub *Subscription
}

func (sw *ProgramSubscription) Recv() (*ProgramResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*ProgramResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *ProgramSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}

// SignatureSubscribe subscribes to a transaction signature to receive
// notification when the transaction is confirmed On signatureNotification,
// the subscription is automatically cancelled
func (cl *Client) SignatureSubscribe(
	signature solana.Signature,
	commitment rpc.CommitmentType,
) (*SignatureSubscription, error) {

	params := []interface{}{signature.String()}
	conf := map[string]interface{}{}
	if commitment != "" {
		conf["commitment"] = commitment
	}

	genSub, err := cl.subscribe(
		params,
		conf,
		"signatureSubscribe",
		"signatureUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res SignatureResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &SignatureSubscription{
		sub: genSub,
	}, nil
}

type SignatureSubscription struct {
	sub *Subscription
}

func (sw *SignatureSubscription) Recv() (*SignatureResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*SignatureResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *SignatureSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}

// SlotSubscribe subscribes to receive notification anytime a slot is processed by the validator.
func (cl *Client) SlotSubscribe() (*SlotSubscription, error) {
	genSub, err := cl.subscribe(
		nil,
		nil,
		"slotSubscribe",
		"slotUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res SlotResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &SlotSubscription{
		sub: genSub,
	}, nil
}

type SlotSubscription struct {
	sub *Subscription
}

func (sw *SlotSubscription) Recv() (*SlotResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*SlotResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *SlotSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}

// SignatureSubscribe subscribes to receive notification
// anytime a new root is set by the validator.
func (cl *Client) RootSubscribe() (*RootSubscription, error) {
	genSub, err := cl.subscribe(
		nil,
		nil,
		"rootSubscribe",
		"rootUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res RootResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &RootSubscription{
		sub: genSub,
	}, nil
}

type RootSubscription struct {
	sub *Subscription
}

func (sw *RootSubscription) Recv() (*RootResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*RootResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *RootSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}

// VoteSubscribe (UNSTABLE) subscribes to receive notification anytime
// a new vote is observed in gossip.
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
