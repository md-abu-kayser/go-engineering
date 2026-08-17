package main

import "fmt"

func summarizeBuildConstraints() (string, int) {
	topic := "Build Constraints"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBuildConstraints()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
