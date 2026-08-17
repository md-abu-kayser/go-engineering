package main

import "fmt"

func StartupOrder() string {
	const topic = "Startup Order"
	return topic
}

func main() {
	fmt.Println(StartupOrder())
}
