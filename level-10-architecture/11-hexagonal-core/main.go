package main

import "fmt"

func summarizeHexagonalCore() (string, int) {
	topic := "Hexagonal Core"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHexagonalCore()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
