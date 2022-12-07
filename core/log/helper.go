package log

import (
	"github.com/go-kratos/kratos/v2/log"
	ipfsLog "github.com/ipfs/go-log/v2"
)

// Level is a logger level.
type Level int8

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota - 1
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

var (
	helper *log.Helper
)

func SetUp(name string, level Level) {
	helper = log.NewHelper(NewLogger(name))
	// There has to be
	_ = ipfsLog.SetLogLevel(name, level.String()) //nolint:errcheck
}

func Debugf(format string, vals ...interface{}) {
	helper.Debugf(format, vals...)
}

func Infof(format string, vals ...interface{}) {
	helper.Infof(format, vals...)
}

func Errorf(format string, vals ...interface{}) {
	helper.Errorf(format, vals...)
}

func Warnf(format string, vals ...interface{}) {
	helper.Warnf(format, vals...)
}

func IsInit() bool {
	return helper != nil
}
