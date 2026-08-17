package main

import "testing"

func TestLessonConcept_001(t *testing.T) {
	got := "Install And Verify Go"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
