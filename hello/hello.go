package main

import "fmt"

func main() {
	fmt.Println(Hello("Jose", "Spanish"))
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
