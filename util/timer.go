package util

import (
	"fmt"
	"time"
)

// From: https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s completed in %v\n", name, time.Since(start))
	}
}
