package diff

import (
	"fmt"
	"reflect"

	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var traceEnabled = logging.IsTraceEnabled("solana-go", "github.com/dfuse-io/solana-go/diff")
var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/dfuse-io/solana-go/diff", &zlog)
}

type reflectType struct {
	in interface{}
}

func (r reflectType) String() string {
	if r.in == nil {
		return "<nil>"
	}

	valueOf := reflect.ValueOf(r.in)
	return fmt.Sprintf("%s (zero? %t, value %s)", valueOf.Type(), valueOf.IsZero(), r.in)
}
