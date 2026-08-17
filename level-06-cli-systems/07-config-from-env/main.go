package main

import "fmt"

func ConfigFromEnv() string {
	const topic = "Config From Env"
	return topic
}

func main() {
	fmt.Println(ConfigFromEnv())
}
