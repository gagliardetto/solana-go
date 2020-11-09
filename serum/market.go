package serum

import (
	"math/big"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/token"
)

type MarketMeta struct {
	Address    solana.PublicKey `json:"address"`
	Name       string           `json:"name"`
	Deprecated bool             `json:"deprecated"`
	QuoteMint  *token.Mint
	BaseMint   *token.Mint

	MarketV2 *MarketV2
}

func (m *MarketMeta) baseSplTokenMultiplier() *big.Int {
	return solana.DecimalsInBigInt(uint32(m.BaseMint.Decimals))
}

func (m *MarketMeta) quoteSplTokenMultiplier() *big.Int {
	return solana.DecimalsInBigInt(uint32(m.QuoteMint.Decimals))
}

func divideBnToNumber(numerator, denomiator *big.Float) *big.Float {
	return F().Quo(numerator, denomiator)
}

func (m *MarketMeta) PriceLotsToNumber(price *big.Int) *big.Float {
	ratio := I().Mul(I().SetInt64(int64(m.MarketV2.QuoteLotSize)), m.baseSplTokenMultiplier())
	numerator := F().Mul(F().SetInt(price), F().SetInt(ratio))
	denomiator := F().Mul(F().SetFloat64(float64(m.MarketV2.BaseLotSize)), F().SetInt(m.quoteSplTokenMultiplier()))
	v := divideBnToNumber(numerator, denomiator)
	return v
}

func (m *MarketMeta) BaseSizeLotsToNumber(size *big.Int) *big.Float {
	numerator := I().Mul(size, I().SetInt64(int64(m.MarketV2.BaseLotSize)))
	denomiator := m.baseSplTokenMultiplier()
	return F().Quo(F().SetInt(numerator), F().SetInt(denomiator))
}

func (m *MarketMeta) PriceNumberToLots(price *big.Int) *big.Float {
	numerator := I().Mul(price, m.quoteSplTokenMultiplier())
	numerator = I().Mul(numerator, big.NewInt(int64(m.MarketV2.BaseLotSize)))
	denomiator := I().Mul(m.baseSplTokenMultiplier(), I().SetInt64(int64(m.MarketV2.QuoteLotSize)))
	return F().Quo(F().SetInt(numerator), F().SetInt(denomiator))
}

func I() *big.Int {
	return new(big.Int)
}

func F() *big.Float {
	return new(big.Float)
}
