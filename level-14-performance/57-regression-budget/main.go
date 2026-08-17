package main

import "fmt"

func summarizeRegressionBudget() (string, int) {
	topic := "Regression Budget"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRegressionBudget()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
