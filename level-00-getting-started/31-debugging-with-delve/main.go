package main

import "fmt"

func DebuggingWithDelve() string {
	const topic = "Debugging With Delve"
	return topic
}

func main() {
	fmt.Println(DebuggingWithDelve())
}
