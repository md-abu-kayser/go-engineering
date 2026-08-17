package main

import "fmt"

func HealthEndpoint() string {
	const topic = "Health Endpoint"
	return topic
}

func main() {
	fmt.Println(HealthEndpoint())
}
