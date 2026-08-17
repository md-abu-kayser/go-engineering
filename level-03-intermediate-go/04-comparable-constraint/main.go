package main

import "fmt"

func summarizeComparableConstraint() (string, int) {
	topic := "Comparable Constraint"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeComparableConstraint()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
