package main

import "fmt"

func summarizeInterfaceMock() (string, int) {
	topic := "Interface Mock"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInterfaceMock()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
