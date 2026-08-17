package main

import "fmt"

func AuditLogService() string {
	const topic = "Audit Log Service"
	return topic
}

func main() {
	fmt.Println(AuditLogService())
}
