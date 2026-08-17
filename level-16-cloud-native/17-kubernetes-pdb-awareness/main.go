package main

import "fmt"

func summarizeKubernetesPdbAwareness() (string, int) {
	topic := "Kubernetes Pdb Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeKubernetesPdbAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
