package integers

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	for _, c := range casesSum() {
		t.Run(c.name, func(t *testing.T) {
			got := Sum(c.nums)
			if got != c.sum {
				t.Errorf("got '%d', but want '%d', given '%v'", got, c.sum, c.nums)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	for _, c := range casesSumSlice() {
		t.Run(c.name, func(t *testing.T) {
			got := SumSlice(c.nums)
			if got != c.sum {
				t.Errorf("got '%d', but want '%d', given '%v'", got, c.sum, c.nums)
			}
		})
	}
}

func TestSumVar(t *testing.T) {
	for _, c := range casesSumVar() {
		t.Run(c.name, func(t *testing.T) {
			got := SumVar(c.nums...)
			if !reflect.DeepEqual(got, c.sums) {
				t.Errorf("got '%d', but want '%d', given '%v'", got, c.sums, c.nums)
			}
		})
	}
}

func TestSumTail(t *testing.T) {
	for _, c := range casesSumTail() {
		t.Run(c.name, func(t *testing.T) {
			got := SumTail(c.nums...)
			if !reflect.DeepEqual(got, c.sums) {
				t.Errorf("got '%d', but want '%d', given '%v'", got, c.sums, c.nums)
			}
		})
	}
}

type testSum struct {
	name string
	nums [5]int
	sum  int
}

func casesSum() []testSum {
	return []testSum{
		{"zero", [5]int{0, 0, 0, 0, 0}, 0},
		{"positive", [5]int{1, 0, -1, 2, 3}, 5},
		{"negative", [5]int{1, 0, -1, 2, -3}, -1},
	}
}

type testSumSlice struct {
	name string
	nums []int
	sum  int
}

func casesSumSlice() []testSumSlice {
	return []testSumSlice{
		{"empty", []int{}, 0},
		{"positive", []int{1, 0, -1, 2, 3}, 5},
		{"negative", []int{1, 0, -1, 2, -3}, -1},
	}
}

type testSumVar struct {
	name string
	sums []int
	nums [][]int
}

func casesSumVar() []testSumVar {
	return []testSumVar{
		{"none", nil, nil},
		{"one empty", []int{5, 0}, [][]int{
			{1, 0, -1, 2, 3},
			{},
		}},
		{"none empty", []int{1, 5}, [][]int{
			{1},
			{1, 0, -1, 2, 3},
		}},
	}
}

type testSumTail struct {
	name string
	sums []int
	nums [][]int
}

func casesSumTail() []testSumTail {
	return []testSumTail{
		{"none", nil, nil},
		{"one empty", []int{4, 0}, [][]int{
			{1, 0, -1, 2, 3},
			{},
		}},
		{"only one", []int{0, 4}, [][]int{
			{2},
			{1, 0, -1, 2, 3},
		}},
		{"more than one", []int{3, 5}, [][]int{
			{1, 3},
			{1, 1, -1, 2, 3},
		}},
	}
}
