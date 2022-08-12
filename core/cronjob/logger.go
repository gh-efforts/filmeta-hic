package cronjob

import (
	"github.com/bitrainforest/filmeta-hic/core/log"
	"github.com/robfig/cron/v3"
)

type logger struct {
}

func newLogger() cron.Logger {
	return logger{}
}

func (l logger) Info(msg string, keysAndValues ...interface{}) {
	log.Infof(msg, keysAndValues...)
}

func (l logger) Error(err error, msg string, keysAndValues ...interface{}) {
	log.Infof(msg+err.Error(), keysAndValues...)
}
