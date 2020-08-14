package vault

import (
	"fmt"
	"os"

	"github.com/dfuse-io/solana-go/cli"
	"github.com/pkg/errors"
)

type SecretBoxer interface {
	Seal(in []byte) (string, error)
	Open(in string) ([]byte, error)
	WrapType() string
}

func SecretBoxerForType(boxerType string, keypath string) (SecretBoxer, error) {
	switch boxerType {
	case "kms-gcp":
		if keypath == "" {
			return nil, errors.New("missing kms-gcp keypath")
		}
		return NewKMSGCPBoxer(keypath), nil
	case "passphrase":
		var password string
		var err error
		if envVal := os.Getenv("SLNC_GLOBAL_INSECURE_VAULT_PASSPHRASE"); envVal != "" {
			password = envVal
		} else {
			password, err = cli.GetDecryptPassphrase()
			if err != nil {
				return nil, err
			}
		}

		return NewPassphraseBoxer(password), nil
	default:
		return nil, fmt.Errorf("unknown secret boxer: %s", boxerType)
	}
}
