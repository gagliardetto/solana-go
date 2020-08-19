package system

import "github.com/dfuse-io/logging"

func init() {
	logging.TestingOverride()
}
