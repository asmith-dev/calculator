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
	"strconv"
)

// Checks if a char is part of a floating-point number, i.e. in {0,1,2,3,4,5,6,7,8,9,.}
func isFloatPart(str string) bool {
	_, err := strconv.ParseInt(str, 10, 8)

	if err == nil || str == "." {
		return true
	} else {
		return false
	}
}

// Checks if the string is an operator
func isOperator(str string) bool {
	ops := [4]string{"+", "-", "*", "/"}

	// Checks the input string against each operator
	for i := 0; i < 4; i++ {
		if ops[i] == str {
			return true
		}
	}

	// If no operator is found, then the string is not an operator
	return false
}

// Checks for consecutive operators and throws an error if applicable
// For the implementation used here, str2 is assumed to be an operator
func checkDoubleOperator(str1 string, str2 string) {
	if isOperator(str1) {
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
		// If true, this appends a token for the number
		if !isFloatPart(tempChar) && tempNum != "" {
			lexed = append(lexed, tempNum)
			tempNum = ""
		}

		// Tokenizes symbols
		// Throws errors for invalid characters and double operators
		switch {
		case tempChar == "(":
			lexed = append(lexed, tempChar)
			parenCount++
		case tempChar == ")":
			lexed = append(lexed, tempChar)
			parenCount--
		case isOperator(tempChar):
			lexed = append(lexed, tempChar)
			if i != 0 {
				checkDoubleOperator(lexed[len(lexed)-2], lexed[len(lexed)-1])
			}
		case isFloatPart(tempChar):
			tempNum += tempChar
		default:
			log.Fatal("Invalid character entered: " + tempChar)
		}

		// Checks parenthesis count and throws errors for invalid values
		if parenCount < 0 {
			log.Fatal("Syntax error: \")\" before \"(\"")
		} else if parenCount != 0 && i == len(str)-1 {
			log.Fatal("Syntax error: unmatched \"(\"")
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
