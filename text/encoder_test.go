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

package text

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nested struct {
	F1 string
	F2 string
}

func TestEncoder_TestStruct(t *testing.T) {
	s := &binaryTestStruct{
		F1:  "abc",
		F2:  -75,
		F3:  99,
		F4:  -231,
		F5:  999,
		F6:  -13231,
		F7:  99999,
		F8:  -23.13,
		F9:  3.92,
		F10: []string{"def", "789"},
		F11: [2]string{"foo", "bar"},
		F12: 0xff,
		F13: []byte{1, 2, 3, 4, 5},
		F14: true,
		F15: Int64(-23),
		F16: Uint64(23),
		F17: JSONFloat64(3.14),
		F18: Uint128{
			Lo: 10,
			Hi: 82,
		},
		F19: Int128{
			Lo: 7,
			Hi: 3,
		},
		F20: Float128{
			Lo: 10,
			Hi: 82,
		},
		F21: Varuint32(999),
		F22: Varint32(-999),
		F23: Bool(true),
		F24: HexBytes([]byte{1, 2, 3, 4, 5}),
		NESTED1: &nested{
			F1: "NF1",
			F2: "NF2",
		},
		NESTED2: &nested{
			F1: "NF1",
			F2: "NF2",
		},
	}

	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(s, nil)
	assert.NoError(t, err)

	assert.Equal(t,
		"0a2062696e617279546573745374727563740a204e4553544544323a200a20206e65737465640a202046313a204e46310a202046323a204e46320a0a204e4553544544313a200a20206e65737465642046313a204e46312046323a204e4632200a2046313a206162630a2046323a202d37350a2046333a2039390a2046343a202d3233310a2046353a203939390a2046363a202d31333233310a2046373a2039393939390a2046383a202d32332e3132393939390a2046393a20332e3932303030300a204631303a200a20205b305d206465660a20205b315d203738390a204631313a200a20205b305d20666f6f0a20205b315d206261720a204631323a203235350a204631333a200a20205b305d20310a20205b315d20320a20205b325d20330a20205b335d20340a20205b345d20350a204631343a20747275650a204631353a202d32330a204631363a2032330a204631373a20332e3134303030300a204631383a20313531323633333031343034343138333233323532320a204631393a2035353334303233323232313132383635343835350a204632303a20313531323633333031343034343138333233323532320a204632313a203939390a204632323a202d3939390a204632333a20747275650a204632343a20303130323033303430350a",
		hex.EncodeToString(buf.Bytes()),
	)
}

type binaryTestStruct struct {
	NESTED2 *nested `bin:"sss" text:"notype"`
	NESTED1 *nested `text:"linear,notype"`
	F1      string
	F2      int16
	F3      uint16
	F4      int32
	F5      uint32
	F6      int64
	F7      uint64
	F8      float32
	F9      float64
	F10     []string
	F11     [2]string
	F12     byte
	F13     []byte
	F14     bool
	F15     Int64
	F16     Uint64
	F17     JSONFloat64
	F18     Uint128
	F19     Int128
	F20     Float128
	F21     Varuint32
	F22     Varint32
	F23     Bool
	F24     HexBytes
}
