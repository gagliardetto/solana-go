package solana

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

func init() {
	zlog, _ = zap.NewDevelopment()
	logging.TestingOverride()
}
