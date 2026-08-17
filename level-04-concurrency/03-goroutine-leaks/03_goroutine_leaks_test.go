package main

import "testing"

func TestLessonConcept_014(t *testing.T) {
	got := "Goroutine Leaks"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
