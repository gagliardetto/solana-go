// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package transfer_hook

import (
	binary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go/programs/common"
)

// TransferHook Struct
type TransferHook struct {
	// Authority that can set the transfer hook program id
	Authority common.PublicKey
	// Program that authorizes the transfer
	ProgramId common.PublicKey
}

const TRANSFER_HOOK_SIZE = 64

func (obj *TransferHook) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Authority); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.ProgramId); err != nil {
		return err
	}
	return nil
}

func (obj *TransferHook) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Authority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.ProgramId); err != nil {
		return err
	}
	return nil
}

// TransferHookAccount Struct
// Indicates that the tokens from this account belong to a mint with a transfer
// hook
type TransferHookAccount struct {
	// Flag to indicate that the account is in the middle of a transfer
	Transferring bool
}

const TRANSFER_HOOK_ACCOUNT_SIZE = 1

func (obj *TransferHookAccount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Transferring); err != nil {
		return err
	}
	return nil
}

func (obj *TransferHookAccount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Transferring); err != nil {
		return err
	}
	return nil
}
