package main

import "fmt"

func summarizeDependencyRules() (string, int) {
	topic := "Dependency Rules"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDependencyRules()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
