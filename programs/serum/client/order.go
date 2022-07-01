package client

import (
	"encoding/binary"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	tkn "github.com/gagliardetto/solana-go/programs/token"
)

func (m *Market) SubmitOrderV1(payerPrivateKey *solana.PrivateKey, ownerPrivateKey *solana.PrivateKey, order *serum.InstructionNewOrder) (err error) {
	err = nil
	payer := payerPrivateKey.PublicKey()
	owner := ownerPrivateKey.PublicKey()

	keyMap := make(map[string]*solana.PrivateKey)

	list := []*solana.AccountMeta{
		solana.NewAccountMeta(m.MarketAddress, false, false),
		solana.NewAccountMeta(m.MarketAddress, false, false),
		solana.NewAccountMeta(m.Market.RequestQueue, false, false),
		solana.NewAccountMeta(payer, false, false),
		solana.NewAccountMeta(owner, false, false),
		solana.NewAccountMeta(m.BaseVault.PublicKey(), false, false),
		solana.NewAccountMeta(m.QuoteVault.PublicKey(), false, false),
		solana.NewAccountMeta(tkn.ProgramID, false, false),
		solana.NewAccountMeta(solana.SysVarRentPubkey, false, false),
	}

	err = order.SetAccounts(list)
	if err != nil {
		return
	}
	typeID := serum.InstructionDefVariant.TypeID(serum.Instruction_NewOrder)

	gi := &genericInstruction{
		programId: m.DexPID,
		list:      list,
		instruction: &serum.Instruction{
			Version: 0,
			BaseVariant: bin.BaseVariant{
				TypeID: bin.TypeIDFromUint32(typeID.Uint32(), binary.LittleEndian), // according to https://github.com/project-serum/serum-ts/blob/ca175c189fa9d3d8e41a9ad96884b874637c69f6/packages/serum/src/instructions.js#L34
				Impl:   order,
			},
		},
	}
	err = m.client.send_tx_generic(keyMap, []solana.Instruction{
		gi,
	})
	if err != nil {
		return
	}
	return
}

/*
Market             *solana.AccountMeta `text:"linear,notype"`
	OpenOrders         *solana.AccountMeta `text:"linear,notype"`
	RequestQueue       *solana.AccountMeta `text:"linear,notype"`
	Payer              *solana.AccountMeta `text:"linear,notype"`
	Owner              *solana.AccountMeta `text:"linear,notype"` // The owner of the open orders, i.e. the trader
	CoinVault          *solana.AccountMeta `text:"linear,notype"`
	PCVault            *solana.AccountMeta `text:"linear,notype"`
	SPLTokenProgram    *solana.AccountMeta `text:"linear,notype"`
	Rent               *solana.AccountMeta `text:"linear,notype"`
	SRMDiscountAccount *solana.AccountMeta `text:"linear,notype"`
*/

/*
let owner = new Account('...');
let payer = new PublicKey('...'); // spl-token account
await market.placeOrder(connection, {
  owner,
  payer,
  side: 'buy', // 'buy' or 'sell'
  price: 123.45,
  size: 17.0,
  orderType: 'limit', // 'limit', 'ioc', 'postOnly'
});
*/
