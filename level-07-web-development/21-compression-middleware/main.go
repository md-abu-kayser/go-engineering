package main

import "fmt"

func CompressionMiddleware() string {
	const topic = "Compression Middleware"
	return topic
}

func main() {
	fmt.Println(CompressionMiddleware())
}
