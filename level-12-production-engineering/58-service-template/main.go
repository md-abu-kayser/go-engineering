package main

import "fmt"

func ServiceTemplate() string {
	const topic = "Service Template"
	return topic
}

func main() {
	fmt.Println(ServiceTemplate())
}
