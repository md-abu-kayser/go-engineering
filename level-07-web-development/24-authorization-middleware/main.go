package main

import "fmt"

func AuthorizationMiddleware() string {
	const topic = "Authorization Middleware"
	return topic
}

func main() {
	fmt.Println(AuthorizationMiddleware())
}
