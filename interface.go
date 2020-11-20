package solana

type AccountSettable interface {
	SetAccounts(accounts []*AccountMeta, instructionActIdx []uint8) error
}
