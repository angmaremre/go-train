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
	markalar := [...]string{"Adidas", "Flo", "Nike"}   // slice
	markalarArray := []string{"Adidas", "Flo", "Nike"} // array

	fmt.Println("Slice sizes is: ", len(markalar))
	fmt.Println("Array sizes is: ", len(markalarArray))

	for index, marka := range markalar {
		fmt.Println(index, marka)
	}

	myCustomer := customer{adi: "Müşteriad", soyadi: "Müşterisoyad"}
	fmt.Println(myCustomer)
	fmt.Println(myCustomer.adi)
}

type customer struct {
	adi    string
	soyadi string
}

func sayHello(greeting, name string) string {
	return greeting + " " + name
}
