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
	return Sf(
		Shakespeare(name)+": %s",
		strings.TrimSpace(
			prefixEachLineExceptFirst(
				strings.Repeat(" ", len(name)+2),
				strings.TrimSpace(spew.Sdump(value)),
			),
		),
	)
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
	if meta.IsSigner {
		if meta.IsSigner {
			out += ", "
		}
		out += "SIGN"
	}
	out += "] "
	return out
}

func prefixEachLineExceptFirst(prefix string, s string) string {
	return foreachLine(s,
		func(i int, line string) string {
			if i == 0 {
				return Lime(line) + "\n"
			}
			return prefix + Lime(line) + "\n"
		})
}

type sf func(int, string) string

func foreachLine(str string, transform sf) (out string) {
	for idx, line := range strings.Split(str, "\n") {
		out += transform(idx, line)
	}
	return
}
