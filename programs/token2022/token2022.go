package token2022

import (
	"github.com/gagliardetto/solana-go"
)

var (
	/*
		The Token-2022 Program, also known as Token Extensions, is a superset of the functionality provided by the Token Program.

		For more information, see the [Token-2022 Program Documentation](https://spl.solana.com/token-2022).
	*/
	ProgramID = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
)

type TokenInstruction byte
type MetadataPointerInstruction byte

const (
	Initialize MetadataPointerInstruction = 0
	Update     MetadataPointerInstruction = 1
)

const (
	InitializeMint TokenInstruction = iota
	InitializeAccount
	InitializeMultisig
	Transfer
	Approve
	Revoke
	SetAuthority
	MintTo
	Burn
	CloseAccount
	FreezeAccount
	ThawAccount
	TransferChecked
	ApproveChecked
	MintToChecked
	BurnChecked
	InitializeAccount2
	SyncNative
	InitializeAccount3
	InitializeMultisig2
	InitializeMint2
	GetAccountDataSize
	InitializeImmutableOwner
	AmountToUiAmount
	UiAmountToAmount
	InitializeMintCloseAuthority
	TransferFeeExtension
	ConfidentialTransferExtension
	DefaultAccountStateExtension
	Reallocate
	MemoTransferExtension
	CreateNativeMint
	InitializeNonTransferableMint
	InterestBearingMintExtension
	CpiGuardExtension
	InitializePermanentDelegate
	TransferHookExtension
	MetadataPointerExtension
	GroupPointerExtension
	GroupMemberPointerExtension
	ScaledUiAmountExtension
	PausableExtension
)
