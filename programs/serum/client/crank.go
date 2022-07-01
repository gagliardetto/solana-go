package client

import (
	"encoding/binary"
	"errors"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/gagliardetto/solana-go/util"
)

const (
	MARKET_STATE_SPAN          uint64 = 388
	MARKET_STATE_MARKET_OFFSET uint64 = 1 * 8
)

type marketLookup struct {
	open *util.LinkedList[*serum.OpenOrders]
}

const MAX_INSTUCTION_SIZE = 4

// call the consume event instruction
func (c *Client) Crank(dexPid solana.PublicKey, crankAuthority solana.PublicKey, limit uint16) (err error) {
	err = nil
	keyMap := make(map[string]*solana.PrivateKey)

	//x := &serum.ConsumeEventsAccounts{}

	orderMap, err := c.GetOpenOrdersForAllMarkets(dexPid)
	if err != nil {
		return
	}
	if len(orderMap) == 0 {
		err = errors.New("no open orders")
		return
	}

	marketMap := make(map[string]*serum.MarketV2)
	instructionList := make([]solana.Instruction, MAX_INSTUCTION_SIZE)
	i_inst := 0
	var maddr solana.PublicKey

	for mstr, list := range orderMap {

		market, present := marketMap[mstr]
		if !present {
			maddr, err = solana.PublicKeyFromBase58(mstr)
			if err != nil {
				return
			}
			market, err = c.GetMarket(maddr)
			if err != nil {
				return
			}
		}

		metaList := make([]*solana.AccountMeta, 4+list.Size)
		length := len(metaList)
		metaList[length-1-3] = solana.NewAccountMeta(market.OwnAddress, true, false)
		metaList[length-1-2] = solana.NewAccountMeta(market.EventQueue, true, false)
		// should be blank
		metaList[length-1-1] = solana.NewAccountMeta(market.EventQueue, true, false)
		// should be blank
		metaList[length-1-0] = solana.NewAccountMeta(market.EventQueue, true, false)

		err = list.Iterate(func(index int, v *openOrderItem) error {
			metaList[index] = solana.NewAccountMeta(v.address, true, false)
			return nil
		})
		if err != nil {
			return
		}
		inst := &serum.InstructionConsumeEvents{Limit: limit}

		typeID := serum.InstructionDefVariant.TypeID(serum.Instruction_ConsumeEvents)

		instructionList[i_inst] = &genericInstruction{
			programId: dexPid,
			list:      metaList,
			instruction: &serum.Instruction{
				Version: 0,
				BaseVariant: bin.BaseVariant{
					TypeID: bin.TypeIDFromUint32(typeID.Uint32(), binary.LittleEndian), // according to https://github.com/project-serum/serum-ts/blob/ca175c189fa9d3d8e41a9ad96884b874637c69f6/packages/serum/src/instructions.js#L34
					Impl:   inst,
				},
			},
		}
		i_inst++
		if i_inst == len(instructionList) {
			i_inst = 0
			err = c.send_tx_generic(keyMap, instructionList)
			if err != nil {
				return
			}
			instructionList = make([]solana.Instruction, MAX_INSTUCTION_SIZE)
		}

	}
	if 0 < i_inst {
		err = c.send_tx_generic(keyMap, instructionList)
		if err != nil {
			return
		}
	}

	return
}
