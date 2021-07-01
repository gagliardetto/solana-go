package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
)

// the result field will be a dictionary of validator identities,
// as base-58 encoded strings, and their corresponding leader
// slot indices as values (indices are relative to the first
// slot in the requested epoch)
type GetLeaderScheduleResult map[string][]bin.Uint64

// GetLeaderSchedule returns the leader schedule for an epoch.
func (cl *Client) GetLeaderSchedule(
	ctx context.Context,
	epoch *uint64, // Fetch the leader schedule for the epoch that corresponds to the provided slot. If unspecified, the leader schedule for the current epoch is fetched
	commitment CommitmentType,
	// TODO: is identity a pubkey?
	identity string, // Only return results for this validator identity (base-58 encoded)
) (out *GetLeaderScheduleResult, err error) {
	params := []interface{}{}
	if epoch != nil {
		params = append(params, epoch)
	}
	obj := M{}
	if commitment != "" {
		obj["commitment"] = commitment
	}
	if identity != "" {
		obj["identity"] = identity
	}
	if len(obj) > 0 {
		params = append(params, obj)
	}
	err = cl.rpcClient.CallFor(&out, "getLeaderSchedule", params...)
	if err != nil {
		return nil, err
	}
	// TODO: check that this behaviour is implemented everywhere:
	if out == nil {
		return nil, ErrNotFound
	}
	return
}
