package main

import "testing"

func TestLessonConcept_002(t *testing.T) {
	got := "Semantic Import Versioning"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
