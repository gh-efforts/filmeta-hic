package threading

import (
	"fmt"

	"github.com/bitrainforest/filmeta-hic/core/log"
)

func GoSafe(fn func()) {
	defer func() {
		if p := recover(); p != nil {
			if log.IsInit() {
				log.Warnf("[GoSafe] err: %v", p)
			}
			fmt.Println("GoSafe happens panic:", p)
		}
	}()
	go fn()
}
