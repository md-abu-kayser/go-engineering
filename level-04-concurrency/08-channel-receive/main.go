package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Channel Receive"
	value, ok := <-ch
	fmt.Printf("%q ok=%t\n", value, ok)
}
