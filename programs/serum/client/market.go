package client

import (
	"bytes"
	"errors"
	"log"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/util"
)

type Market struct {
	client        *Client
	MarketAddress solana.PublicKey
	DexPID        solana.PublicKey
	Market        *serum.MarketV2
	BaseVault     *solana.PrivateKey
	QuoteVault    *solana.PrivateKey
}

type genericInstruction struct {
	programId   solana.PublicKey
	list        []*solana.AccountMeta
	instruction *serum.Instruction
}

func (mi *genericInstruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := bin.NewBinEncoder(buf).Encode(mi.instruction)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (mi *genericInstruction) ProgramID() solana.PublicKey {
	log.Printf("dex program id=%s", mi.programId.String())
	return mi.programId
}

func (mi *genericInstruction) Accounts() []*solana.AccountMeta {
	return mi.list
}

type openOrderItem struct {
	order   *serum.OpenOrders
	address solana.PublicKey
}

// look up the Dex program and find all open orders
func (c *Client) GetOpenOrdersForAllMarkets(dexPid solana.PublicKey) (ans map[string]*util.LinkedList[*openOrderItem], err error) {
	err = nil
	// market state stored on serum-dex/src/state.rs#L300
	accountList, err := c.rpc.GetProgramAccountsWithOpts(
		c.ctx,
		dexPid,
		&rpc.GetProgramAccountsOpts{
			Commitment: c.commitment,
			Encoding:   "base64",
			Filters:    []rpc.RPCFilter{
				//{
				//	Memcmp: &rpc.RPCFilterMemcmp{
				//		Offset: MARKET_STATE_MARKET_OFFSET,
				//		Bytes:  solana.Base58(market.String()[:]),
				//	},
				//},
				//{
				//	DataSize: MARKET_STATE_SPAN,
				//},
			},
		},
	)
	if err != nil {
		return
	}

	if len(accountList) == 0 {
		err = errors.New("no accounts fetched")
		return
	}
	ans = make(map[string]*util.LinkedList[*openOrderItem])

	for i := 0; i < len(accountList); i++ {
		//accountList[]
		info := accountList[i]
		var data []byte
		data, err = getBinaryDataFromAccount(info.Account)
		if err != nil {
			return
		}
		openOrders := new(serum.OpenOrders)
		err = bin.NewBinDecoder(data).Decode(openOrders)
		if err != nil {
			return
		}

		openOrderList, present := ans[openOrders.Market.String()]
		if !present {
			openOrderList = util.NewLinkedList[*openOrderItem]()
		}
		openOrderList.Append(&openOrderItem{
			address: accountList[i].Pubkey,
			order:   openOrders,
		})
	}

	return
}
