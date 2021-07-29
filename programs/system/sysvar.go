package system

import "github.com/gagliardetto/solana-go"

// See more here: https://github.com/solana-labs/solana/blob/master/docs/src/developing/runtime-facilities/sysvars.md

// From https://github.com/solana-labs/solana/blob/94ab0eb49f1bce18d0a157dfe7a2bb1fb39dbe2c/docs/src/developing/runtime-facilities/sysvars.md
var (
	// The Clock sysvar contains data on cluster time,
	// including the current slot, epoch, and estimated wall-clock Unix timestamp.
	// It is updated every slot.
	SysVarClockPubkey = solana.MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111")

	// The EpochSchedule sysvar contains epoch scheduling constants that are set in genesis,
	// and enables calculating the number of slots in a given epoch,
	// the epoch for a given slot, etc.
	// (Note: the epoch schedule is distinct from the leader schedule)
	SysVarEpochSchedulePubkey = solana.MustPublicKeyFromBase58("SysvarEpochSchedu1e111111111111111111111111")

	// The Fees sysvar contains the fee calculator for the current slot.
	// It is updated every slot, based on the fee-rate governor.
	SysVarFeesPubkey = solana.MustPublicKeyFromBase58("SysvarFees111111111111111111111111111111111")

	// The Instructions sysvar contains the serialized instructions in a Message while that Message is being processed.
	// This allows program instructions to reference other instructions in the same transaction.
	SysVarInstructionsPubkey = solana.MustPublicKeyFromBase58("Sysvar1nstructions1111111111111111111111111")

	// The RecentBlockhashes sysvar contains the active recent blockhashes as well as their associated fee calculators.
	// It is updated every slot.
	// Entries are ordered by descending block height,
	// so the first entry holds the most recent block hash,
	// and the last entry holds an old block hash.
	SysVarRecentBlockHashesPubkey = solana.MustPublicKeyFromBase58("SysvarRecentB1ockHashes11111111111111111111")

	// The Rent sysvar contains the rental rate.
	// Currently, the rate is static and set in genesis.
	// The Rent burn percentage is modified by manual feature activation.
	SysVarRentPubkey = solana.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")

	// The SlotHashes sysvar contains the most recent hashes of the slot's parent banks.
	// It is updated every slot.
	SysVarSlotHashesPubkey = solana.MustPublicKeyFromBase58("SysvarS1otHashes111111111111111111111111111")

	// The SlotHistory sysvar contains a bitvector of slots present over the last epoch. It is updated every slot.
	SysVarSlotHistoryPubkey = solana.MustPublicKeyFromBase58("SysvarS1otHistory11111111111111111111111111")

	// The StakeHistory sysvar contains the history of cluster-wide stake activations and de-activations per epoch.
	// It is updated at the start of every epoch.
	SysVarStakeHistoryPubkey = solana.MustPublicKeyFromBase58("SysvarStakeHistory1111111111111111111111111")

	//
	SysVarRewardsPubkey = solana.MustPublicKeyFromBase58("SysvarRewards111111111111111111111111111111")
)
