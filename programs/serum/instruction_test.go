package serum

import (
	"encoding/hex"
	"testing"

	bin "github.com/dfuse-io/binary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeInstruction(t *testing.T) {
	tests := []struct {
		name              string
		hexData           string
		expectInstruction *Instruction
	}{
		{
			name:    "New Order",
			hexData: "000900000001000000b80600000000000010eb09000000000000000000168106e091da511601000000",
			expectInstruction: &Instruction{
				BaseVariant: bin.BaseVariant{
					TypeID: 9,
					Impl: &InstructionNewOrderV2{
						Side:              SideAsk,
						LimitPrice:        1720,
						MaxQuantity:       650000,
						OrderType:         OrderTypeLimit,
						ClientID:          1608306862011613462,
						SelfTradeBehavior: SelfTradeBehaviorCancelProvide,
					},
				},
				Version: 0,
			},
		},
		{
			name:    "Match Order",
			hexData: "0002000000ffff",
			expectInstruction: &Instruction{
				BaseVariant: bin.BaseVariant{
					TypeID: 2,
					Impl: &InstructionMatchOrder{
						Limit: 65535,
					},
				},
				Version: 0,
			},
		},
		{
			name:    "Settle Funds",
			hexData: "0005000000",
			expectInstruction: &Instruction{
				BaseVariant: bin.BaseVariant{
					TypeID: 5,
					Impl:   &InstructionSettleFunds{},
				},
				Version: 0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := hex.DecodeString(test.hexData)
			require.NoError(t, err)
			var instruction *Instruction
			err = bin.NewBinDecoder(data).Decode(&instruction)
			require.NoError(t, err)
			assert.Equal(t, test.expectInstruction, instruction)
		})
	}
}
