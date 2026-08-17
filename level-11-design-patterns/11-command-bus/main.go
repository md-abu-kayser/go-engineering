package main

import "fmt"

func CommandBus() string {
	const topic = "Command Bus"
	return topic
}

func main() {
	fmt.Println(CommandBus())
}
