package ws

import (
	"context"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// Block until a transaction has achieved a specified level of commitment
func (cl *Client) WaitSig(ctx context.Context, sig solana.Signature, commitment rpc.CommitmentType) error {
	sub, err := cl.SignatureSubscribe(sig, commitment)
	if err != nil {
		return err
	}
	streamC := sub.RecvStream()
	doneC := ctx.Done()
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
			x, ok := d.(*SignatureResult)
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
