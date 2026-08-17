package main

import "fmt"

func summarizePasswordResetToken() (string, int) {
	topic := "Password Reset Token"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePasswordResetToken()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
