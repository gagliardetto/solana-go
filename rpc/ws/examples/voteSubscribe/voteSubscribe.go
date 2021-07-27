package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {
	client, err := ws.Connect(context.Background(), rpc.EndpointWS_MainNetBeta)
	if err != nil {
		panic(err)
	}

	// NOTE: this subscription must be enabled by the node you're connecting to.
	// This subscription is disabled by default.
	sub, err := client.VoteSubscribe()
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
}
