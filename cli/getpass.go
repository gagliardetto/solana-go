package cli

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func GetPassword(input string) (string, error) {
	fd := os.Stdin.Fd()
	fmt.Printf(input)
	pass, err := terminal.ReadPassword(int(fd))
	fmt.Println("")
	return string(pass), err
}
