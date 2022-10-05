/*
Calculator made in Go.

Gets an expression from the user, evaluates it, and returns the result.
Expression will be evaluated using the order of operations.
Valid data type will be float64.
Valid operations will be +, -, *, and /.
Parentheses will be acceptable.
*/

package main

import (
	"fmt"
	"log"
)

// Python-like implementation for getting user input in a condensed format.
// Might only call this function once, in which case I will remove it later.
func input(str string) string {
	var response string
	fmt.Print(str)

	_, err := fmt.Scanln(&response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func main() {
	// Demonstrates usage of the input function.
	getInput := input("Enter an equation (without spaces): ")
	fmt.Println(getInput) // temporary usage of the equation
}
