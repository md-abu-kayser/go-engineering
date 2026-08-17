package main

import "fmt"

func summarizeCryptoRand() (string, int) {
	topic := "Crypto Rand"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCryptoRand()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
