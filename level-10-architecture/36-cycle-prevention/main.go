// Lesson 36: Cycle Prevention
//
// Goal: Keep a use case dependent on a consumer-owned port, allowing the
// composition root to choose infrastructure without leaking it inward.
package main

import "fmt"

const lesson = "Cycle Prevention"

type eventPublisher interface {
	Publish(string) error
}

type memoryPublisher struct {
	events []string
}

func (p *memoryPublisher) Publish(event string) error {
	p.events = append(p.events, event)
	return nil
}

func registerUser(publisher eventPublisher, name string) error {
	return publisher.Publish("user.registered:" + name)
}

func main() {
	publisher := &memoryPublisher{}
	if err := registerUser(publisher, "Asha"); err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("published: %v\n", publisher.events)
}