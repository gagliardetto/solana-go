package format

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
	. "github.com/gagliardetto/solana-go/text"
)

func Program(name string, programID solana.PublicKey) string {
	return IndigoBG("Program") + ": " + Bold(name) + " " + text.ColorizeBG(programID.String())
}

func Instruction(name string) string {
	return Purple(Bold("Instruction")) + ": " + Bold(name)
}

func Param(name string, value interface{}) string {
	return Sf(CC(Shakespeare(name), ": %s"), Lime(spew.Sdump(value)))
}

func Account(name string, pubKey solana.PublicKey) string {
	return Shakespeare(name) + ": " + text.ColorizeBG(pubKey.String())
}
