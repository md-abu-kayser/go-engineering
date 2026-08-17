package main

import "fmt"

func summarizeKubernetesProbes() (string, int) {
	topic := "Kubernetes Probes"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeKubernetesProbes()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
