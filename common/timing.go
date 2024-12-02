package common

import (
	"fmt"
	"time"
)

func Time(name string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%s: %v\n", name, time.Since(start))
}
