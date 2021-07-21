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
	stdjson "encoding/json"
	"fmt"
	"time"

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
	Value *BlockhashResult `json:"value"`
}

type BlockhashResult struct {
	Blockhash     solana.Hash   `json:"blockhash"`
	FeeCalculator FeeCalculator `json:"feeCalculator"`
}

type FeeCalculator struct {
	LamportsPerSignature bin.Uint64 `json:"lamportsPerSignature"`
}

type GetConfirmedBlockResult struct {
	Blockhash solana.Hash `json:"blockhash"`

	// could be zeroes if ledger was clean-up and this is unavailable
	PreviousBlockhash solana.Hash `json:"previousBlockhash"`

	ParentSlot   bin.Uint64            `json:"parentSlot"`
	Transactions []TransactionWithMeta `json:"transactions"`
	Signatures   []solana.Signature    `json:"signatures"`
	Rewards      []BlockReward         `json:"rewards"`
	BlockTime    *bin.Uint64           `json:"blockTime,omitempty"`
}

type BlockReward struct {
	// The public key of the account that received the reward.
	Pubkey solana.PublicKey `json:"pubkey"`

	// Number of reward lamports credited or debited by the account, as a i64.
	Lamports bin.Int64 `json:"lamports"`

	// Account balance in lamports after the reward was applied.
	PostBalance bin.Uint64 `json:"postBalance"`

	// Type of reward: "Fee", "Rent", "Voting", "Staking".
	RewardType RewardType `json:"rewardType"`

	// Vote account commission when the reward was credited,
	// only present for voting and staking rewards.
	Commission *uint8 `json:"commission,omitempty"`
}

type RewardType string

const (
	RewardTypeFee     RewardType = "Fee"
	RewardTypeRent    RewardType = "Rent"
	RewardTypeVoting  RewardType = "Voting"
	RewardTypeStaking RewardType = "Staking"
)

type TransactionWithMeta struct {
	// Transaction status metadata object
	Meta        *TransactionMeta    `json:"meta,omitempty"`
	Transaction *solana.Transaction `json:"transaction"`
}

type TransactionParsed struct {
	Meta        *TransactionMeta   `json:"meta,omitempty"`
	Transaction *ParsedTransaction `json:"transaction"`
}

type TokenBalance struct {
	// Index of the account in which the token balance is provided for.
	AccountIndex uint16 `json:"accountIndex"`

	// Pubkey of the token's mint.
	Mint          solana.PublicKey `json:"mint"`
	UiTokenAmount *UiTokenAmount   `json:"uiTokenAmount"`
}

type UiTokenAmount struct {
	// Raw amount of tokens as a string, ignoring decimals.
	Amount string `json:"amount"`

	// TODO: <number> == bin.Int64 ???
	// Number of decimals configured for token's mint.
	Decimals uint8 `json:"decimals"`

	// DEPRECATED: Token amount as a float, accounting for decimals.
	UiAmount *bin.JSONFloat64 `json:"uiAmount"`

	// Token amount as a string, accounting for decimals.
	UiAmountString string `json:"uiAmountString"`
}

type TransactionMeta struct {
	// Error if transaction failed, null if transaction succeeded.
	// https://github.com/solana-labs/solana/blob/master/sdk/src/transaction.rs#L24
	Err interface{} `json:"err"`

	// Fee this transaction was charged
	Fee bin.Uint64 `json:"fee"`

	// Array of u64 account balances from before the transaction was processed
	PreBalances []bin.Uint64 `json:"preBalances"`

	// Array of u64 account balances after the transaction was processed
	PostBalances []bin.Uint64 `json:"postBalances"`

	// List of inner instructions or omitted if inner instruction recording
	// was not yet enabled during this transaction
	InnerInstructions []InnerInstruction `json:"innerInstructions,omitempty"`

	// List of token balances from before the transaction was processed
	// or omitted if token balance recording was not yet enabled during this transaction
	PreTokenBalances []TokenBalance `json:"preTokenBalances"`

	// List of token balances from after the transaction was processed
	// or omitted if token balance recording was not yet enabled during this transaction
	PostTokenBalances []TokenBalance `json:"postTokenBalances"`

	// Array of string log messages or omitted if log message
	// recording was not yet enabled during this transaction
	LogMessages []string `json:"logMessages"`

	// DEPRECATED: Transaction status.
	Status DeprecatedTransactionMetaStatus `json:"status"`

	Rewards []BlockReward `json:"rewards,omitempty"`
}

type InnerInstruction struct {
	// TODO: <number> == bin.Int64 ???
	// Index of the transaction instruction from which the inner instruction(s) originated
	Index uint16 `json:"index"`

	// Ordered list of inner program instructions that were invoked during a single transaction instruction.
	Instructions []solana.CompiledInstruction `json:"instructions"`
}

// 	Ok  interface{} `json:"Ok"`  // <null> Transaction was successful
// 	Err interface{} `json:"Err"` // Transaction failed with TransactionError
type DeprecatedTransactionMetaStatus M

type TransactionSignature struct {
	// Error if transaction failed, nil if transaction succeeded.
	Err interface{} `json:"err"`

	// Memo associated with the transaction, nil if no memo is present.
	Memo *string `json:"memo"`

	// Transaction signature.
	Signature solana.Signature `json:"signature"`

	// The slot that contains the block with the transaction.
	Slot bin.Uint64 `json:"slot,omitempty"`

	// Estimated production time, as Unix timestamp (seconds since the Unix epoch)
	// of when transaction was processed. Nil if not available.
	BlockTime *UnixTimeSeconds `json:"blockTime,omitempty"`

	ConfirmationStatus ConfirmationStatusType `json:"confirmationStatus,omitempty"`
}

type GetAccountInfoResult struct {
	RPCContext
	Value *Account `json:"value"`
}

type Account struct {
	// Number of lamports assigned to this account
	Lamports bin.Uint64 `json:"lamports"`

	// Pubkey of the program this account has been assigned to
	Owner solana.PublicKey `json:"owner"`

	// Data associated with the account, either as encoded binary data or JSON format {<program>: <state>}, depending on encoding parameter
	Data *DataBytesOrJSON `json:"data"`

	// Boolean indicating if the account contains a program (and is strictly read-only)
	Executable bool `json:"executable"`

	// The epoch at which this account will next owe rent
	RentEpoch bin.Uint64 `json:"rentEpoch"`
}

type DataBytesOrJSON struct {
	rawDataEncoding solana.EncodingType
	asDecodedBinary solana.Data
	asJSON          stdjson.RawMessage
}

func (dt DataBytesOrJSON) MarshalJSON() ([]byte, error) {
	if dt.rawDataEncoding == solana.EncodingJSONParsed || dt.rawDataEncoding == solana.EncodingJSON {
		return json.Marshal(dt.asJSON)
	}
	return json.Marshal(dt.asDecodedBinary)
}

func (wrap *DataBytesOrJSON) UnmarshalJSON(data []byte) error {

	if len(data) == 0 || (len(data) == 4 && string(data) == "null") {
		// TODO: is this an error?
		return nil
	}

	firstChar := data[0]

	switch firstChar {
	// Check if first character is `[`, standing for a JSON array.
	case '[':
		// It's base64 (or similar)
		{
			err := wrap.asDecodedBinary.UnmarshalJSON(data)
			if err != nil {
				return err
			}
			wrap.rawDataEncoding = wrap.asDecodedBinary.Encoding
		}
	case '{':
		// It's JSON, most likely.
		// TODO: is it always JSON???
		{
			// Store raw bytes, and unmarshal on request.
			wrap.asJSON = data
			wrap.rawDataEncoding = solana.EncodingJSONParsed
		}
	default:
		return fmt.Errorf("Unknown kind: %v", data)
	}

	return nil
}

// GetBinary returns the decoded bytes if the encoding is
// "base58", "base64", or "base64+zstd".
func (dt *DataBytesOrJSON) GetBinary() []byte {
	return dt.asDecodedBinary.Content
}

// GetRawJSON returns a stdjson.RawMessage when the data
// encoding is "jsonParsed".
func (dt *DataBytesOrJSON) GetRawJSON() stdjson.RawMessage {
	return dt.asJSON
}

type DataSlice struct {
	Offset *uint64 `json:"offset,omitempty"`
	Length *uint64 `json:"length,omitempty"`
}
type GetProgramAccountsOpts struct {
	Commitment CommitmentType `json:"commitment,omitempty"`

	Encoding solana.EncodingType `json:"encoding,omitempty"`

	// Limit the returned account data
	DataSlice *DataSlice `json:"dataSlice,omitempty"`

	// Filter on accounts, implicit AND between filters.
	// Filter results using various filter objects;
	// account must meet all filter criteria to be included in results.
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
	Offset uint64        `json:"offset"`
	Bytes  solana.Base58 `json:"bytes"`
}

type CommitmentType string

const (
	CommitmentMax          CommitmentType = "max"
	CommitmentRecent       CommitmentType = "recent"
	CommitmentRoot         CommitmentType = "root"
	CommitmentSingle       CommitmentType = "single"
	CommitmentSingleGossip CommitmentType = "singleGossip"
)

// Parsed Transaction
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
	ProgramIDIndex uint16           `json:"programIdIndex"`
}

type InstructionInfo struct {
	Info            map[string]interface{} `json:"info"`
	InstructionType string                 `json:"type"`
}

func (p *ParsedInstruction) IsParsed() bool {
	return p.Parsed != nil
}

type M map[string]interface{}

// Unix timestamp (seconds since the Unix epoch)
type UnixTimeSeconds int64

func (res UnixTimeSeconds) Time() time.Time {
	return time.Unix(int64(res), 0)
}
