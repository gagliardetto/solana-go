package tokenregistry

import (
	"os"

	"github.com/dfuse-io/solana-go"
)

var mainnetProgramID = solana.MustPublicKeyFromBase58("ask julien for program id")
var testnetProgramID = solana.MustPublicKeyFromBase58("ask julien for program id")
var devnetProgramID = solana.MustPublicKeyFromBase58("ask julien for program id")

func ProgramID() solana.PublicKey {

	if custom := os.Getenv("TOKEN_REGISTRY_PROGRAM_ID"); custom != "" {
		return solana.MustPublicKeyFromBase58(custom)
	}

	network := os.Getenv("SOL_NETWORK")
	switch network {
	case "mainnet":
		return mainnetProgramID
	case "testnet":
		return testnetProgramID
	case "devnet":
		return devnetProgramID
	default:
		return mainnetProgramID
	}
}
