package main

import "fmt"

func UnixDomainSockets() string {
	const topic = "Unix Domain Sockets"
	return topic
}

func main() {
	fmt.Println(UnixDomainSockets())
}
