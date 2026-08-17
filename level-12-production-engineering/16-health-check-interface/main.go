package main

import "fmt"

func HealthCheckInterface() string {
	const topic = "Health Check Interface"
	return topic
}

func main() {
	fmt.Println(HealthCheckInterface())
}
