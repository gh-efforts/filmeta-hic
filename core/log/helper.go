package log

import "github.com/go-kratos/kratos/v2/log"

var (
	helper *log.Helper
)

func SetUp(name string) {
	helper = log.NewHelper(NewLogger(name))
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
