// Lesson 27: Config Reload Awareness
//
// Goal: Separate deployment configuration from application behavior and make
// readiness a small, testable decision rather than an implicit side effect.
package main

import "fmt"

const lesson = "Config Reload Awareness"

type deployment struct {
	Name     string
	Replicas int
	Database bool
}

func (d deployment) Ready() bool {
	return d.Name != "" && d.Replicas > 0 && d.Database
}

func main() {
	service := deployment{Name: "catalog", Replicas: 3, Database: true}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("%s ready: %t\n", service.Name, service.Ready())
}