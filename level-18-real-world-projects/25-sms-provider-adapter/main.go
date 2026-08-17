package main

import "fmt"

func SmsProviderAdapter() string {
	const topic = "Sms Provider Adapter"
	return topic
}

func main() {
	fmt.Println(SmsProviderAdapter())
}
