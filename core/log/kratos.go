package log

import (
	k "github.com/go-kratos/kratos/v2"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"os"
)

func LoggerKratosOption(serverName, version string) k.Option {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.name", serverName,
		"service.version", version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	return k.Logger(logger)
}
