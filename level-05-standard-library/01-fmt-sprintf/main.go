package main

import "fmt"

func summarizeFmtSprintf() (string, int) {
	topic := "Fmt Sprintf"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFmtSprintf()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
