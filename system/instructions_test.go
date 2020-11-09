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

package system

import (
	"bytes"
	"testing"

	"github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSystemInstructions(t *testing.T) {
	t.Skip()
	ins1 := &SystemInstruction{
		Variant: &CreateAccount{
			Lamports: 125,
			Space:    120,
			Owner:    solana.MustPublicKeyFromBase58("4JuGp6UkTewQXG1tJpYY1dxW1H9yS6sSeCDc1FSdWKNR"),
		}}

	buf := &bytes.Buffer{}
	err := struc.Pack(buf, ins1)
	require.NoError(t, err)
	assert.Equal(t, []byte{0, 1, 2, 3}, buf.Bytes())

	out := SystemInstruction{}

	require.NoError(t, struc.Unpack(bytes.NewReader(buf.Bytes()), &out))
	assert.Equal(t, "hello", out.String(), out.Variant)

	// tests := []struct{
	//     name string
	//     input string
	//     expect string
	// }{
	//     {
	//         name: "name",
	//         input: "input",
	//         expect: "expect",
	//     },
	// }

	// for _, test := range tests {
	//     t.Run(test.name, func(t *testing.T) {
	//         res := (test.in)
	//         assert.Equal(t, test.expect, res)
	//     })
	// }
}
