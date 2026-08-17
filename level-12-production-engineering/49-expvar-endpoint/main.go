package main

import "fmt"

func ExpvarEndpoint() string {
	const topic = "Expvar Endpoint"
	return topic
}

func main() {
	fmt.Println(ExpvarEndpoint())
}
