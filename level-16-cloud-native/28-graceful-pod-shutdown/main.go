package main

import "fmt"

func GracefulPodShutdown() string {
	const topic = "Graceful Pod Shutdown"
	return topic
}

func main() {
	fmt.Println(GracefulPodShutdown())
}
