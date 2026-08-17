package main

import "fmt"

func summarizeCliPasswordGenerator() (string, int) {
	topic := "Cli Password Generator"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCliPasswordGenerator()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
