package main

import "fmt"

func summarizeKubernetesService() (string, int) {
	topic := "Kubernetes Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeKubernetesService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
