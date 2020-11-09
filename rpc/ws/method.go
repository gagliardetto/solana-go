package ws

import (
	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/rpc"
)

func (c *Client) ProgramSubscribe(programId solana.PublicKey, commitment rpc.CommitmentType) (*Subscription, error) {
	return c.subscribe([]interface{}{programId.String()}, "programSubscribe", "programUnsubscribe", commitment, ProgramResult{})
}

func (c *Client) AccountSubscribe(account solana.PublicKey, commitment rpc.CommitmentType) (*Subscription, error) {
	return c.subscribe([]interface{}{account.String()}, "accountSubscribe", "accountUnsubscribe", commitment, AccountResult{})
}
