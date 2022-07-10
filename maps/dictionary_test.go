package dictionary

import "fmt"

func ExampleSearch() {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	fmt.Println(got)
	// Output:
	// this is just a test
}

func Search(dictionary map[string]string, s string) (value string) {
	return dictionary[s]
}
