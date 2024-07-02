package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	hello()
	makeQuote()
	fmt.Println(sum(1, 2))

	// get user input
	var a, b int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter another number: ")
	fmt.Scanln(&b)
	fmt.Println("The sum of", a, "and", b, "is", sum(a, b))

	fmt.Scanln()

}

func hello() {
	fmt.Println("Hello, World!")
}

func sum(a int, b int) int {
	return a + b
}

func makeQuote() {
	fmt.Println(quote.Go())
}
