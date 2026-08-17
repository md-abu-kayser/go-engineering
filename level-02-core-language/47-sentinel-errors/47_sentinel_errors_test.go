package main

import "testing"

func TestLessonConcept_010(t *testing.T) {
	got := "Sentinel Errors"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
