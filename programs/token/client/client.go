package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type Client struct {
	ctx        context.Context
	ws         *ws.Client
	rpc        *rpc.Client
	commitment rpc.CommitmentType
}

func Create(ctx context.Context, rpcClient *rpc.Client, wsClient *ws.Client, commitment rpc.CommitmentType) *Client {
	return &Client{ws: wsClient, rpc: rpcClient, ctx: ctx}
}

func (c *Client) latest_blockhash() (solana.Hash, error) {
	latestBlockhashResult, err := c.rpc.GetLatestBlockhash(c.ctx, c.commitment)
	if err != nil {
		return solana.Hash{}, err
	}

	if latestBlockhashResult.Value == nil {
		return solana.Hash{}, errors.New("blank block hash")
	}

	return latestBlockhashResult.Value.Blockhash, nil
}

// block until we get notification from a validator that the transaction has reached c.commitment state
func (c *Client) block_on_tx_processing(sig solana.Signature) error {
	sub, err := c.ws.SignatureSubscribe(sig, c.commitment)
	if err != nil {
		return err
	}
	streamC := sub.RecvStream()
	doneC := c.ctx.Done()
	errorC := sub.CloseSignal()

	defer sub.Unsubscribe()

out:
	for {
		select {
		case <-doneC:
			break out
		case err = <-errorC:
			break out
		case d := <-streamC:
			x, ok := d.(*ws.SignatureResult)
			if !ok {
				err = errors.New("bad type")
			}

			if x.Value.Err != nil {
				err = fmt.Errorf("%+v", x.Value.Err)
			}

			break out
		}
	}

	return err
}
