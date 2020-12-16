package solana

import (
	"crypto"
	"crypto/ed25519"
	crypto_rand "crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mr-tron/base58"
)

type PrivateKey []byte

func MustPrivateKeyFromBase58(in string) PrivateKey {
	out, err := PrivateKeyFromBase58(in)
	if err != nil {
		panic(err)
	}
	return out
}

func PrivateKeyFromBase58(privkey string) (PrivateKey, error) {
	res, err := base58.Decode(privkey)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PrivateKeyFromSolanaKeygenFile(file string) (PrivateKey, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read keygen file: %w", err)
	}

	var values []uint8
	err = json.Unmarshal(content, &values)
	if err != nil {
		return nil, fmt.Errorf("decode keygen file: %w", err)
	}

	return PrivateKey([]byte(values)), nil
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
		return out, fmt.Errorf("decode: %w", err)
	}

	if len(val) != 32 {
		return out, fmt.Errorf("invalid length, expected 32, got %d", len(val))
	}

	copy(out[:], val)
	return
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p *PublicKey) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*p, err = PublicKeyFromBase58(s)
	if err != nil {
		return fmt.Errorf("invalid public key %q: %w", s, err)
	}
	return
}

var zeroPublicKey = PublicKey{}

func (p PublicKey) IsZero() bool {
	return p == zeroPublicKey
}

func (p PublicKey) String() string {
	return base58.Encode(p[:])
}
