package cmd

import (
	"bytes"
	"fmt"
	"log"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/token"
	"github.com/lunixbochs/struc"
)

func decode(owner solana.PublicKey, data []byte) (interface{}, error) {
	bdx, _ := solana.PublicKeyFromBase58("BdxDnkFufu8tjAE5gdPkWdjGfQ3Lz2v6ozfiBDMKxDFW")
	tkn, _ := solana.PublicKeyFromBase58("TokenSVp5gheXUvJ6jGWGeCsgPKgnE3YgdGKRVCMY9o")
	switch owner {
	case bdx, tkn:
		return decodeAsToken(data)
	}
	return nil, nil
}

func decodeAsToken(data []byte) (out interface{}, err error) {
	reader := bytes.NewReader(data)

	switch len(data) {
	case 120:
		var tokenAcct token.Account
		if err := struc.Unpack(reader, &tokenAcct); err != nil {
			return nil, fmt.Errorf("failed unpacking: %w", err)
		}

		return tokenAcct, nil
	case 40:
		var mint token.Mint
		if err := struc.Unpack(reader, &mint); err != nil {
			log.Fatalln("failed unpack", err)
		}

		return mint, nil

		// cnt, _ := json.MarshalIndent(mint, "", "  ")
		// fmt.Println(string(cnt))
	case 7777:
		// decode the Multisig struct
	}
	return
}
