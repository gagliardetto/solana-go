// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pumpfun

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Global struct {
	Initialized                 bool
	Authority                   ag_solanago.PublicKey
	FeeRecipient                ag_solanago.PublicKey
	InitialVirtualTokenReserves uint64
	InitialVirtualSolReserves   uint64
	InitialRealTokenReserves    uint64
	TokenTotalSupply            uint64
	FeeBasisPoints              uint64
}

var GlobalDiscriminator = [8]byte{167, 232, 232, 177, 200, 108, 114, 127}

func (obj Global) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(GlobalDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized` param:
	err = encoder.Encode(obj.Initialized)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `FeeRecipient` param:
	err = encoder.Encode(obj.FeeRecipient)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualTokenReserves` param:
	err = encoder.Encode(obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualSolReserves` param:
	err = encoder.Encode(obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialRealTokenReserves` param:
	err = encoder.Encode(obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `TokenTotalSupply` param:
	err = encoder.Encode(obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Serialize `FeeBasisPoints` param:
	err = encoder.Encode(obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Global) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(GlobalDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[167 232 232 177 200 108 114 127]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized`:
	err = decoder.Decode(&obj.Initialized)
	if err != nil {
		return err
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `FeeRecipient`:
	err = decoder.Decode(&obj.FeeRecipient)
	if err != nil {
		return err
	}
	// Deserialize `InitialVirtualTokenReserves`:
	err = decoder.Decode(&obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialVirtualSolReserves`:
	err = decoder.Decode(&obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialRealTokenReserves`:
	err = decoder.Decode(&obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `TokenTotalSupply`:
	err = decoder.Decode(&obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Deserialize `FeeBasisPoints`:
	err = decoder.Decode(&obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	return nil
}

type BondingCurve struct {
	VirtualTokenReserves uint64
	VirtualSolReserves   uint64
	RealTokenReserves    uint64
	RealSolReserves      uint64
	TokenTotalSupply     uint64
	Complete             bool
}

var BondingCurveDiscriminator = [8]byte{23, 183, 248, 55, 96, 216, 172, 96}

func (obj BondingCurve) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BondingCurveDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualTokenReserves` param:
	err = encoder.Encode(obj.VirtualTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `VirtualSolReserves` param:
	err = encoder.Encode(obj.VirtualSolReserves)
	if err != nil {
		return err
	}
	// Serialize `RealTokenReserves` param:
	err = encoder.Encode(obj.RealTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `RealSolReserves` param:
	err = encoder.Encode(obj.RealSolReserves)
	if err != nil {
		return err
	}
	// Serialize `TokenTotalSupply` param:
	err = encoder.Encode(obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Serialize `Complete` param:
	err = encoder.Encode(obj.Complete)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BondingCurve) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BondingCurveDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[23 183 248 55 96 216 172 96]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualTokenReserves`:
	err = decoder.Decode(&obj.VirtualTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `VirtualSolReserves`:
	err = decoder.Decode(&obj.VirtualSolReserves)
	if err != nil {
		return err
	}
	// Deserialize `RealTokenReserves`:
	err = decoder.Decode(&obj.RealTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `RealSolReserves`:
	err = decoder.Decode(&obj.RealSolReserves)
	if err != nil {
		return err
	}
	// Deserialize `TokenTotalSupply`:
	err = decoder.Decode(&obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Deserialize `Complete`:
	err = decoder.Decode(&obj.Complete)
	if err != nil {
		return err
	}
	return nil
}

type LastWithdraw struct {
	LastWithdrawTimestamp int64
}

var LastWithdrawDiscriminator = [8]byte{203, 18, 220, 103, 120, 145, 187, 2}

func (obj LastWithdraw) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(LastWithdrawDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `LastWithdrawTimestamp` param:
	err = encoder.Encode(obj.LastWithdrawTimestamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LastWithdraw) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(LastWithdrawDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[203 18 220 103 120 145 187 2]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LastWithdrawTimestamp`:
	err = decoder.Decode(&obj.LastWithdrawTimestamp)
	if err != nil {
		return err
	}
	return nil
}
