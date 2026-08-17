package main

import "fmt"

func summarizeCodeGenerationDesign() (string, int) {
	topic := "Code Generation Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCodeGenerationDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
