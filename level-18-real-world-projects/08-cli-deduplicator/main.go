// Lesson 08: Cli Deduplicator
//
// Goal: Assemble a tiny vertical slice: validate a request, apply a domain
// rule, persist it through a focused store, and return the resulting value.
package main

import (
	"fmt"
	"strings"
)

const lesson = "Cli Deduplicator"

type item struct {
	ID   string
	Name string
}

type store struct {
	items map[string]item
}

func (s *store) Create(id, name string) (item, error) {
	name = strings.TrimSpace(name)
	if id == "" || name == "" {
		return item{}, fmt.Errorf("id and name are required")
	}
	created := item{ID: id, Name: name}
	s.items[id] = created
	return created, nil
}

func main() {
	data := store{items: make(map[string]item)}
	created, err := data.Create("item-1", "first item")
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("created: %#v\n", created)
}