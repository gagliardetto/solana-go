package solana

import (
	"crypto"
	"crypto/ed25519"
	crypto_rand "crypto/rand"

	"github.com/mr-tron/base58"
)

type PrivateKey []byte

func PrivateKeyFromBase58(privkey string) (PrivateKey, error) {
	res, err := base58.Decode(privkey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k PrivateKey) String() string {
	return base58.Encode(k)
}

func NewRandomPrivateKey() (PublicKey, PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(crypto_rand.Reader)
	if err != nil {
		return PublicKey{}, nil, err
	}
	var publicKey PublicKey
	copy(publicKey[:], pub)
	return publicKey, PrivateKey(priv), nil
}

func (k PrivateKey) Sign(payload []byte) ([]byte, error) {
	p := ed25519.PrivateKey(k)
	return p.Sign(crypto_rand.Reader, payload, crypto.Hash(0))
}

func (k PrivateKey) PublicKey() PublicKey {
	p := ed25519.PrivateKey(k)
	pub := p.Public().(ed25519.PublicKey)

	var publicKey PublicKey
	copy(publicKey[:], pub)

	return publicKey
}
