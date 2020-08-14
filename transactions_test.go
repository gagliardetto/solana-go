package solana

import (
	"bytes"
	"testing"

	"github.com/lunixbochs/struc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompiledInstructions(t *testing.T) {
	ci := &CompiledInstruction{
		ProgramIDIndex: 5,
		Accounts:       []uint8{2, 5, 8},
		Data:           Base58([]byte{1, 2, 3, 4, 5}),
	}
	buf := &bytes.Buffer{}
	require.NoError(t, struc.Pack(buf, ci))
	assert.Equal(t, []byte{5, 3, 2, 5, 8, 5, 1, 2, 3, 4, 5}, buf.Bytes())
}
