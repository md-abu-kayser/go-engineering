package main

import "fmt"

func summarizeSecretLoading() (string, int) {
	topic := "Secret Loading"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSecretLoading()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
