package main

import "fmt"

func summarizeBcryptPasswords() (string, int) {
	topic := "Bcrypt Passwords"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBcryptPasswords()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
