package ws

import "github.com/gagliardetto/solana-go/rpc"

type SlotsUpdatesResult struct {
	// The parent slot.
	Parent uint64 `json:"parent"`
	// The newly updated slot.
	Slot uint64 `json:"slot"`
	// The Unix timestamp of the update.
	Timestamp *rpc.UnixTimeSeconds `json:"timestamp"`
	// The update type.
	Type SlotsUpdatesType `json:"type"`
}

type SlotsUpdatesType string

const (
	SlotsUpdatesFirstShredReceived     SlotsUpdatesType = "firstShredReceived"
	SlotsUpdatesCompleted              SlotsUpdatesType = "completed"
	SlotsUpdatesCreatedBank            SlotsUpdatesType = "createdBank"
	SlotsUpdatesFrozen                 SlotsUpdatesType = "frozen"
	SlotsUpdatesDead                   SlotsUpdatesType = "dead"
	SlotsUpdatesOptimisticConfirmation SlotsUpdatesType = "optimisticConfirmation"
	SlotsUpdatesRoot                   SlotsUpdatesType = "root"
)

// SlotsUpdatesSubscribe (UNSTABLE) subscribes to receive a notification
// from the validator on a variety of updates on every slot.
//
// This subscription is unstable; the format of this subscription
// may change in the future and it may not always be supported.
func (cl *Client) SlotsUpdatesSubscribe() (*SlotsUpdatesSubscription, error) {
	genSub, err := cl.subscribe(
		nil,
		nil,
		"slotsUpdatesSubscribe",
		"slotsUpdatesUnsubscribe",
		func(msg []byte) (interface{}, error) {
			var res SlotsUpdatesResult
			err := decodeResponseFromMessage(msg, &res)
			return &res, err
		},
	)
	if err != nil {
		return nil, err
	}
	return &SlotsUpdatesSubscription{
		sub: genSub,
	}, nil
}

type SlotsUpdatesSubscription struct {
	sub *Subscription
}

func (sw *SlotsUpdatesSubscription) Recv() (*SlotsUpdatesResult, error) {
	select {
	case d := <-sw.sub.stream:
		return d.(*SlotsUpdatesResult), nil
	case err := <-sw.sub.err:
		return nil, err
	}
}

func (sw *SlotsUpdatesSubscription) Unsubscribe() {
	sw.sub.Unsubscribe()
}
