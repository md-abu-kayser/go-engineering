// Lesson 01: Go Memory Model Overview
//
// Goal: Use safe, public runtime APIs to observe a runtime property. The
// program explains a runtime concern without coupling application code to
// undocumented implementation details.
package main

import (
	"fmt"
	"runtime"
)

const lesson = "Go Memory Model Overview"

func main() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("goroutines: %d heap objects: %d\n", runtime.NumGoroutine(), stats.HeapObjects)
}