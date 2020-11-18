package serum

import (
	"encoding/hex"
	"fmt"
	"testing"

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
	fmt.Println(instruction)
}

func TestString(t *testing.T) {

}
