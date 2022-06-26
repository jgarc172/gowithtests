package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: must enter one character as argument")
		os.Exit(1)
	}
	fmt.Printf("%q\n", Repeat(os.Args[1]))
}

// Repeat returns the character c repeated 5 times
func Repeat(c string) (repeated string) {
	for i := 1; i <= 5; i++ {
		repeated += c
	}
	return repeated
}
