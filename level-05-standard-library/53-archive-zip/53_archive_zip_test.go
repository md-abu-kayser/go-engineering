package main

import "testing"

func TestLessonConcept_020(t *testing.T) {
	got := "Archive Zip"
	if got == "" {
		t.Fatal("lesson topic must not be empty")
	}
}
