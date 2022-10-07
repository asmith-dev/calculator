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

// Tokens for the lexer:

var LP = "LPAREN"
var RP = "RPAREN"
var A = "ADD"
var S = "SUB/NEG"
var M = "MULTI"
var D = "DIV"
var N = "NUM:"

// Checks if a char is part of a floating-point number, i.e. in {0,1,2,3,4,5,6,7,8,9,.}
func isFloatPart(str string) bool {
	_, err := strconv.ParseInt(str, 10, 8)

	if err == nil || str == "." {
		return true
	} else {
		return false
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
			lexed = append(lexed, N+tempNum)
			tempNum = ""
		}

		// Tokenizes symbols
		// Throws error for invalid characters
		switch tempChar {
		case "(":
			lexed = append(lexed, LP)
			parenCount++
		case ")":
			lexed = append(lexed, RP)
			parenCount--
		case "+":
			lexed = append(lexed, A)
		case "-":
			lexed = append(lexed, S)
		case "*":
			lexed = append(lexed, M)
		case "/":
			lexed = append(lexed, D)
		default:
			if !isFloatPart(tempChar) {
				log.Fatal("Invalid character entered: " + tempChar)
			}
		}

		// Checks parenthesis count and throws errors for invalid values
		if parenCount < 0 {
			log.Fatal("Syntax error: \")\" before \"(\"")
		} else if parenCount != 0 && i == len(str)-1 {
			log.Fatal("Syntax error: unmatched \"(\"")
		}

		// Builds numbers char by char. Works for integers and floats
		if isFloatPart(tempChar) {
			tempNum += tempChar
		}

		// If the loop is ending and a number is being built,
		// then appends the number
		if i == len(str)-1 && tempNum != "" {
			lexed = append(lexed, N+tempNum)
		}
	}

	return lexed
}

func main() {
	// Demonstrates usage of the input function.
	getInput := p.Input("Enter an equation (without spaces): ")
	lexedExpression := lexer(getInput)
	fmt.Println(lexedExpression)
}
