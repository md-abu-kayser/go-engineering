package main

import "fmt"

func summarizeDddDomainService() (string, int) {
	topic := "Ddd Domain Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDddDomainService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
