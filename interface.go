package solana

type AccountSettable interface {
	SetAccounts(accounts []PublicKey) error
}
