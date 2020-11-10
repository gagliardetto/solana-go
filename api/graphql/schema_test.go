package graphql

import (
	"io/ioutil"
	"testing"

	"github.com/dfuse-io/solana-go/api/graphql/resolvers"
	"github.com/graph-gophers/graphql-go"
	"github.com/stretchr/testify/require"
)

func TestSchema(t *testing.T) {

	// initialize GraphQL
	cnt, err := ioutil.ReadFile("schema.graphql")
	require.NoError(t, err)

	rootResolver := resolvers.NewRoot()

	_, err = graphql.ParseSchema(
		string(cnt),
		rootResolver,
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	)
	require.NoError(t, err)
}
