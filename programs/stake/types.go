package stake

import (
	ag_binary "github.com/gagliardetto/binary"
)

// StakeAuthorize represents the different types of authorities in the stake program.
type StakeAuthorize ag_binary.BorshEnum

const (
	// Authority to stake
	StakeAuthorizeStaker StakeAuthorize = iota

	// Authority to withdraw
	StakeAuthorizeWithdrawer
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
