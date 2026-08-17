package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		done <- "Scheduler Goroutine Model"
	}()
	wg.Wait()
	fmt.Println(<-done)
}
