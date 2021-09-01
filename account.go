package solana

import (
	"fmt"
)

// Wallet is a wrapper around a PrivateKey
type Wallet struct {
	PrivateKey PrivateKey
}

func NewWallet() *Wallet {
	privateKey, err := NewRandomPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("failed to generate private key: %s", err))
	}
	return &Wallet{
		PrivateKey: privateKey,
	}
}

func WalletFromPrivateKeyBase58(privateKey string) (*Wallet, error) {
	k, err := PrivateKeyFromBase58(privateKey)
	if err != nil {
		return nil, fmt.Errorf("account from private key: private key from b58: %w", err)
	}
	return &Wallet{
		PrivateKey: k,
	}, nil
}

func (a *Wallet) PublicKey() PublicKey {
	return a.PrivateKey.PublicKey()
}

type AccountMeta struct {
	PublicKey  PublicKey
	IsWritable bool
	IsSigner   bool
}

// Meta intializes a new AccountMeta with the provided pubKey.
func Meta(
	pubKey PublicKey,
) *AccountMeta {
	return &AccountMeta{
		PublicKey: pubKey,
	}
}

// WRITE sets IsWritable to true.
func (meta *AccountMeta) WRITE() *AccountMeta {
	meta.IsWritable = true
	return meta
}

// SIGNER sets IsSigner to true.
func (meta *AccountMeta) SIGNER() *AccountMeta {
	meta.IsSigner = true
	return meta
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

// GetSigners returns the accounts that are signers.
func (slice AccountMetaSlice) GetSigners() []*AccountMeta {
	signers := make([]*AccountMeta, 0)
	for _, ac := range slice {
		if ac.IsSigner {
			signers = append(signers, ac)
		}
	}
	return signers
}
