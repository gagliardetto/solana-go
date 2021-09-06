package rpc

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func (cl *Client) SendAndConfirmTransactionWithOpts(
	ctx context.Context,
	transaction *solana.Transaction,
	skipPreflight bool, // if true, skip the preflight transaction checks (default: false)
	preflightCommitment CommitmentType, // optional; Commitment level to use for preflight (default: "finalized").
) (signature solana.Signature, err error) {

	sig, err := cl.SendTransactionWithOpts(
		ctx,
		transaction,
		skipPreflight,
		preflightCommitment,
	)
	if err != nil {
		return sig, err
	}

	client, err := ws.Connect(context.Background(), TestNet_WS)
	if err != nil {
		panic(err)
	}

	txSig := solana.MustSignatureFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	sub, err := client.SignatureSubscribe(
		txSig,
		"",
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			panic(err)
		}
		spew.Dump(got)
	}
	return
}
