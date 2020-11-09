package cmd

import (
	"os"
	"strings"

	zapbox "github.com/dfuse-io/slnc/zap-box"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/dfuse-io/logging"
)

var zlog *zap.Logger

func init() {
	logging.Register("github.com/dfuse-io/cmd/cmd/cmd/commands", &zlog)
}

func SetupLogger(debug bool) {
	if debug {
		zlog, err := zap.NewDevelopment()
		if err == nil {
			logging.Set(zlog)
		}
		// Hijack standard Golang `log` and redirect it to our common logger
		zap.RedirectStdLogAt(zlog, zap.DebugLevel)

	}

	// Fine-grain customization
	//
	// Note that `zapbox.WithLevel` used below does not work in all circumstances! See
	// https://github.com/uber-go/zap/issues/581#issuecomment-600641485 for details.
	if value := os.Getenv("WARN"); value != "" {
		changeLoggersLevel(value, zap.WarnLevel)
	}

	if value := os.Getenv("INFO"); value != "" {
		changeLoggersLevel(value, zap.InfoLevel)
	}

	if value := os.Getenv("DEBUG"); value != "" {
		changeLoggersLevel(value, zap.DebugLevel)
	}
}

func createLogger(serviceName string, verbosity int, logLevel zapcore.Level) *zap.Logger {
	opts := []zap.Option{zap.AddCaller()}
	logStdoutWriter := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(zapbox.NewEncoder(verbosity), logStdoutWriter, logLevel)
	return zap.New(consoleCore, opts...).Named(serviceName)
}

func changeLoggersLevel(inputs string, level zapcore.Level) {
	for _, input := range strings.Split(inputs, ",") {
		logging.Extend(overrideLoggerLevel(level), input)
	}
}

func overrideLoggerLevel(level zapcore.Level) logging.LoggerExtender {
	return func(current *zap.Logger) *zap.Logger {
		return current.WithOptions(zapbox.WithLevel(level))
	}
}
