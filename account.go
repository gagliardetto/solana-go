package solana

import (
	"fmt"
)

type Account struct {
	PrivateKey PrivateKey
}

func NewAccount() *Account {
	_, privateKey, err := NewRandomPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	return &Account{
		PrivateKey: privateKey,
	}
}

func AccountFromPrivateKeyBase58(privateKey string) (*Account, error) {
	k, err := PrivateKeyFromBase58(privateKey)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Account{
		PrivateKey: k,
	}, nil
}

func (a *Account) PublicKey() PublicKey {
	return a.PrivateKey.PublicKey()
}

type AccountMeta struct {
	PublicKey  PublicKey
	IsWritable bool
	IsSigner   bool
}

func NewAccountMeta(
	pubKey PublicKey,
	WRITE bool,
	SIGNER bool,
) *AccountMeta {
	return &AccountMeta{
		PublicKey:  pubKey,
		IsWritable: WRITE,
		IsSigner:   SIGNER,
	}
}

func (a *AccountMeta) less(act *AccountMeta) bool {
	if a.IsSigner && !act.IsSigner {
		return true
	} else if !a.IsSigner && act.IsSigner {
		return false
	}

	if a.IsWritable {
		return true
	}
	return false
}

type AccountMetaSlice []*AccountMeta

func (slice *AccountMetaSlice) Append(account *AccountMeta) {
	*slice = append(*slice, account)
}

func (slice *AccountMetaSlice) SetAccounts(accounts []*AccountMeta) error {
	*slice = accounts
	return nil
}

func (slice AccountMetaSlice) GetAccounts() []*AccountMeta {
	return slice
}
