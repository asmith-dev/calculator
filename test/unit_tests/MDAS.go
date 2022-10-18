/*
	Demonstrates calculating a simple expression, like 2+2, with calc().

	Then, in main(), a method is given for updating an algebraic expression
	while also enforcing the order of operations.
*/

package main

import (
	"fmt"
	"strconv"
)

// OPS must have the chars in order of MDAS
var OPS = "*/+-"

// Calculates the result of a [3]string array
// where the first/last indexes are floats and the middle is an operator
func calc(expr []string) string {
	var1, _ := strconv.ParseFloat(expr[0], 64)
	var2, _ := strconv.ParseFloat(expr[2], 64)

	// Checks which operator is used and returns the calculated value
	switch expr[1] {
	case "*":
		return strconv.FormatFloat(var1*var2, 'f', -1, 64)
	case "/":
		return strconv.FormatFloat(var1/var2, 'f', -1, 64)
	case "+":
		return strconv.FormatFloat(var1+var2, 'f', -1, 64)
	case "-":
		return strconv.FormatFloat(var1-var2, 'f', -1, 64)
	}

	// Returns an error string, since this code should never be reached
	return "ERR"
}

func main() {
	// Example expression
	expr := []string{"5.2", "*", "6", "-", "88", "/", "1"}

	// The current operator of MDAS, which indexes OPS
	currentOP := 0

	var result string

	// Prints the initial expression
	fmt.Println(expr)

	// Continues until the result remains, i.e. len(expr)=1
	for len(expr) > 1 {
		for i := 0; i < len(expr); i++ {
			// Checks until it finds the current operator of MDAS
			if expr[i] == string(OPS[currentOP]) {
				result = calc(expr[i-1 : i+2])

				// Updates the expression based on the location of the operator
				switch i {
				case 1:
					expr = append([]string{result}, expr[3:]...)
				case len(expr) - 2:
					expr = append(expr[:len(expr)-3], result)
				default:
					expr = append(append(expr[:i-1], result), expr[i+2:]...)
				}

				// Prints an updated expression and breaks the loop
				fmt.Println(expr)
				break
			}

			// If the loop is ending and still hasn't found the current operator,
			// then it assumes that that operator is no longer present and moves to the next one.
			if i == len(expr)-1 {
				currentOP++
			}
		}
	}
}
