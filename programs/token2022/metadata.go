package token2022

import (
	"bytes"
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

type TokenInstruction byte
type MetadataPointerInstruction byte

const (
    MetadataPointerExtension TokenInstruction = 39
    Initialize               MetadataPointerInstruction = 0
	Update				   MetadataPointerInstruction = 1
)


type TokenMetadata struct {
    // The authority that can sign to update the metadata
    UpdateAuthority *solana.PublicKey;
    // The associated mint, used to counter spoofing to be sure that metadata belongs to a particular mint
    Mint solana.PublicKey;
    // The longer name of the token
    Name string;
    // The shortened symbol for the token
    Symbol string;
    // The URI pointing to richer metadata
    Uri string;
    // Any additional metadata about the token as key-value pairs
    AdditionalMetadata map[string]string;
}

func fixCodecSize(_ func([]byte) []byte, size int) func([]byte) []byte {
    return func(b []byte) []byte {
        if len(b) > size {
            return b[:size]
        }
        padded := make([]byte, size)
        copy(padded, b)
        return padded
    }
}

func getBytesCodec() func([]byte) []byte {
    return func(b []byte) []byte {
        return b
    }
}

func getStringCodec() func(string) []byte {
    return func(s string) []byte {
        return []byte(s)
    }
}


func getMapCodec(codec func(string) []byte) func(map[string]string) []byte {
    return func(m map[string]string) []byte {
        var buf bytes.Buffer
        for k, v := range m {
            buf.Write(codec(k))
            buf.Write(codec(v))
        }
        return buf.Bytes()
    }
}

// Convert the metadata to a byte array
func (meta *TokenMetadata) Pack() []byte {
    // If no updateAuthority given, set it to the None/Zero PublicKey for encoding
     updateAuthority,_ := solana.PublicKeyFromBase58("11111111111111111111111111111111")
	if meta.UpdateAuthority != nil {
		updateAuthority = *meta.UpdateAuthority
	}
	var buf bytes.Buffer
	buf.Write(fixCodecSize(getBytesCodec(), 32)(updateAuthority.Bytes()))
	buf.Write(fixCodecSize(getBytesCodec(), 32)(meta.Mint.Bytes()))
	buf.Write(getStringCodec()(meta.Name))
	buf.Write(getStringCodec()(meta.Symbol))
	buf.Write(getStringCodec()(meta.Uri))
	buf.Write(getMapCodec(getStringCodec())(meta.AdditionalMetadata))
	return buf.Bytes()
}


// Use this in conjuntion with GetMinimumBalanceForRentExemption to calculate the lamports needed to create the account
func (meta *TokenMetadata) LenForLamports() uint64 {
	// Size of MetadataExtension 2 bytes for type, 2 bytes for length
	 metadataExtension := 2 + 2;
	 // Size of Mint Account with extension
	 mintLen :=  234
	return uint64(metadataExtension)+ uint64(len(meta.Pack())) + uint64(mintLen)
}



func CreateInitializeMetadataPointerInstruction(
    mint solana.PublicKey,
    authority *solana.PublicKey,
    metadataAddress *solana.PublicKey,
) solana.Instruction {
	programID := ProgramID

    pointerData := initializeMetadataPointerData{
        Instruction:              MetadataPointerExtension,
        MetadataPointerInstruction: Initialize,
        Authority:                *authority,
        MetadataAddress:          *metadataAddress,
    }

	ix := &createInitializeMetadataPointerInstruction{
		programID: programID,
		accounts: []*solana.AccountMeta{
			solana.Meta(mint).WRITE(),
		},
		data: pointerData.encode(),
	}

	return ix
}

type createInitializeMetadataPointerInstruction struct {
	programID solana.PublicKey
	accounts []*solana.AccountMeta
	data []byte
}

func (inst *createInitializeMetadataPointerInstruction) ProgramID() solana.PublicKey {
	return inst.programID
}

func (inst *createInitializeMetadataPointerInstruction) Accounts() (out []*solana.AccountMeta) {
	return inst.accounts
}

func (inst *createInitializeMetadataPointerInstruction) Data() ([]byte, error) {
	return inst.data, nil
}

type initializeMetadataPointerData struct {
    Instruction              TokenInstruction
    MetadataPointerInstruction MetadataPointerInstruction
    Authority                solana.PublicKey
    MetadataAddress          solana.PublicKey
}

func (data *initializeMetadataPointerData) encode() []byte {
    var buf bytes.Buffer
    binary.Write(&buf, binary.LittleEndian, data.Instruction)
    binary.Write(&buf, binary.LittleEndian, data.MetadataPointerInstruction)
    buf.Write(data.Authority.Bytes())
    buf.Write(data.MetadataAddress.Bytes())
    return buf.Bytes()
}

func CreateUpdateMetadataPointerInstruction(
    mint solana.PublicKey,
    authority solana.PublicKey,
    metadataAddress *solana.PublicKey,
    multiSigners []interface{},
) solana.Instruction {
	programID := ProgramID

    keys := addSigners([]*solana.AccountMeta{
        solana.Meta(mint).WRITE(),
    }, authority, multiSigners)

	pointerData := updateMetadataPointerData{
        Instruction:              MetadataPointerExtension,
        MetadataPointerInstruction: Update,
        MetadataAddress:          *metadataAddress,
    }

	ix := &createUpdateMetadataPointerInstruction{
		programID: programID,
		accounts: keys,
		data: pointerData.encode(),
	}
	return ix
}

func addSigners(keys []*solana.AccountMeta, authority solana.PublicKey, multiSigners []interface{}) []*solana.AccountMeta {
    keys = append(keys, solana.Meta(authority).SIGNER())
    for _, signer := range multiSigners {
        switch v := signer.(type) {
        case solana.PublicKey:
            keys = append(keys, solana.Meta(v).SIGNER())
        case solana.PrivateKey:
            keys = append(keys, solana.Meta(v.PublicKey()).SIGNER())
        }
    }
    return keys
}

type updateMetadataPointerData struct {
    Instruction              TokenInstruction
    MetadataPointerInstruction MetadataPointerInstruction
    MetadataAddress          solana.PublicKey
}

func (data *updateMetadataPointerData) encode() []byte {
    var buf bytes.Buffer
    buf.WriteByte(byte(data.Instruction))
    buf.WriteByte(byte(data.MetadataPointerInstruction))
    buf.Write(data.MetadataAddress.Bytes())
    return buf.Bytes()
}

type createUpdateMetadataPointerInstruction struct {
	programID solana.PublicKey
	accounts []*solana.AccountMeta
	data []byte
}

func (inst *createUpdateMetadataPointerInstruction) ProgramID() solana.PublicKey {
	return inst.programID
}

func (inst *createUpdateMetadataPointerInstruction) Accounts() (out []*solana.AccountMeta) {
	return inst.accounts
}

func (inst *createUpdateMetadataPointerInstruction) Data() ([]byte, error) {
	return inst.data, nil
}