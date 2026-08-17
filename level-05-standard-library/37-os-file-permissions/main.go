package main

import "fmt"

func summarizeOsFilePermissions() (string, int) {
	topic := "Os File Permissions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOsFilePermissions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
