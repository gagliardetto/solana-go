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

package serum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetSeqNum(t *testing.T) {
	seqNum, err := GetSeqNum("0000000000000eedffffffffffa78933", SideBid)
	require.NoError(t, err)
	assert.Equal(t, uint64(5797580), seqNum)

	seqNum, err = GetSeqNum("0000000000000eed00000000005876cc", SideAsk)
	require.NoError(t, err)
	assert.Equal(t, uint64(5797580), seqNum)

	seqNum, err = GetSeqNum("0000000000000840ffffffffffacdefd", SideBid)
	require.NoError(t, err)
	assert.Equal(t, uint64(5447938), seqNum)
}

func Test_PriceLotsToNumber(t *testing.T) {
	price, err := GetSeqNum("0000000000000eedffffffffffa78933", SideBid)
	require.NoError(t, err)
	assert.Equal(t, uint64(5797580), price)

	price, err = GetSeqNum("0000000000000eed00000000005876cc", SideAsk)
	require.NoError(t, err)
	assert.Equal(t, uint64(5797580), price)
}
