package serum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	bin "github.com/dfuse-io/binary"

	"github.com/stretchr/testify/require"
)

func TestDecoder_Slabs(t *testing.T) {

	//zlog, _ := zap.NewDevelopment()
	//EnableDebugLogging(zlog)

	rawSlabs := []string{
		"0100000035000000010babffffffffff4105000000000000400000003f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"0200000014060000b2cea5ffffffffff23070000000000005ae01b52d00a090c6dc6fce8e37a225815cff2223a99c6dfdad5aae56d3db670e62c000000000000140b0fadcf8fcebf",
		"030000003400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	}

	for _, s := range rawSlabs {
		cnt, err := hex.DecodeString(s)
		require.NoError(t, err)

		decoder := bin.NewDecoder(cnt)
		var slab *Slab
		err = decoder.Decode(&slab)
		require.NoError(t, err)

		json, err := json.MarshalIndent(slab, "", "   ")
		require.NoError(t, err)
		fmt.Println(string(json))

		//require.Equal(t, 0, decoder.remaining())

	}
}
