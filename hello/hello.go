package main

import "fmt"

func main() {
	fmt.Println(Hello("Jose", "Spanish"))
}

func Hello(name string, lang string) string {
	greeting := "Hello"
	if name == "" {
		name = "world"
	}
	if lang == "Spanish" {
		greeting = "Hola"
	}
	return greeting + ", " + name
}
