package main

import "fmt"

func summarizeWorkspaceFiles() (string, int) {
	topic := "Workspace Files"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeWorkspaceFiles()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
