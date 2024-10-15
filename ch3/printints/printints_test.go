package main

import (
	"testing"
)

func TestAddComma(t *testing.T) {
	t.Run("fraction", func(t *testing.T) {
		sample := "123.456"
		got := FractionPart(sample)
		want := ".456"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
