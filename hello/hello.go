package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: provide two arguments")
		os.Exit(1)
	}
	fmt.Println(Hello(os.Args[1], os.Args[2]))
}

func Hello(name string, lang string) string {
	greeting := prefix(lang)
	if name == "" {
		name = "world"
	}
	return greeting + ", " + name
}

func prefix(lang string) string {
	switch lang {
	case "Spanish":
		return "Hola"
	case "French":
		return "Bonjour"
	default: // English
		return "Hello"
	}
}
