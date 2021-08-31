package format

import (
	"strings"

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
	return Sf(CC(Shakespeare(name), ": %s"), Lime(strings.TrimSpace(spew.Sdump(value))))
}

func Account(name string, pubKey solana.PublicKey) string {
	return Shakespeare(name) + ": " + text.ColorizeBG(pubKey.String())
}

func Meta(name string, meta *solana.AccountMeta) string {
	out := Shakespeare(name) + ": " + text.ColorizeBG(meta.PublicKey.String())
	out += " ["
	if meta.IsWritable {
		out += "WRITE"
	}
	if meta.IsWritable {
		if meta.IsWritable {
			out += ", "
		}
		out += "SIGN"
	}
	out += "] "
	return out
}
