// Copyright 2024 github.com/cordialsys
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stake

import (
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_treeout "github.com/gagliardetto/treeout"

	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
)

type Lockup struct {
	UnixTimestamp *int64
	Epoch         *uint64
	Custodian     *ag_solanago.PublicKey
}

func (obj *Lockup) SetUnixTimestamp(unixTimestamp int64) *Lockup {
	obj.UnixTimestamp = &unixTimestamp
	return obj
}

func (obj *Lockup) SetEpoch(epoch uint64) *Lockup {
	obj.Epoch = &epoch
	return obj
}

func (obj *Lockup) SetCustodian(custodian ag_solanago.PublicKey) *Lockup {
	obj.Custodian = &custodian
	return obj
}

func (obj *Lockup) Validate() error {
	if obj.UnixTimestamp == nil {
		return errors.New("UnixTimestamp parameter is not set")
	}
	if obj.Epoch == nil {
		return errors.New("epoch parameter is not set")
	}
	if obj.Custodian == nil {
		return errors.New("custodian parameter is not set")
	}
	return nil
}

func (obj *Lockup) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	if err := encoder.Encode(obj.UnixTimestamp); err != nil {
		return err
	}
	if err := encoder.Encode(obj.Epoch); err != nil {
		return err
	}
	if err := encoder.Encode(obj.Custodian); err != nil {
		return err
	}
	return nil
}

func (obj *Lockup) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	if err := decoder.Decode(&obj.UnixTimestamp); err != nil {
		return err
	}
	if err := decoder.Decode(&obj.Epoch); err != nil {
		return err
	}
	if err := decoder.Decode(&obj.Custodian); err != nil {
		return err
	}
	return nil
}

func (obj *Lockup) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Param("UnixTimestamp", obj.UnixTimestamp))
	parent.Child(ag_format.Param("Epoch", obj.Epoch))
	parent.Child(ag_format.Account("Custodian", *obj.Custodian))
}

func NewLockup() *Lockup {
	return &Lockup{}
}
