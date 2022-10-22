/*
	Calculator made in Go.

	Gets an expression from the user, evaluates it, and prints the result.
	Expressions are evaluated using the order of operations,
	except minuses are interpreted as addition of negatives.
	Supports the use of negatives, floats, and parentheses, but not exponents.
	Valid operations will be +, -, *, and /.
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
var OPS = "*/+-"
var FLT = "1234567890."

type Calculator struct {
	input    string
	inputEnd int
	lexed    []string
	pc       int   // parenthesis count
	ec       []int // expression count per level
	tree     [][][]string
}

func newCalculator(inp string) Calculator {
	return Calculator{
		input:    inp,
		inputEnd: len(inp) - 1,
		ec:       []int{0},
		tree:     [][][]string{{{}}},
	}
}

// Appends a string to the lexer
func (c *Calculator) appendLexed(str string) {
	c.lexed = append(c.lexed, str)
}

// Gets the expression count at a certain level
func (c *Calculator) exprCount() int {
	return c.ec[c.pc]
}

// Ensures there is enough levels in the tree and the expression counter
func (c *Calculator) isEnoughLevels() {
	if len(c.tree) < c.pc+2 {
		c.tree = append(c.tree, [][]string{{}})
		c.ec = append(c.ec, 0)
	}
}

// Appends to the current expression either a reference or a value from the lexer
func (c *Calculator) appendCurrentExpr(i int, ref bool) {
	var entry string

	// Ensures there is enough expressions appended for the current level
	if len(c.tree[c.pc]) < c.exprCount()+1 {
		c.tree[c.pc] = append(c.tree[c.pc], []string{})
	}

	if ref {
		// Entry is a reference, giving the slice position of the associated expression
		entry = "expr:" + strconv.Itoa(c.pc+1) + "." + strconv.Itoa(c.ec[c.pc+1])
	} else {
		// Entry is the current entry in the lexer
		entry = c.lexed[i]
	}

	// Appends to the current expression
	c.tree[c.pc][c.exprCount()] = append(c.tree[c.pc][c.exprCount()], entry)
}

// Checks if a char is a parenthesis and changes the count of parentheses accordingly
func (c *Calculator) isParen(str string) bool {
	switch str {
	case "(":
		c.pc++
		return true
	case ")":
		c.pc--
		return true
	default:
		return false
	}
}

// Checks for consecutive operators and throws an error if applicable
func (c *Calculator) checkDoubleOperator(i int) {
	if strings.ContainsAny(OPS, c.lexed[i]) && strings.ContainsAny(OPS, c.lexed[i+1]) {
		log.Fatal("Syntax error: double operator used \"" + c.lexed[i] + "\" and \"" + c.lexed[i+1] + "\"")
	}
}

// Checks for a number adjacent to a parenthesis in one of two forms: "3(" or ")3"
func (c *Calculator) checkNumParen(i int) {
	// If err is nil, then the associated str is a float
	_, err1 := strconv.ParseFloat(c.lexed[i], 64)
	_, err2 := strconv.ParseFloat(c.lexed[i+1], 64)

	// Conditional for checking for specifically the forms "5(" and ")5"
	numParen := err1 == nil && c.lexed[i+1] == "("
	parenNum := c.lexed[i] == ")" && err2 == nil

	if numParen || parenNum {
		log.Fatal("Syntax error: cannot put \"" + c.lexed[i] + "\" before \"" + c.lexed[i+1] + "\"")
	}
}

// Checks for "()" and ")("
func (c *Calculator) checkOppositeParens(i int) {
	xParens := c.lexed[i] == ")" && c.lexed[i+1] == "("
	oParens := c.lexed[i] == "(" && c.lexed[i+1] == ")"

	if xParens || oParens {
		log.Fatal("Syntax error: cannot put \"" + c.lexed[i] + "\" next to \"" + c.lexed[i+1] + "\"")
	}
}

// Checks for hanging operators, i.e. 6+4/ or *9-3
func (c *Calculator) checkHangingOperator() {
	for i := 0; i < len(c.tree); i++ {
		for j := 0; j < len(c.tree[i]); j++ {
			// Checks for non-minus beginning operators
			if strings.Contains(OPS, c.tree[i][j][0]) && c.tree[i][j][0] != "-" {
				log.Fatal("Syntax error: expression " + strconv.Itoa(i) + "." + strconv.Itoa(j) +
					" cannot begin with \"" + c.tree[i][j][0] + "\"")
			}

			// Checks for ending operators
			if strings.Contains(OPS, c.tree[i][j][len(c.tree[i][j])-1]) {
				log.Fatal("Syntax error: expression " + strconv.Itoa(i) + "." + strconv.Itoa(j) +
					" cannot end with \"" + c.tree[i][j][len(c.tree[i][j])-1] + "\"")
			}
		}
	}
}

// Separates a given expression into numbers, operators, and parentheses
// Error-handling for invalid symbols and invalid parenthesis count
func (c *Calculator) lexer() {
	c.pc = 0
	var tempNum string

	// This loop iterates through the expression,
	// evaluates each char, and appends to the lexed slice.
	for i := 0; i < len(c.input); i++ {
		tempChar := string(c.input[i])

		// Checks if a number was being made, but is now ended
		// If true, this appends the number
		if !strings.ContainsAny(FLT, tempChar) && tempNum != "" {
			c.appendLexed(tempNum)
			tempNum = ""
		}

		// Parses symbols and builds numbers
		// Throws errors for invalid characters
		if c.isParen(tempChar) || strings.ContainsAny(OPS, tempChar) {
			c.appendLexed(tempChar)
		} else if strings.ContainsAny(FLT, tempChar) {
			tempNum += tempChar
		} else {
			log.Fatal("Invalid character entered: " + tempChar)
		}

		// Checks parenthesis count and throws errors for invalid values
		if c.pc < 0 {
			log.Fatal("Syntax error: \")\" before \"(\"")
		} else if c.pc != 0 && i == c.inputEnd {
			log.Fatal("Syntax error: unmatched \"(\"")
		}

		// If the loop is ending and a number is being built,
		// then appends the number
		if i == c.inputEnd && tempNum != "" {
			c.appendLexed(tempNum)
		}
	}

	// Various checks for syntax errors
	for i := 0; i < len(c.lexed)-1; i++ {
		c.checkDoubleOperator(i)
		c.checkNumParen(i)
		c.checkOppositeParens(i)
	}
}

// Organizes the parsed user-entered expression into a tree of expressions based on parenthesis
func (c *Calculator) parser() {
	c.pc = 0

	// Iterates through the lexed expression,
	// allocating each expression to a list of expressions at different levels of the tree
	// and then each level is an appended slice of expressions.
	// The entire equation is level 0, the 1st level is the first layer of parentheses, etc.
	for i := 0; i < len(c.lexed); i++ {
		if c.lexed[i] == "(" {
			// Ensures there is enough levels appended in "tree" and "ec"
			c.isEnoughLevels()

			// Appends a reference
			c.appendCurrentExpr(i, true)

			c.pc++
		} else if c.lexed[i] == ")" {
			// Ending parenthesis marks the end of an expression,
			// so this increments the expression count accordingly
			c.ec[c.pc]++

			c.pc--
		} else if c.pc == 0 {
			// The 0th level only ever needs one expression,
			// thus it gets dealt with separately.
			c.tree[0][0] = append(c.tree[0][0], c.lexed[i])
		} else {
			// This is the command to generally add code from the lexer to the tree
			c.appendCurrentExpr(i, false)
		}
	}

	// Checking for hanging operators, i.e. 5+6* or /6*9+5
	c.checkHangingOperator()
}

// Calculates the result of a [3]string
// where the first/last indexes are floats and the middle is an operator
func calcSimple(expr []string) string {
	var1, err1 := strconv.ParseFloat(expr[0], 64)
	p.HandleError(err1)
	var2, err2 := strconv.ParseFloat(expr[2], 64)
	p.HandleError(err2)

	// Checks which operator is used and returns the calculated value
	switch expr[1] {
	case "*":
		return strconv.FormatFloat(var1*var2, 'f', -1, 64)
	case "/":
		return strconv.FormatFloat(var1/var2, 'f', -1, 64)
	case "+":
		return strconv.FormatFloat(var1+var2, 'f', -1, 64)
	}

	// Returns an error string, since this code should never be reached
	return "ERR"
}

// Calculates an expression using the order of operations,
// except subtraction implies adding a negative
func (c *Calculator) simplifyExpr(i int, j int) {
	expr := &c.tree[i][j]

	// Expressions beginning with a minus imply the first value is negative
	if (*expr)[0] == "-" {
		*expr = append([]string{(*expr)[0] + (*expr)[1]}, (*expr)[2:]...)
	}

	// The current operator, used to index OPS
	currentOP := 0

	var result string

	// Continues until the result remains, i.e. len(expr)=1
	for len(*expr) > 1 {
		for i := 0; i < len(*expr); i++ {
			// Replaces "minus a positive" with "plus a negative"
			// and "minus a negative" with "plus a positive"
			if (*expr)[i] == "-" {
				(*expr)[i] = "+"
				if string((*expr)[i+1][0]) == "-" {
					(*expr)[i+1] = (*expr)[i+1][1:]
				} else {
					(*expr)[i+1] = "-" + (*expr)[i+1]
				}
			}

			// Checks until it finds the current operator of OPS, excluding "-"
			if (*expr)[i] == string(OPS[:3][currentOP]) {
				result = calcSimple((*expr)[i-1 : i+2])

				// Updates the expression based on the location of the operator
				switch i {
				case 1:
					*expr = append([]string{result}, (*expr)[3:]...)
				case len(*expr) - 2:
					*expr = append((*expr)[:len(*expr)-3], result)
				default:
					*expr = append(append((*expr)[:i-1], result), (*expr)[i+2:]...)
				}

				break
			}

			// If the loop is ending and still hasn't found the current operator,
			// then it assumes that that operator is no longer present and moves to the next one.
			if i == len(*expr)-1 {
				currentOP++
			}
		}
	}
}

// Calculates the result based on the entire parsed expression
func (c *Calculator) calc() {
	var eRef []string

	// Calculates expressions based on a descending hierarchy,
	// i.e. from most parenthetical to least
	for i := len(c.tree) - 1; i > -1; i-- {
		for j := 0; j < len(c.tree[i]); j++ {
			// Iterates through the current expression,
			// checks for references to expressions at higher levels,
			// and replaces them accordingly
			for k := 0; k < len(c.tree[i][j]); k++ {
				if strings.Contains(c.tree[i][j][k], "expr") {
					eRef = strings.Split(c.tree[i][j][k][5:], ".")
					ref1, err1 := strconv.ParseInt(eRef[0], 10, 0)
					p.HandleError(err1)
					ref2, err2 := strconv.ParseInt(eRef[1], 10, 0)
					p.HandleError(err2)

					c.tree[i][j][k] = c.tree[ref1][ref2][0]
				}
			}

			// Replaces expressions with their results
			c.simplifyExpr(i, j)
		}
	}
}

func main() {
	// Gets the expression from the user, parses it, organizes it into a tree
	// and calculates the result and prints it
	c := newCalculator(p.Input("Enter an equation (without spaces): "))
	c.lexer()
	c.parser()
	c.calc()
	fmt.Println("Answer is:", c.tree[0][0][0])
}
