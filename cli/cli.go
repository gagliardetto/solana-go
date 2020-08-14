package cli

import (
	"fmt"

	"github.com/pkg/errors"
)

func GetDecryptPassphrase() (string, error) {
	passphrase, err := GetPassword("Enter passphrase to decrypt your vault: ")
	if err != nil {
		return "", fmt.Errorf("reading password: %s", err)
	}

	return passphrase, nil
}
func GetEncryptPassphrase() (string, error) {
	passphrase, err := GetPassword("Enter passphrase to encrypt your vault: ")
	if err != nil {
		return "", fmt.Errorf("reading password: %s", err)
	}

	passphraseConfirm, err := GetPassword("Confirm passphrase: ")
	if err != nil {
		return "", fmt.Errorf("reading confirmation password: %s", err)
	}

	if passphrase != passphraseConfirm {
		fmt.Println()
		return "", errors.New("passphrase mismatch!")
	}
	return passphrase, nil

}
