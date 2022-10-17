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
	"strings"
)

// OPS and FLT are some helpful tokens
var OPS = "+-*/"
var FLT = "1234567890."

// Checks if a char is a parenthesis and changes the count of parentheses accordingly
func isParen(str string, count *int8) bool {
	switch str {
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

// Checks for a number adjacent to a parenthesis in one of two forms: "3(" or ")3"
func checkNumParen(str1 string, str2 string) {
	// If err is nil, then the associated str is a float
	_, err1 := strconv.ParseFloat(str1, 64)
	_, err2 := strconv.ParseFloat(str2, 64)

	// Conditional for checking for specifically the forms "5(" and ")5"
	numParen := err1 == nil && str2 == "("
	parenNum := str1 == ")" && err2 == nil

	if numParen || parenNum {
		log.Fatal("Syntax error: cannot put \"" + str1 + "\" before \"" + str2 + "\"")
	}
}

// Checks for "()" and ")("
func checkOppositeParens(str1 string, str2 string) {
	xParens := str1 == ")" && str2 == "("
	oParens := str1 == "(" && str2 == ")"

	if xParens || oParens {
		log.Fatal("Syntax error: cannot put \"" + str1 + "\" next to \"" + str2 + "\"")
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

		// If the loop is ending and a number is being built,
		// then appends the number
		if i == len(str)-1 && tempNum != "" {
			lexed = append(lexed, tempNum)
		}
	}

	// Various checks for syntax errors
	for i := 0; i < len(lexed)-1; i++ {
		checkDoubleOperator(lexed[i], lexed[i+1])
		checkNumParen(lexed[i], lexed[i+1])
		checkOppositeParens(lexed[i], lexed[i+1])
	}

	return lexed
}

// Parses a lexed expression into a hierarchy of expressions based on parenthesis
func parser(lexed []string) [][][]string {
	parsed := [][][]string{{{}}}
	ec := []int{0} // expression count
	var pc int     // parenthesis count

	// Iterates through the lexed expression,
	// allocating each expression to a list of expressions at different levels
	// and then each level is an appended slice of expressions.
	// The entire equation is level 0, the 1st level is the first layer of parentheses, etc.
	for i := 0; i < len(lexed); i++ {
		if lexed[i] == "(" {
			// Ensures there is enough levels appended in "parsed" and "ec"
			if len(parsed) < pc+2 {
				parsed = append(parsed, [][]string{{}})
				ec = append(ec, 0)
			}

			// Ensures there is enough expressions appended for the current level
			if len(parsed[pc]) < ec[pc]+1 {
				parsed[pc] = append(parsed[pc], []string{})
			}

			// Appends a reference, giving the slice position of the associated expression
			eRef := "expr:" + strconv.Itoa(pc+1) + "." + strconv.Itoa(ec[pc+1])
			parsed[pc][ec[pc]] = append(parsed[pc][ec[pc]], eRef)

			pc++
		} else if lexed[i] == ")" {
			// Ending parenthesis marks the end of an expression,
			// so this increments the expression count accordingly
			ec[pc]++

			pc--
		} else if pc == 0 {
			// The 0th level only ever needs one expression,
			// thus it gets dealt with separately.
			parsed[0][0] = append(parsed[0][0], lexed[i])
		} else {
			// Ensures there is enough expressions appended for the current level
			if len(parsed[pc]) < ec[pc]+1 {
				parsed[pc] = append(parsed[pc], []string{})
			}

			// This is the command to generally add code from the lexer to the parser
			parsed[pc][ec[pc]] = append(parsed[pc][ec[pc]], lexed[i])
		}
	}

	// Checking for hanging operators, i.e. 5+6* or /6*9+5
	for i := 0; i < len(parsed); i++ {
		for j := 0; j < len(parsed[i]); j++ {
			if strings.Contains(OPS, parsed[i][j][0]) && parsed[i][j][0] != "-" {
				log.Fatal("Syntax error: expression " + strconv.Itoa(i) + "." + strconv.Itoa(j) +
					" cannot begin with \"" + parsed[i][j][0] + "\"")
			}
			if strings.Contains(OPS, parsed[i][j][len(parsed[i][j])-1]) {
				log.Fatal("Syntax error: expression " + strconv.Itoa(i) + "." + strconv.Itoa(j) +
					" cannot end with \"" + parsed[i][j][len(parsed[i][j])-1] + "\"")
			}
		}
	}

	return parsed
}

func main() {
	// Gets the expression from the user, runs it through the lexer and the parser,
	// and prints the lexed and parsed expressions if no errors are found
	getInput := p.Input("Enter an equation (without spaces): ")
	lexedExpression := lexer(getInput)
	fmt.Println(lexedExpression)
	parsedExpression := parser(lexedExpression)
	fmt.Println(parsedExpression)
}
