package threading

import (
	"github.com/bitrainforest/filmeta-hic/core/log"
	"runtime/debug"
)

func GoSafe(fn func()) {
	go safe(fn)
}
func safe(fn func()) {
	defer func() {
		if p := recover(); p != nil {
			if log.IsInit() {
				log.Errorf("[GoSafe] err: %s\n%s", p, string(debug.Stack()))
			}
		}
	}()
	fn()
}
