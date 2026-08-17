package main

import "fmt"

func summarizeUpsertSql() (string, int) {
	topic := "Upsert Sql"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUpsertSql()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
