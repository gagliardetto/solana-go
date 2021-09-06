package text

import (
	"io"

	"github.com/gagliardetto/treeout"
)

type TreeEncoder struct {
	output io.Writer
	*treeout.Tree
}

type EncodableToTree interface {
	EncodeToTree(parent treeout.Branches)
}

func NewTreeEncoder(w io.Writer, doc string) *TreeEncoder {
	return &TreeEncoder{
		output: w,
		Tree:   treeout.New(doc),
	}
}

func (enc *TreeEncoder) WriteString(s string) (int, error) {
	return enc.output.Write([]byte(s))
}
