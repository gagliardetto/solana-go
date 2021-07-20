package ws

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type ProgramResult struct {
	Context struct {
		Slot uint64
	} `json:"context"`
	Value rpc.KeyedAccount `json:"value"`
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
