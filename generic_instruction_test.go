package solana

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInstruction(t *testing.T) {
	progID := MemoProgramID
	accounts := []*AccountMeta{
		Meta(SPLAssociatedTokenAccountProgramID).SIGNER().WRITE(),
	}
	data := []byte("hello world")

	ins := NewInstruction(progID, accounts, data)

	require.Equal(t, progID, ins.ProgramID())
	require.Equal(t, accounts, ins.Accounts())
	{
		got, err := ins.Data()
		require.NoError(t, err)
		require.Equal(t, data, got)
	}
}
