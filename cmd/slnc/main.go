package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/dfuse-io/solana-go"
)

func main() {
	ctx := context.Background()

	c := solana.NewClient("http://api.mainnet-beta.solana.com/rpc")
	resp, err := c.GetAccountInfo(ctx, "14e9wAw5bMKUZC9vmV3t6axNhpiJsNWBXJBH2xfzFsws", "")
	if err != nil {
		log.Fatalln("failed", err)
	}

	spew.Dump(resp)
	fmt.Println("Resp:", resp.Value)
}
