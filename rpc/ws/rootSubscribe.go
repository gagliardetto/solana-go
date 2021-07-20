package ws

import bin "github.com/dfuse-io/binary"

type RootResult bin.Uint64

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
