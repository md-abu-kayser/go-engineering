package main

import "fmt"

func summarizeBlankImports() (string, int) {
	topic := "Blank Imports"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBlankImports()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
