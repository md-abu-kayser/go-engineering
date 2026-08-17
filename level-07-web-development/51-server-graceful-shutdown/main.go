package main

import "fmt"

func ServerGracefulShutdown() string {
	const topic = "Server Graceful Shutdown"
	return topic
}

func main() {
	fmt.Println(ServerGracefulShutdown())
}
