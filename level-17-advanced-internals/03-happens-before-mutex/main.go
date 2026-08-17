package main

import "fmt"

func summarizeHappensBeforeMutex() (string, int) {
	topic := "Happens Before Mutex"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHappensBeforeMutex()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
