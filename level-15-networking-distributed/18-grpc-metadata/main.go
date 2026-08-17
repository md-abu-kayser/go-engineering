package main

import "fmt"

func GrpcMetadata() string {
	const topic = "Grpc Metadata"
	return topic
}

func main() {
	fmt.Println(GrpcMetadata())
}
