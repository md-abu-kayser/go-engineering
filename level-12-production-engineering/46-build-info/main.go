package main

import "fmt"

func BuildInfo() string {
	const topic = "Build Info"
	return topic
}

func main() {
	fmt.Println(BuildInfo())
}
