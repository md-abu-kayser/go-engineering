package main

import "fmt"

func summarizeGoVersionAndToolchain() (string, int) {
	topic := "Go Version And Toolchain"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoVersionAndToolchain()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
