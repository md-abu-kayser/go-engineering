package main

import "fmt"

func RecoveryMiddleware() string {
	const topic = "Recovery Middleware"
	return topic
}

func main() {
	fmt.Println(RecoveryMiddleware())
}
