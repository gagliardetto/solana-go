package tokenregistry

import (
	"testing"

	bin "github.com/dfuse-io/binary"
	"github.com/stretchr/testify/require"
)

func TestAsciiBytes_String(t *testing.T) {
	require.Equal(t, "ABC", AsciiString([]byte{65, 66, 67, 0, 0, 0, 0, 0}))
}

func TestAsciiBytes_Bin_Decode(t *testing.T) {
	d := []byte{
		108, 111, 103, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		110, 97, 109, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		115, 121, 109, 98, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	var m *TokenMeta

	err := bin.NewDecoder(d).Decode(&m)
	require.NoError(t, err)

	require.Equal(t, "logo", m.Logo.String())
	require.Equal(t, "name", m.Name.String())
	require.Equal(t, "symb", m.Symbol.String())
}
