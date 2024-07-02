package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	hello()
	makeQuote()
	fmt.Println(sum(1, 2))

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
