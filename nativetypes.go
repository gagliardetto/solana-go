package solana

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"

	"github.com/lunixbochs/struc"
	"github.com/mr-tron/base58"
)

type Padding []byte

type Hash PublicKey

///

type Signature [64]byte

func SignatureFromBase58(in string) (out Signature, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return
	}

	if len(val) != 64 {
		err = fmt.Errorf("invalid length, expected 64, got %d", len(val))
		return
	}
	copy(out[:], val)
	return
}

func (p Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}
func (p *Signature) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	dat, err := base58.Decode(s)
	if err != nil {
		return err
	}

	if len(dat) != 64 {
		return errors.New("invalid data length for public key")
	}

	target := Signature{}
	copy(target[:], dat)
	*p = target
	return
}

func (p Signature) String() string {
	return base58.Encode(p[:])
}

///

type PublicKey [32]byte

func MustPublicKeyFromBase58(in string) PublicKey {
	out, _ := PublicKeyFromBase58(in)
	return out
}

func PublicKeyFromBase58(in string) (out PublicKey, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return
	}

	if len(val) != 32 {
		err = fmt.Errorf("invalid length, expected 32, got %d", len(val))
		return
	}
	copy(out[:], val)
	return
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}
func (p *PublicKey) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	dat, err := base58.Decode(s)
	if err != nil {
		return err
	}

	if len(dat) != 32 {
		return errors.New("invalid data length for public key")
	}

	target := PublicKey{}
	copy(target[:], dat)
	*p = target
	return
}

func (p PublicKey) String() string {
	return base58.Encode(p[:])
}

///

type Base58 []byte

func (t Base58) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(t))
}

func (t *Base58) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	*t, err = base58.Decode(s)
	return
}

func (t Base58) String() string {
	return base58.Encode(t)
}

///

type Data []byte

func (t Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"data":     []byte(t),
		"encoding": "base64",
	})
}

func (t *Data) UnmarshalJSON(data []byte) (err error) {
	var in []string
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}

	if len(in) != 2 {
		return fmt.Errorf("invalid length for solana.Data, expected 2, found %d", len(in))
	}

	if in[1] == "base64" {
		*t, err = base64.StdEncoding.DecodeString(in[0])
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unsupported encoding %s", in[1])
}

func (t Data) String() string {
	return base64.StdEncoding.EncodeToString(t)
}

///
type U128 big.Int

func (u U128) BigInt() *big.Int {
	v := big.Int(u)
	return &v
}

func (u U128) Pack(p []byte, opt *struc.Options) (int, error) {
	panic("implement me")
}

func (u *U128) Unpack(r io.Reader, length int, opt *struc.Options) error {
	buf := make([]byte, 16)
	reader := &ByteWrapper{r}
	for i := 0; i < 16; i++ {
		b, err := reader.ReadByte()
		if err != nil {
			return err
		}
		buf[16-i-1] = b
	}
	big := new(big.Int).SetBytes(buf)
	*u = U128(*big)
	return nil
}

func (u U128) Size(opt *struc.Options) int {
	return 16
}

func (u U128) String() string {
	v := big.Int(u)
	return v.String()
}

type U64 uint64

func (i U64) MarshalJSON() (data []byte, err error) {
	if i > 0xffffffff {
		encodedInt, err := json.Marshal(uint64(i))
		if err != nil {
			return nil, err
		}
		data = append([]byte{'"'}, encodedInt...)
		data = append(data, '"')
		return data, nil
	}
	return json.Marshal(uint64(i))
}

func (i *U64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty value")
	}

	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}

		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}

		*i = U64(val)

		return nil
	}

	var v uint64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = U64(v)

	return nil
}

///

type ByteWrapper struct {
	io.Reader
}

func (w *ByteWrapper) ReadByte() (byte, error) {
	var b [1]byte
	_, err := w.Read(b[:])
	return b[0], err
}

/// Varuint16
type Varuint16 uint16

func (v Varuint16) Pack(buf []byte, opt *struc.Options) (int, error) {
	x := uint64(v)
	i := 0
	for x >= 0x80 {
		buf[i] = byte(x) | 0x80
		x >>= 7
		i++
	}
	buf[i] = byte(x)
	return i + 1, nil
	// JAVASCRIPT
	// let rem_len = len;
	// for (;;) {
	//   let elem = rem_len & 0x7f;
	//   rem_len >>= 7;
	//   if (rem_len == 0) {
	//     bytes.push(elem);
	//     break;
	//   } else {
	//     elem |= 0x80;
	//     bytes.push(elem);
	//   }
	// }

	// RUST
	//     // Pass a non-zero value to serialize_tuple() so that serde_json will
	// // generate an open bracket.
	// let mut seq = serializer.serialize_tuple(1)?;

	// let mut rem_len = self.0;
	// loop {
	//     let mut elem = (rem_len & 0x7f) as u8;
	//     rem_len >>= 7;
	//     if rem_len == 0 {
	//         seq.serialize_element(&elem)?;
	//         break;
	//     } else {
	//         elem |= 0x80;
	//         seq.serialize_element(&elem)?;
	//     }
	// }
	// seq.end()
}

func (v *Varuint16) Unpack(r io.Reader, length int, opt *struc.Options) error {
	res, err := readVaruint16(&ByteWrapper{r})
	if err != nil {
		return err
	}
	*v = Varuint16(res)
	return nil
}
func (v *Varuint16) Size(opt *struc.Options) int {
	// TODO: fix the `Size`, which doesn't reflect.. need to Pack, and return the size here?
	var buf [8]byte
	return binary.PutUvarint(buf[:], uint64(*v))
}
func (v *Varuint16) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

var shortVecOverflow = errors.New("short_vec: varint overflows a 16-bit integer")

func readVaruint16(r io.ByteReader) (uint64, error) {
	// This was identified https://groups.google.com/g/golang-announce/c/NyPIaucMgXo/m/GdsyQP6QAAAJ?pli=1
	// after I copied it here.. and I think we're using the EXACT same construct.. and in our case,
	// we don't want to read more than 3 bytes.
	// FIXME!!!!
	var x uint64
	var s uint
	for i := 0; ; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return x, err
		}
		if b < 0x80 {
			if i > 4 || i == 4 && b > 1 {
				return x, shortVecOverflow
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}
