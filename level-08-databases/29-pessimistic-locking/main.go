package main

import "fmt"

func summarizePessimisticLocking() (string, int) {
	topic := "Pessimistic Locking"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePessimisticLocking()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
