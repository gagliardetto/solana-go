package ws

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type AccountResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		rpc.Account
	} `json:"value"`
}

// AccountSubscribe subscribes to an account to receive notifications
// when the lamports or data for a given account public key changes.
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

// AccountSubscribe subscribes to an account to receive notifications
// when the lamports or data for a given account public key changes.
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
