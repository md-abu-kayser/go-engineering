package main

import "fmt"

func summarizeFilePermissions() (string, int) {
	topic := "File Permissions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFilePermissions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
