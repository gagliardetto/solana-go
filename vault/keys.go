package vault

import (
	"crypto"
	"crypto/ed25519"
	crypto_rand "crypto/rand"

	"github.com/mr-tron/base58"
)

type PublicKey []byte

func (p PublicKey) String() string {
	return base58.Encode(p)
}

type PrivateKey []byte

func PrivateKeyFromBase58(privkey string) (PrivateKey, error) {
	res, err := base58.Decode(privkey)
	if err != nil {
		return nil, err
	}
	return PrivateKey(res), nil
}

func (p PrivateKey) String() string {
	return base58.Encode(p)
}

func NewRandomPrivateKey() (PublicKey, PrivateKey, error) {
	pub, priv, err := ed25519.GenerateKey(crypto_rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return PublicKey(pub), PrivateKey(priv), nil
}

func (k PrivateKey) Sign(payload []byte) ([]byte, error) {
	p := ed25519.PrivateKey(k)
	return p.Sign(crypto_rand.Reader, payload, crypto.Hash(0))
}

func (k PrivateKey) PublicKey() PublicKey {
	p := ed25519.PrivateKey(k)
	res := p.Public()
	return PublicKey(res.([]byte))
}
