package main

import "fmt"

func summarizeDatabaseSqlPing() (string, int) {
	topic := "Database Sql Ping"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDatabaseSqlPing()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
