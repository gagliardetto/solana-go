package web3kit

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestGetTokenMetadata(t *testing.T) {
	ctx := context.Background()
	client := rpc.New(rpc.MainNetBeta_RPC)
	metadata, err := GetTokenMetadata(ctx, client, solana.MPK("HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC"), solana.Token2022ProgramID, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("metadata: %+v", metadata)
}
