// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type PoolAccount struct {
	Owner            ag_solanago.PublicKey
	Vault            ag_solanago.PublicKey
	Mint             ag_solanago.PublicKey
	AuthorityBump    uint8
	IsActive         bool
	AmpInitialFactor uint16
	AmpTargetFactor  uint16
	RampStartTs      int64
	RampStopTs       int64
	SwapFee          uint64
	Tokens           []PoolToken
	PendingOwner     *ag_solanago.PublicKey `bin:"optional"`
}

var PoolAccountDiscriminator = [8]byte{116, 210, 187, 119, 196, 196, 52, 137}

func (obj PoolAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PoolAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Vault` param:
	err = encoder.Encode(obj.Vault)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `AuthorityBump` param:
	err = encoder.Encode(obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Serialize `IsActive` param:
	err = encoder.Encode(obj.IsActive)
	if err != nil {
		return err
	}
	// Serialize `AmpInitialFactor` param:
	err = encoder.Encode(obj.AmpInitialFactor)
	if err != nil {
		return err
	}
	// Serialize `AmpTargetFactor` param:
	err = encoder.Encode(obj.AmpTargetFactor)
	if err != nil {
		return err
	}
	// Serialize `RampStartTs` param:
	err = encoder.Encode(obj.RampStartTs)
	if err != nil {
		return err
	}
	// Serialize `RampStopTs` param:
	err = encoder.Encode(obj.RampStopTs)
	if err != nil {
		return err
	}
	// Serialize `SwapFee` param:
	err = encoder.Encode(obj.SwapFee)
	if err != nil {
		return err
	}
	// Serialize `Tokens` param:
	err = encoder.Encode(obj.Tokens)
	if err != nil {
		return err
	}
	// Serialize `PendingOwner` param (optional):
	{
		if obj.PendingOwner == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PendingOwner)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *PoolAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PoolAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[116 210 187 119 196 196 52 137]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Vault`:
	err = decoder.Decode(&obj.Vault)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityBump`:
	err = decoder.Decode(&obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Deserialize `IsActive`:
	err = decoder.Decode(&obj.IsActive)
	if err != nil {
		return err
	}
	// Deserialize `AmpInitialFactor`:
	err = decoder.Decode(&obj.AmpInitialFactor)
	if err != nil {
		return err
	}
	// Deserialize `AmpTargetFactor`:
	err = decoder.Decode(&obj.AmpTargetFactor)
	if err != nil {
		return err
	}
	// Deserialize `RampStartTs`:
	err = decoder.Decode(&obj.RampStartTs)
	if err != nil {
		return err
	}
	// Deserialize `RampStopTs`:
	err = decoder.Decode(&obj.RampStopTs)
	if err != nil {
		return err
	}
	// Deserialize `SwapFee`:
	err = decoder.Decode(&obj.SwapFee)
	if err != nil {
		return err
	}
	// Deserialize `Tokens`:
	err = decoder.Decode(&obj.Tokens)
	if err != nil {
		return err
	}
	// Deserialize `PendingOwner` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PendingOwner)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type StrategyAccount struct {
	Pool            ag_solanago.PublicKey
	IsActive        bool
	AmpMinFactor    uint16
	AmpMaxFactor    uint16
	RampMinStep     uint16
	RampMaxStep     uint16
	RampMinDuration uint32
	RampMaxDuration uint32
}

var StrategyAccountDiscriminator = [8]byte{97, 218, 254, 239, 248, 146, 59, 42}

func (obj StrategyAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(StrategyAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool` param:
	err = encoder.Encode(obj.Pool)
	if err != nil {
		return err
	}
	// Serialize `IsActive` param:
	err = encoder.Encode(obj.IsActive)
	if err != nil {
		return err
	}
	// Serialize `AmpMinFactor` param:
	err = encoder.Encode(obj.AmpMinFactor)
	if err != nil {
		return err
	}
	// Serialize `AmpMaxFactor` param:
	err = encoder.Encode(obj.AmpMaxFactor)
	if err != nil {
		return err
	}
	// Serialize `RampMinStep` param:
	err = encoder.Encode(obj.RampMinStep)
	if err != nil {
		return err
	}
	// Serialize `RampMaxStep` param:
	err = encoder.Encode(obj.RampMaxStep)
	if err != nil {
		return err
	}
	// Serialize `RampMinDuration` param:
	err = encoder.Encode(obj.RampMinDuration)
	if err != nil {
		return err
	}
	// Serialize `RampMaxDuration` param:
	err = encoder.Encode(obj.RampMaxDuration)
	if err != nil {
		return err
	}
	return nil
}

func (obj *StrategyAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(StrategyAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[97 218 254 239 248 146 59 42]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	err = decoder.Decode(&obj.Pool)
	if err != nil {
		return err
	}
	// Deserialize `IsActive`:
	err = decoder.Decode(&obj.IsActive)
	if err != nil {
		return err
	}
	// Deserialize `AmpMinFactor`:
	err = decoder.Decode(&obj.AmpMinFactor)
	if err != nil {
		return err
	}
	// Deserialize `AmpMaxFactor`:
	err = decoder.Decode(&obj.AmpMaxFactor)
	if err != nil {
		return err
	}
	// Deserialize `RampMinStep`:
	err = decoder.Decode(&obj.RampMinStep)
	if err != nil {
		return err
	}
	// Deserialize `RampMaxStep`:
	err = decoder.Decode(&obj.RampMaxStep)
	if err != nil {
		return err
	}
	// Deserialize `RampMinDuration`:
	err = decoder.Decode(&obj.RampMinDuration)
	if err != nil {
		return err
	}
	// Deserialize `RampMaxDuration`:
	err = decoder.Decode(&obj.RampMaxDuration)
	if err != nil {
		return err
	}
	return nil
}

type VaultAccount struct {
	Admin ag_solanago.PublicKey

	// PDAofpoolprogramsseededbyvaultaddress
	WithdrawAuthority ag_solanago.PublicKey

	// bumpseedofwithdraw_authorityPDA
	WithdrawAuthorityBump uint8

	// bumpseedofvault_authorityPDA
	AuthorityBump  uint8
	IsActive       bool
	Beneficiary    ag_solanago.PublicKey
	BeneficiaryFee uint64
	PendingAdmin   *ag_solanago.PublicKey `bin:"optional"`
}

var VaultAccountDiscriminator = [8]byte{230, 251, 241, 83, 139, 202, 93, 28}

func (obj VaultAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(VaultAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `WithdrawAuthority` param:
	err = encoder.Encode(obj.WithdrawAuthority)
	if err != nil {
		return err
	}
	// Serialize `WithdrawAuthorityBump` param:
	err = encoder.Encode(obj.WithdrawAuthorityBump)
	if err != nil {
		return err
	}
	// Serialize `AuthorityBump` param:
	err = encoder.Encode(obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Serialize `IsActive` param:
	err = encoder.Encode(obj.IsActive)
	if err != nil {
		return err
	}
	// Serialize `Beneficiary` param:
	err = encoder.Encode(obj.Beneficiary)
	if err != nil {
		return err
	}
	// Serialize `BeneficiaryFee` param:
	err = encoder.Encode(obj.BeneficiaryFee)
	if err != nil {
		return err
	}
	// Serialize `PendingAdmin` param (optional):
	{
		if obj.PendingAdmin == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PendingAdmin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *VaultAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(VaultAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[230 251 241 83 139 202 93 28]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawAuthority`:
	err = decoder.Decode(&obj.WithdrawAuthority)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawAuthorityBump`:
	err = decoder.Decode(&obj.WithdrawAuthorityBump)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityBump`:
	err = decoder.Decode(&obj.AuthorityBump)
	if err != nil {
		return err
	}
	// Deserialize `IsActive`:
	err = decoder.Decode(&obj.IsActive)
	if err != nil {
		return err
	}
	// Deserialize `Beneficiary`:
	err = decoder.Decode(&obj.Beneficiary)
	if err != nil {
		return err
	}
	// Deserialize `BeneficiaryFee`:
	err = decoder.Decode(&obj.BeneficiaryFee)
	if err != nil {
		return err
	}
	// Deserialize `PendingAdmin` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PendingAdmin)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
