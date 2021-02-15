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
