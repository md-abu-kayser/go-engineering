package main

import "fmt"

func TransactionContext() string {
	const topic = "Transaction Context"
	return topic
}

func main() {
	fmt.Println(TransactionContext())
}
