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

package serum

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"

	"github.com/klauspost/compress/zstd"
	"github.com/streamingfast/logging"
	"github.com/stretchr/testify/require"
)

func init() {
	logging.TestingOverride()
}

func writeCompressedFile(t *testing.T, filename string, data []byte) {
	writeFile(t, filename, zstEncoder.EncodeAll(data, nil))
}

func readCompressedFile(t *testing.T, file string) []byte {
	data := readFile(t, file)

	out, err := zstDecoder.DecodeAll(data, nil)
	require.NoError(t, err)

	return out
}

func readHexFile(t *testing.T, file string) []byte {
	data := readFile(t, file)

	out, err := hex.DecodeString(string(data))
	require.NoError(t, err)

	return out
}

var zstEncoder, _ = zstd.NewWriter(nil)
var zstDecoder, _ = zstd.NewReader(nil)

func readFile(t *testing.T, file string) []byte {
	data, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	return data
}

func writeFile(t *testing.T, filename string, data []byte) {
	require.NoError(t, ioutil.WriteFile(filename, data, os.ModePerm), "unable to write file %s", filename)
}

func writeJSONFile(t *testing.T, filename string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "  ")
	require.NoError(t, err)

	require.NoError(t, ioutil.WriteFile(filename, out, os.ModePerm), "unable to write JSON file %s", filename)
}

func readJSONFile(t *testing.T, file string, v interface{}) {
	data, err := ioutil.ReadFile(file)
	require.NoError(t, err)

	require.NoError(t, json.Unmarshal(data, v))
	return
}
