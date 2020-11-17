package serum

import (
	"fmt"

	bin "github.com/dfuse-io/binary"
	"github.com/dfuse-io/solana-go"
)

var DEX_PROGRAM_ID = solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")

func init() {
	solana.RegisterInstructionDecoder(DEX_PROGRAM_ID, registryDecodeInstruction)
}

func registryDecodeInstruction(accounts []solana.PublicKey, rawInstruction *solana.CompiledInstruction) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, rawInstruction)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []solana.PublicKey, compiledInstruction *solana.CompiledInstruction) (*Instruction, error) {
	var inst *Instruction
	if err := bin.NewDecoder(compiledInstruction.Data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}

	if v, ok := inst.Impl.(AccountSettable); ok {
		err := v.setAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instructions: %w", err)
		}

	}
	return inst, nil
}

type AccountSettable interface {
	setAccounts(accounts []solana.PublicKey) error
}

type Instruction struct {
	bin.BaseVariant
	Version uint8
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

	err = encoder.WriteUint32(i.TypeID)
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
})

type InitializeMarketAccounts struct {
	Market        solana.AccountMeta
	SPLCoinToken  solana.AccountMeta
	SPLPriceToken solana.AccountMeta
	CoinMint      solana.AccountMeta
	PriceMint     solana.AccountMeta
}

type InstructionInitializeMarket struct {
	BaseLotSize        uint64
	QuoteLotSize       uint64
	FeeRateBps         uint16
	VaultSignerNonce   uint64
	QuoteDustThreshold uint64

	Accounts *InitializeMarketAccounts `bin="-"`
}

func (i *InstructionInitializeMarket) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insuficient account, Initialize Market requires at-least 8 accounts not %d", len(accounts))
	}
	i.Accounts = &InitializeMarketAccounts{
		Market:        solana.AccountMeta{accounts[0], false, true},
		SPLCoinToken:  solana.AccountMeta{accounts[5], false, true},
		SPLPriceToken: solana.AccountMeta{accounts[6], false, true},
		CoinMint:      solana.AccountMeta{accounts[7], false, false},
		PriceMint:     solana.AccountMeta{accounts[8], false, false},
	}
	return nil
}

type NewOrderAccounts struct {
	Market             solana.AccountMeta
	OpenOrders         solana.AccountMeta
	RequestQueue       solana.AccountMeta
	Payer              solana.AccountMeta
	Owner              solana.AccountMeta
	CoinVault          solana.AccountMeta
	PCVault            solana.AccountMeta
	SPLTokenProgram    solana.AccountMeta
	Rent               solana.AccountMeta
	SRMDiscountAccount *solana.AccountMeta
}

type InstructionNewOrder struct {
	Side        uint32
	LimitPrice  uint64
	MaxQuantity uint64
	OrderType   uint32
	ClientID    uint64

	Accounts *NewOrderAccounts `bin="-"`
}

func (i *InstructionNewOrder) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insuficient account, New Order requires at-least 10 accounts not %d", len(accounts))
	}
	i.Accounts = &NewOrderAccounts{
		Market:          solana.AccountMeta{accounts[0], false, true},
		OpenOrders:      solana.AccountMeta{accounts[1], false, true},
		RequestQueue:    solana.AccountMeta{accounts[2], false, true},
		Payer:           solana.AccountMeta{accounts[3], false, true},
		Owner:           solana.AccountMeta{accounts[4], true, false},
		CoinVault:       solana.AccountMeta{accounts[5], false, true},
		PCVault:         solana.AccountMeta{accounts[6], false, true},
		SPLTokenProgram: solana.AccountMeta{accounts[7], false, false},
		Rent:            solana.AccountMeta{accounts[8], false, false},
	}

	if len(accounts) >= 10 {
		i.Accounts.SRMDiscountAccount = &solana.AccountMeta{accounts[9], false, true}
	}

	return nil
}

type MatchOrderAccounts struct {
	Market            solana.AccountMeta
	RequestQueue      solana.AccountMeta
	EventQueue        solana.AccountMeta
	Bids              solana.AccountMeta
	Asks              solana.AccountMeta
	CoinFeeReceivable solana.AccountMeta
	PCFeeReceivable   solana.AccountMeta
}

type InstructionMatchOrder struct {
	Limit uint16

	Accounts *MatchOrderAccounts `bin:"-"`
}

func (i *InstructionMatchOrder) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 7 {
		return fmt.Errorf("insuficient account, Match Order requires at-least 7 accounts not %d", len(accounts))
	}
	i.Accounts = &MatchOrderAccounts{
		Market:            solana.AccountMeta{accounts[0], false, true},
		RequestQueue:      solana.AccountMeta{accounts[1], false, true},
		EventQueue:        solana.AccountMeta{accounts[2], false, true},
		Bids:              solana.AccountMeta{accounts[3], false, true},
		Asks:              solana.AccountMeta{accounts[4], false, true},
		CoinFeeReceivable: solana.AccountMeta{accounts[5], false, true},
		PCFeeReceivable:   solana.AccountMeta{accounts[6], false, true},
	}
	return nil
}

type ConsumeEventsAccounts struct {
	OpenOrders        []solana.AccountMeta
	Market            solana.AccountMeta
	EventQueue        solana.AccountMeta
	CoinFeeReceivable solana.AccountMeta
	PCFeeReceivable   solana.AccountMeta
}

type InstructionConsumeEvents struct {
	Limit uint16

	Accounts *ConsumeEventsAccounts `bin="-"`
}

func (i *InstructionConsumeEvents) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 4 {
		return fmt.Errorf("insuficient account, Consume Events requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &ConsumeEventsAccounts{
		Market:            solana.AccountMeta{accounts[len(accounts)-4], false, true},
		EventQueue:        solana.AccountMeta{accounts[len(accounts)-3], false, true},
		CoinFeeReceivable: solana.AccountMeta{accounts[len(accounts)-2], false, true},
		PCFeeReceivable:   solana.AccountMeta{accounts[len(accounts)-1], false, true},
	}

	for itr := 0; itr < len(accounts)-4; itr++ {
		i.Accounts.OpenOrders = append(i.Accounts.OpenOrders, solana.AccountMeta{accounts[itr], false, true})
	}

	return nil
}

type CancelOrderAccounts struct {
	Market       solana.AccountMeta
	OpenOrders   solana.AccountMeta
	RequestQueue solana.AccountMeta
	Owner        solana.AccountMeta
}

type InstructionCancelOrder struct {
	Side          uint32
	OrderID       bin.Uint128
	OpenOrders    solana.PublicKey
	OpenOrderSlot uint8

	Accounts *CancelOrderAccounts `bin="-"`
}

func (i *InstructionCancelOrder) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 4 {
		return fmt.Errorf("insuficient account, Cancel Order requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderAccounts{
		Market:       solana.AccountMeta{accounts[0], false, false},
		OpenOrders:   solana.AccountMeta{accounts[1], false, true},
		RequestQueue: solana.AccountMeta{accounts[2], false, true},
		Owner:        solana.AccountMeta{accounts[3], true, false},
	}

	return nil
}

type SettleFundsAccounts struct {
	Market           solana.AccountMeta
	OpenOrders       solana.AccountMeta
	Owner            solana.AccountMeta
	CoinVault        solana.AccountMeta
	PCVault          solana.AccountMeta
	CoinWallet       solana.AccountMeta
	PCWallet         solana.AccountMeta
	Signer           solana.AccountMeta
	SPLTokenProgram  solana.AccountMeta
	ReferrerPCWallet *solana.AccountMeta
}

type InstructionSettleFunds struct {
	Accounts *SettleFundsAccounts `bin="-"`
}

func (i *InstructionSettleFunds) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 9 {
		return fmt.Errorf("insuficient account, Settle Funds requires at-least 10 accounts not %d", len(accounts))
	}
	i.Accounts = &SettleFundsAccounts{
		Market:          solana.AccountMeta{accounts[0], false, true},
		OpenOrders:      solana.AccountMeta{accounts[1], false, true},
		Owner:           solana.AccountMeta{accounts[2], true, false},
		CoinVault:       solana.AccountMeta{accounts[3], false, true},
		PCVault:         solana.AccountMeta{accounts[4], false, true},
		CoinWallet:      solana.AccountMeta{accounts[5], false, true},
		PCWallet:        solana.AccountMeta{accounts[6], false, true},
		Signer:          solana.AccountMeta{accounts[7], false, false},
		SPLTokenProgram: solana.AccountMeta{accounts[8], false, false},
	}

	if len(accounts) >= 10 {
		i.Accounts.ReferrerPCWallet = &solana.AccountMeta{accounts[9], false, true}
	}

	return nil
}

type CancelOrderByClientIdAccounts struct {
	Market       solana.AccountMeta
	OpenOrders   solana.AccountMeta
	RequestQueue solana.AccountMeta
	Owner        solana.AccountMeta
}

type InstructionCancelOrderByClientId struct {
	ClientID uint64

	Accounts *CancelOrderByClientIdAccounts
}

func (i *InstructionCancelOrderByClientId) setAccounts(accounts []solana.PublicKey) error {
	if len(accounts) < 4 {
		return fmt.Errorf("insuficient account, Cancel Order By Client Id requires at-least 4 accounts not %d", len(accounts))
	}
	i.Accounts = &CancelOrderByClientIdAccounts{
		Market:       solana.AccountMeta{accounts[0], false, false},
		OpenOrders:   solana.AccountMeta{accounts[1], false, true},
		RequestQueue: solana.AccountMeta{accounts[2], false, true},
		Owner:        solana.AccountMeta{accounts[3], true, false},
	}

	return nil
}

type SideLayoutType string

const (
	SideLayoutTypeUnknown SideLayoutType = "UNKNOWN"
	SideLayoutTypeBid     SideLayoutType = "BID"
	SideLayoutTypeAsk     SideLayoutType = "ASK"
)

type SideLayout uint32

func (s SideLayout) getSide() SideLayoutType {
	switch s {
	case 0:
		return SideLayoutTypeBid
	case 1:
		return SideLayoutTypeAsk
	}
	return SideLayoutTypeUnknown
}

type OrderType string

const (
	OrderTypeUnknown           OrderType = "UNKNOWN"
	OrderTypeLimit             OrderType = "LIMIT"
	OrderTypeImmediateOrCancel OrderType = "IMMEDIATE_OR_CANCEL"
	OrderTypePostOnly          OrderType = "POST_ONLY"
)

type OrderTypeLayout uint32

func (o OrderTypeLayout) getOrderType() OrderType {
	switch o {
	case 0:
		return OrderTypeLimit
	case 1:
		return OrderTypeImmediateOrCancel
	case 2:
		return OrderTypePostOnly
	}
	return OrderTypeUnknown
}
