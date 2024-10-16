package main

import "testing"

func TestAnagram(t *testing.T) {
	s1 := "sample"
	s2 := "sAmple"

	got := Anagram(s1, s2)
	want := false
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
