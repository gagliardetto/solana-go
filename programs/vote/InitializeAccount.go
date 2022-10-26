// Copyright 2021 github.com/gagliardetto
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

package vote

import (
	"github.com/gagliardetto/solana-go"
)

type InitializeAccount struct {
	// TODO

	// [0] = [WRITE] VoteAccount
	// ··········· Vote account to vote with
	//
	// [1] = [] SysVarSlotHashes
	// ··········· Slot hashes sysvar
	//
	// [2] = [] SysVarClock
	// ··········· Clock sysvar
	//
	// [3] = [SIGNER] VoteAuthority
	// ··········· New validator identity (node_pubkey)
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}
