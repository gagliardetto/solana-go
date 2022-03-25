package client_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
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

	rpcClient := rpc.New("http://127.0.0.1:48899", nil)
	result, err := rpcClient.GetBlock(ctx, 12994433)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("block %s %d", result.Blockhash.String(), *result.BlockHeight)
	//logSub, err := wsClient.LogsSubscribe(ws.LogsSubscribeFilterAll, rpc.CommitmentProcessed)
	//if err != nil {
	//t.Fatal(err)
	//}

	//logC := logSub.RecvStream()
	//logCloseC := logSub.CloseSignal()

	slotSub, err := wsClient.SlotSubscribe()
	if err != nil {
		t.Fatal(err)
	}

	slotC := slotSub.RecvStream()
	slotCloseC := slotSub.CloseSignal()

	doneC := ctx.Done()

out:
	for {
		select {
		case <-doneC:
			break out
		/*case x := <-logC:
		l := x.(*ws.LogResult)
		for i := 0; i < len(l.Value.Logs); i++ {
			log.Print(l.Value.Logs[i])
		}*/
		case x := <-slotC:
			s := x.(*ws.SlotResult)
			log.Printf("slot=%d", s.Slot)
			go printBlock(ctx, s.Slot)
		//case err = <-logCloseC:
		//		break out
		case err = <-slotCloseC:
			break out
		}
	}
	if err != nil {
		t.Fatal(err)
	}

}

func printBlock(ctx context.Context, slot uint64) {
	rpcClient := rpc.New("http://127.0.0.1:48899", nil)
	result, err := rpcClient.GetBlock(ctx, slot)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("block %s %d", result.Blockhash.String(), *result.BlockHeight)
}

func TestList(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	bpf, present := os.LookupEnv("SERUM_DEX_BPF_FILE_PATH")
	if !present {
		t.Fatal("no bpf")
	}
	ctx, cancel := context.WithCancel(context.Background())
	tv, err := util.SetupTestValidator(ctx, rpc.CommitmentFinalized, bpf)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cancel()
	})

	err = tv.Airdrop(100 * solana.LAMPORTS_PER_SOL)
	if err != nil {
		t.Fatal(err)
	}

	//tokenClient := tknclient.Create(ctx, rpcClient, wsClient, rpc.CommitmentFinalized)
	//tokenClient.CreateToken(&tknclient.TokenArgs{})

	//serumClient := client.Create(ctx, rpcClient, wsClient, rpc.CommitmentFinalized)

	//serumClient.List(&client.ListArgs{})

	time.Sleep(20 * time.Second)

}
