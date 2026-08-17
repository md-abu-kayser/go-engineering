package main

import "fmt"

func summarizeNosqlConsistencyModel() (string, int) {
	topic := "Nosql Consistency Model"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNosqlConsistencyModel()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
