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

func NewTreeEncoder(w io.Writer, docs ...string) *TreeEncoder {
	return &TreeEncoder{
		output: w,
		Tree:   treeout.New(docs...),
	}
}

func (e *TreeEncoder) WriteString(s string) error {
	_, err := e.output.Write([]byte(s))
	return err
}
