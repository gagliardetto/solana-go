package serum

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	bin "github.com/dfuse-io/binary"
	"github.com/stretchr/testify/require"
)

func TestDecodeInstruction(t *testing.T) {
	x := `00020000000500`
	data, err := hex.DecodeString(x)
	require.NoError(t, err)
	var instruction *Instruction
	err = bin.NewDecoder(data).Decode(&instruction)
	require.NoError(t, err)
	assert.Equal(t, instruction.Version, uint8(0))
}
