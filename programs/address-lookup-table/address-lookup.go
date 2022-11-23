package addresslookuptable

import (
	"context"
	"fmt"
	"math"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// The serialized size of lookup table metadata.
const LOOKUP_TABLE_META_SIZE = 56

// DecodeAddressLookupTableState decodes the given account bytes into a AddressLookupTableState.
func DecodeAddressLookupTableState(data []byte) (*AddressLookupTableState, error) {
	decoder := bin.NewBinDecoder(data)
	var state AddressLookupTableState
	if err := state.UnmarshalWithDecoder(decoder); err != nil {
		return nil, err
	}
	return &state, nil
}

func GetAddressLookupTable(
	ctx context.Context,
	rpcClient *rpc.Client,
	address solana.PublicKey,
) (*AddressLookupTableState, error) {
	account, err := rpcClient.GetAccountInfo(ctx, address)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("account not found")
	}
	return DecodeAddressLookupTableState(account.GetBinary())
}

func GetAddressLookupTableStateWithOpts(
	ctx context.Context,
	rpcClient *rpc.Client,
	address solana.PublicKey,
	opts *rpc.GetAccountInfoOpts,
) (*AddressLookupTableState, error) {
	account, err := rpcClient.GetAccountInfoWithOpts(ctx, address, opts)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, fmt.Errorf("account not found")
	}
	return DecodeAddressLookupTableState(account.GetBinary())
}

type AddressLookupTableState struct {
	TypeIndex                  uint32
	DeactivationSlot           uint64
	LastExtendedSlot           uint64
	LastExtendedSlotStartIndex uint8
	Authority                  *solana.PublicKey
	Addresses                  solana.PublicKeySlice
}

func (a *AddressLookupTableState) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	if a.TypeIndex, err = decoder.ReadUint32(bin.LE); err != nil {
		return fmt.Errorf("failed to decode TypeIndex: %w", err)
	}
	if a.DeactivationSlot, err = decoder.ReadUint64(bin.LE); err != nil {
		return err
	}
	if a.LastExtendedSlot, err = decoder.ReadUint64(bin.LE); err != nil {
		return err
	}
	if a.LastExtendedSlotStartIndex, err = decoder.ReadUint8(); err != nil {
		return err
	}
	has, err := decoder.ReadBool()
	if err != nil {
		return err
	}
	if has {
		var auth solana.PublicKey
		if _, err := decoder.Read(auth[:]); err != nil {
			return err
		}
		a.Authority = &auth
	} else {
		err = decoder.Discard(32)
		if err != nil {
			return err
		}
	}
	serializedAddressesLen := decoder.Len() - LOOKUP_TABLE_META_SIZE
	numSerializedAddresses := serializedAddressesLen / 32
	if serializedAddressesLen%32 != 0 {
		return fmt.Errorf("lookup table is invalid")
	}

	decoder.SetPosition(decoder.Position() + 2)
	a.Addresses = make(solana.PublicKeySlice, numSerializedAddresses)

	for i := 0; i < numSerializedAddresses; i++ {
		var address solana.PublicKey
		numRead, err := decoder.Read(address[:])
		if err != nil {
			return fmt.Errorf("failed to read addresses[%d]: %w", i, err)
		}
		if numRead != 32 {
			return fmt.Errorf("failed to read addresses[%d]: expected to read 32, but read %d", i, numRead)
		}
		a.Addresses[i] = address
	}
	if decoder.Remaining() != 0 {
		return fmt.Errorf("failed to read all addresses: remaining %d bytes", decoder.Remaining())
	}
	return nil
}

func (a AddressLookupTableState) MarshalWithEncoder(encoder *bin.Encoder) error {
	if err := encoder.WriteUint32(a.TypeIndex, bin.LE); err != nil {
		return err
	}
	if err := encoder.WriteUint64(a.DeactivationSlot, bin.LE); err != nil {
		return err
	}
	if err := encoder.WriteUint64(a.LastExtendedSlot, bin.LE); err != nil {
		return err
	}
	if err := encoder.WriteUint8(a.LastExtendedSlotStartIndex); err != nil {
		return err
	}
	if a.Authority != nil {
		if err := encoder.WriteBool(true); err != nil {
			return err
		}
		if _, err := encoder.Write(a.Authority[:]); err != nil {
			return err
		}
	} else {
		if err := encoder.WriteBool(false); err != nil {
			return err
		}
		if _, err := encoder.Write(make([]byte, 32)); err != nil {
			return err
		}
	}
	if _, err := encoder.Write(make([]byte, 2)); err != nil {
		return err
	}
	for _, address := range a.Addresses {
		if _, err := encoder.Write(address[:]); err != nil {
			return err
		}
	}
	return nil
}

type KeyedAddressLookupTable struct {
	Key   solana.PublicKey
	State AddressLookupTableState
}

func NewKeyedAddressLookupTable(key solana.PublicKey) *KeyedAddressLookupTable {
	return &KeyedAddressLookupTable{
		Key: key,
	}
}

func (a AddressLookupTableState) IsActive() bool {
	return a.DeactivationSlot == math.MaxUint64
}

func (a *KeyedAddressLookupTable) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	if err := a.State.UnmarshalWithDecoder(decoder); err != nil {
		return err
	}
	return nil
}

func (a KeyedAddressLookupTable) MarshalWithEncoder(encoder *bin.Encoder) error {
	if err := a.State.MarshalWithEncoder(encoder); err != nil {
		return err
	}
	return nil
}
