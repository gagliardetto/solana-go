package confidential

import "errors"

// Proof-generation errors, corresponding to TokenProofGenerationError in
// spl-token-2022's confidential-transfer proof-generation crate.
//
// Amounts above MaxAmount are rejected up front as ErrIllegalAmountBitLength.
// Rust SDK defers to the range proof instead.
var (
	ErrNotEnoughFunds         = errors.New("zk: not enough funds in account")
	ErrIllegalAmountBitLength = errors.New("zk: amount has illegal bit length")
	ErrFeeCalculation         = errors.New("zk: fee calculation failed")
	ErrBalanceOverflow        = errors.New("zk: balance overflows a uint64")
)
