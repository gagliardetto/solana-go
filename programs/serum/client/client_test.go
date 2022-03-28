package client_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum/client"
	tknclient "github.com/gagliardetto/solana-go/programs/token/client"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/util"
	"github.com/joho/godotenv"
)

func TestExternal(t *testing.T) {
	ctx := context.Background()
	wsClient, err := ws.Connect(ctx, "ws://127.0.0.1:48899", nil)
	if err != nil {
		t.Fatal(err)
	}

	logSub, err := wsClient.LogsSubscribe(ws.LogsSubscribeFilterAll, rpc.CommitmentProcessed)
	if err != nil {
		t.Fatal(err)
	}

	logC := logSub.RecvStream()
	logCloseC := logSub.CloseSignal()

	slotSub, err := wsClient.SlotSubscribe()
	if err != nil {
		t.Fatal(err)
	}

	slotC := slotSub.RecvStream()
	slotCloseC := slotSub.CloseSignal()

	doneC := ctx.Done()

	currentTime := time.Now()

	solana.SignatureFromBase58("")
out:
	for {
		select {
		case <-doneC:
			break out
		case x := <-logC:
			l := x.(*ws.LogResult)
			for i := 0; i < len(l.Value.Logs); i++ {
				log.Print(l.Value.Logs[i])
			}
		case x := <-slotC:
			s := x.(*ws.SlotResult)
			newTime := time.Now()
			diff := newTime.UnixMilli() - currentTime.UnixMilli()
			currentTime = newTime
			go printBlock(ctx, s.Slot, diff)
		case err = <-logCloseC:
			break out
		case err = <-slotCloseC:
			break out
		}
	}
	if err != nil {
		t.Fatal(err)
	}

}

func printBlock(ctx context.Context, slot uint64, diff int64) {
	rpcClient := rpc.New("http://127.0.0.1:48899", nil)
	result, err := rpcClient.GetBlock(ctx, slot)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("blockHash=%s diff seconds=%d", result.Blockhash, diff)
}

func TestBasic(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	bpf, present := os.LookupEnv("SERUM_DEX_BPF_FILE_PATH")
	if !present {
		t.Fatal("no bpf")
	}
	walletFilePath, present := os.LookupEnv("SERUM_DEX_WALLET")
	if !present {
		t.Fatal("no bpf")
	}
	ctx, cancel := context.WithCancel(context.Background())
	tv, err := util.SetupTestValidator(ctx, rpc.CommitmentFinalized, bpf, walletFilePath, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cancel()
	})

	bal := 100 * solana.LAMPORTS_PER_SOL
	err = tv.Airdrop(bal)
	if err != nil {
		t.Fatal(err)
	}

	checkBal, err := tv.Balance()
	if err != nil {
		t.Fatal(err)
	}
	if bal != checkBal {
		t.Fatal("balance does not match")
	}

	//tokenClient := tknclient.Create(ctx, rpcClient, wsClient, rpc.CommitmentFinalized)
	//tokenClient.CreateToken(&tknclient.TokenArgs{})

	//serumClient := client.Create(ctx, rpcClient, wsClient, rpc.CommitmentFinalized)

	//serumClient.List(&client.ListArgs{})

	time.Sleep(20 * time.Second)

}

func TestMintToken(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	bpf, present := os.LookupEnv("SERUM_DEX_BPF_FILE_PATH")
	if !present {
		t.Fatal("no bpf")
	}
	programWalletFilePath, present := os.LookupEnv("SERUM_DEX_WALLET")
	if !present {
		t.Fatal("no bpf")
	}
	ctx, cancel := context.WithCancel(context.Background())
	tv, err := util.SetupTestValidator(ctx, rpc.CommitmentFinalized, bpf, programWalletFilePath, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cancel()
	})

	programIdPrivateKey, err := solana.PrivateKeyFromSolanaKeygenFile(programWalletFilePath)
	if err != nil {
		t.Fatal(err)
	}
	dexId := programIdPrivateKey.PublicKey()

	bal := 100 * solana.LAMPORTS_PER_SOL
	err = tv.Airdrop(bal)
	if err != nil {
		t.Fatal(err)
	}

	checkBal, err := tv.Balance()
	if err != nil {
		t.Fatal(err)
	}
	if bal != checkBal {
		t.Fatal("balance does not match")
	}

	tokenClient := tknclient.Create(ctx, tv.Rpc, tv.Ws, rpc.CommitmentFinalized)
	mint_USD, err := tokenClient.CreateToken(&tknclient.TokenArgs{
		Decimals:        2,
		InitialSupply:   10000,
		MintAuthority:   tv.PublicKey(),
		FreezeAuthority: tv.PublicKey(),
	}, tv.PrivateKey, tv.PublicKey())
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("mint usd=%+v", mint_USD)
	if mint_USD.State.MintAuthority == nil {
		t.Fatal("no mint authority")
	}

	mint_JPY, err := tokenClient.CreateToken(&tknclient.TokenArgs{
		Decimals:        0,
		InitialSupply:   10000 * 100,
		MintAuthority:   tv.PublicKey(),
		FreezeAuthority: tv.PublicKey(),
	}, tv.PrivateKey, tv.PublicKey())
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("mint jpy=%+v", mint_JPY.State)

	serumClient := client.Create(ctx, tv.Rpc, tv.Ws, rpc.CommitmentFinalized)

	result, err := serumClient.List(&client.ListArgs{
		DEX_PID:            dexId,
		Payer:              tv.PrivateKey,
		BaseMint:           mint_USD.Address,
		QuoteMint:          mint_JPY.Address,
		BaseLotSize:        1000,
		QuoteLotSize:       10,
		FeeRateBps:         10,
		QuoteDustThreshold: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("market address=%s", result.MarketAddress.String())

}
