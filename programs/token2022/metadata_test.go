package token2022

import (
	"context"
	"strings"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	sendandconfirmtransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func TestCreateInitializeMetadataPointerInstruction(t *testing.T) {
	t.Skip()
	rpcUrl, err := getRpcUrl()
	if err != nil {
		t.Fatal(err)
	}
	defer clearMockchain(rpcUrl)
	rpcClient := rpc.New(rpcUrl)

	payer, err := solana.NewRandomPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := rpcClient.RequestAirdrop(
		context.Background(),
		payer.PublicKey(),
		solana.LAMPORTS_PER_SOL,
		rpc.CommitmentConfirmed,
	); err != nil {
		t.Fatal(err)
	}

	mint, err := solana.NewRandomPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	metadata := TokenMetadata{
		UpdateAuthority: payer.PublicKey().ToPointer(),
		Mint:            mint.PublicKey(),
		Name:            "OPOS",
		Symbol:          "OPOSSSS",
		Uri:             "https://raw.githubusercontent.com/solana-developers/opos-asset/main/assets/DeveloperPortal/metadata.json",
		AdditionalMetadata: map[string]string{
			"description": "Only Possible On Solana",
		},
	}

	lamports, err := rpcClient.GetMinimumBalanceForRentExemption(context.Background(), metadata.LenForLamports(), rpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}

	createAccountIx := system.NewCreateAccountInstruction(
		lamports,
		metadata.LenForLamports(),
		ProgramID,
		payer.PublicKey(),
		mint.PublicKey(),
	)

	initializeMetadataIx := CreateInitializeMetadataPointerInstruction(
		mint.PublicKey(),
		payer.PublicKey(),
		mint.PublicKey(),
	)

	token.SetProgramID(ProgramID)
	initializeMintIx := token.NewInitializeMintInstructionBuilder().
		SetDecimals(2).
		SetMintAuthority(payer.PublicKey()).
		SetMintAccount(mint.PublicKey()).
		SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey)

	recentBlockhash, err := rpcClient.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := solana.NewTransaction([]solana.Instruction{
		createAccountIx.Build(),
		initializeMetadataIx,
		initializeMintIx.Build(),
	}, recentBlockhash.Value.Blockhash)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		return &payer
	}); err != nil {
		t.Fatal(err)
	}

	wsUrl := strings.Replace(rpcUrl, "https://", "wss://", 1)
	wsClient, err := ws.Connect(context.Background(), wsUrl)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := sendandconfirmtransaction.SendAndConfirmTransaction(context.TODO(), rpcClient, wsClient, tx); err != nil {
		t.Fatal(err)
	}
}
