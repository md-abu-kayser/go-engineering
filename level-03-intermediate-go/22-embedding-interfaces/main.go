package main

import "fmt"

func summarizeEmbeddingInterfaces() (string, int) {
	topic := "Embedding Interfaces"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEmbeddingInterfaces()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
