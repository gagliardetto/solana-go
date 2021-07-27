package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// GetRecentPerformanceSamples returns a list of recent performance samples,
// in reverse slot order. Performance samples are taken every 60 seconds
// and include the number of transactions and slots that occur in a given time window.
func (cl *Client) GetRecentPerformanceSamples(
	ctx context.Context,
	limit *uint,
) (out []*GetRecentPerformanceSamplesResult, err error) {
	params := []interface{}{}
	if limit != nil {
		params = append(params, limit)
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getRecentPerformanceSamples", params)
	return
}

type GetRecentPerformanceSamplesResult struct {
	// Slot in which sample was taken at.
	Slot bin.Uint64 `json:"slot"`

	// Number of transactions in sample.
	NumTransactions bin.Uint64 `json:"numTransactions"`

	// Number of slots in sample.
	NumSlots bin.Uint64 `json:"numSlots"`

	// Number of seconds in a sample window.
	SamplePeriodSecs uint16 `json:"samplePeriodSecs"`
}
