package dictionary

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleDictionary() {
	dictionary := Dictionary{"test": "this is just a test"}
	got, _ := dictionary.Search("test")
	fmt.Println(got)
	// Output:
	// this is just a test
}

func TestDictionary(t *testing.T) {
	for _, c := range casesSuccess() {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.dict.Search(c.term)
			if got != c.value || err != nil {
				t.Errorf("got '%v', expected '%v', '%v'", got, c.value, err)
			}
		})
	}
	for _, c := range casesError() {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.dict.Search(c.term)
			if got != c.value || err == nil {
				t.Errorf("got '%v', expected '%v', '%v'", got, c.value, err)
			}
		})
	}
}

type caseDictionary struct {
	name  string
	dict  Dictionary
	term  string
	value string
	err   error
}

func casesSuccess() []caseDictionary {
	return []caseDictionary{
		{"found",
			Dictionary{"test1": "this is just a test", "test2": "another test"},
			"test1",
			"this is just a test",
			nil,
		},
	}
}

func casesError() []caseDictionary {
	return []caseDictionary{
		{"empty",
			Dictionary{},
			"test",
			"",
			errors.New("could not find the word you were looking for"),
		},
		{"not found",
			Dictionary{"term": "exists here"},
			"test",
			"",
			errors.New("could not find the word you were looking for"),
		},
	}
}

type Dictionary map[string]string

func (d Dictionary) Search(s string) (value string, err error) {
	value, ok := d[s]
	if !ok {
		err = errors.New("could not find the word you were looking for")
	}
	return
}
