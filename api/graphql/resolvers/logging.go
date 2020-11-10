package resolvers

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/dfuse/solana-go/api/graphql/resolvers", &zlog)
}
