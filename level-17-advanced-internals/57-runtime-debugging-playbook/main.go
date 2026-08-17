package main

import "fmt"

func summarizeRuntimeDebuggingPlaybook() (string, int) {
	topic := "Runtime Debugging Playbook"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRuntimeDebuggingPlaybook()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
