package main

import "fmt"

func OutboxTable() string {
	const topic = "Outbox Table"
	return topic
}

func main() {
	fmt.Println(OutboxTable())
}
