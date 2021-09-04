package token

import (
	"bytes"
	"strconv"
	"testing"

	ag_gofuzz "github.com/google/gofuzz"
	ag_require "github.com/stretchr/testify/require"
)

func TestEncodeDecode_Transfer(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Transfer"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Transfer)
				fu.Fuzz(params)
				params.Accounts = nil
				params.Signers = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				//
				got := new(Transfer)
				err = decodeT(got, buf.Bytes())
				got.Accounts = nil
				params.Signers = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}
