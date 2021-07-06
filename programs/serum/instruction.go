package serum

import (
	"encoding/binary"
	"fmt"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text"
)

func init() {
	solana.RegisterInstructionDecoder(DEXProgramIDV2, registryDecodeInstruction)
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

var InstructionDefVariant = bin.NewVariantDefinition(bin.Uint32TypeIDEncoding, []bin.VariantType{
	{Name: "initialize_market", Type: (*InstructionInitializeMarket)(nil)},
	{Name: "new_order", Type: (*InstructionNewOrder)(nil)},
	{Name: "match_orders", Type: (*InstructionMatchOrder)(nil)},
	{Name: "consume_events", Type: (*InstructionConsumeEvents)(nil)},
	{Name: "cancel_order", Type: (*InstructionCancelOrder)(nil)},
	{Name: "settle_funds", Type: (*InstructionSettleFunds)(nil)},
	{Name: "cancel_order_by_client_id", Type: (*InstructionCancelOrderByClientId)(nil)},
	{Name: "disable_market", Type: (*InstructionDisableMarketAccounts)(nil)},
	{Name: "sweep_fees", Type: (*InstructionSweepFees)(nil)},
	{Name: "new_order_v2", Type: (*InstructionNewOrderV2)(nil)},

	// Added in DEX V3
	{Name: "new_order_v3", Type: (*InstructionNewOrderV3)(nil)},
	{Name: "cancel_order_v2", Type: (*InstructionCancelOrderV2)(nil)},
	{Name: "cancel_order_by_client_id_v2", Type: (*InstructionCancelOrderByClientIdV2)(nil)},
	{Name: "send_take", Type: (*InstructionSendTake)(nil)},
})

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

func (i *InstructionInitializeMarket) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insufficient account, Initialize Market requires at-least 8 accounts not %d", len(accounts))
	}
	i.Accounts = &InitializeMarketAccounts{
		Market:        accounts[0],
		SPLCoinToken:  accounts[5],
		SPLPriceToken: accounts[6],
		CoinMint:      accounts[7],
		PriceMint:     accounts[8],
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

// InstructionNewOrder seems to be unused after DEX v3 (unconfirmed claim)
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
		return fmt.Errorf("insufficient account, New Order requires at-least 10 accounts not %d", len(accounts))
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

// InstructionMatchOrder seems to be unused after DEX v3 (unconfirmed claim)
type InstructionMatchOrder struct {
	Limit uint16

	Accounts *MatchOrderAccounts `bin:"-"`
}

func (i *InstructionMatchOrder) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 7 {
		return fmt.Errorf("insufficient account, Match Order requires at-least 7 accounts not %d\n", len(accounts))
	}
	i.Accounts = &MatchOrderAccounts{
		Market:            accounts[0],
		RequestQueue:      accounts[1],
		EventQueue:        accounts[2],
		Bids:              accounts[3],
		Asks:              accounts[4],
		CoinFeeReceivable: accounts[5],
		PCFeeReceivable:   accounts[6],
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

func (i *InstructionConsumeEvents) SetAccounts(accounts []*solana.AccountMeta) error {
	l := len(accounts)
	if l < 4 {
		return fmt.Errorf("insufficient account, Consume Events requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &ConsumeEventsAccounts{
		Market:            accounts[l-4],
		EventQueue:        accounts[l-3],
		CoinFeeReceivable: accounts[l-2],
		PCFeeReceivable:   accounts[l-1],
	}

	for idx := 0; idx < l-4; idx++ {
		i.Accounts.OpenOrders = append(i.Accounts.OpenOrders, accounts[idx])
	}

	return nil
}

type CancelOrderAccounts struct {
	Market       *solana.AccountMeta `text:"linear,notype"`
	OpenOrders   *solana.AccountMeta `text:"linear,notype"`
	RequestQueue *solana.AccountMeta `text:"linear,notype"`
	Owner        *solana.AccountMeta `text:"linear,notype"`
}

// InstructionCancelOrder seems to be unused after DEX v3 (unconfirmed claim)
type InstructionCancelOrder struct {
	Side          Side
	OrderID       bin.Uint128
	OpenOrders    solana.PublicKey
	OpenOrderSlot uint8

	Accounts *CancelOrderAccounts `bin:"-"`
}

func (i *InstructionCancelOrder) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 4 {
		return fmt.Errorf("insufficient account, Cancel Order requires at-least 4 accounts not %d\n", len(accounts))
	}
	i.Accounts = &CancelOrderAccounts{
		Market:       accounts[0],
		OpenOrders:   accounts[1],
		RequestQueue: accounts[2],
		Owner:        accounts[3],
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

func (i *InstructionSettleFunds) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insufficient account, Settle Funds requires at-least 10 accounts not %d", len(accounts))
	}
	i.Accounts = &SettleFundsAccounts{
		Market:          accounts[0],
		OpenOrders:      accounts[1],
		Owner:           accounts[2],
		CoinVault:       accounts[3],
		PCVault:         accounts[4],
		CoinWallet:      accounts[5],
		PCWallet:        accounts[6],
		Signer:          accounts[7],
		SPLTokenProgram: accounts[8],
	}

	if len(accounts) >= 10 {
		i.Accounts.ReferrerPCWallet = accounts[9]
	}

	return nil
}

type CancelOrderByClientIdAccounts struct {
	Market       *solana.AccountMeta `text:"linear,notype"`
	OpenOrders   *solana.AccountMeta `text:"linear,notype"`
	RequestQueue *solana.AccountMeta `text:"linear,notype"`
	Owner        *solana.AccountMeta `text:"linear,notype"`
}

// InstructionCancelOrderByClientId seems to be unused after DEX v3 (unconfirmed claim)
type InstructionCancelOrderByClientId struct {
	ClientID uint64

	Accounts *CancelOrderByClientIdAccounts `bin:"-"`
}

func (i *InstructionCancelOrderByClientId) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 4 {
		return fmt.Errorf("insufficient account, Cancel Order By Client Id requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderByClientIdAccounts{
		Market:       accounts[0],
		OpenOrders:   accounts[1],
		RequestQueue: accounts[2],
		Owner:        accounts[3],
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

func (i *InstructionDisableMarketAccounts) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 2 {
		return fmt.Errorf("insufficient account, Disable Market requires at-least 2 accounts not %d", len(accounts))
	}

	i.Accounts = &DisableMarketAccounts{
		Market:           accounts[0],
		DisableAuthority: accounts[1],
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

func (i *InstructionSweepFees) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 6 {
		return fmt.Errorf("insufficient account, Sweep Fees requires at-least 6 accounts not %d", len(accounts))
	}

	i.Accounts = &SweepFeesAccounts{
		Market:               accounts[0],
		PCVault:              accounts[1],
		FeeSweepingAuthority: accounts[2],
		FeeReceivableAccount: accounts[3],
		VaultSigner:          accounts[4],
		SPLTokenProgram:      accounts[5],
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

	// Added in DEX V3

	SelfTradeBehaviorAbortTransaction
)

// InstructionNewOrderV2 seems to be unused after DEX v3 (unconfirmed claim)
type InstructionNewOrderV2 struct {
	Side              Side
	LimitPrice        uint64
	MaxQuantity       uint64
	OrderType         OrderType
	ClientID          uint64
	SelfTradeBehavior SelfTradeBehavior

	Accounts *NewOrderV2Accounts `bin:"-"`
}

func (i *InstructionNewOrderV2) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insufficient account, New Order V2 requires at-least 9 accounts + 1 optional not %d", len(accounts))
	}

	i.Accounts = &NewOrderV2Accounts{
		Market:          accounts[0],
		OpenOrders:      accounts[1],
		RequestQueue:    accounts[2],
		Payer:           accounts[3],
		Owner:           accounts[4],
		CoinVault:       accounts[5],
		PCVault:         accounts[6],
		SPLTokenProgram: accounts[7],
		RentSysvar:      accounts[8],
	}

	if len(accounts) == 10 {
		i.Accounts.FeeDiscount = accounts[9]
	}

	return nil
}

// DEX V3 Support

type NewOrderV3Accounts struct {
	Market          *solana.AccountMeta `text:"linear,notype"` // the market
	OpenOrders      *solana.AccountMeta `text:"linear,notype"` // the OpenOrders account to use
	RequestQueue    *solana.AccountMeta `text:"linear,notype"` // the request queue
	EventQueue      *solana.AccountMeta `text:"linear,notype"` // the event queue
	Bidder          *solana.AccountMeta `text:"linear,notype"` // bids
	Asker           *solana.AccountMeta `text:"linear,notype"` // asks
	Payer           *solana.AccountMeta `text:"linear,notype"` // the (coin or price currency) account paying for the order
	Owner           *solana.AccountMeta `text:"linear,notype"` // owner of the OpenOrders account
	CoinVault       *solana.AccountMeta `text:"linear,notype"` // coin vault
	PCVault         *solana.AccountMeta `text:"linear,notype"` // pc vault
	SPLTokenProgram *solana.AccountMeta `text:"linear,notype"` // spl token program
	RentSysvar      *solana.AccountMeta `text:"linear,notype"` // the rent sysvar
	FeeDiscount     *solana.AccountMeta `text:"linear,notype"` // (optional) the (M)SRM account used for fee discounts
}

type InstructionNewOrderV3 struct {
	Side                             Side
	LimitPrice                       uint64
	MaxCoinQuantity                  uint64
	MaxNativePCQuantityIncludingFees uint64
	SelfTradeBehavior                SelfTradeBehavior
	OrderType                        OrderType
	ClientOrderID                    uint64
	Limit                            uint16

	Accounts *NewOrderV3Accounts `bin:"-"`
}

func (i *InstructionNewOrderV3) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 13 {
		return fmt.Errorf("insufficient account, New Order V3 requires at-least 13 accounts not %d", len(accounts))
	}

	i.Accounts = &NewOrderV3Accounts{
		Market:          accounts[0],
		OpenOrders:      accounts[1],
		RequestQueue:    accounts[2],
		EventQueue:      accounts[3],
		Bidder:          accounts[4],
		Asker:           accounts[5],
		Payer:           accounts[6],
		Owner:           accounts[7],
		CoinVault:       accounts[8],
		PCVault:         accounts[9],
		SPLTokenProgram: accounts[10],
		RentSysvar:      accounts[11],
		FeeDiscount:     accounts[12],
	}

	return nil
}

type CancelOrderV2Accounts struct {
	Market     *solana.AccountMeta `text:"linear,notype"` // 0. `[writable]` market
	Bids       *solana.AccountMeta `text:"linear,notype"` // 1. `[writable]` bids
	Asks       *solana.AccountMeta `text:"linear,notype"` // 2. `[writable]` asks
	OpenOrders *solana.AccountMeta `text:"linear,notype"` // 3. `[writable]` OpenOrders
	Owner      *solana.AccountMeta `text:"linear,notype"` // 4. `[signer]` the OpenOrders owner
	EventQueue *solana.AccountMeta `text:"linear,notype"` // 5. `[writable]` event_q
}

type InstructionCancelOrderV2 struct {
	Side    Side
	OrderID bin.Uint128

	Accounts *CancelOrderV2Accounts `bin:"-"`
}

func (i *InstructionCancelOrderV2) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 6 {
		return fmt.Errorf("insufficient account, Cancel Order V2 requires at-least 6 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderV2Accounts{
		Market:     accounts[0],
		Bids:       accounts[1],
		Asks:       accounts[2],
		OpenOrders: accounts[3],
		Owner:      accounts[4],
		EventQueue: accounts[5],
	}

	return nil
}

type CancelOrderByClientIdV2Accounts struct {
	Market     *solana.AccountMeta `text:"linear,notype"` // 0. `[writable]` market
	Bids       *solana.AccountMeta `text:"linear,notype"` // 1. `[writable]` bids
	Asks       *solana.AccountMeta `text:"linear,notype"` // 2. `[writable]` asks
	OpenOrders *solana.AccountMeta `text:"linear,notype"` // 3. `[writable]` OpenOrders
	Owner      *solana.AccountMeta `text:"linear,notype"` // 4. `[signer]` the OpenOrders owner
	EventQueue *solana.AccountMeta `text:"linear,notype"` // 5. `[writable]` event_q
}

type InstructionCancelOrderByClientIdV2 struct {
	ClientID uint64

	Accounts *CancelOrderByClientIdV2Accounts `bin:"-"`
}

func (i *InstructionCancelOrderByClientIdV2) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 6 {
		return fmt.Errorf("insufficient account, Cancel Order By Client Id V2 requires at-least 6 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderByClientIdV2Accounts{
		Market:     accounts[0],
		Bids:       accounts[1],
		Asks:       accounts[2],
		OpenOrders: accounts[3],
		Owner:      accounts[4],
		EventQueue: accounts[5],
	}

	return nil
}

// InstructionSendTakeAccounts defined from comment in serum-dex contract code, was never able to validate it's correct
type InstructionSendTakeAccounts struct {
	Market     *solana.AccountMeta `text:"linear,notype"` // 0. `[writable]` market
	Bids       *solana.AccountMeta `text:"linear,notype"` // 1. `[writable]` bids
	Asks       *solana.AccountMeta `text:"linear,notype"` // 2. `[writable]` asks
	OpenOrders *solana.AccountMeta `text:"linear,notype"` // 3. `[writable]` OpenOrders
	Owner      *solana.AccountMeta `text:"linear,notype"` // 4. `[]`
}

type InstructionSendTake struct {
	Side                             Side
	LimitPrice                       uint64
	MaxCoinQuantity                  uint64
	MaxNativePCQuantityIncludingFees uint64
	MinCoinQuantity                  uint64
	MinNativePCQuantity              uint64
	Limit                            uint16

	Accounts *InstructionSendTakeAccounts `bin:"-"`
}

func (i *InstructionSendTake) SetAccounts(accounts []*solana.AccountMeta) error {
	if len(accounts) < 5 {
		return fmt.Errorf("insufficient account, Send Take requires at-least 5 accounts not %d", len(accounts))
	}
	i.Accounts = &InstructionSendTakeAccounts{
		Market:     accounts[0],
		Bids:       accounts[1],
		Asks:       accounts[2],
		OpenOrders: accounts[3],
		Owner:      accounts[4],
	}
	return nil
}
