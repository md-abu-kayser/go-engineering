package main

import "fmt"

func SecretRedaction() string {
	const topic = "Secret Redaction"
	return topic
}

func main() {
	fmt.Println(SecretRedaction())
}
