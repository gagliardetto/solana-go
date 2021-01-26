package serum

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
)

var _10b = big.NewInt(10)

var decimalsBigInt = []*big.Int{
	new(big.Int).Exp(_10b, big.NewInt(1), nil),
	new(big.Int).Exp(_10b, big.NewInt(2), nil),
	new(big.Int).Exp(_10b, big.NewInt(3), nil),
	new(big.Int).Exp(_10b, big.NewInt(4), nil),
	new(big.Int).Exp(_10b, big.NewInt(5), nil),
	new(big.Int).Exp(_10b, big.NewInt(6), nil),
	new(big.Int).Exp(_10b, big.NewInt(7), nil),
	new(big.Int).Exp(_10b, big.NewInt(8), nil),
	new(big.Int).Exp(_10b, big.NewInt(9), nil),
	new(big.Int).Exp(_10b, big.NewInt(10), nil),
	new(big.Int).Exp(_10b, big.NewInt(11), nil),
	new(big.Int).Exp(_10b, big.NewInt(12), nil),
	new(big.Int).Exp(_10b, big.NewInt(13), nil),
	new(big.Int).Exp(_10b, big.NewInt(14), nil),
	new(big.Int).Exp(_10b, big.NewInt(15), nil),
	new(big.Int).Exp(_10b, big.NewInt(16), nil),
	new(big.Int).Exp(_10b, big.NewInt(17), nil),
	new(big.Int).Exp(_10b, big.NewInt(18), nil),
}

func decimalMultiplier(decimal uint) *big.Int {
	if decimal == 0 {
		return new(big.Int).Exp(_10b, big.NewInt(0), nil)
	}
	if decimal <= uint(len(decimalsBigInt)) {
		return decimalsBigInt[decimal-1]
	}
	return new(big.Int).Exp(_10b, big.NewInt(int64(decimal)), nil)
}

func I() *big.Int {
	return new(big.Int)
}

func F() *big.Float {
	return new(big.Float)
}

func divideBnToNumber(numerator, denomiator *big.Float) *big.Float {
	return F().Quo(numerator, denomiator)
}

func GetPrice(orderId string) (uint64, error) {
	d, err := hex.DecodeString(orderId)
	if err != nil {
		return 0, fmt.Errorf("unable to decode order ID: %w", err)
	}

	if len(d) < 8 {
		return 0, fmt.Errorf("order ID too short expecting atleast 8 bytes got %d", len(d))
	}

	return binary.BigEndian.Uint64(d[:8]), nil
}

func GetSeqNum(orderId string, side Side) (uint64, error) {
	d, err := hex.DecodeString(orderId)
	if err != nil {
		return 0, fmt.Errorf("unable to decode order ID: %w", err)
	}

	if len(d) < 16 {
		return 0, fmt.Errorf("order ID too short expecting atleast 8 bytes got %d", len(d))
	}

	v := binary.BigEndian.Uint64(d[8:])

	if side == SideBid {
		return ^v, nil
	}

	return v, nil
}

func PriceLotsToNumber(price, baseLotSize, quoteLotSize, baseDecimals, quoteDecimals uint64) *big.Float {
	baseMultiplier := F().SetInt(decimalMultiplier(uint(baseDecimals)))
	quoteMultiplier := F().SetInt(decimalMultiplier(uint(quoteDecimals)))
	numerator := F().Mul(F().Mul(F().SetUint64(price), F().SetUint64(quoteLotSize)), baseMultiplier)
	denominator := F().Mul(F().SetUint64(baseLotSize), quoteMultiplier)
	return F().Quo(numerator, denominator)
}
