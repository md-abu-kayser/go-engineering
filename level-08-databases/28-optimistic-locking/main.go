// Lesson 28: Optimistic Locking
//
// Goal: Demonstrate the repository boundary with an in-memory implementation
// whose update is explicit, validated, and easy to replace in a real system.
package main

import (
	"errors"
	"fmt"
)

const lesson = "Optimistic Locking"

type account struct {
	ID      string
	Balance int
}

type accountStore struct {
	rows map[string]account
}

func (s *accountStore) Credit(id string, amount int) (account, error) {
	if amount <= 0 {
		return account{}, errors.New("credit must be positive")
	}
	current := s.rows[id]
	current.ID = id
	current.Balance += amount
	s.rows[id] = current
	return current, nil
}

func main() {
	store := accountStore{rows: make(map[string]account)}
	updated, err := store.Credit("acct-7", 25)
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("%s balance: %d\n", updated.ID, updated.Balance)
}