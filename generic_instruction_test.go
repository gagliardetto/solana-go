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

package solana

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInstruction(t *testing.T) {
	progID := MemoProgramID
	accounts := []*AccountMeta{
		Meta(SPLAssociatedTokenAccountProgramID).SIGNER().WRITE(),
	}
	data := []byte("hello world")

	ins := NewInstruction(progID, accounts, data)

	require.Equal(t, progID, ins.ProgramID())
	require.Equal(t, accounts, ins.Accounts())
	{
		got, err := ins.Data()
		require.NoError(t, err)
		require.Equal(t, data, got)
	}
}
