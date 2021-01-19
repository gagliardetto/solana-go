package storage

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mr-tron/base58"

	"github.com/dfuse-io/solana-go"

	bin "github.com/dfuse-io/binary"
	"github.com/stretchr/testify/require"
)

func Test_LoadAccount(t *testing.T) {

	file, err := os.Open("./test_data/59665756.902451748") // For read access.
	require.NoError(t, err)

	data, err := ioutil.ReadAll(file)
	require.NoError(t, err)

	decoder := bin.NewDecoder(data)

	for {
		var s *StoredMeta
		err := decoder.Decode(&s)
		require.NoError(t, err)
		fmt.Println("Version:", s.Version)
		fmt.Println("Pub key:", s.PubKey)
		fmt.Println("Owner:", s.Owner)
		fmt.Println("Executable", s.Executable)
		fmt.Println("Data Length", s.DataLength)
		//fmt.Println("Data", hex.EncodeToString(s.Data))

		rem := int(s.DataLength) % 8
		rem = 8 - rem

		for i := 0; i < rem; i++ {
			_, err := decoder.ReadByte()
			require.NoError(t, err)
		}

		fmt.Println("remaining:", decoder.Remaining())
		if decoder.Remaining() == 128 {
			fmt.Println(hex.EncodeToString(data[len(data)-decoder.Remaining():]))
		}
		if !decoder.HasRemaining() {
			break
		}
		//err = text.NewEncoder(os.Stdout).Encode(err, nil)
		//require.NoError(t, err)
	}
	fmt.Println("all good")

}

func TestFoo(t *testing.T) {
	k := solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")
	fmt.Println(hex.EncodeToString(k[:]))
}
func TestFoo2(t *testing.T) {
	data, err := hex.DecodeString(strings.ReplaceAll("c8 49 cb fe e9 8c 36 2c 71 7f e8 ad 19 8d 43 e2 6d b4 96 5c 28 ab 60 f1 ae b9 59 7b cd 3e fc f4", " ", ""))
	require.NoError(t, err)

	fmt.Println(base58.Encode(data))

}
