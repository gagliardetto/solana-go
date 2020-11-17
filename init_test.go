package solana

import (
	"go.uber.org/zap"
)

func init() {
	zlog, _ = zap.NewDevelopment()
}
