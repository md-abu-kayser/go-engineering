package main

import "fmt"

func summarizeImports() (string, int) {
	topic := "Imports"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeImports()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
