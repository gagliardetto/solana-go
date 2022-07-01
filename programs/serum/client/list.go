package client

import (
	"encoding/binary"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/gagliardetto/solana-go/programs/system"
	tkn "github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

func (c *Client) create_generic_account(space uint64, payer solana.PublicKey) (*system.CreateAccount, error) {

	minLamports, err := c.rpc.GetMinimumBalanceForRentExemption(c.ctx, space, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, err
	}
	return system.NewCreateAccountInstructionBuilder().SetSpace(space).SetLamports(minLamports).SetFundingAccount(payer), nil
}

func (c *Client) create_token_account(payer solana.PublicKey, newAccount solana.PublicKey) (*system.Instruction, error) {
	prefix, err := c.create_generic_account(uint64(165), payer)
	if err != nil {
		return nil, err
	}
	return prefix.SetOwner(tkn.ProgramID).SetNewAccount(newAccount).Build(), nil
}

func (c *Client) create_market_account(dexPid solana.PublicKey, payer solana.PublicKey, newAccount solana.PublicKey) (*system.Instruction, error) {
	// 388
	space := 5 + 8 + 32 + 8 + 32 + 32 + 32 + 8 + 8 + 32 + 8 + 8 + 8 + 32 + 32 + 32 + 32 + 8 + 8 + 8 + 8 + 7
	prefix, err := c.create_generic_account(uint64(space), payer)
	if err != nil {
		return nil, err
	}

	return prefix.SetOwner(dexPid).SetNewAccount(newAccount).Build(), nil
}

func (c *Client) create_request_queue(dexPid solana.PublicKey, payer solana.PublicKey, newAccount solana.PublicKey) (*system.Instruction, error) {

	prefix, err := c.create_generic_account(uint64(5120+12), payer)
	if err != nil {
		return nil, err
	}

	return prefix.SetOwner(dexPid).SetNewAccount(newAccount).Build(), nil
}

func (c *Client) create_event_queue(dexPid solana.PublicKey, payer solana.PublicKey, newAccount solana.PublicKey) (*system.Instruction, error) {

	prefix, err := c.create_generic_account(uint64(262144+12), payer)
	if err != nil {
		return nil, err
	}

	return prefix.SetOwner(dexPid).SetNewAccount(newAccount).Build(), nil
}

func (c *Client) create_bidask_account(dexPid solana.PublicKey, payer solana.PublicKey, newAccount solana.PublicKey) (*system.Instruction, error) {
	prefix, err := c.create_generic_account(uint64(65536+12), payer)
	if err != nil {
		return nil, err
	}
	return prefix.SetOwner(dexPid).SetNewAccount(newAccount).Build(), nil
}

func (c *Client) initialize_token_account(newAccount solana.PublicKey, mint solana.PublicKey, owner solana.PublicKey) *tkn.Instruction {
	return tkn.NewInitializeAccountInstructionBuilder().SetAccount(newAccount).SetMintAccount(mint).SetOwnerAccount(owner).Build()
}

type ListArgs struct {
	DEX_PID            solana.PublicKey
	Payer              solana.PrivateKey
	BaseMint           solana.PublicKey
	QuoteMint          solana.PublicKey
	BaseLotSize        uint64
	QuoteLotSize       uint64
	FeeRateBps         uint16
	QuoteDustThreshold uint64
}

// List a new token pair on Serum
func (c *Client) List(args *ListArgs) (result *Market, err error) {
	result = new(Market)
	m := new(marketInfo)
	payer := args.Payer.PublicKey()

	m.privateKeyMap = make(map[string]*solana.PrivateKey)

	m.privateKeyMap[payer.String()] = &args.Payer

	marketPrivateKey, err := m.create_private_key()
	if err != nil {
		return
	}
	m.market = marketPrivateKey.PublicKey()

	requestQueuePrivateKey, err := m.create_private_key()
	if err != nil {
		return
	}
	m.requestQueue = requestQueuePrivateKey.PublicKey()

	eventQueuePrivateKey, err := m.create_private_key()
	if err != nil {
		return
	}
	m.eventQueue = eventQueuePrivateKey.PublicKey()

	bidsPrivateKey, err := m.create_private_key()
	if err != nil {
		return
	}
	m.bids = bidsPrivateKey.PublicKey()

	asksPrivateKey, err := m.create_private_key()
	if err != nil {
		return
	}
	m.asks = asksPrivateKey.PublicKey()
	m.privateKeyMap[m.asks.String()] = asksPrivateKey

	baseVault, err := m.create_private_key()
	if err != nil {
		return
	}
	m.baseVault = baseVault

	quoteVault, err := m.create_private_key()
	if err != nil {
		return
	}
	m.quoteVault = quoteVault

	vaultOwner, vaultNonce, err := getVaultOwnerAndNonce(args.DEX_PID, m.market)
	if err != nil {
		return
	}

	inst_base_vault, err := c.create_token_account(payer, m.baseVault.PublicKey())
	if err != nil {
		return
	}

	inst_quote_vault, err := c.create_token_account(payer, m.quoteVault.PublicKey())
	if err != nil {
		return
	}
	inst_base_vault_initialize := c.initialize_token_account(m.baseVault.PublicKey(), args.BaseMint, *vaultOwner)
	inst_quote_vault_initialize := c.initialize_token_account(m.quoteVault.PublicKey(), args.QuoteMint, *vaultOwner)

	err = c.send_tx_for_listing(m, []solana.Instruction{
		inst_base_vault, inst_quote_vault, inst_base_vault_initialize, inst_quote_vault_initialize,
	})
	if err != nil {
		return
	}

	inst_market, err := c.create_market_account(args.DEX_PID, payer, m.market)
	if err != nil {
		return
	}
	inst_request_queue, err := c.create_request_queue(args.DEX_PID, payer, m.requestQueue)
	if err != nil {
		return
	}
	inst_event_queue, err := c.create_event_queue(args.DEX_PID, payer, m.eventQueue)
	if err != nil {
		return
	}

	inst_bid, err := c.create_bidask_account(args.DEX_PID, payer, m.bids)
	if err != nil {
		return
	}
	inst_ask, err := c.create_bidask_account(args.DEX_PID, payer, m.asks)
	if err != nil {
		return
	}
	ima := new(serum.InstructionInitializeMarket)
	ima.BaseLotSize = args.BaseLotSize
	ima.QuoteLotSize = args.QuoteLotSize
	ima.FeeRateBps = args.FeeRateBps
	ima.VaultSignerNonce = vaultNonce
	ima.QuoteDustThreshold = args.QuoteDustThreshold

	list := []*solana.AccountMeta{
		/// 0. `[writable]` the market to initialize
		solana.NewAccountMeta(m.market, true, false),
		/// 1. `[writable]` zeroed out request queue
		solana.NewAccountMeta(m.requestQueue, true, false),
		/// 2. `[writable]` zeroed out event queue
		solana.NewAccountMeta(m.eventQueue, true, false),
		/// 3. `[writable]` zeroed out bids
		solana.NewAccountMeta(m.bids, true, false),
		/// 4. `[writable]` zeroed out asks
		solana.NewAccountMeta(m.asks, true, false),
		/// 5. `[writable]` spl-token account for the coin currency
		solana.NewAccountMeta(m.baseVault.PublicKey(), true, false),
		/// 6. `[writable]` spl-token account for the price currency
		solana.NewAccountMeta(m.quoteVault.PublicKey(), true, false),
		/// 7. `[]` coin currency Mint
		solana.NewAccountMeta(args.BaseMint, false, false),
		/// 8. `[]` price currency Mint
		solana.NewAccountMeta(args.QuoteMint, false, false),
		/// 9. `[]` the rent sysvar
		solana.NewAccountMeta(solana.SysVarRentPubkey, false, false),
		/// 10. `[]` open orders market authority (optional)
		//solana.NewAccountMeta(authority, false, false),
		/// 11. `[]` prune authority (optional, requires open orders market authority)
		//solana.NewAccountMeta(authority_prune, false, false),
		/// 12. `[]` crank authority (optional, requires prune authority)
		//solana.NewAccountMeta(authority_crank, false, false),
	}
	mi := &genericInstruction{
		programId: args.DEX_PID,
		list:      list,
		instruction: &serum.Instruction{
			Version: 0,
			BaseVariant: bin.BaseVariant{
				TypeID: bin.TypeIDFromUint32(0, binary.LittleEndian), // according to https://github.com/project-serum/serum-ts/blob/ca175c189fa9d3d8e41a9ad96884b874637c69f6/packages/serum/src/instructions.js#L34
				Impl:   ima,
			},
		},
	}

	err = c.send_tx_for_listing(m, []solana.Instruction{
		inst_market, inst_request_queue, inst_event_queue, inst_bid, inst_ask, mi,
	})
	if err != nil {
		return
	}

	marketAddress := marketPrivateKey.PublicKey()

	result.client = c
	result.DexPID = args.DEX_PID
	result.MarketAddress = marketAddress
	result.BaseVault = baseVault
	result.QuoteVault = quoteVault
	result.Market, err = c.GetMarket(m.market)
	if err != nil {
		return
	}
	return
}

type marketInfo struct {
	market        solana.PublicKey
	requestQueue  solana.PublicKey
	eventQueue    solana.PublicKey
	bids          solana.PublicKey
	asks          solana.PublicKey
	baseVault     *solana.PrivateKey
	quoteVault    *solana.PrivateKey
	privateKeyMap map[string]*solana.PrivateKey
}

func (m *marketInfo) create_private_key() (*solana.PrivateKey, error) {
	key, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	m.privateKeyMap[key.PublicKey().String()] = &key
	return &key, nil
}

// get the current market information for a token trading pair
func (c *Client) GetMarket(market solana.PublicKey) (*serum.MarketV2, error) {
	result, err := c.rpc.GetAccountInfo(c.ctx, market)
	if err != nil {
		return nil, err
	}
	data, err := getBinaryDataFromAccountInfo(result)
	if err != nil {
		return nil, err
	}
	ans := new(serum.MarketV2)
	err = bin.NewBinDecoder(data).Decode(ans)
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func (c *Client) send_tx_for_listing(m *marketInfo, instructions []solana.Instruction) error {
	return c.send_tx_generic(m.privateKeyMap, instructions)
}

func (c *Client) send_tx_generic(keyMap map[string]*solana.PrivateKey, instructions []solana.Instruction) error {
	blockhash, err := c.latest_blockhash()
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction(
		instructions,
		blockhash,
	)
	if err != nil {
		return err
	}
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		otherKey, present := keyMap[key.String()]
		if present {
			return otherKey
		}
		return nil
	})
	if err != nil {
		return err
	}
	sig, err := c.rpc.SendTransactionWithOpts(c.ctx, tx, true, c.commitment)
	if err != nil {
		return err
	}
	err = c.block_on_tx_processing(sig)
	if err != nil {
		return err
	}
	return nil
}
