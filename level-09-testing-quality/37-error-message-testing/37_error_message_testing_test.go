package main

import "testing"

func TestLessonConcept_032(t *testing.T) {
	got := "Error Message Testing"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
