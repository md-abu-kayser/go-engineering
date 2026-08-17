package main

import "fmt"

func ContextTimeout() string {
	const topic = "Context Timeout"
	return topic
}

func main() {
	fmt.Println(ContextTimeout())
}
