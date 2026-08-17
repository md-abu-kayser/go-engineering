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
		done <- "Parallel Reduce"
	}()
	wg.Wait()
	fmt.Println(<-done)
}
