package serum

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lunixbochs/struc"
	"github.com/stretchr/testify/require"
)

func TestMarketDecode(t *testing.T) {
	s := `c2VydW0DAAAAAAAAAGdTY0MQ+JzjRQitPPQw7a6jFaO3QEtNbJqFwHlHoaDdAAAAAAAAAAAGm4hX/quBhPtof2NGGMA12sQ53BrrO1WYoPAAAAAAAZqhgojuhD2D9j0JH/1UU78OyY17yIzxSctOkEdQqtVn3J9Uzkd6sWTIS0jwubpXWBCWAdTAPAO1WBTXCUKhAiUA8Vb/JRoAAAAAAAAAAAAAe3KroxoWcpqkx0i8u5vTEf6p3L/ywwDq7Ktzy9emFhswTe1qDgAAAMiPz0sAAAAAZAAAAAAAAADe8TmrC1DwGtEuKniQQk1igbzy3FWrl5XOfqClpgrXPn8QBfRYuHutaT3LATJRSLiB32YOH7aWMK3FLYuxLSTFXlqAgmTHudbh5oQuYmowZgURANtyvD/GmOSv6N1nBjDpvaPvZD875suiOGn2gX+hOv6zwukfAaxjJxOl1QIRVQDh9QUAAAAAZAAAAAAAAAAAAAAAAAAAAPdrAgAAAAAAcGFkZGluZw==`
	cnt, err := base64.StdEncoding.DecodeString(s)
	require.NoError(t, err)

	var m MarketV2
	require.NoError(t, struc.Unpack(bytes.NewReader(cnt), &m))

	cnt, _ = json.MarshalIndent(m, "", "  ")
	fmt.Println("MAMA", string(cnt))
}
