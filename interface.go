package solana

type AccountSettable interface {
	SetAccounts(accounts []*AccountMeta) error
}
