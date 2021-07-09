// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type Context struct {
	Slot bin.Uint64 `json:"slot"`
}

type RPCContext struct {
	Context Context `json:"context,omitempty"`
}

type GetBalanceResult struct {
	RPCContext
	Value bin.Uint64 `json:"value"`
}

type GetRecentBlockhashResult struct {
	RPCContext
	Value BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     solana.Hash   `json:"blockhash"`
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature bin.Uint64 `json:"lamportsPerSignature"`
}

type GetConfirmedBlockResult struct {
	Blockhash         solana.Hash           `json:"blockhash"`
	PreviousBlockhash solana.Hash           `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        bin.Uint64            `json:"parentSlot"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []BlockReward         `json:"rewards"`
	BlockTime         bin.Uint64            `json:"blockTime,omitempty"`
}

type BlockReward struct {
	Pubkey      solana.PublicKey `json:"pubkey"`      // The public key, as base-58 encoded string, of the account that received the reward
	Lamports    bin.Int64        `json:"lamports"`    // number of reward lamports credited or debited by the account, as a i64
	PostBalance bin.Uint64       `json:"postBalance"` // account balance in lamports after the reward was applied
	RewardType  RewardType       `json:"rewardType"`  // type of reward: "fee", "rent", "voting", "staking"
}

type RewardType string

const (
	RewardTypeFee     RewardType = "Fee"
	RewardTypeRent    RewardType = "Rent"
	RewardTypeVoting  RewardType = "Voting"
	RewardTypeStaking RewardType = "Staking"
)

type TransactionWithMeta struct {
	Meta        *TransactionMeta    `json:"meta,omitempty"` // transaction status metadata object
	Transaction *solana.Transaction `json:"transaction"`
}

type TransactionParsed struct {
	Transaction *ParsedTransaction `json:"transaction"`
	Meta        *TransactionMeta   `json:"meta,omitempty"`
}

type TokenBalance struct {
	// TODO: <number> == bin.Int64 ???
	AccountIndex  uint8            `json:"accountIndex"` // Index of the account in which the token balance is provided for.
	Mint          solana.PublicKey `json:"mint"`         // Pubkey of the token's mint.
	UiTokenAmount *UiTokenAmount   `json:"uiTokenAmount"`
}

type UiTokenAmount struct {
	Amount string `json:"amount"` // Raw amount of tokens as a string, ignoring decimals.
	// TODO: <number> == bin.Int64 ???
	Decimals       uint8            `json:"decimals"`       // Number of decimals configured for token's mint.
	UiAmount       *bin.JSONFloat64 `json:"uiAmount"`       // DEPRECATED: Token amount as a float, accounting for decimals.
	UiAmountString string           `json:"uiAmountString"` // Token amount as a string, accounting for decimals.
}

type TransactionMeta struct {
	Err               interface{}        `json:"err"`                         // Error if transaction failed, null if transaction succeeded. https://github.com/solana-labs/solana/blob/master/sdk/src/transaction.rs#L24
	Fee               bin.Uint64         `json:"fee"`                         // fee this transaction was charged
	PreBalances       []bin.Uint64       `json:"preBalances"`                 //  array of u64 account balances from before the transaction was processed
	PostBalances      []bin.Uint64       `json:"postBalances"`                // array of u64 account balances after the transaction was processed
	InnerInstructions []InnerInstruction `json:"innerInstructions,omitempty"` // List of inner instructions or omitted if inner instruction recording was not yet enabled during this transaction

	PreTokenBalances  []TokenBalance `json:"preTokenBalances"`  // List of token balances from before the transaction was processed or omitted if token balance recording was not yet enabled during this transaction
	PostTokenBalances []TokenBalance `json:"postTokenBalances"` // List of token balances from after the transaction was processed or omitted if token balance recording was not yet enabled during this transaction

	LogMessages []string `json:"logMessages"` // array of string log messages or omitted if log message recording was not yet enabled during this transaction

	Status DeprecatedTransactionMetaStatus `json:"status"` // DEPRECATED: Transaction status.

	Rewards []BlockReward `json:"rewards,omitempty"`
}

type InnerInstruction struct {
	// TODO: <number> == bin.Int64 ???
	Index        uint8                        `json:"index"`        // Index of the transaction instruction from which the inner instruction(s) originated
	Instructions []solana.CompiledInstruction `json:"instructions"` // Ordered list of inner program instructions that were invoked during a single transaction instruction.
}

// 	Ok  interface{} `json:"Ok"`  // <null> Transaction was successful
// 	Err interface{} `json:"Err"` // Transaction failed with TransactionError
type DeprecatedTransactionMetaStatus M

type TransactionSignature struct {
	Err                interface{}            `json:"err"`                 // Error if transaction failed, null if transaction succeeded
	Memo               *string                `json:"memo"`                // Memo associated with the transaction, null if no memo is present
	Signature          solana.Signature       `json:"signature"`           // transaction signature as base-58 encoded string
	Slot               bin.Uint64             `json:"slot,omitempty"`      // The slot that contains the block with the transaction
	BlockTime          bin.Int64              `json:"blockTime,omitempty"` // estimated production time, as Unix timestamp (seconds since the Unix epoch) of when transaction was processed. null if not available.
	ConfirmationStatus ConfirmationStatusType `json:"confirmationStatus,omitempty"`
}

type GetAccountInfoResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	Lamports   bin.Uint64       `json:"lamports"`   // number of lamports assigned to this account
	Owner      solana.PublicKey `json:"owner"`      // base-58 encoded Pubkey of the program this account has been assigned to
	Data       solana.Data      `json:"data"`       // data associated with the account, either as encoded binary data or JSON format {<program>: <state>}, depending on encoding parameter
	Executable bool             `json:"executable"` // boolean indicating if the account contains a program (and is strictly read-only)
	RentEpoch  bin.Uint64       `json:"rentEpoch"`  // the epoch at which this account will next owe rent
}

type DataSlice struct {
	Offset *int `json:"offset,omitempty"`
	Length *int `json:"length,omitempty"`
}
type GetProgramAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	Encoding EncodingType `json:"encoding,omitempty"`

	DataSlice *DataSlice `json:"dataSlice,omitempty"` // limit the returned account data

	// Filter on accounts, implicit AND between filters
	Filters []RPCFilter `json:"filters,omitempty"` // filter results using various filter objects; account must meet all filter criteria to be included in results

	// TODO:
	// WithContext *bool `json:"withContext,omitempty"` // wrap the result in an RpcResponse JSON object.
}

type GetProgramAccountsResult []*KeyedAccount

type KeyedAccount struct {
	Pubkey  solana.PublicKey `json:"pubkey"`
	Account *Account         `json:"account"`
}

type GetConfirmedSignaturesForAddress2Opts struct {
	Limit  uint64 `json:"limit,omitempty"`
	Before string `json:"before,omitempty"`
	Until  string `json:"until,omitempty"`
}

type GetConfirmedSignaturesForAddress2Result []*TransactionSignature

type RPCFilter struct {
	Memcmp   *RPCFilterMemcmp `json:"memcmp,omitempty"`
	DataSize bin.Uint64       `json:"dataSize,omitempty"`
}

type RPCFilterMemcmp struct {
	Offset int           `json:"offset"`
	Bytes  solana.Base58 `json:"bytes"`
}

type CommitmentType string

const (
	CommitmentMax          = CommitmentType("max")
	CommitmentRecent       = CommitmentType("recent")
	CommitmentRoot         = CommitmentType("root")
	CommitmentSingle       = CommitmentType("single")
	CommitmentSingleGossip = CommitmentType("singleGossip")
)

/// Parsed Transaction

type ParsedTransaction struct {
	Signatures []solana.Signature `json:"signatures"`
	Message    Message            `json:"message"`
}

type Message struct {
	AccountKeys     []solana.PublicKey   `json:"accountKeys"`
	RecentBlockhash solana.Hash          `json:"recentBlockhash"`
	Instructions    []ParsedInstruction  `json:"instructions"`
	Header          solana.MessageHeader `json:"header"`
}

type AccountKey struct {
	PublicKey solana.PublicKey `json:"pubkey"`
	Signer    bool             `json:"signer"`
	Writable  bool             `json:"writable"`
}

type ParsedInstruction struct {
	Accounts       []bin.Int64      `json:"accounts,omitempty"`
	Data           solana.Base58    `json:"data,omitempty"`
	Parsed         *InstructionInfo `json:"parsed,omitempty"`
	Program        string           `json:"program,omitempty"`
	ProgramIDIndex bin.Int64        `json:"programIdIndex"`
}

type InstructionInfo struct {
	Info            map[string]interface{} `json:"info"`
	InstructionType string                 `json:"type"`
}

func (p *ParsedInstruction) IsParsed() bool {
	return p.Parsed != nil
}

type M map[string]interface{}

type EncodingType string

const (
	EncodingBase58     EncodingType = "base58"      // limited to Account data of less than 129 bytes
	EncodingBase64     EncodingType = "base64"      // will return base64 encoded data for Account data of any size
	EncodingBase64Zstd EncodingType = "base64+zstd" // compresses the Account data using Zstandard and base64-encodes the result

	// attempts to use program-specific state parsers to
	// return more human-readable and explicit account state data.
	// If "jsonParsed" is requested but a parser cannot be found,
	// the field falls back to "base64" encoding, detectable when the data field is type <string>.
	// Cannot be used if specifying dataSlice parameters (offset, length).
	EncodingJSONParsed EncodingType = "jsonParsed"

	EncodingJSON EncodingType = "json"
)
