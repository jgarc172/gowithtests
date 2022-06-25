package main

import "testing"

func TestHello(t *testing.T) {
	for _, test := range cases() {
		t.Run(test.name, func(t *testing.T) {
			got := Hello(test.input1, test.input2)
			if got != test.want {
				t.Errorf("got %q  want %q", got, test.want)
			}
		})
	}
}

type test struct {
	name   string
	input1 string
	input2 string
	want   string
}

func cases() []test {
	return []test{
		{"empty name", "", "", "Hello, world"},
		{"empty lang", "Jose", "", "Hello, Jose"},
		{"Spanish", "Jose", "Spanish", "Hola, Jose"},
		{"French", "Jose", "French", "Bonjour, Jose"},
	}
}
