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

package cmd

import (
	"github.com/spf13/viper"

	"github.com/dfuse-io/solana-go/api/graphql"
	"github.com/spf13/cobra"
)

var apiGraphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "start serving graphql api",
	RunE: func(cmd *cobra.Command, args []string) error {
		SetupLogger(viper.GetBool("global-debug"))

		server := graphql.NewServer("/Users/cbillett/devel/dfuse/go/solana-go/api/graphql/schema.graphql", ":9000")
		zlog.Info("serving ...")
		return server.Launch()

	},
}

func init() {
	apiCmd.AddCommand(apiGraphqlCmd)
}
