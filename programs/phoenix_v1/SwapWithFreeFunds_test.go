// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"bytes"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_SwapWithFreeFunds(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("SwapWithFreeFunds"+strconv.Itoa(i), func(t *testing.T) {
			{
				{
					{
						params := new(SwapWithFreeFunds)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(PostOnly)
						fu.Fuzz(tmp)
						params.SetOrderPacket(tmp)
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(SwapWithFreeFunds)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						ag_require.Equal(t, params, got)
					}
					{
						params := new(SwapWithFreeFunds)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(Limit)
						fu.Fuzz(tmp)
						params.SetOrderPacket(tmp)
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(SwapWithFreeFunds)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						ag_require.Equal(t, params, got)
					}
					{
						params := new(SwapWithFreeFunds)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(ImmediateOrCancel)
						fu.Fuzz(tmp)
						params.SetOrderPacket(tmp)
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(SwapWithFreeFunds)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						ag_require.Equal(t, params, got)
					}
				}
			}
		})
	}
}
