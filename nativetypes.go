// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solana

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

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
type ByteWrapper struct {
	io.Reader
}

func (w *ByteWrapper) ReadByte() (byte, error) {
	var b [1]byte
	_, err := w.Read(b[:])
	return b[0], err
}
