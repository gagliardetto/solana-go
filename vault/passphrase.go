package vault

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

type PassphraseBoxer struct {
	passphrase string
}

func NewPassphraseBoxer(password string) *PassphraseBoxer {
	return &PassphraseBoxer{
		passphrase: password,
	}
}

func (b *PassphraseBoxer) WrapType() string {
	return "passphrase"
}

func (b *PassphraseBoxer) Seal(in []byte) (string, error) {
	var nonce [nonceLength]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		return "", err
	}

	salt := make([]byte, saltLength)
	if _, err := crypto_rand.Read(salt); err != nil {
		return "", err
	}
	secretKey := deriveKey(b.passphrase, salt)
	prefix := append(salt, nonce[:]...)

	cipherText := secretbox.Seal(prefix, in, &nonce, &secretKey)

	return base64.RawStdEncoding.EncodeToString(cipherText), nil
}

func (b *PassphraseBoxer) Open(in string) ([]byte, error) {
	buf, err := base64.RawStdEncoding.DecodeString(in)
	if err != nil {
		return []byte{}, err
	}

	salt := make([]byte, saltLength)
	copy(salt, buf[:saltLength])
	var nonce [nonceLength]byte
	copy(nonce[:], buf[saltLength:nonceLength+saltLength])

	secretKey := deriveKey(b.passphrase, salt)
	decrypted, ok := secretbox.Open(nil, buf[nonceLength+saltLength:], &nonce, &secretKey)
	if !ok {
		return []byte{}, fmt.Errorf("failed to decrypt")
	}
	return decrypted, nil
}
