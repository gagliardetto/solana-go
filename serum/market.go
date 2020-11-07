package serum

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

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

func divideBnToNumber(numerator, denomiator *big.Int) *big.Int {
	quotient := new(big.Int).Quo(numerator, denomiator)
	rem := new(big.Int).Mod(numerator, denomiator)
	gcd := new(big.Int).GCD(nil, nil, rem, denomiator)
	n1 := new(big.Int).Quo(rem, gcd)
	d1 := new(big.Int).Quo(denomiator, gcd)
	div1 := new(big.Int).Quo(n1, d1)
	return new(big.Int).Add(quotient, div1)
}

func (m *MarketMeta) PriceLotsToNumber(price *big.Int) *big.Int {
	ratio := new(big.Int).Mul(big.NewInt(int64(m.MarketV2.QuoteLotSize)), m.baseSplTokenMultiplier())
	numerator := new(big.Int).Mul(price, ratio)
	denomiator := new(big.Int).Mul(big.NewInt(int64(m.MarketV2.BaseLotSize)), m.quoteSplTokenMultiplier())
	return divideBnToNumber(numerator, denomiator)
}

func (m *MarketMeta) BaseSizeLotsToNumber(size *big.Int) *big.Float {
	numerator := new(big.Int).Mul(size, big.NewInt(int64(m.MarketV2.BaseLotSize)))
	denomiator := m.baseSplTokenMultiplier()
	return new(big.Float).Quo(new(big.Float).SetInt(numerator), new(big.Float).SetInt(denomiator))
}

func (m *MarketMeta) PriceNumberToLots(price *big.Int) *big.Float {
	numerator := new(big.Int).Mul(price, m.quoteSplTokenMultiplier())
	numerator = new(big.Int).Mul(numerator, big.NewInt(int64(m.MarketV2.BaseLotSize)))

	denomiator := new(big.Int).Mul(m.baseSplTokenMultiplier(), big.NewInt(int64(m.MarketV2.QuoteLotSize)))
	return new(big.Float).Quo(new(big.Float).SetInt(numerator), new(big.Float).SetInt(denomiator))
}

func KnownMarket() ([]*MarketMeta, error) {
	f, err := os.Open("serum/markets.json")
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve known markets: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var markets []*MarketMeta
	err = dec.Decode(&markets)
	if err != nil {
		return nil, fmt.Errorf("unable to decode known markets: %w", err)
	}
	return markets, nil
}
