package main

import "fmt"

func SecureDefaults() string {
	const topic = "Secure Defaults"
	return topic
}

func main() {
	fmt.Println(SecureDefaults())
}
