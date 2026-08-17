package main

import "fmt"

func summarizeBlueGreenDeployment() (string, int) {
	topic := "Blue Green Deployment"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBlueGreenDeployment()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
