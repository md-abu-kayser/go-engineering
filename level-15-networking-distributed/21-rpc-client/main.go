package main

import "fmt"

func RpcClient() string {
	const topic = "Rpc Client"
	return topic
}

func main() {
	fmt.Println(RpcClient())
}
