package vote

import (
	"encoding/binary"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/stretchr/testify/require"
)

func TestVoteStateUpdate_UnmarshalRejectsOversizedCount(t *testing.T) {
	// A huge declared lockout count with no payload must return an error, not panic.
	for _, count := range []uint64{1, 1 << 40, ^uint64(0)} {
		data := make([]byte, 8)
		binary.LittleEndian.PutUint64(data, count)
		var u VoteStateUpdate
		require.NotPanics(t, func() {
			err := u.UnmarshalWithDecoder(bin.NewBinDecoder(data))
			require.Error(t, err)
		})
	}
}
