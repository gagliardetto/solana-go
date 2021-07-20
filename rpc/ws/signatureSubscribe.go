package ws

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type SignatureResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value struct {
		Err interface{} `json:"err"`
	} `json:"value"`
}

// SignatureSubscribe subscribes to a transaction signature to receive
// notification when the transaction is confirmed On signatureNotification,
// the subscription is automatically cancelled
func (cl *Client) SignatureSubscribe(
	signature solana.Signature, // Transaction Signature.
	commitment rpc.CommitmentType, // (optional)
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
