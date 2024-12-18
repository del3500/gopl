package main

import (
	"testing"
)

func TestAddComma(t *testing.T) {
	sample := "123.456"

	t.Run("fraction", func(t *testing.T) {
		got, err := FractionPart(sample)
		if err != nil {
			t.Error(err)
		}
		want := ".456"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("integer", func(t *testing.T) {
		got := IntegerPart(sample)
		want := "123"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("add comma", func(t *testing.T) {
		str := "93232182136.7345435375"
		intpart := IntegerPart(str)
		fracpart, err := FractionPart(str)
		if err != nil {
			t.Error(err)
		}
		got := AddComma(intpart, fracpart)
		want := "93,232,182,136.7345435375"
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
