package main

import (
	"concurrency_practice/patterns"
	"time"
)

func main() {

	go patterns.TestPatterns()
	time.Sleep(100 * time.Millisecond)
}
