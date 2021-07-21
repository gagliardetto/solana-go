package solana

import (
	"strconv"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMustHashFromBase58(t *testing.T) {
	require.Panics(t, func() {
		MustHashFromBase58("toto")
	})
}

func TestHashFromBase58(t *testing.T) {
	in := "uoEAQCWCKjV9ecsBvngctJ7upNBZX7hpN4SfdR6TaUz"
	out, err := HashFromBase58(in)
	assert.NoError(t, err)
	assert.Equal(t, in, out.String())

	jsonOut, err := out.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, strconv.Quote(in), string(jsonOut))

	var ha Hash
	assert.True(t, ha.IsZero())
	err = ha.UnmarshalJSON(jsonOut)
	assert.NoError(t, err)
	assert.Equal(t, out, ha)
	assert.True(t, out.Equals(ha))
	assert.False(t, out.Equals(Hash{}))
	assert.False(t, out.IsZero())
}

func TestSignatureFromBase58(t *testing.T) {
	in := "gD3jeeaPNiyuJvTKXNEv1gntazWEkvpocofEmrz2rL6Fi4prWSsBH6a9SrwyZEatAozyMsnK2fnk3APXNFxD2Mq"
	out, err := SignatureFromBase58(in)
	assert.NoError(t, err)
	assert.Equal(t, in, out.String())

	jsonOut, err := out.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, strconv.Quote(in), string(jsonOut))

	var sig Signature
	assert.True(t, sig.IsZero())
	err = sig.UnmarshalJSON(jsonOut)
	assert.NoError(t, err)
	assert.Equal(t, out, sig)
	assert.True(t, out.Equals(sig))
	assert.False(t, out.Equals(Signature{}))
	assert.False(t, out.IsZero())
}

func TestMustSignatureFromBase58(t *testing.T) {
	require.Panics(t, func() {
		MustSignatureFromBase58("toto")
	})
}

func TestBase58(t *testing.T) {
	in := "RYcCwZg97M2jet84ttG8"

	out, err := base58.Decode(in)
	assert.NoError(t, err)
	assert.Equal(t, in, Base58(out).String())

	jsonOut, err := Base58(out).MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, strconv.Quote(in), string(jsonOut))

	var b58 Base58
	err = b58.UnmarshalJSON(jsonOut)
	assert.NoError(t, err)
}

func TestData_base64(t *testing.T) {
	val := "dGVzdA=="
	in := `["` + val + `", "base64"]`

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte("test"),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase64,
		data.Encoding,
	)

	assert.Equal(t,
		[]interface{}{
			val,
			"base64",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_base64_empty(t *testing.T) {
	val := ""
	in := `["", "base64"]`

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte(""),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase64,
		data.Encoding,
	)

	assert.Equal(t,
		[]interface{}{
			val,
			"base64",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_base64_zstd(t *testing.T) {
	val := "KLUv/QQAWQAAaGVsbG8td29ybGTcLcaB"
	in := `["` + val + `", "base64+zstd"]`

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte("hello-world"),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase64Zstd,
		data.Encoding,
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

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte(""),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase64Zstd,
		data.Encoding,
	)

	assert.Equal(t,
		[]interface{}{
			"",
			"base64+zstd",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_base58(t *testing.T) {
	val := "3yZe7d"
	in := `["` + val + `", "base58"]`

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte("test"),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase58,
		data.Encoding,
	)

	assert.Equal(t,
		[]interface{}{
			val,
			"base58",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

func TestData_base58_empty(t *testing.T) {
	val := ""
	in := `["", "base58"]`

	var data Data
	err := data.UnmarshalJSON([]byte(in))
	assert.NoError(t, err)

	assert.Equal(t,
		[]byte(""),
		data.Content,
	)

	assert.Equal(t,
		EncodingBase58,
		data.Encoding,
	)

	assert.Equal(t,
		[]interface{}{
			val,
			"base58",
		},
		mustJSONToInterface(mustAnyToJSON(data)),
	)
}

// mustAnyToJSON marshals the provided variable
// to JSON bytes.
func mustAnyToJSON(raw interface{}) []byte {
	out, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return out
}

// mustJSONToInterface unmarshals the provided JSON bytes
// into an `interface{}` type variable, and returns it.
func mustJSONToInterface(rawJSON []byte) interface{} {
	var out interface{}
	err := json.Unmarshal(rawJSON, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func TestMustPublicKeyFromBase58(t *testing.T) {
	require.Panics(t, func() {
		MustPublicKeyFromBase58("toto")
	})
}
