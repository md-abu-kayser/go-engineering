package main

import "fmt"

func summarizeEditorWorkflow() (string, int) {
	topic := "Editor Workflow"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEditorWorkflow()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
