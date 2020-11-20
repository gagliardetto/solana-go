package serum

import "math/big"

func I() *big.Int {
	return new(big.Int)
}

func F() *big.Float {
	return new(big.Float)
}

func divideBnToNumber(numerator, denomiator *big.Float) *big.Float {
	return F().Quo(numerator, denomiator)
}
