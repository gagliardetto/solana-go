// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_cp

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AmmConfig struct {
	// Bump to identify PDA
	Bump uint8

	// Status to control if new pool can be create
	DisableCreatePool bool

	// Config index
	Index uint16

	// The trade fee, denominated in hundredths of a bip (10^-6)
	TradeFeeRate uint64

	// The protocol fee
	ProtocolFeeRate uint64

	// The fund fee, denominated in hundredths of a bip (10^-6)
	FundFeeRate uint64

	// Fee for create a new pool
	CreatePoolFee uint64

	// Address of the protocol fee owner
	ProtocolOwner ag_solanago.PublicKey

	// Address of the fund fee owner
	FundOwner ag_solanago.PublicKey

	// padding
	Padding [16]uint64
}

var AmmConfigDiscriminator = [8]byte{218, 244, 33, 104, 203, 203, 43, 111}

func (obj AmmConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AmmConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `DisableCreatePool` param:
	err = encoder.Encode(obj.DisableCreatePool)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeRate` param:
	err = encoder.Encode(obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeRate` param:
	err = encoder.Encode(obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Serialize `FundFeeRate` param:
	err = encoder.Encode(obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Serialize `CreatePoolFee` param:
	err = encoder.Encode(obj.CreatePoolFee)
	if err != nil {
		return err
	}
	// Serialize `ProtocolOwner` param:
	err = encoder.Encode(obj.ProtocolOwner)
	if err != nil {
		return err
	}
	// Serialize `FundOwner` param:
	err = encoder.Encode(obj.FundOwner)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmmConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AmmConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[218 244 33 104 203 203 43 111]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `DisableCreatePool`:
	err = decoder.Decode(&obj.DisableCreatePool)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `TradeFeeRate`:
	err = decoder.Decode(&obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeRate`:
	err = decoder.Decode(&obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `FundFeeRate`:
	err = decoder.Decode(&obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `CreatePoolFee`:
	err = decoder.Decode(&obj.CreatePoolFee)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolOwner`:
	err = decoder.Decode(&obj.ProtocolOwner)
	if err != nil {
		return err
	}
	// Deserialize `FundOwner`:
	err = decoder.Decode(&obj.FundOwner)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type ObservationState struct {
	// Whether the ObservationState is initialized
	Initialized bool

	// the most-recently updated index of the observations array
	ObservationIndex uint16
	PoolId           ag_solanago.PublicKey

	// observation array
	Observations [100]Observation

	// padding for feature update
	Padding [4]uint64
}

var ObservationStateDiscriminator = [8]byte{122, 174, 197, 53, 129, 9, 165, 132}

func (obj ObservationState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ObservationStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized` param:
	err = encoder.Encode(obj.Initialized)
	if err != nil {
		return err
	}
	// Serialize `ObservationIndex` param:
	err = encoder.Encode(obj.ObservationIndex)
	if err != nil {
		return err
	}
	// Serialize `PoolId` param:
	err = encoder.Encode(obj.PoolId)
	if err != nil {
		return err
	}
	// Serialize `Observations` param:
	err = encoder.Encode(obj.Observations)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ObservationState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ObservationStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[122 174 197 53 129 9 165 132]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized`:
	err = decoder.Decode(&obj.Initialized)
	if err != nil {
		return err
	}
	// Deserialize `ObservationIndex`:
	err = decoder.Decode(&obj.ObservationIndex)
	if err != nil {
		return err
	}
	// Deserialize `PoolId`:
	err = decoder.Decode(&obj.PoolId)
	if err != nil {
		return err
	}
	// Deserialize `Observations`:
	err = decoder.Decode(&obj.Observations)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

type PoolState struct {
	// Which config the pool belongs
	AmmConfig ag_solanago.PublicKey

	// pool creator
	PoolCreator ag_solanago.PublicKey

	// Token A
	Token0Vault ag_solanago.PublicKey

	// Token B
	Token1Vault ag_solanago.PublicKey

	// Pool tokens are issued when A or B tokens are deposited.
	// Pool tokens can be withdrawn back to the original A or B token.
	LpMint ag_solanago.PublicKey

	// Mint information for token A
	Token0Mint ag_solanago.PublicKey

	// Mint information for token B
	Token1Mint ag_solanago.PublicKey

	// token_0 program
	Token0Program ag_solanago.PublicKey

	// token_1 program
	Token1Program ag_solanago.PublicKey

	// observation account to store oracle data
	ObservationKey ag_solanago.PublicKey
	AuthBump       uint8

	// Bitwise representation of the state of the pool
	// bit0, 1: disable deposit(vaule is 1), 0: normal
	// bit1, 1: disable withdraw(vaule is 2), 0: normal
	// bit2, 1: disable swap(vaule is 4), 0: normal
	Status         uint8
	LpMintDecimals uint8

	// mint0 and mint1 decimals
	Mint0Decimals uint8
	Mint1Decimals uint8

	// lp mint supply
	LpSupply uint64

	// The amounts of token_0 and token_1 that are owed to the liquidity provider.
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64
	FundFeesToken0     uint64
	FundFeesToken1     uint64

	// The timestamp allowed for swap in the pool.
	OpenTime uint64

	// padding for future updates
	Padding [32]uint64
}

var PoolStateDiscriminator = [8]byte{247, 237, 227, 245, 215, 195, 222, 70}

func (obj PoolState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PoolStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `AmmConfig` param:
	err = encoder.Encode(obj.AmmConfig)
	if err != nil {
		return err
	}
	// Serialize `PoolCreator` param:
	err = encoder.Encode(obj.PoolCreator)
	if err != nil {
		return err
	}
	// Serialize `Token0Vault` param:
	err = encoder.Encode(obj.Token0Vault)
	if err != nil {
		return err
	}
	// Serialize `Token1Vault` param:
	err = encoder.Encode(obj.Token1Vault)
	if err != nil {
		return err
	}
	// Serialize `LpMint` param:
	err = encoder.Encode(obj.LpMint)
	if err != nil {
		return err
	}
	// Serialize `Token0Mint` param:
	err = encoder.Encode(obj.Token0Mint)
	if err != nil {
		return err
	}
	// Serialize `Token1Mint` param:
	err = encoder.Encode(obj.Token1Mint)
	if err != nil {
		return err
	}
	// Serialize `Token0Program` param:
	err = encoder.Encode(obj.Token0Program)
	if err != nil {
		return err
	}
	// Serialize `Token1Program` param:
	err = encoder.Encode(obj.Token1Program)
	if err != nil {
		return err
	}
	// Serialize `ObservationKey` param:
	err = encoder.Encode(obj.ObservationKey)
	if err != nil {
		return err
	}
	// Serialize `AuthBump` param:
	err = encoder.Encode(obj.AuthBump)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `LpMintDecimals` param:
	err = encoder.Encode(obj.LpMintDecimals)
	if err != nil {
		return err
	}
	// Serialize `Mint0Decimals` param:
	err = encoder.Encode(obj.Mint0Decimals)
	if err != nil {
		return err
	}
	// Serialize `Mint1Decimals` param:
	err = encoder.Encode(obj.Mint1Decimals)
	if err != nil {
		return err
	}
	// Serialize `LpSupply` param:
	err = encoder.Encode(obj.LpSupply)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeesToken0` param:
	err = encoder.Encode(obj.ProtocolFeesToken0)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeesToken1` param:
	err = encoder.Encode(obj.ProtocolFeesToken1)
	if err != nil {
		return err
	}
	// Serialize `FundFeesToken0` param:
	err = encoder.Encode(obj.FundFeesToken0)
	if err != nil {
		return err
	}
	// Serialize `FundFeesToken1` param:
	err = encoder.Encode(obj.FundFeesToken1)
	if err != nil {
		return err
	}
	// Serialize `OpenTime` param:
	err = encoder.Encode(obj.OpenTime)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PoolState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PoolStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[247 237 227 245 215 195 222 70]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `AmmConfig`:
	err = decoder.Decode(&obj.AmmConfig)
	if err != nil {
		return err
	}
	// Deserialize `PoolCreator`:
	err = decoder.Decode(&obj.PoolCreator)
	if err != nil {
		return err
	}
	// Deserialize `Token0Vault`:
	err = decoder.Decode(&obj.Token0Vault)
	if err != nil {
		return err
	}
	// Deserialize `Token1Vault`:
	err = decoder.Decode(&obj.Token1Vault)
	if err != nil {
		return err
	}
	// Deserialize `LpMint`:
	err = decoder.Decode(&obj.LpMint)
	if err != nil {
		return err
	}
	// Deserialize `Token0Mint`:
	err = decoder.Decode(&obj.Token0Mint)
	if err != nil {
		return err
	}
	// Deserialize `Token1Mint`:
	err = decoder.Decode(&obj.Token1Mint)
	if err != nil {
		return err
	}
	// Deserialize `Token0Program`:
	err = decoder.Decode(&obj.Token0Program)
	if err != nil {
		return err
	}
	// Deserialize `Token1Program`:
	err = decoder.Decode(&obj.Token1Program)
	if err != nil {
		return err
	}
	// Deserialize `ObservationKey`:
	err = decoder.Decode(&obj.ObservationKey)
	if err != nil {
		return err
	}
	// Deserialize `AuthBump`:
	err = decoder.Decode(&obj.AuthBump)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `LpMintDecimals`:
	err = decoder.Decode(&obj.LpMintDecimals)
	if err != nil {
		return err
	}
	// Deserialize `Mint0Decimals`:
	err = decoder.Decode(&obj.Mint0Decimals)
	if err != nil {
		return err
	}
	// Deserialize `Mint1Decimals`:
	err = decoder.Decode(&obj.Mint1Decimals)
	if err != nil {
		return err
	}
	// Deserialize `LpSupply`:
	err = decoder.Decode(&obj.LpSupply)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeesToken0`:
	err = decoder.Decode(&obj.ProtocolFeesToken0)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeesToken1`:
	err = decoder.Decode(&obj.ProtocolFeesToken1)
	if err != nil {
		return err
	}
	// Deserialize `FundFeesToken0`:
	err = decoder.Decode(&obj.FundFeesToken0)
	if err != nil {
		return err
	}
	// Deserialize `FundFeesToken1`:
	err = decoder.Decode(&obj.FundFeesToken1)
	if err != nil {
		return err
	}
	// Deserialize `OpenTime`:
	err = decoder.Decode(&obj.OpenTime)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	return nil
}
