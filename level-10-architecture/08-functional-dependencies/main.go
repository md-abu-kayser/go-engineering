package main

import "fmt"

func summarizeFunctionalDependencies() (string, int) {
	topic := "Functional Dependencies"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFunctionalDependencies()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
