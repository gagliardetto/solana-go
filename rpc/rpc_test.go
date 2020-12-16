package rpc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockJSONRPCServer struct {
	*httptest.Server
	body []byte
}

func mockJSONRPC(t *testing.T, response interface{}) (mock *mockJSONRPCServer, close func()) {
	mock = &mockJSONRPCServer{
		Server: httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			var err error
			mock.body, err = ioutil.ReadAll(req.Body)
			require.NoError(t, err)

			var responseBody []byte
			if v, ok := response.(json.RawMessage); ok {
				responseBody = v
			} else {
				responseBody, err = json.Marshal(response)
				require.NoError(t, err)
			}

			rw.Write(responseBody)
		})),
	}

	return mock, func() { mock.Close() }
}

func (s *mockJSONRPCServer) RequestBodyAsJSON(t *testing.T) (out string) {
	return string(s.body)
}

func (s *mockJSONRPCServer) RequestBody(t *testing.T) (out map[string]interface{}) {
	err := json.Unmarshal(s.body, &out)
	require.NoError(t, err)

	return out
}
