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
	"github.com/dfuse-io/solana-go"
)

type Context struct {
	Slot bin.Uint64
}

type RPCContext struct {
	Context Context `json:"context,omitempty"`
}

type GetBalanceResult struct {
	RPCContext
	Value bin.Uint64 `json:"value"`
}

type GetSlotResult bin.Uint64

type GetRecentBlockhashResult struct {
	RPCContext
	Value BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     solana.PublicKey `json:"blockhash"` /* make this a `Hash` type, which is a copy of the PublicKey` type */
	FeeCalculator FeeCalculator    `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature bin.Uint64 `json:"lamportsPerSignature"`
}

type GetConfirmedBlockResult struct {
	Blockhash         solana.PublicKey      `json:"blockhash"`
	PreviousBlockhash solana.PublicKey      `json:"previousBlockhash"` // could be zeroes if ledger was clean-up and this is unavailable
	ParentSlot        bin.Uint64            `json:"parentSlot"`
	Transactions      []TransactionWithMeta `json:"transactions"`
	Rewards           []BlockReward         `json:"rewards"`
	BlockTime         bin.Uint64            `json:"blockTime,omitempty"`
}

type BlockReward struct {
	Pubkey      solana.PublicKey `json:"pubkey"`      // The public key, as base-58 encoded string, of the account that received the reward
	Lamports    bin.Int64        `json:"lamports"`    // number of reward lamports credited or debited by the account, as a i64
	PostBalance bin.Int64        `json:"postBalance"` // account balance in lamports after the reward was applied
	RewardType  string           `json:"rewardType"`  // type of reward: "fee", "rent", "voting", "staking"
}

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
	UiAmount       *bin.JSONFloat64 `json:"uiAmount"`       // Token amount as a float, accounting for decimals. DEPRECATED
	UiAmountString string           `json:"uiAmountString"` // Token amount as a string, accounting for decimals.
}

type TransactionMeta struct {
	Err               interface{}        `json:"err"`               // Error if transaction failed, null if transaction succeeded. https://github.com/solana-labs/solana/blob/master/sdk/src/transaction.rs#L24
	Fee               bin.Uint64         `json:"fee"`               // fee this transaction was charged
	PreBalances       []bin.Uint64       `json:"preBalances"`       //  array of u64 account balances from before the transaction was processed
	PostBalances      []bin.Uint64       `json:"postBalances"`      // array of u64 account balances after the transaction was processed
	InnerInstructions []InnerInstruction `json:"innerInstructions"` // List of inner instructions or omitted if inner instruction recording was not yet enabled during this transaction

	PreTokenBalances  []TokenBalance `json:"preTokenBalances"`  // List of token balances from before the transaction was processed or omitted if token balance recording was not yet enabled during this transaction
	PostTokenBalances []TokenBalance `json:"postTokenBalances"` // List of token balances from after the transaction was processed or omitted if token balance recording was not yet enabled during this transaction

	LogMessages []string `json:"logMessages"` // array of string log messages or omitted if log message recording was not yet enabled during this transaction

	Status DeprecatedTransactionMetaStatus `json:"status"` // Transaction status; DEPRECATED.

	// TODO: not present in spec doc, but present in json.
	// Rewards []BlockReward `json:"rewards"`
}

type InnerInstruction struct {
	// TODO: <number> == bin.Int64 ???
	Index        uint8                        `json:"index"`        // Index of the transaction instruction from which the inner instruction(s) originated
	Instructions []solana.CompiledInstruction `json:"instructions"` // Ordered list of inner program instructions that were invoked during a single transaction instruction.
}

type DeprecatedTransactionMetaStatus struct {
	Ok  interface{} `json:"Ok"`  // <null> Transaction was successful
	Err interface{} `json:"Err"` // Transaction failed with TransactionError
}

type TransactionSignature struct {
	Err       interface{} `json:"err,omitempty"`
	Memo      string      `json:"memo,omitempty"`
	Signature string      `json:"signature,omitempty"`
	Slot      bin.Uint64  `json:"slot,omitempty"`
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

type GetProgramAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	// Filter on accounts, implicit AND between filters
	Filters []RPCFilter `json:"filters,omitempty"`
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
	AccountKeys     []*AccountKey `json:"accountKeys"`
	RecentBlockhash solana.PublicKey/* TODO: change to Hash */ `json:"recentBlockhash"`
	Instructions    []ParsedInstruction `json:"instructions"`
}

type AccountKey struct {
	PublicKey solana.PublicKey `json:"pubkey"`
	Signer    bool             `json:"signer"`
	Writable  bool             `json:"writable"`
}

type ParsedInstruction struct {
	Accounts  []solana.PublicKey `json:"accounts,omitempty"`
	Data      solana.Base58      `json:"data,omitempty"`
	Parsed    *InstructionInfo   `json:"parsed,omitempty"`
	Program   string             `json:"program,omitempty"`
	ProgramID solana.PublicKey   `json:"programId"`
}

type InstructionInfo struct {
	Info            map[string]interface{} `json:"info"`
	InstructionType string                 `json:"type"`
}

func (p *ParsedInstruction) IsParsed() bool {
	return p.Parsed != nil
}
