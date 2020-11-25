package solana

import (
	"crypto"
	"crypto/ed25519"
	crypto_rand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"

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

func (k PrivateKey) Sign(payload []byte) (Signature, error) {
	p := ed25519.PrivateKey(k)
	signData, err := p.Sign(crypto_rand.Reader, payload, crypto.Hash(0))
	if err != nil {
		return Signature{}, err
	}

	var signature Signature
	copy(signature[:], signData)

	return signature, err
}

func (k PrivateKey) PublicKey() PublicKey {
	p := ed25519.PrivateKey(k)
	pub := p.Public().(ed25519.PublicKey)

	var publicKey PublicKey
	copy(publicKey[:], pub)

	return publicKey
}

///
type PublicKey [32]byte

func (p PublicKey) Equals(pb PublicKey) bool {
	return p.String() == pb.String()
}

func MustPublicKeyFromBase58(in string) PublicKey {
	out, err := PublicKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func PublicKeyFromBase58(in string) (out PublicKey, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return
	}

	if len(val) != 32 {
		err = fmt.Errorf("invalid length, expected 32, got %d", len(val))
		return
	}
	copy(out[:], val)
	return
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p *PublicKey) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	dat, err := base58.Decode(s)
	if err != nil {
		return err
	}

	if len(dat) != 32 {
		return errors.New("invalid data length for public key")
	}

	target := PublicKey{}
	copy(target[:], dat)
	*p = target
	return
}

func (p PublicKey) String() string {
	return base58.Encode(p[:])
}
