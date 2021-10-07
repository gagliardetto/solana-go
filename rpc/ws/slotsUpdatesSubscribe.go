// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ws

import "github.com/gagliardetto/solana-go"

type SlotsUpdatesResult struct {
	// The parent slot.
	Parent uint64 `json:"parent"`
	// The newly updated slot.
	Slot uint64 `json:"slot"`
	// The Unix timestamp of the update.
	Timestamp *solana.UnixTimeSeconds `json:"timestamp"`
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
