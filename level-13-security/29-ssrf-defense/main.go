package main

import "fmt"

func summarizeSsrfDefense() (string, int) {
	topic := "Ssrf Defense"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSsrfDefense()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
