package rpc

import (
	"context"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetSignatureStatusesResult struct {
	RPCContext
	Value []*SignatureStatusesResult `json:"value"`
}

type SignatureStatusesResult struct {
	Slot bin.Uint64 `json:"slot"` // The slot the transaction was processed
	// TODO: what is <usize | null> ???
	Confirmations      *bin.Uint64                     `json:"confirmations"`      // Number of blocks since signature confirmation, null if rooted, as well as finalized by a supermajority of the cluster
	Err                interface{}                     `json:"err"`                // Error if transaction failed, null if transaction succeeded
	ConfirmationStatus ConfirmationStatusType          `json:"confirmationStatus"` // The transaction's cluster confirmation status; either processed, confirmed, or finalized
	Status             DeprecatedTransactionMetaStatus `json:"status"`             // DEPRECATED: Transaction status
}

type ConfirmationStatusType string

const (
	ConfirmationStatusProcessed ConfirmationStatusType = "processed"
	ConfirmationStatusConfirmed ConfirmationStatusType = "confirmed"
	ConfirmationStatusFinalized ConfirmationStatusType = "finalized"
)

// GetSignatureStatuses Returns the statuses of a list of signatures.
// Unless the searchTransactionHistory configuration parameter
// is included,this method only searches the recent status cache
// of signatures, which retains statuses for all active slots plus
// MAX_RECENT_BLOCKHASHES rooted slots.
func (cl *Client) GetSignatureStatuses(
	ctx context.Context,
	searchTransactionHistory bool, // if true, a Solana node will search its ledger cache for any signatures not found in the recent status cache
	transactionSignatures ...solana.Signature, //transaction signatures to confirm
) (out *GetSignatureStatusesResult, err error) {
	params := []interface{}{transactionSignatures}
	if searchTransactionHistory {
		params = append(params, M{"searchTransactionHistory": searchTransactionHistory})
	}
	err = cl.rpcClient.CallFor(&out, "getSignatureStatuses", params)

	if err != nil {
		return nil, err
	}

	if out.Value == nil {
		return nil, ErrNotFound
	}

	return
}
