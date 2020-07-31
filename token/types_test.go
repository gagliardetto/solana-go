package token

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/lunixbochs/struc"
	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	b58data := "SqtzmJArwV2556pK7AdHbHNPVP2L2WaR6zfcFeot94TzGRUyUMEWew558UxnYEGrmm9b9VZY7MS6TCHT5wqtzaA5Vy8ghoFyGmbRNC58CttRf5GzH9wfjCkncyrmKjfevyjrJ2W9XKLgYGth46ctFWzJJXCeHsYwDx1d"
	data, _ := base58.Decode(b58data)

	var out Account
	err := struc.Unpack(bytes.NewReader(data), &out)
	require.NoError(t, err)

	expect := Account{}
	expectJSON, err := json.MarshalIndent(expect, "", "  ")
	require.NoError(t, err)

	outJSON, err := json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)

	assert.JSONEq(t, string(expectJSON), string(outJSON))
}

// func TestMint(t *testing.T) {
// 	mintData := ""
// }
