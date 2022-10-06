/*
	Shows default values for common types in Go
*/

package main

import "fmt"

func main() {
	// Integer default value
	var a int
	fmt.Printf("Integer: \"%d\"\n", a)

	// String default value
	var b string
	fmt.Printf("String: \"%s\"\n", b)

	// Boolean default value
	var c bool
	fmt.Printf("Boolean: \"%t\"\n", c)

	// Float default value
	var d float64
	fmt.Printf("Float: \"%g\"\n", d)
}
