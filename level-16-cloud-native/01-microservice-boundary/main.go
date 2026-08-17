package main

import "fmt"

func MicroserviceBoundary() string {
	const topic = "Microservice Boundary"
	return topic
}

func main() {
	fmt.Println(MicroserviceBoundary())
}
