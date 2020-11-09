// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solana

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVaruint16(t *testing.T) {
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
			el := Varuint16(test.input)

			b := make([]byte, 3)

			size, err := el.Pack(b, nil)
			require.NoError(t, err)
			assert.Equal(t, test.expect, b[:size])

			buf := bytes.NewBuffer(b)
			target := Varuint16(0)
			require.NoError(t, (&target).Unpack(buf, 0, nil))
			assert.Equal(t, test.input, uint16(target))
		})
	}
}
