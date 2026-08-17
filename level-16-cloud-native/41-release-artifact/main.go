package main

import "fmt"

func summarizeReleaseArtifact() (string, int) {
	topic := "Release Artifact"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReleaseArtifact()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
