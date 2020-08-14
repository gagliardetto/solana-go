package solana

import (
	"bytes"
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
		{input: 0x7f, expect: []byte{0x7f}},
		{input: 0x80, expect: []byte{0x80, 0x01}},
		{input: 0xff, expect: []byte{0xff, 0x01}},
		{input: 0x100, expect: []byte{0x80, 0x02}},
		{input: 0x7fff, expect: []byte{0xff, 0xff, 0x01}},
		{input: 0xffff, expect: []byte{0xff, 0xff, 0x03}},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("%d", idx+1), func(t *testing.T) {
			el := ShortVec(test.input)

			b := make([]byte, 3)

			size, err := el.Pack(b, nil)
			require.NoError(t, err)
			assert.Equal(t, test.expect, b[:size])

			buf := bytes.NewBuffer(b)
			target := ShortVec(0)
			require.NoError(t, (&target).Unpack(buf, 0, nil))
			assert.Equal(t, test.input, uint16(target))
		})
	}
}
