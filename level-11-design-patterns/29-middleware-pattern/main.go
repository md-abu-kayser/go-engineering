package main

import "fmt"

func MiddlewarePattern() string {
	const topic = "Middleware Pattern"
	return topic
}

func main() {
	fmt.Println(MiddlewarePattern())
}
