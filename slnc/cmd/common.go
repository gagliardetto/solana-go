package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dfuse-io/solana-go"
	"github.com/dfuse-io/solana-go/vault"
	"github.com/spf13/viper"
)

func getClient() *solana.Client {
	httpHeaders := viper.GetStringSlice("global-http-header")
	api := solana.NewClient(sanitizeAPIURL(viper.GetString("global-rpc-url")))

	for i := 0; i < 25; i++ {
		if val := os.Getenv(fmt.Sprintf("SLNC_GLOBAL_HTTP_HEADER_%d", i)); val != "" {
			httpHeaders = append(httpHeaders, val)
		}
	}

	for _, header := range httpHeaders {
		headerArray := strings.SplitN(header, ": ", 2)
		if len(headerArray) != 2 || strings.Contains(headerArray[0], " ") {
			errorCheck("validating http headers", fmt.Errorf("invalid HTTP Header format"))
		}
		api.SetHeader(headerArray[0], headerArray[1])
	}

	api.Debug = viper.GetBool("global-debug")

	return api
}

func sanitizeAPIURL(input string) string {
	switch input {
	case "devnet":
		return "https://devnet.solana.com"
	case "testnet":
		return "https://testnet.solana.com"
	case "mainnet":
		return "https://api.mainnet-beta.solana.com"
	}
	return strings.TrimRight(input, "/")
}

func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		if strings.HasSuffix(err.Error(), "connection refused") && strings.Contains(err.Error(), defaultRPCURL) {
			fmt.Println("Have you selected a valid Solana JSON-RPC endpoint ? You can use the --rpc-url flag or SLNC_GLOBAL_RPC_URL environment variable.")
		}
		os.Exit(1)
	}
}

func mustGetWallet() *vault.Vault {
	vault, err := setupWallet()
	errorCheck("wallet setup", err)
	return vault
}

func setupWallet() (*vault.Vault, error) {
	walletFile := viper.GetString("global-vault-file")
	if _, err := os.Stat(walletFile); err != nil {
		return nil, fmt.Errorf("wallet file %q missing: %s", walletFile, err)
	}

	v, err := vault.NewVaultFromWalletFile(walletFile)
	if err != nil {
		return nil, fmt.Errorf("loading vault: %s", err)
	}

	boxer, err := vault.SecretBoxerForType(v.SecretBoxWrap, viper.GetString("global-kms-gcp-keypath"))
	if err != nil {
		return nil, fmt.Errorf("secret boxer: %s", err)
	}

	if err := v.Open(boxer); err != nil {
		return nil, err
	}

	return v, nil
}
