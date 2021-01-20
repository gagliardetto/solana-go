package serum

import (
	"encoding/binary"
	"fmt"

	"github.com/dfuse-io/solana-go/text"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

func init() {
	solana.RegisterInstructionDecoder(PROGRAM_ID, registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	// FIXME: can't we dedupe this in some ways? It's copied in all of the programs' folders.
	var inst Instruction
	if err := bin.NewDecoder(data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(solana.AccountSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}

	return &inst, nil
}

type Instruction struct {
	bin.BaseVariant
	Version uint8
}

func (i *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(i.Impl, option)
}

func (i *Instruction) UnmarshalBinary(decoder *bin.Decoder) (err error) {
	i.Version, err = decoder.ReadUint8()
	if err != nil {
		return fmt.Errorf("unable to read version: %w", err)
	}
	return i.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionDefVariant)
}

func (i *Instruction) MarshalBinary(encoder *bin.Encoder) error {
	err := encoder.WriteUint8(i.Version)
	if err != nil {
		return fmt.Errorf("unable to write instruction version: %w", err)
	}

	err = encoder.WriteUint32(i.TypeID, binary.LittleEndian)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(i.Impl)
}

var InstructionDefVariant = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{"initialize_market", (*InstructionInitializeMarket)(nil)},
	{"new_order", (*InstructionNewOrder)(nil)},
	{"match_orders", (*InstructionMatchOrder)(nil)},
	{"consume_events", (*InstructionConsumeEvents)(nil)},
	{"cancel_order", (*InstructionCancelOrder)(nil)},
	{"settle_funds", (*InstructionSettleFunds)(nil)},
	{"cancel_order_by_client_id", (*InstructionCancelOrderByClientId)(nil)},
	{"disable_market", (*InstructionDisableMarketAccounts)(nil)},
	{"sweep_fees", (*InstructionSweepFees)(nil)},
	{"new_order_v2", (*InstructionNewOrderV2)(nil)},
})

type InitializeMarketAccounts struct {
	Market        *solana.AccountMeta `text:"linear,notype"`
	SPLCoinToken  *solana.AccountMeta `text:"linear,notype"`
	SPLPriceToken *solana.AccountMeta `text:"linear,notype"`
	CoinMint      *solana.AccountMeta `text:"linear,notype"`
	PriceMint     *solana.AccountMeta `text:"linear,notype"`
}

type InstructionInitializeMarket struct {
	BaseLotSize        uint64
	QuoteLotSize       uint64
	FeeRateBps         uint16
	VaultSignerNonce   uint64
	QuoteDustThreshold uint64

	Accounts *InitializeMarketAccounts `bin:"-"`
}

func (i *InstructionInitializeMarket) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 9 {
		return fmt.Errorf("insuficient account, Initialize Market requires at-least 8 accounts not %d", len(accounts))
	}
	i.Accounts = &InitializeMarketAccounts{
		Market:        accounts[instructionActIdx[0]],
		SPLCoinToken:  accounts[instructionActIdx[5]],
		SPLPriceToken: accounts[instructionActIdx[6]],
		CoinMint:      accounts[instructionActIdx[7]],
		PriceMint:     accounts[instructionActIdx[8]],
	}
	return nil
}

type NewOrderAccounts struct {
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
}

type InstructionNewOrder struct {
	Side        Side
	LimitPrice  uint64
	MaxQuantity uint64
	OrderType   OrderType
	ClientID    uint64

	Accounts *NewOrderAccounts `bin:"-"`
}

func (i *InstructionNewOrder) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insuficient account, New Order requires at-least 10 accounts not %d", len(accounts))
	}
	i.Accounts = &NewOrderAccounts{
		Market:          accounts[0],
		OpenOrders:      accounts[1],
		RequestQueue:    accounts[2],
		Payer:           accounts[3],
		Owner:           accounts[4],
		CoinVault:       accounts[5],
		PCVault:         accounts[6],
		SPLTokenProgram: accounts[7],
		Rent:            accounts[8],
	}

	if len(accounts) >= 10 {
		i.Accounts.SRMDiscountAccount = accounts[9]
	}

	return nil
}

type MatchOrderAccounts struct {
	Market            *solana.AccountMeta `text:"linear,notype"`
	RequestQueue      *solana.AccountMeta `text:"linear,notype"`
	EventQueue        *solana.AccountMeta `text:"linear,notype"`
	Bids              *solana.AccountMeta `text:"linear,notype"`
	Asks              *solana.AccountMeta `text:"linear,notype"`
	CoinFeeReceivable *solana.AccountMeta `text:"linear,notype"`
	PCFeeReceivable   *solana.AccountMeta `text:"linear,notype"`
}

type InstructionMatchOrder struct {
	Limit uint16

	Accounts *MatchOrderAccounts `bin:"-"`
}

func (i *InstructionMatchOrder) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 7 {
		return fmt.Errorf("insuficient account, Match Order requires at-least 7 accounts not %d\n", len(accounts))
	}
	i.Accounts = &MatchOrderAccounts{
		Market:            accounts[instructionActIdx[0]],
		RequestQueue:      accounts[instructionActIdx[1]],
		EventQueue:        accounts[instructionActIdx[2]],
		Bids:              accounts[instructionActIdx[3]],
		Asks:              accounts[instructionActIdx[4]],
		CoinFeeReceivable: accounts[instructionActIdx[5]],
		PCFeeReceivable:   accounts[instructionActIdx[6]],
	}
	return nil
}

type ConsumeEventsAccounts struct {
	OpenOrders        []*solana.AccountMeta `text:"linear,notype"`
	Market            *solana.AccountMeta   `text:"linear,notype"`
	EventQueue        *solana.AccountMeta   `text:"linear,notype"`
	CoinFeeReceivable *solana.AccountMeta   `text:"linear,notype"`
	PCFeeReceivable   *solana.AccountMeta   `text:"linear,notype"`
}

type InstructionConsumeEvents struct {
	Limit uint16

	Accounts *ConsumeEventsAccounts `bin:"-"`
}

func (i *InstructionConsumeEvents) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 4 {
		return fmt.Errorf("insuficient account, Consume Events requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &ConsumeEventsAccounts{
		Market:            accounts[instructionActIdx[len(instructionActIdx)-4]],
		EventQueue:        accounts[instructionActIdx[len(instructionActIdx)-3]],
		CoinFeeReceivable: accounts[instructionActIdx[len(instructionActIdx)-2]],
		PCFeeReceivable:   accounts[instructionActIdx[len(instructionActIdx)-1]],
	}

	for itr := 0; itr < len(instructionActIdx)-4; itr++ {
		i.Accounts.OpenOrders = append(i.Accounts.OpenOrders, accounts[instructionActIdx[itr]])
	}

	return nil
}

type CancelOrderAccounts struct {
	Market       *solana.AccountMeta `text:"linear,notype"`
	OpenOrders   *solana.AccountMeta `text:"linear,notype"`
	RequestQueue *solana.AccountMeta `text:"linear,notype"`
	Owner        *solana.AccountMeta `text:"linear,notype"`
}

type InstructionCancelOrder struct {
	Side          Side
	OrderID       bin.Uint128
	OpenOrders    solana.PublicKey
	OpenOrderSlot uint8

	Accounts *CancelOrderAccounts `bin:"-"`
}

func (i *InstructionCancelOrder) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 4 {
		return fmt.Errorf("insuficient account, Cancel Order requires at-least 4 accounts not %d\n", len(accounts))
	}
	i.Accounts = &CancelOrderAccounts{
		Market:       accounts[instructionActIdx[0]],
		OpenOrders:   accounts[instructionActIdx[1]],
		RequestQueue: accounts[instructionActIdx[2]],
		Owner:        accounts[instructionActIdx[3]],
	}

	return nil
}

type SettleFundsAccounts struct {
	Market           *solana.AccountMeta `text:"linear,notype"`
	OpenOrders       *solana.AccountMeta `text:"linear,notype"`
	Owner            *solana.AccountMeta `text:"linear,notype"`
	CoinVault        *solana.AccountMeta `text:"linear,notype"`
	PCVault          *solana.AccountMeta `text:"linear,notype"`
	CoinWallet       *solana.AccountMeta `text:"linear,notype"`
	PCWallet         *solana.AccountMeta `text:"linear,notype"`
	Signer           *solana.AccountMeta `text:"linear,notype"`
	SPLTokenProgram  *solana.AccountMeta `text:"linear,notype"`
	ReferrerPCWallet *solana.AccountMeta `text:"linear,notype"`
}

type InstructionSettleFunds struct {
	Accounts *SettleFundsAccounts `bin:"-"`
}

func (i *InstructionSettleFunds) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 9 {
		return fmt.Errorf("insuficient account, Settle Funds requires at-least 10 accounts not %d", len(accounts))
	}
	i.Accounts = &SettleFundsAccounts{
		Market:          accounts[instructionActIdx[0]],
		OpenOrders:      accounts[instructionActIdx[1]],
		Owner:           accounts[instructionActIdx[2]],
		CoinVault:       accounts[instructionActIdx[3]],
		PCVault:         accounts[instructionActIdx[4]],
		CoinWallet:      accounts[instructionActIdx[5]],
		PCWallet:        accounts[instructionActIdx[6]],
		Signer:          accounts[instructionActIdx[7]],
		SPLTokenProgram: accounts[instructionActIdx[8]],
	}

	if len(instructionActIdx) >= 10 {
		i.Accounts.ReferrerPCWallet = accounts[instructionActIdx[9]]
	}

	return nil
}

type CancelOrderByClientIdAccounts struct {
	Market       *solana.AccountMeta `text:"linear,notype"`
	OpenOrders   *solana.AccountMeta `text:"linear,notype"`
	RequestQueue *solana.AccountMeta `text:"linear,notype"`
	Owner        *solana.AccountMeta `text:"linear,notype"`
}

type InstructionCancelOrderByClientId struct {
	ClientID uint64

	Accounts *CancelOrderByClientIdAccounts `bin:"-"`
}

func (i *InstructionCancelOrderByClientId) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 4 {
		return fmt.Errorf("insuficient account, Cancel Order By Client Id requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderByClientIdAccounts{
		Market:       accounts[instructionActIdx[0]],
		OpenOrders:   accounts[instructionActIdx[1]],
		RequestQueue: accounts[instructionActIdx[2]],
		Owner:        accounts[instructionActIdx[3]],
	}

	return nil
}

type DisableMarketAccounts struct {
	Market           *solana.AccountMeta `text:"linear,notype"`
	DisableAuthority *solana.AccountMeta `text:"linear,notype"`
}

type InstructionDisableMarketAccounts struct {
	Accounts *DisableMarketAccounts `bin:"-"`
}

func (i *InstructionDisableMarketAccounts) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 2 {
		return fmt.Errorf("insuficient account, Disable Market requires at-least 2 accounts not %d", len(accounts))
	}

	i.Accounts = &DisableMarketAccounts{
		Market:           accounts[instructionActIdx[0]],
		DisableAuthority: accounts[instructionActIdx[1]],
	}

	return nil
}

type SweepFeesAccounts struct {
	Market               *solana.AccountMeta `text:"linear,notype"`
	PCVault              *solana.AccountMeta `text:"linear,notype"`
	FeeSweepingAuthority *solana.AccountMeta `text:"linear,notype"`
	FeeReceivableAccount *solana.AccountMeta `text:"linear,notype"`
	VaultSigner          *solana.AccountMeta `text:"linear,notype"`
	SPLTokenProgram      *solana.AccountMeta `text:"linear,notype"`
}

type InstructionSweepFees struct {
	Accounts *SweepFeesAccounts `bin:"-"`
}

func (i *InstructionSweepFees) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 6 {
		return fmt.Errorf("insuficient account, Sweep Fees requires at-least 6 accounts not %d", len(accounts))
	}

	i.Accounts = &SweepFeesAccounts{
		Market:               accounts[instructionActIdx[0]],
		PCVault:              accounts[instructionActIdx[1]],
		FeeSweepingAuthority: accounts[instructionActIdx[2]],
		FeeReceivableAccount: accounts[instructionActIdx[3]],
		VaultSigner:          accounts[instructionActIdx[4]],
		SPLTokenProgram:      accounts[instructionActIdx[5]],
	}

	return nil
}

type NewOrderV2Accounts struct {
	Market          *solana.AccountMeta `text:"linear,notype"` // the market
	OpenOrders      *solana.AccountMeta `text:"linear,notype"` // the OpenOrders account to use
	RequestQueue    *solana.AccountMeta `text:"linear,notype"` // the request queue
	Payer           *solana.AccountMeta `text:"linear,notype"` // the (coin or price currency) account paying for the order
	Owner           *solana.AccountMeta `text:"linear,notype"` // owner of the OpenOrders account
	CoinVault       *solana.AccountMeta `text:"linear,notype"` // coin vault
	PCVault         *solana.AccountMeta `text:"linear,notype"` // pc vault
	SPLTokenProgram *solana.AccountMeta `text:"linear,notype"` // spl token program
	RentSysvar      *solana.AccountMeta `text:"linear,notype"` // the rent sysvar
	FeeDiscount     *solana.AccountMeta `text:"linear,notype"` // (optional) the (M)SRM account used for fee discounts
}

type SelfTradeBehavior uint32

const (
	SelfTradeBehaviorDecrementTake = iota
	SelfTradeBehaviorCancelProvide
)

type InstructionNewOrderV2 struct {
	Side              Side
	LimitPrice        uint64
	MaxQuantity       uint64
	OrderType         OrderType
	ClientID          uint64
	SelfTradeBehavior SelfTradeBehavior

	Accounts *NewOrderV2Accounts `bin:"-"`
}

func (i *InstructionNewOrderV2) SetAccounts(accounts []*solana.AccountMeta, instructionActIdx []uint8) error {
	if len(instructionActIdx) < 10 {
		return fmt.Errorf("insuficient account, New Order V2 requires at-least 10 accounts not %d", len(accounts))
	}

	i.Accounts = &NewOrderV2Accounts{
		Market:          accounts[instructionActIdx[0]],
		OpenOrders:      accounts[instructionActIdx[1]],
		RequestQueue:    accounts[instructionActIdx[2]],
		Payer:           accounts[instructionActIdx[3]],
		Owner:           accounts[instructionActIdx[4]],
		CoinVault:       accounts[instructionActIdx[5]],
		PCVault:         accounts[instructionActIdx[6]],
		SPLTokenProgram: accounts[instructionActIdx[7]],
		RentSysvar:      accounts[instructionActIdx[8]],
		FeeDiscount:     accounts[instructionActIdx[9]],
	}

	return nil
}
