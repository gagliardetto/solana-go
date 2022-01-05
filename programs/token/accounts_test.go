package token

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestMint(t *testing.T) {
	account := []byte{
		1, 0, 0, 0,
		5, 234, 156, 241, 108, 228, 17, 152, 241, 164, 153, 55, 200, 140, 55, 10, 148, 212, 175, 255, 137, 181, 186, 203, 142, 244, 94, 99, 36, 187, 120, 247,
		9, 169, 49, 235, 241, 182, 6, 0,
		6,
		1,
		1, 0, 0, 0,
		5, 234, 156, 241, 108, 228, 17, 152, 241, 164, 153, 55, 200, 140, 55, 10, 148, 212, 175, 255, 137, 181, 186, 203, 142, 244, 94, 99, 36, 187, 120, 247,
	}
	{
		dec := bin.NewBinDecoder(account)
		mint := Mint{}

		err := dec.Decode(&mint)
		require.NoError(t, err, spew.Sdump(mint))

		require.Equal(t,
			&Mint{
				MintAuthority:   solana.MustPublicKeyFromBase58("Q6XprfkF8RQQKoQVG33xT88H7wi8Uk1B1CC7YAs69Gi").ToPointer(),
				Supply:          1890000009537801,
				Decimals:        6,
				IsInitialized:   true,
				FreezeAuthority: solana.MustPublicKeyFromBase58("Q6XprfkF8RQQKoQVG33xT88H7wi8Uk1B1CC7YAs69Gi").ToPointer(),
			},
			&mint,
		)

		{
			buf := new(bytes.Buffer)
			err := bin.NewBinEncoder(buf).Encode(mint)
			require.NoError(t, err)
			require.Equal(t, account, buf.Bytes())
		}
	}
}
