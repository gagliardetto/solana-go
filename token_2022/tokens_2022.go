package web3kit

import (
	"bytes"
	"context"
	binary2 "encoding/binary"
	"errors"
	"fmt"
	"reflect"

	binary "github.com/gagliardetto/binary"
	web3 "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token_metadata"
	"github.com/gagliardetto/solana-go/rpc"

	. "github.com/gagliardetto/solana-go/programs/spl_token_2022"

	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/cpi_guard"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/default_account_state"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/immutable_owner"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/metadata_pointer"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/mint_close_authority"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/non_transferable"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/permanent_delegate"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/transfer_fee"
	"github.com/gagliardetto/solana-go/programs/spl_token_2022/extension/transfer_hook"
)

var Token2022 = tokenKit2022{}

type tokenKit2022 struct {
}

var ACCOUNT_TYPE_SIZE = 1
var MULTISIG_SIZE = 355

func (t tokenKit2022) MEMO_TRANSFER_SIZE() int {
	return 1
}
func (t tokenKit2022) INTEREST_BEARING_MINT_CONFIG_STATE_SIZE() int {
	return 52
}
func (t tokenKit2022) TRANSFER_FEE_AMOUNT_SIZE() int {
	return 8
}
func (t tokenKit2022) LENGTH_SIZE() int {
	return 2
}
func (t tokenKit2022) TYPE_SIZE() int {
	return 2
}

func (t tokenKit2022) AddTypeAndLengthToLen(l int) uint64 {
	return t.AddTypeAndLengthToLen2(uint64(l))
}

func (t tokenKit2022) AddTypeAndLengthToLen2(l uint64) uint64 {
	return l + uint64(t.TYPE_SIZE()) + uint64(t.LENGTH_SIZE())
}

func (t tokenKit2022) GetMintLen(extensionTypes []ExtensionType) (uint64, error) {
	return t.GetLen(extensionTypes, MINT_SIZE)
}

func (t tokenKit2022) GetLen(extensionTypes []ExtensionType, baseSize int) (uint64, error) {
	if len(extensionTypes) == 0 {
		return uint64(baseSize), nil
	} else {
		var accountLength = uint64(ACCOUNT_SIZE + ACCOUNT_TYPE_SIZE)
		for _, ext := range extensionTypes {
			ret, err := t.GetTypeLen(ext)
			if err != nil {
				return 0, err
			}
			accountLength += t.AddTypeAndLengthToLen(ret)
		}
		if accountLength == uint64(MULTISIG_SIZE) {
			return accountLength + uint64(t.TYPE_SIZE()), nil
		} else {
			return accountLength, nil
		}
	}
}

func (t tokenKit2022) GetExtensionType(tlvData []byte) ([]ExtensionType, error) {
	var extensionTypes []ExtensionType
	var extensionTypeIndex uint64 = 0
	for {
		if extensionTypeIndex < uint64(len(tlvData)) {
			entryType, err := t.readUint16LE(tlvData, extensionTypeIndex)
			if err != nil {
				return nil, err
			}
			extensionTypes = append(extensionTypes, ExtensionType(entryType))
			entryLength, err := t.readUint16LE(tlvData, extensionTypeIndex+uint64(t.TYPE_SIZE()))
			if err != nil {
				return nil, err
			}
			extensionTypeIndex += t.AddTypeAndLengthToLen2(uint64(entryLength))
		} else {
			break
		}
	}
	return extensionTypes, nil
}

func (t tokenKit2022) GetExtensionData(extension ExtensionType, tlvData []byte) ([]byte, error) {
	var extensionTypeIndex uint64 = 0
	for {
		if t.AddTypeAndLengthToLen2(extensionTypeIndex) <= uint64(len(tlvData)) {
			entryType, err := t.readUint16LE(tlvData, extensionTypeIndex)
			if err != nil {
				return nil, err
			}
			entryLength, err := t.readUint16LE(tlvData, extensionTypeIndex+uint64(t.TYPE_SIZE()))
			if err != nil {
				return nil, err
			}
			var typeIndex = t.AddTypeAndLengthToLen2(extensionTypeIndex)
			if entryType == uint16(extension) {
				return tlvData[typeIndex : typeIndex+uint64(entryLength)], nil
			}
			extensionTypeIndex = typeIndex + uint64(entryLength)
		} else {
			break
		}
	}
	return nil, nil
}

func (t tokenKit2022) GetTypeLen(e ExtensionType) (int, error) {
	switch e {
	case ExtensionTypeUninitialized:
		return 0, nil
	case ExtensionTypeTransferFeeConfig:
		return transfer_fee.TRANSFER_FEE_CONFIG_SIZE, nil
	case ExtensionTypeTransferFeeAmount:
		return t.TRANSFER_FEE_AMOUNT_SIZE(), nil
	case ExtensionTypeMintCloseAuthority:
		return mint_close_authority.MINT_CLOSE_AUTHORITY_SIZE, nil
	case ExtensionTypeConfidentialTransferMint:
		return 97, nil
	case ExtensionTypeConfidentialTransferAccount:
		return 286, nil
	case ExtensionTypeCpiGuard:
		return cpi_guard.CPI_GUARD_SIZE, nil
	case ExtensionTypeDefaultAccountState:
		return default_account_state.DEFAULT_ACCOUNT_STATE_SIZE, nil
	case ExtensionTypeImmutableOwner:
		return immutable_owner.IMMUTABLE_OWNER_SIZE, nil
	case ExtensionTypeMemoTransfer:
		return t.MEMO_TRANSFER_SIZE(), nil
	case ExtensionTypeMetadataPointer:
		return metadata_pointer.METADATA_POINTER_SIZE, nil
	case ExtensionTypeNonTransferable:
		return non_transferable.NON_TRANSFERABLE_SIZE, nil
	case ExtensionTypeInterestBearingConfig:
		return t.INTEREST_BEARING_MINT_CONFIG_STATE_SIZE(), nil
	case ExtensionTypePermanentDelegate:
		return permanent_delegate.PERMANENT_DELEGATE_SIZE, nil
	case ExtensionTypeNonTransferableAccount:
		return non_transferable.NON_TRANSFERABLE_ACCOUNT_SIZE, nil
	case ExtensionTypeTransferHook:
		return transfer_hook.TRANSFER_HOOK_SIZE, nil
	case ExtensionTypeTransferHookAccount:
		return transfer_hook.TRANSFER_HOOK_ACCOUNT_SIZE, nil
	case ExtensionTypeTokenMetadata:
		return 0, fmt.Errorf("cannot get type length for variable extension type:%v", e)
	default:
		return 0, fmt.Errorf("unknown extension type: %v", e)
	}
}

func (t tokenKit2022) IsMintExtension(e ExtensionType) bool {
	switch e {
	case ExtensionTypeTransferFeeConfig:
		fallthrough
	case ExtensionTypeMintCloseAuthority:
		fallthrough
	case ExtensionTypeConfidentialTransferMint:
		fallthrough
	case ExtensionTypeDefaultAccountState:
		fallthrough
	case ExtensionTypeNonTransferable:
		fallthrough
	case ExtensionTypeInterestBearingConfig:
		fallthrough
	case ExtensionTypePermanentDelegate:
		fallthrough
	case ExtensionTypeTransferHook:
		fallthrough
	case ExtensionTypeMetadataPointer:
		fallthrough
	case ExtensionTypeTokenMetadata:
		return true
	case ExtensionTypeUninitialized:
		fallthrough
	case ExtensionTypeTransferFeeAmount:
		fallthrough
	case ExtensionTypeConfidentialTransferAccount:
		fallthrough
	case ExtensionTypeImmutableOwner:
		fallthrough
	case ExtensionTypeMemoTransfer:
		fallthrough
	case ExtensionTypeCpiGuard:
		fallthrough
	case ExtensionTypeNonTransferableAccount:
		fallthrough
	case ExtensionTypeTransferHookAccount:
		return false
	default:
		panic(fmt.Sprintf("Unknown extension type: %v", e))
	}
}

type MintInfo struct {
	Mint
	Address web3.PublicKey // address of the mint
	TlvData []byte         // Additional data for extension
}

func GetMint(
	ctx context.Context,
	client *rpc.Client,
	mint, programId web3.PublicKey,
	opts *rpc.GetAccountInfoOpts,
) (*MintInfo, error) {
	_ = ctx
	info, err := client.GetAccountInfoWithOpts(ctx, mint, opts)
	if err != nil {
		return nil, err
	}
	return UnpackMint(mint, info.Value, programId)
}

func ParseMint(data []byte) (*Mint, error) {
	if len(data) < MINT_SIZE {
		return nil, InvalidAccountSizeErr
	}
	return decodeObject[*Mint](data[0:MINT_SIZE])
}

var (
	TokenInvalidAccountOwnerErr = fmt.Errorf("TokenInvalidAccountOwnerErr")
	InvalidAccountSizeErr       = errors.New("InvalidAccountSizeErr")
	TokenAccountNotFoundErr     = errors.New("TokenAccountNotFoundErr")
	TokenInvalidMintErr         = errors.New("TokenInvalidMintErr")
)

func UnpackMint(mintAddress web3.PublicKey, account *rpc.Account, programId web3.PublicKey) (*MintInfo, error) {
	if account.Owner != programId {
		return nil, TokenInvalidAccountOwnerErr
	}
	data := account.Data.GetBinary()
	raw, err := ParseMint(data)
	if err != nil {
		return nil, err
	}
	var ret = &MintInfo{
		Address: mintAddress,
		Mint:    *raw,
	}
	if len(data) > MINT_SIZE {
		if len(data) <= ACCOUNT_SIZE {
			return nil, InvalidAccountSizeErr
		}
		if len(data) == MULTISIG_SIZE {
			return nil, InvalidAccountSizeErr
		}
		if data[ACCOUNT_SIZE] != 1 /*AccountType.Mint*/ {
			return nil, TokenInvalidMintErr
		}
		ret.TlvData = data[ACCOUNT_SIZE+ACCOUNT_TYPE_SIZE:]
	}
	return ret, nil
}

type TokenAccount struct {
	Account
	Address web3.PublicKey // address of the token
	TlvData []byte         // Additional data for extension
}

func GetTokenAccount(
	ctx context.Context,
	client *rpc.Client,
	account, programId web3.PublicKey,
	opts *rpc.GetAccountInfoOpts,
) (*TokenAccount, error) {
	_ = ctx
	info, err := client.GetAccountInfoWithOpts(ctx, account, opts)
	if err != nil {
		return nil, err
	}
	return UnpackTokenAccount(account, info.Value, programId)
}

func ParseTokenAccount(data []byte) (*Account, error) {
	if len(data) < ACCOUNT_SIZE {
		return nil, InvalidAccountSizeErr
	}
	return decodeObject[*Account](data[0:ACCOUNT_SIZE])
}

func UnpackTokenAccount(tokenAccount web3.PublicKey, info *rpc.Account, programId web3.PublicKey) (*TokenAccount, error) {
	if info == nil {
		return nil, TokenAccountNotFoundErr
	}
	if info.Owner != programId {
		return nil, TokenInvalidAccountOwnerErr
	}
	data := info.Data.GetBinary()
	raw, err := ParseTokenAccount(data)
	if err != nil {
		return nil, err
	}
	var ret = &TokenAccount{
		Address: tokenAccount,
		Account: *raw,
	}
	if len(data) > ACCOUNT_SIZE {
		if len(data) == MULTISIG_SIZE {
			return nil, InvalidAccountSizeErr
		}
		if data[ACCOUNT_SIZE] != 2 /*AccountType.Account*/ {
			return nil, TokenInvalidMintErr
		}
		ret.TlvData = data[ACCOUNT_SIZE+ACCOUNT_TYPE_SIZE:]
	}
	return ret, nil
}

func (t tokenKit2022) readUint16LE(data []byte, index uint64) (uint16, error) {
	var value uint16
	err := binary2.Read(bytes.NewReader(data[index:]), binary2.LittleEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ParseDefaultAccountState Extension: default_account_state
func (t tokenKit2022) ParseDefaultAccountState(data []byte) (*default_account_state.DefaultAccountState, error) {
	return parseExtension[*default_account_state.DefaultAccountState](ExtensionTypeDefaultAccountState, data)
}

// ParseTransferFeeConfig Extension: transfer_fee
func (t tokenKit2022) ParseTransferFeeConfig(data []byte) (*transfer_fee.TransferFeeConfig, error) {
	return parseExtension[*transfer_fee.TransferFeeConfig](ExtensionTypeTransferFeeConfig, data)
}

// GetTokenMetadata Extension: token_metadata
func GetTokenMetadata(
	ctx context.Context,
	connection *rpc.Client,
	mint, programId web3.PublicKey,
	opts *rpc.GetAccountInfoOpts,
) (*token_metadata.TokenMetadata, error) {
	mintInfo, err := GetMint(ctx, connection, mint, programId, opts)
	if err != nil {
		return nil, err
	}
	return ParseTokenMetadata(mintInfo.TlvData)
}

// ParseTokenMetadata Extension: token_metadata
func ParseTokenMetadata(data []byte) (*token_metadata.TokenMetadata, error) {
	return parseExtension[*token_metadata.TokenMetadata](ExtensionTypeTokenMetadata, data)
}

func parseExtension[T binary.BinaryUnmarshaler](extension ExtensionType, data []byte) (T, error) {
	var zero T
	extensionData, err := Token2022.GetExtensionData(extension, data)
	if err != nil {
		return zero, err
	}
	return decodeObject[T](extensionData)
}

func decodeObject[T binary.BinaryUnmarshaler](data []byte) (T, error) {
	var zero T
	if len(data) == 0 {
		return zero, nil
	}
	ret := reflect.New(reflect.TypeOf(zero).Elem()).Interface().(T)
	if err := decode(data, ret); err != nil {
		return zero, err
	}
	return ret, nil
}

func decode[T binary.BinaryUnmarshaler](data []byte, input T) error {
	var decoder = binary.NewDecoderWithEncoding(data, binary.EncodingBorsh)
	err := decoder.Decode(input)
	if err != nil {
		return err
	}
	return nil
}
