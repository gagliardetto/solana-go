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

package rpc

import (
	"errors"
	"net/http"

	"github.com/ybbus/jsonrpc"
)

var ErrNotFound = errors.New("not found")

type Client struct {
	rpcURL    string
	rpcClient CallForClientInterface
	headers   http.Header
}

type CallForClientInterface interface {
	CallFor(out interface{}, method string, params ...interface{}) error
}

func NewClient(rpcURL string) *Client {
	return NewClientWithOpts(rpcURL, nil)
}

func NewClientWithOpts(rpcURL string, opts *jsonrpc.RPCClientOpts) *Client {
	rpcClient := jsonrpc.NewClientWithOpts(rpcURL, opts)
	return &Client{
		rpcURL:    rpcURL,
		rpcClient: rpcClient,
	}
}

func NewWithCustomRPCClient(rpcClient CallForClientInterface) *Client {
	return &Client{
		rpcClient: rpcClient,
	}
}

func (c *Client) SetHeader(k, v string) {
	if c.headers == nil {
		c.headers = http.Header{}
	}
	c.headers.Set(k, v)
}
