package main

import "fmt"

func summarizePprofLabels() (string, int) {
	topic := "Pprof Labels"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePprofLabels()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
