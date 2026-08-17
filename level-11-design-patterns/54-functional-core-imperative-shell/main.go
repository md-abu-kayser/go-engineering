package main

import "fmt"

func summarizeFunctionalCoreImperativeShell() (string, int) {
	topic := "Functional Core Imperative Shell"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFunctionalCoreImperativeShell()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
