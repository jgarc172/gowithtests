package main

import "fmt"

func main() {
	fmt.Println(Hello("Jose"))
}

func Hello(name string) string {
	return "Hello, " + name
}
