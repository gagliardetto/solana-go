// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
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

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsciiBytes_String(t *testing.T) {
	require.Equal(t, "ABC", AsciiString([]byte{65, 66, 67, 0, 0, 0, 0, 0}))
}

func TestLogoFromString(t *testing.T) {
	l, err := LogoFromString("logo")
	require.NoError(t, err)

	require.Equal(t, "logo", l.String())
	require.Equal(t, Logo([64]byte{108, 111, 103, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), l)
}

func TestNameFromString(t *testing.T) {
	l, err := NameFromString("name")
	require.NoError(t, err)

	require.Equal(t, "name", l.String())
	require.Equal(t, Name([32]byte{110, 97, 109, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), l)
}

func TestSymbolFromString(t *testing.T) {
	l, err := SymbolFromString("symb")
	require.NoError(t, err)

	require.Equal(t, "symb", l.String())
	require.Equal(t, Symbol([32]byte{115, 121, 109, 98, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), l)
}

func TestWebsiteFromString(t *testing.T) {
	l, err := WebsiteFromString("webs")
	require.NoError(t, err)

	require.Equal(t, "webs", l.String())
	require.Equal(t, Website([32]byte{119, 101, 98, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), l)
}
