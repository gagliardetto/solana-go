package solana

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {

	a := NewWallet()
	privateKey := a.PrivateKey
	public := a.PublicKey()

	a2, err := WalletFromPrivateKeyBase58(privateKey.String())
	require.NoError(t, err)

	require.Equal(t, privateKey, a2.PrivateKey)
	require.Equal(t, public, a2.PublicKey())
}

func Test_AccountMeta_less(t *testing.T) {
	pkey := MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")
	tests := []struct {
		name   string
		left   *AccountMeta
		right  *AccountMeta
		expect bool
	}{
		{
			name:   "accounts are equal",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			expect: false,
		},
		{
			name:   "left is a signer, right is not a signer",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: false},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			expect: true,
		},
		{
			name:   "left is not a signer, right is a signer",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: false},
			expect: false,
		},
		{
			name:   "left is writable, right is not writable",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: true},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			expect: true,
		},
		{
			name:   "left is not writable, right is writable",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: false},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: false, IsWritable: true},
			expect: false,
		},
		{
			name:   "both are signers and left is writable, right is not writable",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: true},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: false},
			expect: true,
		},
		{
			name:   "both are signers andleft is not writable, right is writable",
			left:   &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: false},
			right:  &AccountMeta{PublicKey: pkey, IsSigner: true, IsWritable: true},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, test.left.less(test.right))
		})
	}
}

func TestAccountMetaSlice(t *testing.T) {
	pkey1 := MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")

	var slice AccountMetaSlice

	setting := []*AccountMeta{
		{PublicKey: pkey1, IsSigner: true, IsWritable: false},
	}
	err := slice.SetAccounts(setting)
	require.NoError(t, err)

	require.Len(t, slice, 1)
	require.Equal(t, setting[0], slice[0])
	require.Equal(t, setting, slice.GetAccounts())

	{
		pkey2 := MustPublicKeyFromBase58("BPFLoaderUpgradeab1e11111111111111111111111")

		meta := NewAccountMeta(pkey2, true, false)
		slice.Append(meta)

		require.Len(t, slice, 2)
		require.Equal(t, meta, slice[1])
		require.Equal(t, meta, slice.GetAccounts()[1])
	}
}

func TestNewAccountMeta(t *testing.T) {
	pkey := MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")

	isWritable := false
	isSigner := true

	out := NewAccountMeta(pkey, isWritable, isSigner)

	require.NotNil(t, out)

	require.Equal(t, isSigner, out.IsSigner)
	require.Equal(t, isWritable, out.IsWritable)
}

func TestMeta(t *testing.T) {
	pkey := MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")

	meta := Meta(pkey)
	require.NotNil(t, meta)
	require.Equal(t, pkey, meta.PublicKey)

	require.False(t, meta.IsSigner)
	require.False(t, meta.IsWritable)

	meta.SIGNER()

	require.True(t, meta.IsSigner)
	require.False(t, meta.IsWritable)

	meta.WRITE()

	require.True(t, meta.IsSigner)
	require.True(t, meta.IsWritable)
}
