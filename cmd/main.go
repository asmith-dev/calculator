/*
	Calculator made in Go.

	Gets an expression from the user, evaluates it, and returns the result.
	Expression will be evaluated using the order of operations.
	Supports the use of negatives and floats.
	Valid operations will be +, -, *, and /.
	Parentheses will be acceptable.
*/

package main

import (
	p "calculator/pkg"
	"fmt"
	"log"
	"strings"
)

// OPS and FLT are some helpful tokens
var OPS = "+-*/"
var FLT = "1234567890."

// Checks if a char is a parenthesis and changes the count of parentheses accordingly
func isParen(char string, count *int8) bool {
	switch char {
	case "(":
		*count++
		return true
	case ")":
		*count--
		return true
	default:
		return false
	}
}

// Checks for consecutive operators and throws an error if applicable
func checkDoubleOperator(str1 string, str2 string) {
	if strings.ContainsAny(OPS, str1) && strings.ContainsAny(OPS, str2) {
		log.Fatal("Syntax error: double operator used \"" + str1 + "\" and \"" + str2 + "\"")
	}
}

// Separates a given expression into tokens
// Error-handling for invalid symbols and invalid parenthesis count
func lexer(str string) []string {
	var lexed []string
	var parenCount int8
	var tempNum string

	// This loop iterates through the expression,
	// evaluates each char, and appends to the lexed slice.
	for i := 0; i < len(str); i++ {
		tempChar := string(str[i])

		// Checks if a number was being made, but is now ended
		// If true, this appends the number
		if !strings.ContainsAny(FLT, tempChar) && tempNum != "" {
			lexed = append(lexed, tempNum)
			tempNum = ""
		}

		// Tokenizes symbols
		// Throws errors for invalid characters
		if isParen(tempChar, &parenCount) || strings.ContainsAny(OPS, tempChar) {
			lexed = append(lexed, tempChar)
		} else if strings.ContainsAny(FLT, tempChar) {
			tempNum += tempChar
		} else {
			log.Fatal("Invalid character entered: " + tempChar)
		}

		// Checks parenthesis count and throws errors for invalid values
		if parenCount < 0 {
			log.Fatal("Syntax error: \")\" before \"(\"")
		} else if parenCount != 0 && i == len(str)-1 {
			log.Fatal("Syntax error: unmatched \"(\"")
		}

		// Various checks for syntax errors
		for i := 0; i < len(lexed)-1; i++ {
			checkDoubleOperator(lexed[i], lexed[i+1])
		}

		// If the loop is ending and a number is being built,
		// then appends the number
		if i == len(str)-1 && tempNum != "" {
			lexed = append(lexed, tempNum)
		}
	}

	return lexed
}

func main() {
	// Gets the expression from the user, runs it through the lexer,
	// and prints the lexed expression if no errors are found
	getInput := p.Input("Enter an equation (without spaces): ")
	lexedExpression := lexer(getInput)
	fmt.Println(lexedExpression)
}
