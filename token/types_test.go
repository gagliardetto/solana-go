package token

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dfuse-io/solana-go"
	"github.com/lunixbochs/struc"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	b58data := "SqtzmJArwV2556pK7AdHbHNPVP2L2WaR6zfcFeot94TzGRUyUMEWew558UxnYEGrmm9b9VZY7MS6TCHT5wqtzaA5Vy8ghoFyGmbRNC58CttRf5GzH9wfjCkncyrmKjfevyjrJ2W9XKLgYGth46ctFWzJJXCeHsYwDx1d"
	data, _ := base58.Decode(b58data)

	//fmt.Println("HEX:", hex.EncodeToString(data))
	// ba71eb12868584549b86f75620e7bb3ac5ef49df3fef0d48ad08e48dfa0fc786  // mint
	// d7a1d0a56e355f17cedd5733e36a0cc9e2caf7a435e3256e4c9bff755f682b5a  // owner
	// 5ece000000000000   // amount
	// 00000000    // is delegate set
	// 0000000000000000000000000000000000000000000000000000000000000000  // delegate
	// 01000000    // is initialized, is native + padding
	// 0000000000000000    // delegate amount
	var out Account
	err := struc.Unpack(bytes.NewReader(data), &out)
	require.NoError(t, err)

	expect := Account{
		Mint:          solana.MustPublicKeyFromBase58("DYoajiN32pjK8zMAa67ScNn2E7EmXrZ6doABRqfSZ63F"),
		Owner:         solana.MustPublicKeyFromBase58("FWjmNcjufwC3QFdcHrAK1yAQkCwJSUAxvVFFgvQ1nAJM"),
		Amount:        solana.U64(52830),
		IsInitialized: true,
	}
	expectJSON, err := json.MarshalIndent(expect, "", "  ")
	require.NoError(t, err)

	outJSON, err := json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)

	assert.JSONEq(t, string(expectJSON), string(outJSON))

	buf := &bytes.Buffer{}
	assert.NoError(t, struc.Pack(buf, out))

	fmt.Println("MAMA", hex.EncodeToString(buf.Bytes()))
	assert.Equal(t, b58data, base58.Encode(buf.Bytes()))
}

// func TestMint(t *testing.T) {
// 	mintData := ""
// }
