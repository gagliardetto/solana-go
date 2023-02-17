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

package computebudget

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRequestUnitsDeprecatedInstruction(t *testing.T) {

	t.Run("should validate max units", func(t *testing.T) {
		_, err := NewRequestUnitsDeprecatedInstruction(2000000, 1000).ValidateAndBuild()
		require.Error(t, err)
	})

	t.Run("should build request units ix", func(t *testing.T) {
		ix, err := NewRequestUnitsDeprecatedInstruction(1400000, 1000).ValidateAndBuild()
		require.Nil(t, err)

		require.Equal(t, ProgramID, ix.ProgramID())
		require.Equal(t, 0, len(ix.Accounts()))

		data, err := ix.Data()
		require.Nil(t, err)
		require.Equal(t, []byte{0x0, 0xc0, 0x5c, 0x15, 0x0, 0xe8, 0x3, 0x0, 0x0}, data)
	})

}
