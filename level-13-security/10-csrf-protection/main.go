package main

import "fmt"

func CsrfProtection() string {
	const topic = "Csrf Protection"
	return topic
}

func main() {
	fmt.Println(CsrfProtection())
}
