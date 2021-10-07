// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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

package solana

import (
	"fmt"
)

// InstructionDecoder receives the AccountMeta FOR THAT INSTRUCTION,
// and not the accounts of the *Message object. Resolve with
// CompiledInstruction.ResolveInstructionAccounts(message) beforehand.
type InstructionDecoder func(instructionAccounts []*AccountMeta, data []byte) (interface{}, error)

var InstructionDecoderRegistry = map[string]InstructionDecoder{}

func RegisterInstructionDecoder(programID PublicKey, decoder InstructionDecoder) {
	pid := programID.String()
	if _, found := InstructionDecoderRegistry[pid]; found {
		panic(fmt.Sprintf("unable to re-register instruction decoder for program %q", pid))
	}

	InstructionDecoderRegistry[pid] = decoder
}

func DecodeInstruction(programID PublicKey, accounts []*AccountMeta, data []byte) (interface{}, error) {
	pid := programID.String()

	decoder, found := InstructionDecoderRegistry[pid]
	if !found {
		return nil, fmt.Errorf("instruction decoder not found for %s", pid)
	}

	return decoder(accounts, data)
}
