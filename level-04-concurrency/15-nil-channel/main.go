package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Nil Channel"
	value, ok := <-ch
	fmt.Printf("%q ok=%t\n", value, ok)
}
