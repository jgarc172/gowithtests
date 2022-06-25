package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		want := "Hello, Jose"
		got := Hello("Jose", "")
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})
	t.Run("empty name", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})
	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola, Jose"
		got := Hello("Jose", "Spanish")
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})
	t.Run("in French", func(t *testing.T) {
		want := "Bonjour, Jose"
		got := Hello("Jose", "French")
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})
}
