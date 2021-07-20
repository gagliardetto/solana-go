package rpc

import (
	"encoding/json"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
)

func TestData_base64_zstd(t *testing.T) {
	val := "KLUv/QQAWQAAaGVsbG8td29ybGTcLcaB"
	in := `["` + val + `", "base64+zstd"]`

	var data DataBytesOrJSON
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte("hello-world"),
		data.GetBinary(),
	)
	assert.Equal(t,
		solana.EncodingBase64Zstd,
		data.asDecodedBinary.Encoding,
	)
	assert.Equal(t,
		[]interface{}{
			val,
			"base64+zstd",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_base64_zstd_empty(t *testing.T) {
	in := `["", "base64+zstd"]`

	var data DataBytesOrJSON
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte(""),
		data.GetBinary(),
	)
	assert.Equal(t,
		solana.EncodingBase64Zstd,
		data.asDecodedBinary.Encoding,
	)
	assert.Equal(t,
		[]interface{}{
			"",
			"base64+zstd",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_jsonParsed(t *testing.T) {
	in := `{"hello":"world"}`

	var data DataBytesOrJSON
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		json.RawMessage(in),
		data.GetRawJSON(),
	)
	assert.Equal(t,
		map[string]interface{}{
			"hello": "world",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_jsonParsed_empty(t *testing.T) {
	in := `{}`

	var data DataBytesOrJSON
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		json.RawMessage(in),
		data.GetRawJSON(),
	)
	assert.Equal(t,
		map[string]interface{}{},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}
