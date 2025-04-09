package myfmt

import (
	"fmt"
)

var delayCallback func()

func PrintDelay(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	delayCallback()
}

func SetDelayCallback(fn func()) {
	delayCallback = fn
}
