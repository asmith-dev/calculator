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
)

// Separates a given expression into tokens and categorizes each.
// Need to implement categorizing and error handling for invalid symbols
func lexer(str string) map[string]string {
	lexed := make(map[string]string) // placeholder for the return statement

	// This for loop will iterate through the expression,
	// evaluate each char, and append to the map
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}

	return lexed
}

func main() {
	// Demonstrates usage of the input function.
	getInput := input("Enter an equation (without spaces): ")
	lexer(getInput)
}
