package rpc

import (
	"context"
	"fmt"

	bin "github.com/dfuse-io/binary"
	"github.com/gagliardetto/solana-go"
)

type GetTransactionOpts struct {
	Encoding solana.EncodingType `json:"encoding,omitempty"`

	// Desired commitment. "processed" is not supported. If parameter not provided, the default is "finalized".
	Commitment CommitmentType `json:"commitment,omitempty"`
}

// GetTransaction returns transaction details for a confirmed transaction.
//
// NEW: This method is only available in solana-core v1.7 or newer.
// Please use `getConfirmedTransaction` for solana-core v1.6
func (cl *Client) GetTransaction(
	ctx context.Context,
	txSig solana.Signature, // transaction signature
	opts *GetTransactionOpts,
) (out *GetTransactionResult, err error) {
	params := []interface{}{txSig}
	if opts != nil {
		obj := M{}
		if opts.Encoding != "" {
			if !solana.IsAnyOfEncodingType(
				opts.Encoding,
				// Valid encodings:
				solana.EncodingJSON,
				// solana.EncodingJSONParsed, // TODO
				solana.EncodingBase58,
				solana.EncodingBase64,
			) {
				return nil, fmt.Errorf("provided encoding is not supported: %s", opts.Encoding)
			}
			obj["encoding"] = opts.Encoding
		}
		if opts.Commitment != "" {
			obj["commitment"] = opts.Commitment
		}
		if len(obj) > 0 {
			params = append(params, obj)
		}
	}
	err = cl.rpcClient.CallForInto(ctx, &out, "getTransaction", params)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrNotFound
	}
	return
}

type GetTransactionResult struct {
	// The slot this transaction was processed in.
	Slot bin.Uint64 `json:"slot"`

	// Estimated production time, as Unix timestamp (seconds since the Unix epoch)
	// of when the transaction was processed.
	// Nil if not available.
	BlockTime *UnixTimeSeconds `json:"blockTime"`

	Transaction *TransactionResultEnvelope `json:"transaction"`
	Meta        *TransactionMeta           `json:"meta,omitempty"`
}

// TransactionResultEnvelope will contain a *ParsedTransaction if the requested encoding is `solana.EncodingJSON`
// (which is also the default when the encoding is not specified),
// or a `solana.Data` in case of EncodingBase58, EncodingBase64.
type TransactionResultEnvelope struct {
	asDecodedBinary     solana.Data
	asParsedTransaction *ParsedTransaction
}

func (wrap TransactionResultEnvelope) MarshalJSON() ([]byte, error) {
	if wrap.asParsedTransaction != nil {
		return json.Marshal(wrap.asParsedTransaction)
	}
	return json.Marshal(wrap.asDecodedBinary)
}

func (wrap *TransactionResultEnvelope) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || (len(data) == 4 && string(data) == "null") {
		// TODO: is this an error?
		return nil
	}

	firstChar := data[0]

	switch firstChar {
	// Check if first character is `[`, standing for a JSON array.
	case '[':
		// It's base64 (or similar)
		{
			err := wrap.asDecodedBinary.UnmarshalJSON(data)
			if err != nil {
				return err
			}
		}
	case '{':
		// It's JSON, most likely.
		{
			return json.Unmarshal(data, &wrap.asParsedTransaction)
		}
	default:
		return fmt.Errorf("Unknown kind: %v", data)
	}

	return nil
}

// GetBinary returns the decoded bytes if the encoding is
// "base58", "base64".
func (dt *TransactionResultEnvelope) GetBinary() []byte {
	return dt.asDecodedBinary.Content
}

// GetRawJSON returns a *ParsedTransaction when the data
// encoding is EncodingJSON.
func (dt *TransactionResultEnvelope) GetParsedTransaction() *ParsedTransaction {
	return dt.asParsedTransaction
}
