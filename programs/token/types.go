package token

import (
	ag_binary "github.com/dfuse-io/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AuthorityType ag_binary.BorshEnum

const (
	// Authority to mint new tokens
	AuthorityMintTokens AuthorityType = iota

	// Authority to freeze any account associated with the Mint
	AuthorityFreezeAccount

	// Owner of a given token account
	AuthorityAccountOwner

	// Authority to close a token account
	AuthorityCloseAccount
)

type AccountState ag_binary.BorshEnum

const (
	// Account is not yet initialized
	Uninitialized AccountState = iota

	// Account is initialized; the account owner and/or delegate may perform permitted operations
	// on this account
	Initialized

	// Account has been frozen by the mint freeze authority. Neither the account owner nor
	// the delegate are able to perform operations on this account.
	Frozen
)

type Multisig struct {
	// Number of signers required
	M uint8

	// Number of valid signers
	N uint8

	// Is `true` if this structure has been initialized
	IsInitialized bool

	// Signer public keys
	Signers [11]ag_solanago.PublicKey
}
