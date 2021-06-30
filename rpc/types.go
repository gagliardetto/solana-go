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
	Pubkey   solana.PublicKey `json:"pubkey"`
	Lamports bin.Uint64       `json:"lamports"`
}

type TransactionWithMeta struct {
	Transaction *solana.Transaction `json:"transaction"`
	Meta        *TransactionMeta    `json:"meta,omitempty"`
}

type TransactionParsed struct {
	Transaction *ParsedTransaction `json:"transaction"`
	Meta        *TransactionMeta   `json:"meta,omitempty"`
}

type TransactionMeta struct {
	Err          interface{}  `json:"err"`
	Fee          bin.Uint64   `json:"fee"`
	PreBalances  []bin.Uint64 `json:"preBalances"`
	PostBalances []bin.Uint64 `json:"postBalances"`
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
