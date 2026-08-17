package main

import "fmt"

func summarizeSyscallBoundary() (string, int) {
	topic := "Syscall Boundary"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSyscallBoundary()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
