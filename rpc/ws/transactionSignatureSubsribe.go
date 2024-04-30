package ws

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var ErrInvalidParams = fmt.Errorf("invalid params")

type TransactionSignatureSubscription struct {
	sub *Subscription
}

type TransactionSignatureResult struct {
	Transaction struct {
		Meta struct {
			LogMessages []string `json:"logMessages"`
		} `json:"meta"`
	} `json:"transaction"`
	Signature solana.Signature `json:"signature"`
}

// TransactionSignatureSubscribe subscribes to a transaction signature. Only Helius rpc nodes support this method.
func (cl *Client) TransactionSignatureSubscribe(
	accountInclude []string,
	accountRequired []string,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscription, error) {
	params := make([]any, 0, 1)
	param := rpc.M{}
	if len(accountInclude) > 0 {
		param["accountInclude"] = accountInclude
	}
	if len(accountRequired) > 0 {
		param["accountRequired"] = accountRequired
	}
	if len(param) == 0 {
		return nil, ErrInvalidParams
	}
	param["vote"] = false
	param["failed"] = false
	params = append(params, param)
	return cl.transactionSignatureSubscribe(
		params,
		commitment,
	)
}

func (cl *Client) transactionSignatureSubscribe(
	params []any,
	commitment rpc.CommitmentType,
) (*TransactionSignatureSubscription, error) {
	conf := map[string]interface{}{}
	conf["transaction_details"] = "full"
	conf["maxSupportedTransactionVersion"] = 0
	if commitment != "" {
		conf["commitment"] = commitment
	}
	genSub, err := cl.subscribe(
		params,
		conf,
		"transactionSubscribe",
		"transactionUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res TransactionSignatureResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)

	if err != nil {
		return nil, err
	}
	return &TransactionSignatureSubscription{
		sub: genSub,
	}, nil
}

func (sw *TransactionSignatureSubscription) Recv() (*LogResult, error) {
	select {
	case d, ok := <-sw.sub.stream:
		if !ok {
			if !ok {
				return nil, ErrSubscriptionClosed
			}
		}
		logResult := &LogResult{}
		logResult.Value.Logs = d.(*TransactionSignatureResult).Transaction.Meta.LogMessages
		logResult.Value.Signature = d.(*TransactionSignatureResult).Signature
		return logResult, nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *TransactionSignatureSubscription) Err() <-chan error {
	return sw.sub.err
}

func (sw *TransactionSignatureSubscription) Response() <-chan *LogResult {
	typedChan := make(chan *LogResult, 1)
	go func(ch chan *LogResult) {
		// TODO: will this subscription yield more than one result?
		d, ok := <-sw.sub.stream
		if !ok {
			return
		}
		logResult := &LogResult{}
		logResult.Value.Logs = d.(*TransactionSignatureResult).Transaction.Meta.LogMessages
		logResult.Value.Signature = d.(*TransactionSignatureResult).Signature
		ch <- logResult
	}(typedChan)
	return typedChan
}

func (sw *TransactionSignatureSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}
