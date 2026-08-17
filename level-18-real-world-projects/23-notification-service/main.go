package main

import "fmt"

func summarizeNotificationService() (string, int) {
	topic := "Notification Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNotificationService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
