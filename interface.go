package solana

type AccountSettable interface {
	SetAccounts(accounts []PublicKey, instructionActIdx []uint8) error
}
