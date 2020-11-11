package ws

import (
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
)

var base64Conf = map[string]interface{}{
	"encoding": "base64",
}

func (c *Client) ProgramSubscribe(programId solana.PublicKey, commitment rpc.CommitmentType) (*Subscription, error) {
	return c.subscribe([]interface{}{programId.String()}, base64Conf, "programSubscribe", "programUnsubscribe", commitment, ProgramResult{})
}

func (c *Client) AccountSubscribe(account solana.PublicKey, commitment rpc.CommitmentType) (*Subscription, error) {
	return c.subscribe([]interface{}{account.String()}, base64Conf, "accountSubscribe", "accountUnsubscribe", commitment, AccountResult{})
}

func (c *Client) SlotSubscribe() (*Subscription, error) {
	return c.subscribe(nil, nil, "slotSubscribe", "slotUnsubscribe", "", SlotResult{})
}
