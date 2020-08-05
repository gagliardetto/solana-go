package solana

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortVec(t *testing.T) {
	tests := []struct {
		input  uint16
		expect []byte
	}{
		{input: 0x0, expect: []byte{0x0}},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			el := ShortVec(test.input)

			b := make([]byte, 3)

			size, err := el.Pack(b, nil)
			require.NoError(t, err)
			assert.Equal(t, test.expect, b[:size])
		})
	}
}
