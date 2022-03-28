package util

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type TestValidatorConfig struct {
	ctx               context.Context
	Rpc               *rpc.Client
	Ws                *ws.Client
	PrivateKey        solana.PrivateKey
	defaultCommitment rpc.CommitmentType
}

func (c *TestValidatorConfig) PublicKey() solana.PublicKey {
	return c.PrivateKey.PublicKey()
}

func (c *TestValidatorConfig) Airdrop(supply uint64) error {
	sig, err := c.Rpc.RequestAirdrop(c.ctx, c.PublicKey(), supply, c.defaultCommitment)
	if err != nil {
		return err
	}
	err = c.Ws.WaitSig(c.ctx, sig, c.defaultCommitment)
	if err != nil {
		return err
	}
	return nil
}

func (c *TestValidatorConfig) Balance() (uint64, error) {
	result, err := c.Rpc.GetBalance(c.ctx, c.PublicKey(), c.defaultCommitment)
	if err != nil {
		return 0, err
	}
	return result.Value, nil
}

// run a test valdiator in a random directory
func SetupTestValidator(ctx context.Context, defaultCommitment rpc.CommitmentType, bpfProgram string, walletFilePath string, printLog bool) (config *TestValidatorConfig, err error) {
	config = &TestValidatorConfig{ctx: ctx, defaultCommitment: defaultCommitment}
	err = nil
	var tmpdir string
	tmpdir, err = os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		return
	}
	config.PrivateKey, err = solana.NewRandomPrivateKey()
	if err != nil {
		return
	}

	cmd := exec.Command(
		"solana-test-validator",
		"--bind-address", "127.0.0.1",
		"--ledger", tmpdir+"/test-ledger",
		"--bpf-program", walletFilePath, bpfProgram,
	)
	if printLog {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}

	err = cmd.Start()
	if err != nil {
		return
	}

	log.Printf("pid=%d", cmd.Process.Pid)
	time.Sleep(10 * time.Second)

	config.Rpc = rpc.New(
		"http://127.0.0.1:8899",
		map[string][]string{},
	)
	config.Ws, err = ws.Connect(ctx, "ws://127.0.0.1:8900", map[string][]string{})
	if err != nil {
		return
	}

	doneC := ctx.Done()
	go func() {
		<-doneC
		cmd.Process.Kill()
		time.Sleep(3 * time.Second)
		err2 := os.Remove(tmpdir)
		if err2 != nil {
			log.Print(err2)
		}

		err2 = cmd.Wait()
		if err2 != nil {
			log.Print(err2)
		}
	}()

	return
}
