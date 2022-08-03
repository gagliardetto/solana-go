package rpc

import (
	"context"
	stdjson "encoding/json"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetAccountInfoWithRpcContext(t *testing.T) {
	type wants struct {
		account    *Account
		rpcContext *RPCContext
	}
	tests := []struct {
		name         string
		responseBody string
		key          solana.PublicKey
		opts         GetAccountInfoOpts
		want         wants
	}{
		{
			name:         "No Data",
			responseBody: `{"context":{"slot":83986106}}`,
			key:          solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932"),
			opts: GetAccountInfoOpts{
				Encoding:   solana.EncodingJSON,
				Commitment: CommitmentMax,
				DataSlice: &DataSlice{
					Offset: uint64Ptr(22),
					Length: uint64Ptr(33),
				},
			},
			want: wants{
				rpcContext: &RPCContext{
					Context: Context{
						Slot: 83986106,
					},
				},
			},
		},
		{
			name:         "Happy",
			responseBody: `{"context":{"slot":83986105},"value":{"data":["dGVzdA==","base64"],"executable":true,"lamports":999999,"owner":"11111111111111111111111111111111","rentEpoch":207}}`,
			key:          solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932"),
			opts: GetAccountInfoOpts{
				Encoding:   solana.EncodingJSON,
				Commitment: CommitmentMax,
				DataSlice: &DataSlice{
					Offset: uint64Ptr(22),
					Length: uint64Ptr(33),
				},
			},
			want: wants{
				account: &Account{
					Lamports:   999999,
					Owner:      solana.MustPublicKeyFromBase58("11111111111111111111111111111111"),
					Data:       DataBytesOrJSONFromBytes([]byte("test")),
					Executable: true,
					RentEpoch:  207,
				},
				rpcContext: &RPCContext{
					Context: Context{
						Slot: 83986105,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, closer := mockJSONRPC(t, stdjson.RawMessage(wrapIntoRPC(tt.responseBody)))
			defer closer()
			client := New(server.URL)

			acct, rpcContext, err := client.GetAccountInfoWithRpcContext(
				context.Background(),
				tt.key,
				&tt.opts,
			)
			require.NoError(t, err)

			assert.Equal(t, tt.want.account, acct)
			assert.Equal(t, tt.want.rpcContext, rpcContext)
		})
	}
}

func TestClient_GetAccountInfoWithRpcContext_exsting(t *testing.T) {
	responseBody := `{"context":{"slot":83986105},"value":{"data":["dGVzdA==","base64"],"executable":true,"lamports":999999,"owner":"11111111111111111111111111111111","rentEpoch":207}}`
	server, closer := mockJSONRPC(t, stdjson.RawMessage(wrapIntoRPC(responseBody)))
	defer closer()
	client := New(server.URL)

	offset := uint64(22)
	length := uint64(33)

	pubkeyString := "7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932"
	pubKey := solana.MustPublicKeyFromBase58(pubkeyString)

	opts := &GetAccountInfoOpts{
		Encoding:   solana.EncodingJSON,
		Commitment: CommitmentMax,
		DataSlice: &DataSlice{
			Offset: &offset,
			Length: &length,
		},
	}
	_, _, err := client.GetAccountInfoWithRpcContext(
		context.Background(),
		pubKey,
		opts,
	)
	require.NoError(t, err)

	assert.Equal(t,
		map[string]interface{}{
			"id":      float64(0),
			"jsonrpc": "2.0",
			"method":  "getAccountInfo",
			"params": []interface{}{
				pubkeyString,
				map[string]interface{}{
					"encoding":   string(solana.EncodingJSON),
					"commitment": string(CommitmentMax),
					"dataSlice": map[string]interface{}{
						"offset": float64(offset),
						"length": float64(length),
					},
				},
			},
		},
		server.RequestBody(t),
	)
}

func uint64Ptr(in uint64) *uint64 {
	return &in
}
