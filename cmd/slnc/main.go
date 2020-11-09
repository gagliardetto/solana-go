package main

import (
	"github.com/dfuse-io/solana-go/cmd/slnc/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
