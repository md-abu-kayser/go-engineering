package main

import "fmt"

func CanaryDeployment() string {
	const topic = "Canary Deployment"
	return topic
}

func main() {
	fmt.Println(CanaryDeployment())
}
