package assert

import (
	"fmt"

	"github.com/pkg/errors"
)

func CheckErr(err error, args ...interface{}) {
	if err == nil {
		return
	}
	e := errors.Wrapf(err, fmt.Sprint(args...))
	panic(e)
}
