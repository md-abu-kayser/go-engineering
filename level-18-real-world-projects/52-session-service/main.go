package main

import "fmt"

func SessionService() string {
	const topic = "Session Service"
	return topic
}

func main() {
	fmt.Println(SessionService())
}
