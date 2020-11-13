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

func registryDecodeInstruction(rawInstruction *solana.CompiledInstruction) (interface{}, error) {
	inst, err := DecodeInstruction(rawInstruction)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(rawInstruction *solana.CompiledInstruction) (*Instruction, error) {
	var inst *Instruction
	if err := bin.NewDecoder(rawInstruction.Data).Decode(&inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction for serum program: %w", err)
	}
	return inst, nil
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

type InstructionInitializeMarket struct {
	BaseLotSize        uint64
	QuoteLotSize       uint64
	FeeRateBps         uint16
	VaultSignerNonce   uint64
	QuoteDustThreshold uint64
}

type InstructionNewOrder struct {
	Side        uint32
	LimitPrice  uint64
	MaxQuantity uint64
	OrderType   uint32
	ClientID    uint64
}

type InstructionMatchOrder struct {
	Limit uint16
}

type InstructionConsumeEvents struct {
	Limit uint16
}

type InstructionCancelOrder struct {
	Side          uint32
	OrderID       bin.Uint128
	OpenOrders    solana.PublicKey
	OpenOrderSlot uint8
}

type InstructionSettleFunds struct {
}

type InstructionCancelOrderByClientId struct {
	ClientID uint64
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

//
//export const INSTRUCTION_LAYOUT = new VersionedLayout(
//0,
//union(u32('instruction')),
//);
//INSTRUCTION_LAYOUT.inner.addVariant(
//0,
//struct([
//u64('baseLotSize'),
//u64('quoteLotSize'),
//u16('feeRateBps'),
//u64('vaultSignerNonce'),
//u64('quoteDustThreshold'),
//]),
//'initializeMarket',
//);
//INSTRUCTION_LAYOUT.inner.addVariant(
//1,
//struct([
//sideLayout('side'),
//u64('limitPrice'),
//u64('maxQuantity'),
//orderTypeLayout('orderType'),
//u64('clientId'),
//]),
//'newOrder',
//);
//INSTRUCTION_LAYOUT.inner.addVariant(2, struct([u16('limit')]), 'matchOrders');
//INSTRUCTION_LAYOUT.inner.addVariant(3, struct([u16('limit')]), 'consumeEvents');
//INSTRUCTION_LAYOUT.inner.addVariant(
//4,
//struct([
//sideLayout('side'),
//u128('orderId'),
//publicKeyLayout('openOrders'),
//u8('openOrdersSlot'),
//]),
//'cancelOrder',
//);
//INSTRUCTION_LAYOUT.inner.addVariant(5, struct([]), 'settleFunds');
//INSTRUCTION_LAYOUT.inner.addVariant(
//6,
//struct([u64('clientId')]),
//'cancelOrderByClientId',
//);
