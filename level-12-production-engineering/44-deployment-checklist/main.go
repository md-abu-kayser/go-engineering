package main

import "fmt"

func summarizeDeploymentChecklist() (string, int) {
	topic := "Deployment Checklist"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDeploymentChecklist()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
