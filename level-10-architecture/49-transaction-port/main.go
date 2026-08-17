package main

import "fmt"

func TransactionPort() string {
	const topic = "Transaction Port"
	return topic
}

func main() {
	fmt.Println(TransactionPort())
}
