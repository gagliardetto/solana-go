package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

type GetRecentPerformanceSamplesResult struct {
	Slot             bin.Uint64 `json:"slot"`             // Slot in which sample was taken at
	NumTransactions  bin.Uint64 `json:"numTransactions"`  // Number of transactions in sample
	NumSlots         bin.Uint64 `json:"numSlots"`         // Number of slots in sample
	SamplePeriodSecs uint16     `json:"samplePeriodSecs"` // Number of seconds in a sample window
}

// GetRecentPerformanceSamples returns a list of recent performance samples,
// in reverse slot order. Performance samples are taken every 60 seconds
// and include the number of transactions and slots that occur in a given time window.
func (cl *Client) GetRecentPerformanceSamples(ctx context.Context, limit *int) (out []*GetRecentPerformanceSamplesResult, err error) {
	params := []interface{}{}
	if limit != nil {
		params = append(params, limit)
	}
	err = cl.rpcClient.CallFor(&out, "getRecentPerformanceSamples", params)
	return
}
