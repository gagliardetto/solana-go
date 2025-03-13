package common

import (
	"github.com/gagliardetto/solana-go"
	web3 "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
)

type AccountMeta = web3.AccountMeta
type AccountMetaSlice []*web3.AccountMeta

type PublicKey = web3.PublicKey

func MustPublicKeyFromBase58(input string) PublicKey {
	return web3.MPK(input)
}

func IsZero(pubkey PublicKey) bool {
	for _, item := range pubkey.Bytes() {
		if item != 0 {
			return false
		}
	}
	return true
}

func As(pubkey PublicKey) solana.PublicKey {
	return solana.PublicKeyFromBytes(pubkey.Bytes())
}

func Meta(
	pubKey PublicKey,
) *web3.AccountMeta {
	return &web3.AccountMeta{
		PublicKey: pubKey,
	}
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

func (slice AccountMetaSlice) Get(index int) *AccountMeta {
	if index >= len(slice) {
		return nil
	}
	return slice[index]
}

func (slice *AccountMetaSlice) Append(account *AccountMeta) {
	*slice = append(*slice, account)
}

func (slice *AccountMetaSlice) SetAccounts(accounts []*AccountMeta) error {
	*slice = accounts
	return nil
}

func (slice AccountMetaSlice) GetAccounts() []*AccountMeta {
	out := make([]*AccountMeta, 0, len(slice))
	for i := range slice {
		if slice[i] != nil {
			out = append(out, slice[i])
		}
	}
	return out
}

func FormatMeta(name string, meta *AccountMeta) string {
	return format.Meta(name, meta)
}

func ConvertMeta(input []*solana.AccountMeta) []*AccountMeta {
	return input
}

type AccountsSettable interface {
	SetAccounts(accounts []*AccountMeta) error
}

type AccountsGettable interface {
	GetAccounts() (accounts []*AccountMeta)
}
