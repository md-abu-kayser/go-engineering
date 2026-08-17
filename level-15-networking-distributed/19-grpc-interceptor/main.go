package main

import "fmt"

func summarizeGrpcInterceptor() (string, int) {
	topic := "Grpc Interceptor"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGrpcInterceptor()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
