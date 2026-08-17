package main

import "fmt"

func summarizeConnMaxLifetime() (string, int) {
	topic := "Conn Max Lifetime"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConnMaxLifetime()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
