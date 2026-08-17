package main

import "fmt"

func summarizeImplicitInterfaces() (string, int) {
	topic := "Implicit Interfaces"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeImplicitInterfaces()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
