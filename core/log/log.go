package log

import (
	"fmt"

	kLog "github.com/go-kratos/kratos/v2/log"
	"github.com/ipfs/go-log/v2"
)

var _ kLog.Logger = (*IpfsLog)(nil)

type IpfsLog struct {
	log *log.ZapEventLogger
}

func NewLogger(name string) kLog.Logger {
	return &IpfsLog{log: log.Logger(name)}
}

func (l IpfsLog) Log(level kLog.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	var (
		msg  string
		data []interface{}
	)

	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == kLog.DefaultMessageKey {
			msg = fmt.Sprint(keyvals[i+1])
			continue
		}
		data = append(data, fmt.Sprint(keyvals[i], keyvals[i+1]))
	}

	switch level {
	case kLog.LevelDebug:
		l.log.Debugf(msg, data...)
	case kLog.LevelInfo:
		l.log.Infof(msg, data...)
	case kLog.LevelWarn:
		l.log.Warnf(msg, data...)
	case kLog.LevelError:
		l.log.Errorf(msg, data...)
	case kLog.LevelFatal:
		l.log.Fatalf(msg, data...)
	}
	return nil
}
