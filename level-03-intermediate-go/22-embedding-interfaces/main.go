// Lesson 22: Embedding Interfaces
//
// Goal: Compose small types so the outer type can reuse focused behavior
// without inheriting hidden state.
package main

import "fmt"

const lesson = "Embedding Interfaces"

type logger struct{}

func (logger) Log(message string) string {
	return "log: " + message
}

type service struct {
	logger
	name string
}

func (s service) Start() string {
	return s.Log(s.name + " started")
}

func main() {
	svc := service{name: "billing"}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(svc.Start())
}