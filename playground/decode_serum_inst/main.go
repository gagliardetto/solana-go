package main

import (
	"encoding/hex"
	"os"

	"github.com/dfuse-io/solana-go/programs/serum"
	"github.com/dfuse-io/solana-go/text"

	bin "github.com/dfuse-io/binary"
)

func main() {
	data, err := hex.DecodeString("00020000000f00")
	if err != nil {
		panic(err)
	}

	var instruction *serum.Instruction
	err = bin.NewDecoder(data).Decode(&instruction)

	text.NewEncoder(os.Stdout).Encode(instruction, nil)

}
