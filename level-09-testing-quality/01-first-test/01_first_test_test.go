package main

import "testing"

func TestLessonConcept_030(t *testing.T) {
	got := "First Test"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
