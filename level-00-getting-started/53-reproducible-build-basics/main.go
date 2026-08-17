package main

import "fmt"

func summarizeReproducibleBuildBasics() (string, int) {
	topic := "Reproducible Build Basics"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReproducibleBuildBasics()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
