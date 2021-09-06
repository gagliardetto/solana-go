package sendandconfirmtransaction

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func SendAndConfirmTransactionWithOpts(
	ctx context.Context,
	rpcClient *rpc.Client,
	wsClient *ws.Client,
	transaction *solana.Transaction,
	skipPreflight bool, // if true, skip the preflight transaction checks (default: false)
	preflightCommitment rpc.CommitmentType, // optional; Commitment level to use for preflight (default: "finalized").
) (signature solana.Signature, err error) {

	sig, err := rpcClient.SendTransactionWithOpts(
		ctx,
		transaction,
		skipPreflight,
		preflightCommitment,
	)
	if err != nil {
		return sig, err
	}

	sub, err := wsClient.SignatureSubscribe(
		sig,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			return sig, err
		}
		spew.Dump(got)
		if got.Value.Err != nil {
			return sig, fmt.Errorf("confirmation error: %v", got.Value.Err)
		} else {
			return sig, nil
		}
	}
}
