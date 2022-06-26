package main

import "testing"

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a")
	if got != want {
		t.Errorf("got %q , but wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
