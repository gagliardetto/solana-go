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
package tokenregistry

type Logo [32]byte
type Name [32]byte
type Symbol [12]byte

type TokenMeta struct {
	Logo   Logo
	Name   Name
	Symbol Symbol
}

func (l Logo) String() string {
	return AsciiString(l[:])
}
func (n Name) String() string {
	return AsciiString(n[:])
}

func (s Symbol) String() string {
	return AsciiString(s[:])
}

func AsciiString(data []byte) string {
	var trimmed []byte
	for _, b := range data {
		if b > 0 {
			trimmed = append(trimmed, b)
		}
	}
	return string(trimmed)
}
