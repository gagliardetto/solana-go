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
	"bytes"

	"github.com/lunixbochs/struc"
)

type AccountMeta struct {
	PublicKey  PublicKey
	IsSigner   bool
	IsWritable bool
}

type Instruction struct {
	ProgramID PublicKey
	Accounts  []AccountMeta
	Data      Base58
}

func NewInstruction(programID PublicKey, accountMetas []AccountMeta, instruction interface{}) (*Instruction, error) {
	buf := &bytes.Buffer{}
	err := struc.Pack(buf, instruction)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		ProgramID: programID,
		Accounts:  accountMetas,
		Data:      Base58(buf.Bytes()),
	}, nil
}
