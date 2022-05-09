package threading

import "fmt"

func GoSafe(fn func()) {
	defer func() {
		if p := recover(); p != nil {
			//todo log
			fmt.Println("GoSafe happens panic:", p)
		}
	}()
	go fn()
}
