package main

import "testing"

func TestSum(t *testing.T) {
	for _, c := range cases() {
		t.Run(c.name, func(t *testing.T) {
			got := Sum(c.nums)
			if got != c.sum {
				t.Errorf("got '%d', but want '%d', given '%v'", got, c.sum, c.nums)
			}
		})
	}
}

type test struct {
	name string
	nums [5]int
	sum  int
}

func cases() []test {
	return []test{
		{"zero", [5]int{0, 0, 0, 0, 0}, 0},
		{"positive", [5]int{1, 0, -1, 2, 3}, 5},
		{"negative", [5]int{1, 0, -1, 2, -3}, -1},
	}
}
