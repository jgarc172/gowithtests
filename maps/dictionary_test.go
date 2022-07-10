package dictionary

import (
	"fmt"
	"testing"
)

func ExampleDictionary() {
	dictionary := Dictionary{"test": "this is just a test"}
	got := dictionary.Search("test")
	fmt.Println(got)
	// Output:
	// this is just a test
}

func TestDictionary(t *testing.T) {
	for _, c := range cases() {
		t.Run(c.name, func(t *testing.T) {
			got := c.dict.Search(c.term)
			if got != c.value {
				t.Errorf("got '%v', expected '%v'", got, c.value)
			}
		})
	}
}

type caseDictionary struct {
	name  string
	dict  Dictionary
	term  string
	value string
}

func cases() []caseDictionary {
	return []caseDictionary{
		{"exists", Dictionary{"test": "this is just a test"}, "test", "this is just a test"},
		{"doesn't exist", Dictionary{}, "test", "t"},
	}
}

type Dictionary map[string]string

func (d Dictionary) Search(s string) (value string) {
	return d[s]
}
