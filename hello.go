package main

import "fmt"

var name string = "Emre"
var surname string = "Korkmaz"

func main() {
	// var greeting string = "Emre"
	name := "Emre"
	greeting := "Hello"
	fmt.Println(sayHello(greeting, name))

	// array slice
	markalar := []string{"Adidas", "Flo"}

	for index, marka := range markalar {
		fmt.Println(index, marka)
	}
}

func sayHello(greeting, name string) string {
	return greeting + " " + name
}
