package main

import "fmt"

func summarizeOrderingGuarantees() (string, int) {
	topic := "Ordering Guarantees"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOrderingGuarantees()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
