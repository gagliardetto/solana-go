package text

import (
	"bytes"
	"fmt"
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

	fmt.Println()
	buf := new(bytes.Buffer)
	enc := NewEncoder(buf)
	err := enc.Encode(s, nil)
	assert.NoError(t, err)

	fmt.Println(string(buf.Bytes()))
	//assert.Equal(t,
	//	"03616263b5ff630019ffffffe703000051ccffffffffffff9f860100000000003d0ab9c15c8fc2f5285c0f4002036465660337383903666f6f03626172ff05010203040501e9ffffffffffffff17000000000000001f85eb51b81e09400a000000000000005200000000000000070000000000000003000000000000000a000000000000005200000000000000e707cd0f01050102030405",
	//	string(buf.Bytes()),
	//)
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
