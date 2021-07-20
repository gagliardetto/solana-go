package ws

type SlotResult struct {
	Parent uint64 `json:"parent"`
	Root   uint64 `json:"root"`
	Slot   uint64 `json:"slot"`
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
