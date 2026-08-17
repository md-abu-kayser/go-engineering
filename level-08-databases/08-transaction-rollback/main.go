package main

import "fmt"

func summarizeTransactionRollback() (string, int) {
	topic := "Transaction Rollback"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTransactionRollback()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
