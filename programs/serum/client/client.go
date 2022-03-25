// talk to serum
package client

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type Client struct {
	ctx        context.Context
	commitment rpc.CommitmentType
	ws         *ws.Client
	rpc        *rpc.Client
}

func Create(ctx context.Context, rpcClient *rpc.Client, wsClient *ws.Client, commitment rpc.CommitmentType) *Client {
	return &Client{ctx: ctx, rpc: rpcClient, ws: wsClient, commitment: commitment}
}

func (c *Client) latest_blockhash() (solana.Hash, error) {
	latestBlockhashResult, err := c.rpc.GetLatestBlockhash(c.ctx, c.commitment)
	if err != nil {
		return solana.Hash{}, err
	}

	if latestBlockhashResult.Value == nil {
		return solana.Hash{}, errors.New("blank block hash")
	}

	return latestBlockhashResult.Value.Blockhash, nil
}

// block until we get notification from a validator that the transaction has reached c.commitment state
func (c *Client) block_on_tx_processing(sig solana.Signature) error {
	sub, err := c.ws.SignatureSubscribe(sig, c.commitment)
	if err != nil {
		return err
	}
	streamC := sub.RecvStream()
	doneC := c.ctx.Done()
	errorC := sub.CloseSignal()

	defer sub.Unsubscribe()

out:
	for {
		select {
		case <-doneC:
			break out
		case err = <-errorC:
			break out
		case d := <-streamC:
			x, ok := d.(*ws.SignatureResult)
			if !ok {
				err = errors.New("bad type")
			}
			if x.Value.Err != nil {
				err = fmt.Errorf("%+v", x.Value.Err)
			}
			break out
		}
	}

	return err
}

func getVaultOwnerAndNonce(dexID solana.PublicKey, market solana.PublicKey) (*solana.PublicKey, uint64, error) {
	nonce := uint64(0)
	var pubkey solana.PublicKey
	var err error
out:
	for nonce < 1000 {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, nonce)
		log.Printf("market=%s nonce=%s dex=%s", market.String(), hex.EncodeToString(b), dexID.String())
		pubkey, err = solana.CreateProgramAddress([][]byte{market.Bytes(), b}, dexID)
		if err != nil {
			log.Print(err)
			nonce++
		} else {
			break out
		}
	}
	if err != nil {
		return nil, 0, err
	}
	log.Printf("nonce=%d pubkey=%s", nonce, pubkey.String())

	return &pubkey, nonce, nil
}

func getBinaryDataFromAccountInfo(result *rpc.GetAccountInfoResult) ([]byte, error) {
	v := result.Value
	if v == nil {
		return nil, errors.New("no value")
	}
	data := v.Data
	if data == nil {
		return nil, errors.New("no data")
	}
	return data.GetBinary(), nil
}

func getBinaryDataFromAccount(v *rpc.Account) ([]byte, error) {
	if v == nil {
		return nil, errors.New("no value")
	}
	data := v.Data
	if data == nil {
		return nil, errors.New("no data")
	}
	return data.GetBinary(), nil
}
